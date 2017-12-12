// ex8.7 mirrors a website to a given depth using multiple goroutines and
// rewrites local links.
package main

import (
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"golang.org/x/net/html"
)

// tokens is a counting semaphore used to
// enforce a limit of 20 concurrent requests.
var tokens = make(chan struct{}, 20)
var maxDepth int
var seen = make(map[string]bool)
var seenLock = sync.Mutex{}
var base *url.URL

func crawl(url string, depth int, wg *sync.WaitGroup) {
	defer wg.Done()

	tokens <- struct{}{} // acquire a token
	urls, err := visit(url)
	<-tokens //release token
	if err != nil {
		log.Printf("visit %s: %s", url, err)
	}

	if depth >= maxDepth {
		return
	}
	for _, link := range urls {
		seenLock.Lock()
		if seen[link] {
			seenLock.Unlock()
			continue
		}
		seen[link] = true
		seenLock.Unlock()
		wg.Add(1)
		go crawl(link, depth+1, wg)
	}
}

// Copied from gopl.io/ch5/outline2.
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}

func linkNodes(n *html.Node) []*html.Node {
	var links []*html.Node
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			links = append(links, n)
		}
	}
	forEachNode(n, visitNode, nil)
	return links
}

func linkURLs(linkNodes []*html.Node, base *url.URL) []string {
	var urls []string
	for _, n := range linkNodes {
		for _, a := range n.Attr {
			if a.Key != "href" {
				continue
			}
			link, err := base.Parse(a.Val)
			// ignore bad and non-local URLs
			if err != nil {
				log.Printf("skipping %q: %s", a.Val, err)
				continue
			}
			if link.Host != base.Host {
				//log.Printf("skipping %q: non-local host", a.Val)
				continue
			}
			urls = append(urls, link.String())
		}
	}
	return urls
}

// rewriteLocalLinks rewrites local links to be relative and links without
// extensions to point to index.html, eg /hi/there -> /hi/there/index.html.
func rewriteLocalLinks(linkNodes []*html.Node, base *url.URL) {
	for _, n := range linkNodes {
		for i, a := range n.Attr {
			if a.Key != "href" {
				continue
			}
			link, err := base.Parse(a.Val)
			if err != nil || link.Host != base.Host {
				continue // ignore bad and non-local URLs
			}
			// Clear fields so the url is formatted as /PATH?QUERY#FRAGMENT
			link.Scheme = ""
			link.Host = ""
			link.User = nil
			a.Val = link.String()
			n.Attr[i] = a
		}
	}
}

func visit(rawurl string) (urls []string, err error) {
	fmt.Println(rawurl)
	resp, err := http.Get(rawurl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("GET %s: %s", rawurl, resp.Status)
	}

	u, err := base.Parse(rawurl)
	if err != nil {
		return nil, err
	}
	if base.Host != u.Host {
		log.Printf("not saving %s: non-local", rawurl)
		return nil, nil
	}

	var body io.Reader
	contentType := resp.Header["Content-Type"]
	if strings.Contains(strings.Join(contentType, ","), "text/html") {
		doc, err := html.Parse(resp.Body)
		resp.Body.Close()
		if err != nil {
			return nil, fmt.Errorf("parsing %s as HTML: %v", u, err)
		}
		nodes := linkNodes(doc)
		urls = linkURLs(nodes, u) // Extract links before they're rewritten.
		rewriteLocalLinks(nodes, u)
		b := &bytes.Buffer{}
		err = html.Render(b, doc)
		if err != nil {
			log.Printf("render %s: %s", u, err)
		}
		body = b
	}
	err = save(resp, body)
	return urls, err
}

// If resp.Body has already been consumed, `body` can be passed and will be
// read instead.
func save(resp *http.Response, body io.Reader) error {
	u := resp.Request.URL
	filename := filepath.Join(u.Host, u.Path)
	if filepath.Ext(u.Path) == "" {
		filename = filepath.Join(u.Host, u.Path, "index.html")
	}
	err := os.MkdirAll(filepath.Dir(filename), 0777)
	if err != nil {
		return err
	}
	fmt.Println("filename:", filename)
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	if body != nil {
		_, err = io.Copy(file, body)
	} else {
		_, err = io.Copy(file, resp.Body)
	}
	if err != nil {
		log.Print("save: ", err)
	}
	// Check for delayed write errors, as mentioned at the end of section 5.8.
	err = file.Close()
	if err != nil {
		log.Print("save: ", err)
	}
	return nil
}

func main() {
	flag.IntVar(&maxDepth, "d", 3, "max crawl depth")
	flag.Parse()
	wg := &sync.WaitGroup{}
	if len(flag.Args()) == 0 {
		fmt.Fprintln(os.Stderr, "usage: mirror URL ...")
		os.Exit(1)
	}
	u, err := url.Parse(flag.Arg(0))
	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid url: %s\n", err)
	}
	base = u
	for _, link := range flag.Args() {
		wg.Add(1)
		go crawl(link, 1, wg)
	}
	wg.Wait()
}

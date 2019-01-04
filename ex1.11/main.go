// ex1.11 fetches some number (10 by default) of the top 1m URLs in parallel
// and reports their times and sizes.
package main

import (
	"archive/zip"
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Number of sites to fetch.
	nSites := 10
	if len(os.Args) > 1 {
		nSites, _ = strconv.Atoi(os.Args[1])
	}

	start := time.Now()
	ch := make(chan string)
	for _, url := range top1mSites(nSites) {
		go fetch(url, ch) // start a goroutine
	}
	for range top1mSites(nSites) {
		fmt.Println(<-ch) // receive from channel ch
	}
	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func printErrAndExit(err error) {
	fmt.Fprintf(os.Stderr, "%v", err)
	os.Exit(1)
}

func top1mSites(count int) []string {
	url := "http://s3.amazonaws.com/alexa-static/top-1m.csv.zip"
	var sites []string

	resp, err := http.Get(url)
	if err != nil {
		printErrAndExit(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		printErrAndExit(err)
	}

	zipReader, err := zip.NewReader(bytes.NewReader(body), int64(len(body)))
	if err != nil {
		printErrAndExit(err)
	}
	f, err := zipReader.File[0].Open() // only one file is downloaded
	if err != nil {
		printErrAndExit(err)
	}
	scan := bufio.NewScanner(f)
	for i := 0; scan.Scan() && i < count; i++ {
		fields := strings.Split(scan.Text(), ",")
		url := "http://" + fields[1]
		sites = append(sites, url)
	}
	f.Close()

	return sites
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}

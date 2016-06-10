// ex5.2 counts the frequency of different tags in an html document on stdin.
package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"golang.org/x/net/html"
)

func tagFreq(r io.Reader) (map[string]int, error) {
	freq := make(map[string]int, 0)
	z := html.NewTokenizer(os.Stdin)
	var err error
	for {
		type_ := z.Next()
		if type_ == html.ErrorToken {
			break
		}
		name, _ := z.TagName()
		if len(name) > 0 {
			freq[string(name)]++
		}
	}
	if err != io.EOF {
		return freq, err
	}
	return freq, nil
}

func main() {
	freq, err := tagFreq(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	for tag, count := range freq {
		fmt.Printf("%4d %s\n", count, tag)
	}
}

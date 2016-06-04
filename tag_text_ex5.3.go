package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"golang.org/x/net/html"
)

func printTagText(r io.Reader, w io.Writer) error {
	z := html.NewTokenizer(os.Stdin)
	var err error
	stack := make([]string, 20)
Tokenize:
	for {
		type_ := z.Next()
		switch type_ {
		case html.ErrorToken:
			break Tokenize
		case html.StartTagToken:
			b, _ := z.TagName()
			stack = append(stack, string(b))
		case html.TextToken:
			cur := stack[len(stack)-1]
			w.Write([]byte(fmt.Sprintf("<%s>", cur)))
			if cur != "script" {
				w.Write(z.Text())
			}
		case html.EndTagToken:
			stack = stack[:len(stack)-1]
		}
	}
	if err != io.EOF {
		return err
	}
	return nil
}

func main() {
	err := printTagText(os.Stdin, os.Stdout)
	if err != nil {
		log.Fatal(err)
	}
}

// ex11.1 computes counts of Unicode characters, and includes tests.
package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"unicode"
)

func charCount(r io.Reader) (
	runes map[rune]int, props map[string]int, sizes map[int]int, invalid int) {
	runes = make(map[rune]int)   // char frequency
	props = make(map[string]int) // unicode.Properties frequency
	sizes = make(map[int]int)    // rune length frequency
	invalid = 0                  // invalid char frequency

	in := bufio.NewReader(r)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatal("charCount: %s", err)
		}
		if r == unicode.ReplacementChar && n == 1 {
			invalid++
			continue
		}
		for name, rangeTable := range unicode.Properties {
			if unicode.In(r, rangeTable) {
				props[name]++
			}
		}
		runes[r]++
		sizes[n]++
	}
}

func main() {
	runes, props, sizes, invalid := charCount(os.Stdin)
	fmt.Println(props)
	fmt.Printf("rune\tcount\n")
	for c, n := range runes {
		fmt.Printf("%q\t%d\n", c, n)
	}
	fmt.Print("\nlen\tcount\n")
	for i, n := range sizes {
		if i > 0 {
			fmt.Printf("%d\t%d\n", i, n)
		}
	}
	fmt.Printf("\n%-30s count\n", "category")
	for cat, n := range props {
		fmt.Printf("%-30s %d\n", cat, n)
	}
	if invalid > 0 {
		fmt.Printf("\n%d invalid UTF-8 characters\n", invalid)
	}
}

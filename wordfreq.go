package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	freq := make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := scanner.Text()
		freq[word]++
	}
	for word, n := range freq {
		fmt.Printf("%-30s %d\n", word, n)
	}
}

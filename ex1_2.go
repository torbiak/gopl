package main

import (
	"fmt"
	"os"
)

func main() {
	for i, a := range os.Args {
		fmt.Println(i, a)
	}
}

package main

import (
	"fmt"
)

// ex4.3
func reverse(ints *[5]int) {
	for i := 0; i < len(ints)/2; i++ {
		end := len(ints) - i - 1
		ints[i], ints[end] = ints[end], ints[i]
	}
}

func main() {
	a := [...]int{1, 2, 3, 4, 5}
	reverse(&a)
	fmt.Println(a)
}

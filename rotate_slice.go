package main

import (
	"fmt"
)

func rotate_ints(ints []int) {
	first := ints[0]
	copy(ints, ints[1:])
	ints[len(ints)-1] = first
}

func main() {
	s := []int{1, 2, 3, 4, 5}
	rotate_ints(s)
	fmt.Println(s)
}

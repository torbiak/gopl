// ex5.19 returns a non-zero value using panic and recover, contradicting the
// function signature.
package main

import (
	"fmt"
)

func weird() (ret string) {
	defer func() {
		recover()
		ret = "hi"
	}()
	panic("omg")
}

func main() {
	fmt.Println(weird())
}

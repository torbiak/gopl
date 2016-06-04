package main

import (
	"fmt"
	"math"
)

func main() {
	max := 200
	logMax := math.Log(float64(max))
	for x := 0; x < max; x++ {
		log := math.Log(float64(x))
		fmt.Println(x, log, logMax, log/logMax, 255-log/logMax*255)
	}
}

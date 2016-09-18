package main

import (
	"fmt"
	"math/rand"

	"github.com/iocat/cs344/maxsubarr"
)

func init() {
	// Choose an algorithm to run
	maxsubarr.Register(maxsubarr.Kadane)
}

func main() {
	var a []int
	// Generate negative + positive numbers
	for i := 0; i < 50000; i++ {
		a = append(a, rand.Intn(1000)-500)
	}
	// Find the range of maximum sub array
	i, j, _ := maxsubarr.FindMax(a)
	fmt.Println(i, j)
}

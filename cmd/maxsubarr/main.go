package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/iocat/cs344/maxsubarr"
)

func init() {
	// Choose an algorithm to run
	maxsubarr.Register(maxsubarr.BruteForce)
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var a []int
	// Generate negative + positive numbers
	for i := 0; i < 10; i++ {
		a = append(a, rand.Intn(20)-10)
	}
	// Find the range of maximum sub array
	i, j, max, err := maxsubarr.FindMax(a)
	if err != nil {
		fmt.Printf("error: %s", err)
		return
	}
	fmt.Printf("Randomized Array: a = %v\n", a)
	fmt.Printf("The Maximum Contiguous Sub Array: a[%d,%d+1] = %v\n", i, j, a[i:j+1])
	fmt.Printf("The sum is: sum(a) = %d\n", max)
}

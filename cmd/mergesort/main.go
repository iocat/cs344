package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/iocat/cs344/sorting/mergesort"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var a []int
	// Generate negative + positive numbers
	for i := 0; i < 100; i++ {
		a = append(a, rand.Intn(20)-10)
	}
	fmt.Printf("Randomized Array: a = %v\n", a)
	mergesort.Sort(a)
	fmt.Printf("The Sorted Array = %v\n", a)
}

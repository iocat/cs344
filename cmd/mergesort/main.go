package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/iocat/cs344/sorting/mergesort"
)

func randomize(count int) []int {
	rand.Seed(time.Now().UnixNano())
	var a []int
	// Generate negative + positive numbers
	for i := 0; i < count; i++ {
		a = append(a, rand.Intn(20)-10)
	}
	return a
}

func serialized() {
	a := randomize(100)
	fmt.Printf("Randomized Array: a = %v\n", a)
	mergesort.Sort(a)
	fmt.Printf("The Sorted Array = %v\n", a)
}

func parallel() {
	b := randomize(100)
	fmt.Printf("Randomized Array: a = %v\n", b)
	mergesort.EmbarrassinglyParallel(b)
	fmt.Printf("The Sorted Array = %v\n", b)
}

func estimateCost(fn func()) time.Duration {
	st := time.Now()
	fn()
	en := time.Now()
	return en.Sub(st)
}

func main() {
	serialized()
}

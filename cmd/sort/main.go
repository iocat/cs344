package main

import (
	"fmt"
	"math/rand"
	"time"

	// sorting "github.com/iocat/cs344/sorting/quicksort"
	// sorting "github.com/iocat/cs344/sorting/bubblesort"
	// sorting "github.com/iocat/cs344/sorting/selectionsort"
	sorting "github.com/iocat/cs344/sorting/insertionsort"
	// sorting "github.com/iocat/cs344/sorting/mergesort"
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

func estimateCost(fn func([]int), arr []int) time.Duration {
	st := time.Now()
	fn(arr)
	en := time.Now()
	return en.Sub(st)
}

func main() {
	Nelems := 20
	a := randomize(Nelems)

	fmt.Printf("Randomized Array: a = %v\n", a)
	sorting.Sort(a)
	fmt.Printf("The Sorted Array = %v\n", a)
}

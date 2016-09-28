package main

import (
	"fmt"

	"github.com/iocat/cs344/popularnum"
)

func main() {
	a := []int{0, 2, 3, 4, 5, 5, 0, 0, 0, 0, 1}
	n, err := popularnum.MooreVoting(a)
	if err != nil {
		panic(err)
	}
	fmt.Println("The original array: ", a)
	fmt.Println("The popular element: ", n)
}

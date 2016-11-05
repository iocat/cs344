package quicksort

import (
	"math/rand"
	"time"
)

// Sort sorts the array using quicksort
// qsort sorts in-place, but not stable.
// Qsort has a non-deterministic running time.
// 		The expected running time is O(n*logn)
// 		The worst case running time is O(n^2), but it rarely occurs.
func Sort(a []int) {
	qsort(a)
}

func qsort(a []int) {
	if len(a) <= 1 {
		return
	}
	pi := partition(a)
	qsort(a[:pi])
	qsort(a[pi+1:])

}

// partition performs partition on the array and
// returns the index of the pivot
func partition(a []int) int {
	var (
		pi = pickpivot(a) /* The pivot index */
	)
	a[0], a[pi], pi = a[pi], a[0], 0 /* Swap the pivot to the front */
	for i := 1; i < len(a); i++ {
		// Swap smaller elements before the pivot
		if a[i] <= a[0] {
			pi++
			a[pi], a[i] = a[i], a[pi]
		}
	}
	a[0], a[pi] = a[pi], a[0] /* Swap the pivot to the middle */
	return pi
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

// pickpivot ramdomly choose a pivot index
func pickpivot(a []int) int {
	return rand.Intn(len(a))
}

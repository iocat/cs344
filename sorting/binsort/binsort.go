package binsort // (.aka bucket sort)
import "fmt"

func rangeCheck(l, u, val int) bool {
	if val > u || val < l {
		return false
	}
	return true
}

// SortWithRange sorts the array using binsort algorithm with the given range.
//
// This is not a comparision-based sort. It uses set data structure to store
// all the possible keys (in-range keys). Performance is gained only if the
// range of keys is known.
//
// Binsort's runtime is θ(n+d) where d = u - l, the range of array's values
func SortWithRange(a []int, l, u int) {
	buckets := make([][]int, u-l+1) /* make buckets */
	/* Store all elements in buckets */
	for i, bid := range a {
		if ok := rangeCheck(l, u, a[i]); !ok {
			panic(fmt.Errorf("value out of range: a[%d]=%d which is not in [%d,%d]", i, bid, l, u))
		}
		bucket := &buckets[bid]
		*bucket = append(*bucket, bid)
	}
	// Traverse the buckets and recover the original array in sorted order
	a = a[0:0]
	for _, bucket := range buckets {
		for i := range bucket {
			a = append(a, bucket[i])
		}
	}
}

// MakeFunc makes the sorting subroutine
func MakeFunc(l, u int) func([]int) {
	return func(a []int) {
		SortWithRange(a, l, u)
	}
}

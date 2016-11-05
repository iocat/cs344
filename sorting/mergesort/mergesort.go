package mergesort

// Sort sorts the array
// Mergesort sorts not-inplace, and is stable. It has a deterministic runtime.
// The average and worst case runtime is O(nlogn)
func Sort(si []int) {
	inarr := make([]int, len(si))
	msort(si, 0, len(si)-1, inarr)
}

// EmbarrassinglyParallel implements parallel mergesort
// NOTE: This algorithm spawns as many as n/2 goroutines to sort the array
// NOTE: it turns out to be not efficicent.
// Claim: the runtime is busy switching
// between goroutines without doing any actual useful work :(
func EmbarrassinglyParallel(si []int) {
	temp := make([]int, len(si))
	done := make(chan struct{})
	go paraMsort(si, 0, len(si)-1, temp, done)
	<-done
}

const (
	// SerializedMergingThreshold holds the maximum number of elements to
	// sort the array serially
	SerializedMergingThreshold = 5
)

// paraMsort sorts the array in a parallel fashion
func paraMsort(list []int, a, b int, temp []int, done chan<- struct{}) {
	mid := (a + b) / 2
	if b-a < SerializedMergingThreshold {
		msort(list, a, b, temp) /* normal mergesort */
	} else if a < b {
		subDone := make(chan struct{}) /* A waiter for sub goroutines */
		go paraMsort(list, a, mid, temp, subDone)
		go paraMsort(list, mid+1, b, temp, subDone)
		<-subDone
		<-subDone
		merge(list, a, b, temp)
	}
	done <- struct{}{}
}

// msort conducts merge sort sequentially
// temp represents the temporary array which is shared among the smaller
// subproblems
// a, b are the range of the sorting subproblem
func msort(list []int, a, b int, temp []int) {
	mid := (b + a) / 2
	if a < b { /* base case */
		msort(list, a, mid, temp)
		msort(list, mid+1, b, temp)
		merge(list, a, b, temp)
	}
}

// merge merges the two sorted segments
func merge(list []int, a, b int, temp []int) {
	copy(temp[a:b+1], list[a:b+1]) /* copy data to temporary space */
	mid := (a + b) / 2
	first, second := temp[a:mid+1], temp[mid+1:b+1] /* slice the temporary space*/
	var (
		i, j = 0, 0 /* i, j addresses the first and second half */
		k    = a    /* k addresses the original array */
	)
	for i < len(first) && j < len(second) {
		if first[i] < second[j] {
			list[k] = first[i]
			k, i = k+1, i+1
		} else {
			list[k] = second[j]
			k, j = k+1, j+1
		}
	}
	for i < len(first) { /* Append the first half to the array*/
		list[k] = first[i] /* if not exhausted */
		k, i = k+1, i+1
	}
	for j < len(second) { /* Append the second half */
		list[k] = second[j]
		k, j = k+1, j+1
	}
}

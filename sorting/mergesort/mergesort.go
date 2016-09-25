package mergesort

// Sort sorts the array
// The average runtime is O(nlogn)
func Sort(si []int) {
	inarr := make([]int, len(si))
	msort(si, 0, len(si)-1, inarr)
}

// msort conducts merge sort
// inarr represents the temporary array which is shared among the smaller
// subproblems
// a, b are the range of the sorting subproblem
func msort(list []int, a, b int, temp []int) {
	if a > b {
		panic("a is supposed to be bigger than b")
	}
	mid := (b + a) / 2
	if a < b { /* base case */
		msort(list, a, mid, temp)
		msort(list, mid+1, b, temp)
	}
	merge(list, a, b, temp)
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

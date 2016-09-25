package mergesort

// Sort sorts the array using mergesort
func Sort(si []int) {
	inarr := make([]int, len(si))
	msort(si, 0, len(si)-1, inarr)
}

// msort conducts merge sort
// inarr represents the index array
// a, b are the range of sorting
func msort(list []int, a, b int, temp []int) {
	if a > b {
		panic("a is supposed to be bigger than b")
	}
	mid := (b + a) / 2
	if a < b {
		msort(list, a, mid, temp)
		msort(list, mid+1, b, temp)
	}
	merge(list, a, b, temp)
}

// merge merges the two sorted segments
func merge(list []int, a, b int, temp []int) {
	copy(temp[a:b+1], list[a:b+1]) /* copy to temporary space */
	mid := (a + b) / 2
	first, second := temp[a:mid+1], temp[mid+1:b+1] /* slice the temporary space*/
	var (
		i, j = 0, 0 /* i, j addresses the first and second array */
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
	for i < len(first) {
		list[k] = first[i]
		k, i = k+1, i+1
	}
	for j < len(second) {
		list[k] = second[j]
		k, j = k+1, j+1
	}
}

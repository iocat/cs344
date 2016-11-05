package bubblesort

// Sort sorts the array using (improved) bubble sort algorithm.
//
// Bubblesort is in-place and stable.
// Best case running time is O(n), for already sorted array.
// Worst case running time and average case running time is O(n^2).
//      + Worst case occurs when the array is sorted in decreasing order.
//      + Average case analysis based on the lower bound of the number
//      of inversions the array can have for every input. The runtime
//      corresponds to how many inversions the algorithm reverses.
func Sort(a []int) {
	var (
		ref = a /* the subarray to bubble the biggest element to the right*/
	)
	for len(ref) > 1 {
		swapped := false /* marks whether there is a swap while traversing ref */
		for i := 0; i < len(ref)-1; i++ {
			if ref[i] >= ref[i+1] { /* bubble the bigger element to the end */
				ref[i], ref[i+1] = ref[i+1], ref[i]
				swapped = true
			}
		}
		if !swapped { /* already sorted array, there is no inversion */
			return
		}
		ref = ref[:len(ref)-1] /* iterate over the n-1 elements */
	}
}

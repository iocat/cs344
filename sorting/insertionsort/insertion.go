package insertionsort

// Sort sorts the array using the insertion sort algorithm
//
// The best case running time is O(n) for already sorted array
// The worst case running time is O(n^2) for array sorted in the opposite order
// The average case running time is O(n^2) which corresponds to the average
// number of inversions to reorder.
func Sort(a []int) {
	if len(a) <= 1 { /* the slice is sorted if there is at most 1 element */
		return
	}
	var (
		ref = a[:2] /* references the sorted subarray (to the left of the original) */
	)
	for { /* move the last element to the correct position*/
		cache := ref[len(ref)-1] /* cache the last element which is not in the correct position*/
		i := len(ref) - 2
		for ; i >= 0; i-- { /* shift all the bigger values to the far right */
			if ref[i] > cache {
				ref[i+1] = ref[i]
			} else {
				break
			}
		}
		ref[i+1] = cache /* move the last element to the correct position*/

		if len(ref) == len(a) {
			return /* stop when the reference array is as big as the original one */
		}
		ref = a[:len(ref)+1]
	}
}

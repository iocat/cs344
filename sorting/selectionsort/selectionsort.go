package selectionsort

// Sort sorts the array using selection sort algorithm
//
// The runtime is θ(n^2).
func Sort(a []int) {
	var (
		ref = a /* ref references the lower end of the original array */
	)
	for len(ref) >= 2 {
		/* Find the index of the smallest element in ref */
		minInd := 0 /* index of min */
		for i := 1; i < len(ref); i++ {
			if ref[i] < ref[minInd] {
				minInd = i
			}
		}
		ref[0], ref[minInd] = ref[minInd], ref[0] /* Swap the smallest to the front */
		ref = ref[1:len(ref)]                     /* recurse over the rest of the smaller sublist */
	}
}

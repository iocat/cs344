package maxsubarr

// FindMax runs the registered/default find maximum submatrix func
var FindMax = Kadane

// Register registers an algorithm
func Register(maxFunc MaxFunc) {
	FindMax = maxFunc
}

// MaxFunc represents a function that finds the maximum subarray
// and returns the beginning and the ending index of the array
// NOTE: the maximum sum subarray would be a[begin:end+1]
// Return index is based on zero based indexing scheme of golang
type MaxFunc func([]int) (int, int, int, error)

// BruteForce is Exported Brute Force solution
var BruteForce = bruteForce

func bruteForce(a []int) (int, int, int, error) {
	var (
		imax, jmax, max int
		n               = len(a)
	)
	if len(a) == 0 {
		return 0, 0, 0, nil
	}
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			temp := 0
			for k := i; k <= j; k++ {
				temp += a[k]
			}
			if temp > max {
				imax, jmax, max = i, j, temp
			}
		}
	}
	return imax, jmax, max, nil
}

func max(x int, y int) int {
	if x >= y {
		return x
	}
	return y
}

// Kadane represents the Kadane algorithm (added returned index)
var Kadane MaxFunc = kadane

// kadane implements kadane algorithm: imagine 2 sub arrays called tempmax and
// curmax
// tempmax dynamically captures a new set of contiguous elements.
// If tempmax results in a bigger subarray sum (comparing to curmax), then curmax
// will be replaced with tempmax
// Otherwise, curmax remains the same
// If tempmax results in a less than 0 set of contiguous elements, tempmax will
// start a new set
func kadane(a []int) (int, int, int, error) {
	var (
		imax, jmax, tempmax int
		isofar              int
		curmax              int
	)
	for j := range a {
		tempmax = max(tempmax+a[j], 0)
		if tempmax == 0 {
			// Starting over
			isofar = j + 1
		}
		curmax = max(curmax, tempmax)
		if curmax == tempmax {
			imax = isofar
			jmax = j
		}
	}

	return imax, jmax, curmax, nil
}

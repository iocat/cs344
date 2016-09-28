package popularnum

import "fmt"

// MooreVoting finds the index of most popular element in the list using the
// Moore's Voting algorithm
// The runtime is O(n), space complexity is O(1)
func MooreVoting(a []int) (int, error) {
	var (
		cand    = 0 // the candidate index
		counter = 0 // the current mode of the candidate
	)
	for i := range a {
		if counter == 0 {
			cand = i
			counter = 1
		} else if a[i] == a[cand] {
			counter++
		} else {
			counter--
		}
	}
	// Recheck the candiate
	counter = 0
	for _, n := range a {
		if n == a[cand] {
			counter++
		}
	}
	if counter > (len(a) / 2) {
		return cand, nil
	}
	return 0, fmt.Errorf("there is no popular element")

}

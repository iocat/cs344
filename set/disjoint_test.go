package set

import "testing"

func (d *DisjointSets) union(a, b int) *DisjointSets {
	d.Union(a, b)
	return d
}

func TestFind(t *testing.T) {
	var ds1 = New(10).union(0, 1).union(0, 2).union(3, 1)
	var tests = []struct {
		ds  *DisjointSets
		in  int
		out int
	}{
		{
			ds:  New(3),
			in:  1,
			out: 1,
		},
		{
			ds:  New(5),
			in:  4,
			out: 4,
		}, {
			ds:  ds1,
			in:  0,
			out: 0,
		}, {
			ds:  ds1,
			in:  2,
			out: 0,
		}, {
			ds:  ds1,
			in:  3,
			out: 0,
		}, {
			ds:  ds1,
			in:  4,
			out: 4,
		},
	}

	for _, tc := range tests {
		if res := tc.ds.Find(tc.in); tc.out != res {
			t.Errorf("invalid find(%d)=%d, expect %d", tc.in, res, tc.out)
		}
	}
}

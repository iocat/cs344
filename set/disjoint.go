package set

import "fmt"

// DisjointSets implements link-by-size forest of disjoint set unions/trees
// Each set knows its parent, the number of disjoint sets is the number of
// root nodes in the forest.
// The index of the disjoint sets should follow the 0-index scheme of slices
type DisjointSets []SetWithSize

// SetWithSize represents a set with a size
type SetWithSize struct {
	// The number of children (size).
	// Default to 1
	nelems int
	// The entry of the parent set.
	// If the set has no parent, parent is -1
	parent int
}

// New initializes a set of disjoint sets
func New(nelems int) *DisjointSets {
	ds := DisjointSets(make([]SetWithSize, nelems))
	for i := range ds {
		ds[i].parent = -1
		ds[i].nelems = 1
	}
	return &ds
}

// Find returns the entry of the set the element belongs to
// Find has no side effects.
func (d *DisjointSets) Find(a int) int {
	var parent = (*d)[a].parent
	if parent == -1 { // if this set has no parent
		return a
	}
	return d.Find(parent)
}

// Union joins two disjoint sets.
func (d *DisjointSets) Union(a int, b int) int {
	if a == b {
		panic(fmt.Errorf("cannot operate union on the same element %s", a))
	}
	var (
		id1      = d.Find(a)
		id2      = d.Find(b)
		parentid int
		child    *SetWithSize
		parent   *SetWithSize
	)
	if (*d)[id1].nelems < (*d)[id2].nelems { // decide which is the child set, which is the parent set
		child, parent, parentid = &((*d)[id1]), &((*d)[id2]), id2
	} else {
		child, parent, parentid = &((*d)[id2]), &((*d)[id1]), id1
	}
	parent.nelems += child.nelems
	child.parent = parentid

	fmt.Println(child, parent, parentid)
	return parentid
}

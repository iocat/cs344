package graph

import (
	"errors"
	"fmt"
)

var (
	// ErrLowBound when the vertex's index is smaller than 0
	ErrLowBound = errors.New("vertex out of range: lower bound violation")
	// ErrUpBound when the vertex's index is bigger than the maximum
	ErrUpBound = errors.New("vertex out of range: upper bound violation")
)

// BoundCheck checks whether the vertex index exceeds the allowed bound by the
// graph. This is an helper function the client graph package can use to check bound
//
// Also see MustBoundCheck
func BoundCheck(g Interface, v int) error {
	switch {
	case v < 0:
		return ErrLowBound
	case v >= g.Nvertices():
		return ErrUpBound
	default:
		return nil
	}
}

// MustBoundCheck is BoundCheck, but panics when error is not nil
func MustBoundCheck(g Interface, v int) {
	if BoundCheck(g, v) != nil {
		panic(fmt.Errorf("vertex %d is out of range, expected range: %d to %d", v, 0, g.Nvertices()-1))
	}
}

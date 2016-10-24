// TODO: no testing, consider testing the package

package adjacentlist

import (
	"errors"
	"fmt"
)

var (
	errLowBound = errors.New("vertex out of range: lower bound violation")
	errUpBound  = errors.New("vertex out of range: upper bound violation")
)

const (
	// MaxVertices is the maximum amount of vertices allowed
	MaxVertices = 1000
)

// boundCheck checks for the lower bound and upper bound error
func (g *Graph) boundCheck(v int) error {
	switch {
	case v < 0:
		return errLowBound
	case v >= g.Nvertices():
		return errUpBound
	default:
		return nil
	}
}

// Graph represents a weighted graph using adjacent list.
// This graph is weighted and does not care about directions of edge.
// The indeces of every vertice are assumed to be in increasing order. The
// number of vertices is scaled automatically with insertions, but limited to
// MaxVertices.
//
// Implementation is **partly**
// based on The Algorithm Design Manual [by Steven S. Skiena].
type Graph struct {
	Nodes [][]EdgeNode /* list of node's edges */
}

// EdgeNode represents the edge
type EdgeNode struct {
	Y      int /* the tail of the edge */
	Weight int /* the weight of the edge */
}

// New creates a new graph initialized to nvertices vertices
func New(nvertices int, directed bool) *Graph {
	if nvertices < 0 || nvertices > MaxVertices {
		panic(fmt.Errorf("invalid number of vertices, expected [0, %d], got %d", MaxVertices, nvertices))
	}
	return &Graph{
		Nodes: make([][]EdgeNode, nvertices, MaxVertices),
	}
}

// Nvertices returns the number of vertices the graph has
func (g *Graph) Nvertices() int {
	return len(g.Nodes)
}

// Weight is the weight function of this graph
// Weight returns (0, false) if the edge does not exist. Otherwise, it returns
// the actual (weight,true).
// If the vertex index is out of bound, the method panics
// This runs in O(|V|)
func (g *Graph) Weight(x, y int) (int, bool) {
	nv := g.Nvertices()
	if g.boundCheck(x) != nil {
		panic(fmt.Errorf("vertex %d is out of range, expected range: %d to %d", x, 0, nv-1))
	}
	for _, edge := range g.Nodes[x] {
		if edge.Y == y {
			return edge.Weight, true
		}
	}
	return 0, false
}

// InDegree returns the in-degree which corresponds to one vertex
// This runs in O(|E|+|V|) because it traverses the entire graph
func (g *Graph) InDegree(v int) int {
	var in = 0
	if g.boundCheck(v) != nil {
		nv := g.Nvertices()
		panic(fmt.Errorf("vertex %d is out of range, expected range: %d to %d", v, 0, nv-1))
	}
	for _, edges := range g.Nodes {
		for _, edge := range edges {
			if edge.Y == v {
				in++
			}
		}
	}
	return in
}

// OutDegree returns the out degree corresponds to one vertex
// This runs in O(1)
func (g *Graph) OutDegree(v int) int {
	nv := g.Nvertices()
	if g.boundCheck(v) != nil {
		panic(fmt.Errorf("vertex %d is out of range, expected range: %d to %d", v, 0, nv-1))
	}
	return len(g.Nodes[v])
}

// checkAndExtend extends the length of the vertex list. The vertex list is
// limited to MaxVertices
// This function runtime is O(|V|) because it might copy the adjacent list
func (g *Graph) checkAndExtend(v int) {
	var (
		nv = g.Nvertices()
	)
	if err := g.boundCheck(v); err == errUpBound {
		newnv := v + 1
		if newnv > MaxVertices {
			panic(fmt.Errorf("need more vertices than allowed, got %d, expected [0,%d]", newnv, MaxVertices))
		}
		more := make([][]EdgeNode, newnv-nv) /* add more vertex list */
		g.Nodes = append(g.Nodes, more...)   /* to the adj list*/
	} else if err == errLowBound {
		panic(fmt.Errorf("vertex %d is out of range, expected range: %d to %d", v, 0, nv-1))
	}
}

// InsertEdge adds an edge to the graph.
// If the inserted edge is directed, only one edge is inserted.
// Otherwise, two edges (x,y) and (y,x) are inserted.
// If the inserted edge is out of bound the total number of vertices is
// scaled up and limited to MaxVertices.
//
// Assume that no new vertex is added, this runs in O(1) in the best case
// scenario ( the edge is added to the end ) and O(|V|) in the worst case
// ( copy the entire adjacent list to a new location )
func (g *Graph) InsertEdge(x, y, weight int, directed bool) {
	newEdge := EdgeNode{
		Y:      y,
		Weight: weight,
	}
	g.checkAndExtend(x)                      /* guarantee the capacity */
	g.checkAndExtend(y)                      /* of the graph */
	g.Nodes[x] = append(g.Nodes[x], newEdge) /* add to the end of the adj list */
	if !directed {
		g.InsertEdge(y, x, weight, true) /* insert the undirected edge */
		return
	}
}

// DeleteEdge deletes an edge. If the inserted edge is directed, only one edge
// is deleted. Otherwise, two edges are deleted.
func (g *Graph) DeleteEdge(x, y int, directed bool) {

}

// TODO: no testing, consider testing the package

package adjacentlist

import (
	"fmt"

	"github.com/iocat/cs344/graph"
)

func init() {
	graph.Loads(New) /* load the package with side effect */
}

// Graph represents a weighted graph implemented using adjacent list.
// This graph is weighted and does not care about directions of edges.
// The indeces of every vertice are assumed to be in increasing order. The
// number of vertices is scaled automatically with insertions, but limited to
// graph.MaxVertices.
//
// Implementation is **partly**
// based on The Algorithm Design Manual [by Steven S. Skiena].
//
// The space complexity is: O(|V| + 2|E|)
type Graph struct {
	Nodes [][]EdgeNode /* list of node's edges */
}

// EdgeNode represents the edge
type EdgeNode struct {
	Y      int /* the head of the edge */
	Weight int /* the weight of the edge */
}

// New creates a new graph initialized to nvertices vertices
func New(nvertices int) (graph.Interface, error) {
	if nvertices < 0 || nvertices > graph.MaxVertices {
		return nil, fmt.Errorf("vertex capacity exceeded: expected [0, %d], got %d", graph.MaxVertices, nvertices)
	}
	return &Graph{
		Nodes: make([][]EdgeNode, nvertices, graph.MaxVertices),
	}, nil
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
	graph.MustBoundCheck(g, x)
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
	graph.MustBoundCheck(g, v)
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
	graph.MustBoundCheck(g, v)
	return len(g.Nodes[v])
}

// checkAndExtend extends the length of the vertex list. The vertex list is
// limited to graph.MaxVertices
// This function runtime is O(|V|) because it might copy the adjacent list
func (g *Graph) checkAndExtend(v int) {
	var (
		nv = g.Nvertices()
	)
	if err := graph.BoundCheck(v); err == graph.ErrUpBound {
		newnv := v + 1
		if newnv > graph.MaxVertices {
			panic(fmt.Errorf("need more vertices than allowed, got %d, expected [0,%d]", newnv, graph.MaxVertices))
		}
		more := make([][]EdgeNode, newnv-nv) /* add more vertex list */
		g.Nodes = append(g.Nodes, more...)   /* to the adj list*/
	} else if err == errLowBound {
		panic(fmt.Errorf("vertex %d is out of range, expected range: %d to %d", v, 0, nv-1))
	}
}

// InsertEdge implements graph.Interface
//
// Assume that no new vertex is added, this runs in O(1) in the best case
// scenario ( the edge is added to the end ) and in O(|V|) in the worst case
// ( copy the entire adjacent list to a new location )
func (g *Graph) InsertEdge(x, y, weight int, directed bool) {
	g.checkAndExtend(x) /* guarantee the length */
	g.checkAndExtend(y) /* of the adjacent list */

	newEdge := EdgeNode{
		Y:      y,
		Weight: weight,
	}
	g.Nodes[x] = append(g.Nodes[x], newEdge) /* add to the end of the adj list */
	if !directed {
		g.InsertEdge(y, x, weight, true) /* insert the undirected edge */
		return
	}
}

// DeleteEdge implements the graph.Interface
func (g *Graph) DeleteEdge(x, y int, directed bool) {
	// TODO
}

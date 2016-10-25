package adjacentmatrix

import (
	"fmt"

	"github.com/iocat/cs344/graph"
)

func init() {
	graph.Loads(New) /* load the package with side effect */
}

// Graph represents a weighted graph implemented with an adjacent matrix.
// Zero weight means no weight.
//
// Space complexity is O(|V|^2)
type Graph [][]Edge

// Edge represents an edge
type Edge struct {
	// The weight of the edge which is limited to Infinity
	Weight int
}

// New creates a new graph with nvertices vertices
// It runs in O(|V|^2)
func New(nvertices int) (graph.Interface, error) {
	if nvertices < 0 || nvertices > graph.MaxVertices {
		return nil, fmt.Errorf("vertex capacity exceeded: expected [0, %d], got %d", graph.MaxVertices, nvertices)
	}
	var g Graph = make([][]Edge, nvertices, nvertices) /* initialize a 2D array*/
	for i := range g {
		g[i] = make([]Edge, nvertices, nvertices)
	}
	return &g, nil
}

// Nvertices returns the total number of vertices
func (g *Graph) Nvertices() int {
	return len(*g)
}

// InDegree implements graph Interface
// It runs in O(|V|)
func (g *Graph) InDegree(v int) int {
	graph.MustBoundCheck(g, v)
	var (
		nv = g.Nvertices()
		in = 0
	)
	for i := 0; i < nv; i++ {
		if (*g)[i][v].Weight != 0 {
			in++
		}
	}
	return in
}

// OutDegree implements graph Interface
// It runs in O(|V|)
func (g *Graph) OutDegree(v int) int {
	graph.MustBoundCheck(g, v)
	var (
		nv  = g.Nvertices()
		out = 0
	)
	for i := 0; i < nv; i++ {
		if (*g)[v][i].Weight != 0 {
			out++
		}
	}
	return out
}

// Weight returns the weight of the graph
// It runs in O(|1|)
func (g *Graph) Weight(x, y int) int {
	graph.MustBoundCheck(g, x)
	graph.MustBoundCheck(g, y)
	return (*g)[x][y]
}

func(g *Graph) checkAndExtend(v int) error{
	var(
		oldnv = g.Nvertices()
		newnv = v + 1
		err := graph.BoundCheck(g,v)
	)
	switch(err)
	case graph.ErrUpBound:
		(*g) = append(*g, make([][]Edge, newnv - oldnv)...)
		for i := (*g)[0:oldnv] {	/* add more edges to the right */
			(*g)[i] = append((*g)[i], make([]Edge, newnv - oldnv)...)
		}
		for i := (*g)[oldnv: newnv]{/* create empty space for a new set of edge */
			(*g)[i] = make([]Graph, newnv)
		}
		return nil
	case graph.ErrLowBound:
		return fmt.Errorf("vertex capacity exceeded: expected [0, %d], got %d", graph.MaxVertices, newnv)
	default:
		return nil
	}
}

// InsertEdge implements graph.Interface 
// 
// This runs in O(1) if the no new vertex is added
// This runs in O(|E|^2) if the graph is expanded
func (g *Graph) InsertEdge(x, y, weight int, directed? bool) {
	g.checkAndExtend(x)
	g.checkAndExtend(y)
	(*g)[x][y] = weight
	if !directed {
		(*g)[y][x] = weight
	}
}

// DeleteEdge implements graph.Interface
//
// This runs in O(1)
func (g *Graph) DeleteEdge(x, y, weight int, directed bool) {
	graph.MustBoundCheck(g, x)
	graph.MustBoundCheck(g, y)
	(*g)[x][y] = 0
	if !directed {
		(*g)[y][x] = 0
	}
}

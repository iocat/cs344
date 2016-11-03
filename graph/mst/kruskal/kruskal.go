package kruskal

import (
	"sort"

	graph "github.com/iocat/cs344/graph/adjacentlist"
	"github.com/iocat/cs344/set"
)

type edges []graph.Edge

func (es edges) Len() int {
	return len(es)
}

func (es edges) Less(a int, b int) bool {
	return es[a].Weight < es[b].Weight
}

func (es edges) Swap(a int, b int) {
	es[a], es[b] = es[b], es[a]
}

// Find finds the minimum spanning tree
// NOTE: Find assumes the graph is non-directed even if the representation of
// the adjacentlist.Graph is directed
func Find(g *graph.Graph) *graph.Graph {
	var (
		mst   *graph.Graph             /* the minimal spanning tree */
		set   = set.New(g.Nvertices()) /* union-find disjoint sets */
		edges = edges(g.Edges())       /* the edges of the original tree */
	)
	mst, err := graph.New(g.Nvertices())
	if err != nil {
		panic(err)
	}
	sort.Sort(edges)             /* sort the edges */
	for _, edge := range edges { /* for each edge*/
		if set.Find(edge.X) != set.Find(edge.Y) { /* check if they do not belong to a same disjoint tree */
			mst.InsertEdge(edge.X, edge.Y, edge.Weight, true) /* add the edge to the MST to form a bigger tree*/
			set.Union(edge.X, edge.Y)                         /* combine 2 disjoint trees that have x and y */
		}
	}
	return mst
}

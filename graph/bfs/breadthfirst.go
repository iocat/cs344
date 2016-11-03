package bfs

import (
	graph "github.com/iocat/cs344/graph/adjacentlist"
)

// Vertex represents a vertex to be visited
type Vertex struct {
	// The source going to the vertex
	Source int
	// The vertex
	Dest int
	// The weight of the edge leading to Dest
	Weight int
	// The sub-edges from that destination
	edges []graph.EdgeNode
}

// VisitFn visits a vertex from the source vertex by an edge with weight
// VisitFn returns whether to stop traversal or not
type VisitFn func(source, toVisit, weight int) bool

// Traverse performs bredth-first search on the graph using source as the starting
// vertex
// This runs in O(|V|+|E|)
// NOTE: that the graph only traverse to all the reachable vertex from source,
// Traverse assumes g is directed
func Traverse(g *graph.Graph, source int, visitFn VisitFn) {
	var (
		q       = make([]Vertex, 0, g.Nvertices()) // the visiting node queue
		visited = make([]bool, g.Nvertices())      // check if the node is visited
	)
	// initialization
	for i := range visited {
		visited[i] = false
	}
	// add the root node
	q = append(q, Vertex{
		edges:  g.Nodes[source],
		Source: -1,
		Dest:   source,
		Weight: 0,
	})
	for len(q) != 0 {
		var toVisit Vertex
		toVisit, q = q[0], q[1:len(q)] // pop the vertex from the queue
		if visited[toVisit.Dest] {
			continue
		}
		if visitFn(toVisit.Source, toVisit.Dest, toVisit.Weight) { // visit the vertex
			break
		}
		visited[toVisit.Dest] = true      // set the vertex as visited
		for _, e := range toVisit.edges { // enque all subedges
			q = append(q, Vertex{
				Source: toVisit.Dest,
				Dest:   e.Y,
				Weight: e.Weight,
				edges:  g.Nodes[e.Y],
			})
		}
	}
}

package main

import (
	"fmt"

	graph "github.com/iocat/cs344/graph/adjacentlist"
	traversal "github.com/iocat/cs344/graph/traversal/bfs"
	// traversal "github.com/iocat/cs344/graph/traversal/dfs"
)

func main() {
	g, err := graph.New(6)
	if err != nil {
		panic(err)
	}
	g.InsertEdge(0, 1, 8, true)
	g.InsertEdge(0, 5, 6, true)
	g.InsertEdge(1, 5, 3, true)
	g.InsertEdge(1, 4, 7, true)
	g.InsertEdge(4, 5, 5, true)
	g.InsertEdge(1, 3, 1, true)
	g.InsertEdge(1, 2, 4, true)
	g.InsertEdge(2, 3, 5, true)
	g.InsertEdge(3, 4, 3, true)
	fmt.Println("Original graph: ", g)
	fmt.Println("DFS: ")
	traversal.Traverse(g, 0, func(source, toVisit, weight int) bool {
		fmt.Printf("visit %d from %d, weight %d\n", toVisit, source, weight)
		return false
	})
}

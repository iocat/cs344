package main

import (
	"fmt"

	graph "github.com/iocat/cs344/graph/adjacentlist"
	mst "github.com/iocat/cs344/graph/mst/prim"
	// Change to switch algorithm
	// mst "github.com/iocat/cs344/graph/mst/kruskal"
)

func main() {
	g, err := graph.New(6)
	if err != nil {
		panic(err)
	}
	g.InsertEdge(0, 1, 8, false)
	g.InsertEdge(0, 5, 6, false)
	g.InsertEdge(1, 5, 3, false)
	g.InsertEdge(1, 4, 7, false)
	g.InsertEdge(4, 5, 5, false)
	g.InsertEdge(1, 3, 1, false)
	g.InsertEdge(1, 2, 4, false)
	g.InsertEdge(2, 3, 5, false)
	g.InsertEdge(3, 4, 3, false)
	fmt.Println("Original graph(non-directed): ", g)
	fmt.Println("MST: ", mst.Find(g))
}

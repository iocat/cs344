package dijkstra

import (
	"container/heap"
	"fmt"
	"math"

	graph "github.com/iocat/cs344/graph/adjacentlist"
)

const infinite = math.MaxInt32

// Vertex represents the non-visited vertex which is stored in the minimum
// cost priority heap
type Vertex struct {
	ID     int /* The id of this node */
	Parent int /* The ID of the parent vertex, -1 for the source */
	TotalC int /* The total cost from source node */
}

// toVisitQ represents the priority min heap
type toVisitQ struct {
	unvisited []*Vertex
	// The inverted index to map the ID of the vertex to the
	// actual location on the heap
	invIndex map[int]int
}

func (q toVisitQ) Len() int {
	return len(q.unvisited)
}

func (q toVisitQ) Less(a, b int) bool {
	return q.unvisited[a].TotalC < q.unvisited[b].TotalC
}

func (q toVisitQ) Swap(a, b int) {
	q.invIndex[q.unvisited[a].ID], q.invIndex[q.unvisited[b].ID] = b, a
	q.unvisited[a], q.unvisited[b] = q.unvisited[b], q.unvisited[a]
}

func (q *toVisitQ) Push(v interface{}) {
	q.unvisited = append(q.unvisited, v.(*Vertex))
	q.invIndex[v.(*Vertex).ID] = len(q.unvisited) - 1
}

func (q *toVisitQ) Pop() interface{} {
	var val = q.unvisited[len(q.unvisited)-1]
	q.unvisited = q.unvisited[0 : len(q.unvisited)-1]
	delete(q.invIndex, val.ID)
	return val
}

func (q toVisitQ) Index(vertexID int) (int, bool) {
	index, ok := q.invIndex[vertexID]
	if !ok {
		return -1, false
	}
	return index, true
}

// initDijkstra initializes the heap
func initDijkstra(q *toVisitQ, sqt []*Vertex, n int, root int) {
	if root > n {
		panic(fmt.Errorf("root cannot be bigger than n (%d>%d)", root, n))
	}
	for i := 0; i < n; i++ {
		newV := Vertex{
			ID:     i,
			Parent: -1,
			TotalC: infinite,
		}
		if i == root {
			newV.TotalC = 0 // make the root node popped first
		}
		heap.Push(q, &newV)
		sqt[i] = &newV
	}
}

// SPTree represents the shortest path tree
type SPTree []*Vertex

func (t SPTree) String() string {
	var res string
	var sep string
	for i := range t {
		if i == 0 {
			sep = ""
		} else {
			sep = ","
		}
		if t[i].Parent == -1 {
			res = fmt.Sprintf("%s%s (root %d)", res, sep, t[i].ID)
		} else {
			res = fmt.Sprintf("%s%s (%d->%d, Σw: %d)", res, sep, t[i].Parent, t[i].ID, t[i].TotalC)
		}
	}
	return res
}

// Path returns the shortest path from the source vertex to a chosen vertex
func (t SPTree) Path(toVertex int) SPTree {
	temp := t[toVertex]
	sp := make(SPTree, 0, len(t))
	for temp.Parent != -1 { // while not reached the root
		sp = append(sp, temp)
		temp = t[temp.Parent]
	}
	for i := 0; i < len(sp)/2; i++ {
		sp[i], sp[len(sp)-1-i] = sp[len(sp)-1-i], sp[i]
	}
	return sp
}

// Find builds the shortest path tree using Dijkstra algorithm.
// A greedy algorithm that gradually builds the shortest path tree
//
// The runtime is O((|V|+|E|)*log|E|) where |V|log|E| accounts for heap.Fix
// whenever we check an out-going edge. And |E|log|E| accounts for heap.Pop
// to get the smallest cost vertex.
func Find(g *graph.Graph, source int) SPTree {
	var (
		n = g.Nvertices()
		q = &toVisitQ{
			unvisited: make([]*Vertex, 0, n),
			invIndex:  make(map[int]int),
		} /* the priority queue that store unvisited vertices */
		spt = make(SPTree, n) /* the shortest path tree */
	)
	initDijkstra(q, spt, n, source)
	for q.Len() != 0 {
		visit := heap.Pop(q).(*Vertex)
		outEdges := g.Nodes[visit.ID]
		for _, edge := range outEdges {
			yid := edge.Y
			yInvID, ok := q.Index(edge.Y) // find the inverted index of Y
			if !ok {
				continue /* skip the visited vertex */
			}
			if newC := edge.Weight + visit.TotalC; newC < spt[yid].TotalC { /* we found a better weight */
				spt[yid].TotalC = newC     /* reset the shortest cost to this vertex*/
				spt[yid].Parent = visit.ID /* set the parent of this vertex */
				heap.Fix(q, yInvID)        /* try to move this to the smallest position */
			}
		}
	}
	return spt
}

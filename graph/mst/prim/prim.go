package prim

import (
	"container/heap"
	"math"
	"math/rand"

	graph "github.com/iocat/cs344/graph/adjacentlist"
)

const infinite = math.MaxInt32

// visitedNode represents the node to visit each time. Each node corresponds to
// a small edge ( blue edge ) at a particular instace of the MST
type visitedNode struct {
	from   int /* the source node */
	node   int /* the current node */
	weight int /* the minimal weight that connects from <-> node */
}

// visitedQueue implements the heap.Interface
// it also supports fast retrieval of node's index
type visitedQueue struct {
	// the priority queue as heap
	nodes []visitedNode
	// m is the inverted index, which maps the node to the index in the priority
	// queue in order to have O(1) retrieval of node
	m map[int]int
}

func (q visitedQueue) Len() int {
	return len(q.nodes)
}

func (q *visitedQueue) Swap(a, b int) {
	q.m[q.nodes[a].node], q.m[q.nodes[b].node] = b, a
	q.nodes[a], q.nodes[b] = q.nodes[b], q.nodes[a]
}

func (q visitedQueue) Less(a, b int) bool {
	return q.nodes[a].weight < q.nodes[b].weight
}

func (q *visitedQueue) Push(x interface{}) {
	q.m[x.(visitedNode).node] = len(q.nodes)
	q.nodes = append(q.nodes, x.(visitedNode))
}

func (q *visitedQueue) Pop() interface{} {
	x := q.nodes[q.Len()-1]
	q.nodes = q.nodes[:q.Len()-1]
	delete(q.m, x.node)
	return x
}

// init initializes the queue
func (q *visitedQueue) init(n int, root int) {
	for i := 0; i < n; i++ {
		node := visitedNode{
			from: -1,
			node: i,
		}
		if i == root {
			node.weight = 0
		} else {
			node.weight = infinite
		}
		heap.Push(q, node)
	}
}

// getNode returns the index of the node in the priority queue, the second
// return value is whether the node exists or not
func (q *visitedQueue) getIndex(node int) (int, bool) {
	index, ok := q.m[node]
	return index, ok
}

// Find finds the minimal spanning tree using Prim's algorithm
// Idea: grow the MST and apply blue rule to each vertices
func Find(g *graph.Graph) *graph.Graph {
	var (
		n        = g.Nvertices()
		mst, err = graph.New(n)
		// a random node to make the root
		root = rand.Intn(n)

		q = &visitedQueue{
			nodes: make([]visitedNode, 0, n),
			m:     make(map[int]int),
		}
	)
	if err != nil {
		panic(err)
	}

	q.init(n, root) /* Initialize the visited priority queue */
	for q.Len() != 0 {
		/* grab the vertex connected by the locally smallest weight */
		toVisit := heap.Pop(q).(visitedNode)
		if toVisit.node != root {
			// Add the smallest weight to the MST for a non-root node visit
			mst.InsertEdge(toVisit.from, toVisit.node, toVisit.weight, false)
		}
		for _, edge := range g.Nodes[toVisit.node] { /* for each outgoing edges */
			headIndex, ok := q.getIndex(edge.Y) /* get the tail's index in the visit queue */
			if !ok {
				continue /* skip the node already visited */
			}
			if edge.Weight < q.nodes[headIndex].weight { /* we found a better weight to non-visited vertex */
				q.nodes[headIndex].weight = edge.Weight /* we replace */
				q.nodes[headIndex].from = toVisit.node  /* the edge with the smaller one*/
				heap.Fix(q, headIndex)                  /* try to move this edge to the minimum pos, if possible */
			}
		}
	}
	return mst
}

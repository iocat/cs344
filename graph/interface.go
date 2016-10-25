package graph

const (
	// MaxVertices is the maximum size of the adjacent matrix
	MaxVertices = 1000
)

// Interface represents a graph interface
type Interface interface {
	// returns the total number of vertices in the graph
	Nvertices() int
	// Weight returns the weight of the edge between 2 vertices. If
	// no edge exists, weight returns false as the second argument
	Weight(x, y int) (int, bool)
	// InDegree returns the in-degree of a vertex
	InDegree(v int) int
	// OutDegree returns the out-degree of a vertex
	OutDegree(v int) int
	// InsertEdge adds an edge to the graph.
	// If the inserted edge is directed, only one edge is inserted.
	// Otherwise, two edges (x,y) and (y,x) are inserted.
	// If the inserted edge is out of bound the total number of vertices is
	// scaled up and limited to graph.MaxVertices.
	InsertEdge(x, y, weight int, directed bool)
	// DeleteEdge removes an edge. If the inserted edge is directed, only one edge
	// is deleted. Otherwise, two edges are deleted.
	DeleteEdge(x, y int, directed bool)
}

type createGraphFunc func(nvertices) (Interface, error)

var pickedGraphCreator createGraphFunc

// Loads loads the graph constructor for this package
func Loads(graphCFn createGraphFunc) {
	pickedGraphCreator = graphCFn
}

// New creates a new graph
func New(nvertices int) (Interface, error) {
	if pickedGraphCreator == nil {
		panic("did not load the graph constructor function: please import one (either adjacentlist or adjacentmatrix)")
	}
	return pickedGraphCreator(nvertices)
}

package graphs

import "github.com/x1n13y84issmd42/dm/graphs/nodes"

// IGraph is an interface for unweighted graphs.
type IGraph interface {
	AddEdge(v1 nodes.INode, v2 nodes.INode)
	Node(nID nodes.NodeID) nodes.INode
}

// IWGraph is an interface for weighted graphs.
type IWGraph interface {
	AddEdge(v1 nodes.INode, w float64, v2 nodes.INode)
	Node(nID nodes.NodeID) nodes.INode
}

// DAGraph is an unweighted directed graph represented as an adjacency list.
type DAGraph struct {
	A *nodes.AdjacencyList
}

// UAGraph is an unweighted undirected graph represented as an adjacency list.
type UAGraph struct {
	A *nodes.AdjacencyList
}

// WDAGraph is a weighted directed graph represented as an adjacency list.
type WDAGraph struct {
	A *nodes.AdjacencyList
}

// WUAGraph is a weighted undirected graph represented as an adjacency list.
type WUAGraph struct {
	A *nodes.AdjacencyList
}

// NewDAGraph creates a new DAGraph instance.
func NewDAGraph() *DAGraph {
	return &DAGraph{
		A: nodes.NewAdjacencyList(),
	}
}

// NewUAGraph creates a new UAGraph instance.
func NewUAGraph() *UAGraph {
	return &UAGraph{
		A: nodes.NewAdjacencyList(),
	}
}

// AddEdge creates an edge between v1 and v2 nodes.
func (graph *DAGraph) AddEdge(v1 nodes.INode, v2 nodes.INode) {
	graph.A.AddEdge(v1, v2)
}

// Node returns a node instance by it's ID.
func (graph *DAGraph) Node(nID nodes.NodeID) nodes.INode {
	return graph.A.GetNode(nID)
}

// AddEdge creates an edge between v1 and v2 nodes.
func (graph *UAGraph) AddEdge(v1 nodes.INode, v2 nodes.INode) {
	graph.A.AddEdge(v1, v2)
	graph.A.AddEdge(v2, v1)
}

// Node returns a node instance by it's ID.
func (graph *UAGraph) Node(nID nodes.NodeID) nodes.INode {
	return graph.A.GetNode(nID)
}

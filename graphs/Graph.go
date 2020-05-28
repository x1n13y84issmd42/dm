package graphs

import "github.com/x1n13y84issmd42/dm/graphs/nodes"

// IGraph is an interface for unweighted graphs.
type IGraph interface {
	nodes.IAdjacency
	AddEdge(v1 nodes.Node, v2 nodes.Node)
	Node(nID nodes.NodeID) nodes.Node
	OutEdges(nID nodes.NodeID) []nodes.IEdge
}

// IWGraph is an interface for weighted graphs.
type IWGraph interface {
	nodes.IWAdjacency
	AddEdge(v1 nodes.Node, w float64, v2 nodes.Node)
	Node(nID nodes.NodeID) nodes.Node
	OutEdges(nID nodes.NodeID) []nodes.IEdge
}

// DAGraph is an unweighted directed graph represented as an adjacency list.
type DAGraph struct {
	A nodes.IAdjacency
}

// UAGraph is an unweighted undirected graph represented as an adjacency list.
type UAGraph struct {
	A nodes.IAdjacency
}

// WDAGraph is a weighted directed graph represented as an adjacency list.
type WDAGraph struct {
	A nodes.IAdjacency
}

// WUAGraph is a weighted undirected graph represented as an adjacency list.
type WUAGraph struct {
	A nodes.IAdjacency
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
func (graph *DAGraph) AddEdge(v1 nodes.Node, v2 nodes.Node) {
	graph.A.AddEdge(v1, v2)
}

// Node returns a node instance by it's ID.
func (graph *DAGraph) Node(nID nodes.NodeID) nodes.Node {
	return graph.A.Node(nID)
}

// OutEdges returns a list of outbound edges for a node defined by nID.
func (graph *DAGraph) OutEdges(nID nodes.NodeID) []nodes.IEdge {
	return graph.A.OutEdges(nID)
}

// AdjacentNodes returns a list of adjacent nodes for a node defined by nID.
func (graph *DAGraph) AdjacentNodes(nID nodes.NodeID) *nodes.Nodes {
	return graph.A.AdjacentNodes(nID)
}

// AddEdge creates an edge between v1 and v2 nodes.
func (graph *UAGraph) AddEdge(v1 nodes.Node, v2 nodes.Node) {
	graph.A.AddEdge(v1, v2)
	graph.A.AddEdge(v2, v1)
}

// Node returns a node instance by it's ID.
func (graph *UAGraph) Node(nID nodes.NodeID) nodes.Node {
	return graph.A.Node(nID)
}

// OutEdges returns a list of outbound edges for a node defined by nID.
func (graph *UAGraph) OutEdges(nID nodes.NodeID) []nodes.IEdge {
	return graph.A.OutEdges(nID)
}

// AdjacentNodes returns a list of adjacent nodes for a node defined by nID.
func (graph *UAGraph) AdjacentNodes(nID nodes.NodeID) *nodes.Nodes {
	return graph.A.AdjacentNodes(nID)
}

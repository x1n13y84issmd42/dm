package graph

import "github.com/x1n13y84issmd42/dm/graph/nodes"

// IGraph is an interface for unweighted graph.
type IGraph interface {
	nodes.NodeAccess
	AddEdge(v1 nodes.Node, v2 nodes.Node)
}

// IWGraph is an interface for weighted graph.
type IWGraph interface {
	nodes.NodeAccess
	AddEdge(v1 nodes.Node, w float64, v2 nodes.Node)
}

// DGraph is an unweighted directed graph.
type DGraph struct {
	A nodes.IAdjacency
}

// UGraph is an unweighted undirected graph.
type UGraph struct {
	A nodes.IAdjacency
}

// WDGraph is a weighted directed graph.
type WDGraph struct {
	A nodes.IAdjacency
}

// WUGraph is a weighted undirected graph.
type WUGraph struct {
	A nodes.IAdjacency
}

// NewDGraph creates a new DGraph instance.
func NewDGraph() *DGraph {
	return &DGraph{
		A: nodes.NewAdjacencyList(),
	}
}

// NewUGraph creates a new UGraph instance.
func NewUGraph() *UGraph {
	return &UGraph{
		A: nodes.NewAdjacencyList(),
	}
}

// AddEdge creates an edge between v1 and v2 nodes.
func (graph *DGraph) AddEdge(v1 nodes.Node, v2 nodes.Node) {
	graph.A.AddEdge(v1, v2)
}

// Node returns a node instance by it's ID.
func (graph *DGraph) Node(nID nodes.NodeID) nodes.Node {
	return graph.A.Node(nID)
}

// OutEdges returns a list of outbound edges for a node defined by nID.
func (graph *DGraph) OutEdges(nID nodes.NodeID) []nodes.IEdge {
	return graph.A.OutEdges(nID)
}

// AdjacentNodes returns a list of adjacent nodes for a node defined by nID.
func (graph *DGraph) AdjacentNodes(nID nodes.NodeID) *nodes.Nodes {
	return graph.A.AdjacentNodes(nID)
}

// AddEdge creates an edge between v1 and v2 nodes.
func (graph *UGraph) AddEdge(v1 nodes.Node, v2 nodes.Node) {
	graph.A.AddEdge(v1, v2)
	graph.A.AddEdge(v2, v1)
}

// Node returns a node instance by it's ID.
func (graph *UGraph) Node(nID nodes.NodeID) nodes.Node {
	return graph.A.Node(nID)
}

// OutEdges returns a list of outbound edges for a node defined by nID.
func (graph *UGraph) OutEdges(nID nodes.NodeID) []nodes.IEdge {
	return graph.A.OutEdges(nID)
}

// AdjacentNodes returns a list of adjacent nodes for a node defined by nID.
func (graph *UGraph) AdjacentNodes(nID nodes.NodeID) *nodes.Nodes {
	return graph.A.AdjacentNodes(nID)
}

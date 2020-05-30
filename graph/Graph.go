package graph

import (
	"github.com/x1n13y84issmd42/dm/graph/contract"
	"github.com/x1n13y84issmd42/dm/graph/storage"
)

// DGraph is an unweighted directed graph.
type DGraph struct {
	A contract.IAdjacency
}

// UGraph is an unweighted undirected graph.
type UGraph struct {
	A contract.IAdjacency
}

// WDGraph is a weighted directed graph.
type WDGraph struct {
	A contract.IAdjacency
}

// WUGraph is a weighted undirected graph.
type WUGraph struct {
	A contract.IAdjacency
}

// NewDGraph creates a new DGraph instance.
func NewDGraph() *DGraph {
	return &DGraph{
		A: storage.NewAdjacencyList(),
	}
}

// NewUGraph creates a new UGraph instance.
func NewUGraph() *UGraph {
	return &UGraph{
		A: storage.NewAdjacencyList(),
	}
}

// AddEdge creates an edge between v1 and v2 nodes.
func (graph *DGraph) AddEdge(v1 contract.Node, v2 contract.Node) {
	graph.A.AddEdge(v1, v2)
}

// Node returns a node instance by it's ID.
func (graph *DGraph) Node(nID contract.NodeID) contract.Node {
	return graph.A.Node(nID)
}

// OutEdges returns a list of outbound edges for a node defined by nID.
func (graph *DGraph) OutEdges(nID contract.NodeID) []contract.Edge {
	return graph.A.OutEdges(nID)
}

// AdjacentNodes returns a list of adjacent nodes for a node defined by nID.
func (graph *DGraph) AdjacentNodes(nID contract.NodeID) contract.Nodes {
	return graph.A.AdjacentNodes(nID)
}

// AddEdge creates an edge between v1 and v2 nodes.
func (graph *UGraph) AddEdge(v1 contract.Node, v2 contract.Node) {
	graph.A.AddEdge(v1, v2)
	graph.A.AddEdge(v2, v1)
}

// Node returns a node instance by it's ID.
func (graph *UGraph) Node(nID contract.NodeID) contract.Node {
	return graph.A.Node(nID)
}

// OutEdges returns a list of outbound edges for a node defined by nID.
func (graph *UGraph) OutEdges(nID contract.NodeID) []contract.Edge {
	return graph.A.OutEdges(nID)
}

// AdjacentNodes returns a list of adjacent nodes for a node defined by nID.
func (graph *UGraph) AdjacentNodes(nID contract.NodeID) contract.Nodes {
	return graph.A.AdjacentNodes(nID)
}

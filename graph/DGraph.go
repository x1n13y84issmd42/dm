package graph

import (
	"github.com/x1n13y84issmd42/dm/graph/contract"
	"github.com/x1n13y84issmd42/dm/graph/iterator"
	"github.com/x1n13y84issmd42/dm/graph/storage"
)

// DEdge is a pair of nodes making an undirected edge.
type DEdge struct {
	A contract.Node
	B contract.Node
}

// Reverse creates a new edge by swapping the receiver's nodes.
func (e DEdge) Reverse() contract.Edge {
	return DEdge{
		A: e.B,
		B: e.A,
	}
}

// DGraph is an unweighted directed graph.
type DGraph struct {
	A contract.IAdjacency
}

// NewDGraph creates a new DGraph instance.
// Provided nodes will be added pairwise as edges.
func NewDGraph(nodes ...contract.Node) *DGraph {
	g := &DGraph{
		A: storage.NewAdjacencyList(),
	}

	for i := 0; i < len(nodes); i += 2 {
		g.AddEdge(nodes[i], nodes[i+1])
	}

	return g
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
func (graph *DGraph) OutEdges(nID contract.NodeID) []DEdge {
	res := []DEdge{}
	n := graph.Node(nID)
	if n != nil {
		for na := range graph.AdjacentNodes(nID).Range() {
			res = append(res, DEdge{
				A: n,
				B: na,
			})
		}
	}

	return res
}

// AdjacentNodes returns a list of adjacent nodes for a node defined by nID.
func (graph *DGraph) AdjacentNodes(nID contract.NodeID) contract.Nodes {
	return graph.A.AdjacentNodes(nID)
}

// DFS returns a DFS node iterator.
func (graph *DGraph) DFS(nID contract.NodeID, traverse contract.Traversal) contract.NChannel {
	return iterator.DFS(graph, nID, traverse)
}

// BFS returns a BFS node iterator.
func (graph *DGraph) BFS(nID contract.NodeID) contract.NChannel {
	return iterator.BFS(graph, nID)
}

package storage

import (
	"github.com/x1n13y84issmd42/dm/graph/collection"
	"github.com/x1n13y84issmd42/dm/graph/contract"
)

// adjList is a mapping between node IDs and their instances.
type adjList map[contract.NodeID]contract.Nodes

// AdjacencyList is a list of nodes adjacent to other nodes.
type AdjacencyList struct {
	Nodes contract.Nodes
	List  adjList
}

// NewAdjacencyList creates a new AdjacencyList instance.
func NewAdjacencyList() *AdjacencyList {
	return &AdjacencyList{
		Nodes: collection.NewNodes(),
		List:  adjList{},
	}
}

// AddEdge adds a v2 to the adjacency list of v1.
func (list *AdjacencyList) AddEdge(v1 contract.Node, v2 contract.Node) {
	v1ID := v1.ID()
	v2ID := v2.ID()

	if list.List[v1ID] == nil {
		list.List[v1ID] = collection.NewNodes()
	}

	for n := range list.List[v1ID].Range() {
		if n.ID() == v2ID {
			return
		}
	}

	list.Nodes.Add(v1)
	list.Nodes.Add(v2)

	list.List[v1ID].Add(v2)
}

// AdjacentNodes returns a set of nodes adjacent to n.
func (list *AdjacencyList) AdjacentNodes(nID contract.NodeID) contract.Nodes {
	if list.List[nID] != nil {
		return list.List[nID]
	}

	return collection.NewNodes()
}

// Node returns a node instance by it's ID.
func (list *AdjacencyList) Node(nID contract.NodeID) contract.Node {
	if list.Nodes.Has(nID) {
		return list.Nodes.Get(nID)
	}

	return nil
}

// OutEdges returns a list of outbound edges for a node defined by nID.
func (list *AdjacencyList) OutEdges(nID contract.NodeID) []contract.Edge {
	res := []contract.Edge{}
	// if list.Nodes.Has(nID) {
	// 	n := list.Nodes.Get(nID)
	// 	for na := range n.Adjacent().Range() {
	// 		res = append(res, contract.Edge{
	// 			A: n,
	// 			B: na,
	// 		})
	// 	}
	// }

	return res
}

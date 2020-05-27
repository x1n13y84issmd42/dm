package nodes

// adjList is a mapping between node IDs and their instances.
type adjList map[NodeID]*Nodes

// AdjacencyList is a list of nodes adjacent to other nodes.
type AdjacencyList struct {
	Nodes Nodes
	List  adjList
}

func (list *AdjacencyList) init() {
	if list.List == nil {
		list.Nodes = Nodes{}
		list.List = adjList{}
	}
}

// AddEdge adds a v2 to the adjacency list of v1.
func (list *AdjacencyList) AddEdge(v1 INode, v2 INode) {
	list.init()
	if list.List[v1.ID()] == nil {
		list.List[v1.ID()] = &Nodes{}
	}

	for nID := range *list.List[v1.ID()] {
		if nID == v2.ID() {
			return
		}
	}

	list.Nodes.Add(v1)
	list.Nodes.Add(v2)

	list.List[v1.ID()].Add(v2)
}

// AdjacentNodes returns a set of nodes adjacent to n.
func (list *AdjacencyList) AdjacentNodes(n INode) Nodes {
	list.init()
	if list.List[n.ID()] != nil {
		return *list.List[n.ID()]
	}

	return Nodes{}
}

// GetNode returns a node instance by it's ID.
func (list *AdjacencyList) GetNode(nID NodeID) INode {
	list.init()
	if list.Nodes.Has(nID) {
		return list.Nodes[nID]
	}

	return nil
}

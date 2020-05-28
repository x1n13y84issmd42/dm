package nodes

// adjList is a mapping between node IDs and their instances.
type adjList map[NodeID]*Nodes

// AdjacencyList is a list of nodes adjacent to other nodes.
type AdjacencyList struct {
	Nodes *Nodes
	List  adjList
}

// NewAdjacencyList creates a new AdjacencyList instance.
func NewAdjacencyList() *AdjacencyList {
	return &AdjacencyList{
		Nodes: NewNodes(),
		List:  adjList{},
	}
}

// AddEdge adds a v2 to the adjacency list of v1.
func (list *AdjacencyList) AddEdge(v1 Node, v2 Node) {
	if list.List[v1.ID()] == nil {
		list.List[v1.ID()] = NewNodes()
	}

	for n := range list.List[v1.ID()].Range() {
		if n.ID() == v2.ID() {
			return
		}
	}

	list.Nodes.Add(v1)
	list.Nodes.Add(v2)

	list.List[v1.ID()].Add(v2)
}

// AdjacentNodes returns a set of nodes adjacent to n.
func (list *AdjacencyList) AdjacentNodes(nID NodeID) *Nodes {
	if list.List[nID] != nil {
		return list.List[nID]
	}

	return NewNodes()
}

// Node returns a node instance by it's ID.
func (list *AdjacencyList) Node(nID NodeID) Node {
	if list.Nodes.Has(nID) {
		return list.Nodes.Get(nID)
	}

	return nil
}

// OutEdges returns a list of outbound edges for a node defined by nID.
func (list *AdjacencyList) OutEdges(nID NodeID) []IEdge {
	res := []IEdge{}
	// if list.Nodes.Has(nID) {
	// 	n := list.Nodes.Get(nID)
	// 	for na := range n.Adjacent().Range() {
	// 		res = append(res, Edge{
	// 			A: n,
	// 			B: na,
	// 		})
	// 	}
	// }

	return res
}

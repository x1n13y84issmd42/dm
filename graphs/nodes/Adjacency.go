package nodes

// IAdjacency is an interface to access node's adjacency info.
type IAdjacency interface {
	AddEdge(v1 Node, v2 Node)
	Node(nID NodeID) Node
	AdjacentNodes(nID NodeID) *Nodes
	OutEdges(nID NodeID) []IEdge
}

// IWAdjacency the same as IAdjacency, but for weighted nodes.
type IWAdjacency interface {
	AddEdge(v1 Node, w float64, v2 Node)
	Node(nID NodeID) Node
	AdjacentNodes(nID NodeID) *Nodes
	OutEdges(nID NodeID) []IEdge
}

package nodes

// IAdjacency is a lower-level interface to access basic
type IAdjacency interface {
	AdjacentNodes(n INode) *Nodes
	GetNode(nID NodeID) INode
}

// AdjacencyWList is a weighted adjacency list. Used for weighted edges.
// type AdjacencyWList map[INode][]IWNode

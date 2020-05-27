package nodes

//NodeID is a value, unique within a graph, which determines a node's identity.
type NodeID string

// INode is an intefrace for basic directed graph nodes.
type INode interface {
	ID() NodeID
	Adjacent() *Nodes

	SetAdjacency(a IAdjacency)
}

// Node is a directed graph node.
type Node struct {
	INode
	A IAdjacency
}

// SetAdjacency sets the graph-specific adjacency implementation.
func (node *Node) SetAdjacency(a IAdjacency) {
	node.A = a
}

// Adjacent returns all the adjacent nodes.
func (node *Node) Adjacent() *Nodes {
	return node.A.AdjacentNodes(node)
}

package nodes

//NodeID is a value, unique within a graph, which determines a node's identity.
type NodeID string

// INode is an intefrace for basic directed graph nodes.
type INode interface {
	ID() NodeID
	Adjacent() Nodes

	SetAdjacency(a IAdjacency)

	// DEPRECATED
	Add(n INode) INode
}

// Node is a directed graph node.
type Node struct {
	INode
	A IAdjacency

	// DEPRECATED
	ChildNodes Nodes
}

// init is DEPRECATED
func (node *Node) init() {
	if node.ChildNodes == nil {
		node.ChildNodes = Nodes{}
	}
}

// Add adds a child node. And it's DEPRECATED.
func (node *Node) Add(n INode) INode {
	node.init()
	node.ChildNodes.Add(n)
	return n
}

// NumChildren returns number of child nodes. And it's DEPRECATED.
func (node *Node) NumChildren() int {
	node.init()
	return node.ChildNodes.Count()
}

// SetAdjacency sets the graph-specific adjacency implementation.
func (node *Node) SetAdjacency(a IAdjacency) {
	node.A = a
}

// Adjacent returns all the adjacent nodes.
func (node *Node) Adjacent() Nodes {
	node.init()
	var inode interface{}
	inode = node
	return node.A.AdjacentNodes(inode.(INode))
}

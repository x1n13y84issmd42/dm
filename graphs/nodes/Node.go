package nodes

//NodeID is a value, unique within a graph, which determines a node's identity.
type NodeID string

// INode is an intefrace for basic directed graph nodes.
type INode interface {
	ID() NodeID
	Adjacent() Nodes

	Add(n INode) INode
}

// Node is a directed graph node.
type Node struct {
	ChildNodes Nodes
}

func (node *Node) init() {
	if node.ChildNodes == nil {
		node.ChildNodes = Nodes{}
	}
}

// Add adds a child node.
func (node *Node) Add(n INode) INode {
	node.init()
	node.ChildNodes.Add(n)
	return n
}

// NumChildren returns number of child nodes.
func (node *Node) NumChildren() int {
	node.init()
	return node.ChildNodes.Count()
}

// Adjacent returns all the child nodes.
func (node *Node) Adjacent() Nodes {
	node.init()
	return node.ChildNodes
}

// GhostNode is a Flyweight node used with various representations of graphs.
type GhostNode struct {
}

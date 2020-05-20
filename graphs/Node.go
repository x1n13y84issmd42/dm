package graphs

// INode is an intefrace for basic graph nodes.
type INode interface {
	Add(n INode) INode
	Child(i int) INode
	NumChildren() int
	Children() []INode
}

// Node is a graph node.
type Node struct {
	ChildNodes []INode
}

// Add adds a child node.
func (node *Node) Add(n INode) INode {
	node.ChildNodes = append(node.ChildNodes, n)
	return n
}

// Child returns an i-th child node, if present, nil otherwise.
func (node *Node) Child(i int) INode {
	if i < len(node.ChildNodes) {
		return node.ChildNodes[i]
	}

	return nil
}

// NumChildren returns number of child nodes.
func (node *Node) NumChildren() int {
	return len(node.ChildNodes)
}

// Children returns all the child nodes.
func (node *Node) Children() []INode {
	return node.ChildNodes
}

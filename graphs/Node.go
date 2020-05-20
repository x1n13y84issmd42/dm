package graphs

// INode is an intefrace for basic graph nodes.
type INode interface {
	Add(n INode) INode
	Child(i int) INode
	NumChildren() int
}

// Node is a graph node.
type Node struct {
	Children []INode
}

// Add adds a child node.
func (node *Node) Add(n INode) INode {
	node.Children = append(node.Children, n)
	return n
}

// Child returns an i-th child node, if present, nil otherwise.
func (node *Node) Child(i int) INode {
	if i < len(node.Children) {
		return node.Children[i]
	}

	return nil
}

// NumChildren adds a child node.
func (node *Node) NumChildren() int {
	return len(node.Children)
}

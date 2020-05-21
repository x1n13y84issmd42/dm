package graphs

// IDNode is an intefrace for basic directed graph nodes.
type IDNode interface {
	Add(n IDNode) IDNode
	Child(i int) IDNode
	NumChildren() int
	Children() []IDNode
}

// DNode is a directed graph node.
type DNode struct {
	ChildNodes []IDNode
}

// Add adds a child node.
func (node *DNode) Add(n IDNode) IDNode {
	node.ChildNodes = append(node.ChildNodes, n)
	return n
}

// Child returns an i-th child node, if present, nil otherwise.
func (node *DNode) Child(i int) IDNode {
	if i < len(node.ChildNodes) {
		return node.ChildNodes[i]
	}

	return nil
}

// NumChildren returns number of child nodes.
func (node *DNode) NumChildren() int {
	return len(node.ChildNodes)
}

// Children returns all the child nodes.
func (node *DNode) Children() []IDNode {
	return node.ChildNodes
}

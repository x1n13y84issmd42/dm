package ut

import "github.com/x1n13y84issmd42/dm/graphs/nodes"

// TestNode ...
type TestNode struct {
	Name string
}

// ID ...
func (node *TestNode) ID() nodes.NodeID {
	return nodes.NodeID(node.Name)
}

// Node ...
func Node(name string) *TestNode {
	node := TestNode{
		Name: name,
	}

	return &node
}

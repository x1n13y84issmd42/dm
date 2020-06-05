package ut

import "github.com/x1n13y84issmd42/gog/graph/contract"

// TestNode ...
type TestNode struct {
	Name string
}

// ID ...
func (node *TestNode) ID() contract.NodeID {
	return contract.NodeID(node.Name)
}

// Node ...
func Node(name string) *TestNode {
	node := TestNode{
		Name: name,
	}

	return &node
}

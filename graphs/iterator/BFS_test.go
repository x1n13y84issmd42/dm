package iterator

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/x1n13y84issmd42/dm/graphs/nodes"
)

type TestNode struct {
	nodes.Node
	Name string
}

func (node *TestNode) ID() nodes.NodeID {
	return nodes.NodeID(node.Name)
}

func Test_BFS_1(T *testing.T) {
	root := TestNode{
		Name: "L0",
	}

	root.Add(&TestNode{
		Name: "L10",
	}).Add(&TestNode{
		Name: "L20",
	}).Add(&TestNode{
		Name: "L30",
	})

	root.Add(&TestNode{
		Name: "L11",
	}).Add(&TestNode{
		Name: "L21",
	})

	root.Add(&TestNode{
		Name: "L12",
	})

	expected := "L0L10L11L12L20L21L30"
	actual := ""

	for node := range BFS(&root) {
		actual = fmt.Sprintf("%s%s", actual, node.(*TestNode).Name)
	}

	assert.Equal(T, expected, actual)
}

func Test_BFS_Loop(T *testing.T) {
	root := TestNode{
		Name: "L0",
	}

	root.Add(&TestNode{
		Name: "L10",
	}).Add(&root)

	expected := "L0L10"
	actual := ""

	for node := range BFS(&root) {
		actual = fmt.Sprintf("%s%s", actual, node.(*TestNode).Name)
	}

	assert.Equal(T, expected, actual)
}
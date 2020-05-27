package graphs

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/x1n13y84issmd42/dm/graphs/iterator"
	"github.com/x1n13y84issmd42/dm/graphs/nodes"
	"github.com/x1n13y84issmd42/dm/graphs/ut"
)

type TestNode struct {
	nodes.Node
	Name string
}

func (node *TestNode) ID() nodes.NodeID {
	return nodes.NodeID(node.Name)
}

func Test_BFS_1(T *testing.T) {
	dg := DAGraph{}

	root := ut.Node("L0")
	L10 := ut.Node("L10")
	L20 := ut.Node("L20")
	L30 := ut.Node("L30")
	L11 := ut.Node("L11")
	L21 := ut.Node("L21")

	dg.AddEdge(root, L10)
	dg.AddEdge(L10, L20)
	dg.AddEdge(L20, L30)

	dg.AddEdge(root, L11)
	dg.AddEdge(L11, L21)

	dg.AddEdge(root, ut.Node("L12"))

	expected := "L0L10L11L12L20L21L30"
	actual := ""

	//TODO: Sometimes it produces invalid order and fails.
	for node := range iterator.BFS(root) {
		actual = fmt.Sprintf("%s%s", actual, node.(*ut.TestNode).Name)
	}

	assert.Equal(T, expected, actual)
}

func Test_BFS_Loop(T *testing.T) {
	root := ut.Node("L0")
	L10 := ut.Node("L10")

	g := DAGraph{}
	g.AddEdge(root, L10)
	g.AddEdge(L10, root)

	expected := "L0L10"
	actual := ""

	for node := range iterator.BFS(root) {
		actual = fmt.Sprintf("%s%s", actual, node.(*ut.TestNode).Name)
	}

	assert.Equal(T, expected, actual)
}

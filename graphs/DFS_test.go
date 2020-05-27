package graphs

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/x1n13y84issmd42/dm/graphs/iterator"
	"github.com/x1n13y84issmd42/dm/graphs/ut"
)

func Test_DFS_1(T *testing.T) {
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

	expected := "L0L10L20L30L11L21L12"
	actual := ""

	//TODO: Sometimes it produces invalid order and fails.
	for node := range iterator.DFS(root) {
		actual = fmt.Sprintf("%s%s", actual, node.(*ut.TestNode).Name)
	}

	assert.Equal(T, expected, actual)
}

func Test_DFS_Loop(T *testing.T) {
	root := ut.Node("L0")
	L10 := ut.Node("L10")

	g := DAGraph{}
	g.AddEdge(root, L10)
	g.AddEdge(L10, root)

	expected := "L0L10"
	actual := ""

	for node := range iterator.DFS(root) {
		actual = fmt.Sprintf("%s%s", actual, node.(*ut.TestNode).Name)
	}

	assert.Equal(T, expected, actual)
}

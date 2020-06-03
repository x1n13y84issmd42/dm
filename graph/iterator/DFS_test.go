package iterator_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/x1n13y84issmd42/dm/graph"
	"github.com/x1n13y84issmd42/dm/graph/iterator"
	"github.com/x1n13y84issmd42/dm/graph/ut"
)

func Test_DFS_PreOrder(T *testing.T) {
	g := graph.NewDGraph()

	root := ut.Node("L0")
	L10 := ut.Node("L10")
	L20 := ut.Node("L20")
	L30 := ut.Node("L30")
	L11 := ut.Node("L11")
	L21 := ut.Node("L21")

	g.AddEdge(root, L10)
	g.AddEdge(L10, L20)
	g.AddEdge(L20, L30)

	g.AddEdge(root, L11)
	g.AddEdge(L11, L21)

	g.AddEdge(root, ut.Node("L12"))

	expected := "L0L10L20L30L11L21L12"
	actual := ""

	for node := range g.DFS("L0", iterator.PreOrder) {
		actual = fmt.Sprintf("%s%s", actual, node.(*ut.TestNode).Name)
	}

	assert.Equal(T, expected, actual)
}

func Test_DFS_PostOrder(T *testing.T) {
	g := graph.NewDGraph()

	root := ut.Node("L0")
	L10 := ut.Node("L10")
	L20 := ut.Node("L20")
	L30 := ut.Node("L30")
	L11 := ut.Node("L11")
	L21 := ut.Node("L21")

	g.AddEdge(root, L10)
	g.AddEdge(L10, L20)
	g.AddEdge(L20, L30)

	g.AddEdge(root, L11)
	g.AddEdge(L11, L21)

	g.AddEdge(root, ut.Node("L12"))

	expected := "L30L20L10L21L11L12L0"
	actual := ""

	for node := range g.DFS("L0", iterator.PostOrder) {
		actual = fmt.Sprintf("%s%s", actual, node.(*ut.TestNode).Name)
	}

	assert.Equal(T, expected, actual)
}

func Test_DFS_Loop(T *testing.T) {
	root := ut.Node("L0")
	L10 := ut.Node("L10")

	g := graph.NewDGraph()
	g.AddEdge(root, L10)
	g.AddEdge(L10, root)

	expected := "L0L10"
	actual := ""

	for node := range g.DFS("L0", iterator.PreOrder) {
		actual = fmt.Sprintf("%s%s", actual, node.(*ut.TestNode).Name)
	}

	assert.Equal(T, expected, actual)
}

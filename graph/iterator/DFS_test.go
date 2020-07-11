package iterator_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/x1n13y84issmd42/gog/graph"
	"github.com/x1n13y84issmd42/gog/graph/iterator"
	"github.com/x1n13y84issmd42/gog/graph/ut"
)

func Test_DFS_PreOrder(T *testing.T) {
	g := graph.NewDGraph(
		ut.Node("L0"),
		ut.Node("L10"),
		ut.Node("L20"),
		ut.Node("L30"),
		ut.Node("L11"),
		ut.Node("L21"),
		ut.Node("L12"),
	)

	g.AddEdge("L0", "L10")
	g.AddEdge("L10", "L20")
	g.AddEdge("L20", "L30")

	g.AddEdge("L0", "L11")
	g.AddEdge("L11", "L21")

	g.AddEdge("L0", "L12")

	expected := "L0L10L20L30L11L21L12"
	actual := ""

	for node := range g.DFS("L0", iterator.PreOrder) {
		actual = fmt.Sprintf("%s%s", actual, node.(*ut.TestNode).Name)
	}

	assert.Equal(T, expected, actual)
}

func Test_DFS_PostOrder(T *testing.T) {
	g := graph.NewDGraph(
		ut.Node("L0"),
		ut.Node("L10"),
		ut.Node("L20"),
		ut.Node("L30"),
		ut.Node("L11"),
		ut.Node("L21"),
		ut.Node("L12"),
	)

	g.AddEdge("L0", "L10")
	g.AddEdge("L10", "L20")
	g.AddEdge("L20", "L30")

	g.AddEdge("L0", "L11")
	g.AddEdge("L11", "L21")

	g.AddEdge("L0", "L12")

	expected := "L30L20L10L21L11L12L0"
	actual := ""

	for node := range g.DFS("L0", iterator.PostOrder) {
		actual = fmt.Sprintf("%s%s", actual, node.(*ut.TestNode).Name)
	}

	assert.Equal(T, expected, actual)
}

func Test_DFS_Loop(T *testing.T) {
	g := graph.NewDGraph(
		ut.Node("L0"),
		ut.Node("L10"),
	)

	g.AddEdge("L0", "L10")
	g.AddEdge("L10", "L0")

	expected := "L0L10"
	actual := ""

	for node := range g.DFS("L0", iterator.PreOrder) {
		actual = fmt.Sprintf("%s%s", actual, node.(*ut.TestNode).Name)
	}

	assert.Equal(T, expected, actual)
}

func Test_RDFS(T *testing.T) {
	g := graph.NewDGraph(
		ut.Node("L0"),
		ut.Node("L-10"),
		ut.Node("L-11"),
		ut.Node("L-21"),
		ut.Node("L-20"),
		ut.Node("L-12"),
		ut.Node("L10"),
		ut.Node("L11"),
	)

	g.AddEdge("L-21", "L-10")
	g.AddEdge("L-20", "L-10")
	g.AddEdge("L-10", "L0")
	g.AddEdge("L-11", "L0")
	g.AddEdge("L-12", "L0")
	g.AddEdge("L0", "L10")
	g.AddEdge("L0", "L11")

	expected := "L0L-10L-20L-21L-11L-12"
	actual := ""

	for node := range g.RDFS("L0", iterator.PreOrder) {
		actual = fmt.Sprintf("%s%s", actual, node.(*ut.TestNode).Name)
	}

	assert.Equal(T, expected, actual)
}

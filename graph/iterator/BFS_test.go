package iterator_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/x1n13y84issmd42/gog/graph"
	"github.com/x1n13y84issmd42/gog/graph/ut"
)

func Test_BFS_1(T *testing.T) {
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

	expected := "L0L10L11L12L20L21L30"
	actual := ""

	for node := range g.BFS("L0") {
		actual = fmt.Sprintf("%s%s", actual, node.(*ut.TestNode).Name)
	}

	assert.Equal(T, expected, actual)
}

func Test_BFS_Loop(T *testing.T) {
	g := graph.NewDGraph(ut.Node("L0"), ut.Node("L10"))
	g.AddEdge("L0", "L10")
	g.AddEdge("L10", "L0")

	expected := "L0L10"
	actual := ""

	for node := range g.BFS("L0") {
		actual = fmt.Sprintf("%s%s", actual, node.(*ut.TestNode).Name)
	}

	assert.Equal(T, expected, actual)
}

func Test_RBFS(T *testing.T) {
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

	expected := "L0L-10L-11L-12L-20L-21"
	actual := ""

	for node := range g.RBFS("L0") {
		actual = fmt.Sprintf("%s%s", actual, node.(*ut.TestNode).Name)
	}

	assert.Equal(T, expected, actual)
}

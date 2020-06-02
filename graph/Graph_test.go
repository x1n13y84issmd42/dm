package graph_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/x1n13y84issmd42/dm/graph"
	"github.com/x1n13y84issmd42/dm/graph/collection"
	"github.com/x1n13y84issmd42/dm/graph/contract"
	"github.com/x1n13y84issmd42/dm/graph/iterator"
	"github.com/x1n13y84issmd42/dm/graph/ut"
)

func CreateTestGraph(graph contract.Graph) {
	root := ut.Node("R007")

	graph.AddEdge(root, ut.Node("A"))
	graph.AddEdge(root, ut.Node("B"))
	graph.AddEdge(ut.Node("C"), root)
	graph.AddEdge(ut.Node("D"), root)
}

func Test_DGraph(T *testing.T) {
	g := graph.NewDGraph()

	CreateTestGraph(g)

	expected := collection.NewNodes()
	expected.Add(g.Node("A"))
	expected.Add(g.Node("B"))

	actual := g.AdjacentNodes("R007")

	assert.Equal(T, expected, actual)
}

// Iterators are tested in the iterator package.
// This exists just to have coverage on graphs' methods.
func Test_Cover_DGraph_Iterators(T *testing.T) {
	dg := graph.NewDGraph(ut.Node("A"), ut.Node("B"))
	dg.DFS("A", iterator.PreOrder)
	dg.BFS("A")

	ug := graph.NewUGraph(ut.Node("A"), ut.Node("B"))
	ug.DFS("A", iterator.PreOrder)
	ug.BFS("A")
}

func Test_DEdge_Reverse(T *testing.T) {
	e := graph.DEdge{
		A: ut.Node("A"),
		B: ut.Node("B"),
	}

	expected := graph.DEdge{
		A: e.B,
		B: e.A,
	}

	assert.Equal(T, expected, e.Reverse())
}

func Test_UEdge_Reverse(T *testing.T) {
	e := graph.UEdge{
		A: ut.Node("A"),
		B: ut.Node("B"),
	}

	expected := graph.UEdge{
		A: e.A,
		B: e.B,
	}

	assert.Equal(T, expected, e.Reverse())
}

func Test_UGraph(T *testing.T) {
	g := graph.NewUGraph()

	CreateTestGraph(g)

	expected := collection.NewNodes()
	expected.Add(g.Node("A"))
	expected.Add(g.Node("B"))
	expected.Add(g.Node("C"))
	expected.Add(g.Node("D"))

	actual := g.AdjacentNodes("R007")

	assert.Equal(T, expected, actual)
}

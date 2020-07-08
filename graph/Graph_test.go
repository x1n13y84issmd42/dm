package graph_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/x1n13y84issmd42/gog/graph"
	"github.com/x1n13y84issmd42/gog/graph/collection"
	"github.com/x1n13y84issmd42/gog/graph/contract"
	"github.com/x1n13y84issmd42/gog/graph/iterator"
	"github.com/x1n13y84issmd42/gog/graph/ut"
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

	T.Run("AdjacentNodes", func(T *testing.T) {
		assert.Equal(T,
			collection.NewNodes(g.Node("A"), g.Node("B")),
			g.AdjacentNodes("R007"),
		)
	})

	T.Run("Nodes", func(T *testing.T) {
		assert.Equal(T,
			collection.NewNodes(
				g.Node("R007"),
				g.Node("A"),
				g.Node("B"),
				g.Node("C"),
				g.Node("D"),
			),
			g.Nodes(),
		)
	})

	T.Run("UpstreamNodes", func(T *testing.T) {
		assert.Equal(T,
			collection.NewNodes(g.Node("C"), g.Node("D")),
			g.UpstreamNodes("R007"),
		)
	})

	T.Run("OutEdges", func(T *testing.T) {
		expected := []graph.DEdge{
			{
				A: g.Node("R007"),
				B: g.Node("A"),
			},
			{
				A: g.Node("R007"),
				B: g.Node("B"),
			},
		}

		assert.Equal(T,
			expected,
			g.OutEdges("R007"),
		)
	})

	T.Run("InEdges", func(T *testing.T) {
		expected := []graph.DEdge{
			{
				A: g.Node("C"),
				B: g.Node("R007"),
			},
			{
				A: g.Node("D"),
				B: g.Node("R007"),
			},
		}

		assert.Equal(T,
			expected,
			g.InEdges("R007"),
		)
	})
}

func Test_UGraph(T *testing.T) {
	g := graph.NewUGraph()
	CreateTestGraph(g)

	T.Run("AdjacentNodes", func(T *testing.T) {
		expected := collection.NewNodes(g.Node("A"), g.Node("B"), g.Node("C"), g.Node("D"))
		actual := g.AdjacentNodes("R007")
		assert.Equal(T, expected, actual)
	})

	T.Run("Nodes", func(T *testing.T) {
		assert.Equal(T,
			collection.NewNodes(
				g.Node("R007"),
				g.Node("A"),
				g.Node("B"),
				g.Node("C"),
				g.Node("D"),
			),
			g.Nodes(),
		)
	})

	T.Run("UpstreamNodes", func(T *testing.T) {
		expected := collection.NewNodes(g.Node("A"), g.Node("B"), g.Node("C"), g.Node("D"))
		actual := g.UpstreamNodes("R007")
		assert.Equal(T, expected, actual)
	})

	T.Run("OutEdges", func(T *testing.T) {
		expected := []graph.UEdge{
			{
				A: g.Node("R007"),
				B: g.Node("A"),
			},
			{
				A: g.Node("R007"),
				B: g.Node("B"),
			},
			{
				A: g.Node("R007"),
				B: g.Node("C"),
			},
			{
				A: g.Node("R007"),
				B: g.Node("D"),
			},
		}

		assert.Equal(T,
			expected,
			g.OutEdges("R007"),
		)
	})

	T.Run("InEdges", func(T *testing.T) {
		expected := []graph.UEdge{
			{
				A: g.Node("A"),
				B: g.Node("R007"),
			},
			{
				A: g.Node("B"),
				B: g.Node("R007"),
			},
			{
				A: g.Node("C"),
				B: g.Node("R007"),
			},
			{
				A: g.Node("D"),
				B: g.Node("R007"),
			},
		}

		assert.Equal(T,
			expected,
			g.InEdges("R007"),
		)
	})
}

// Iterators are tested in the iterator package.
// This exists just to have coverage on graphs' methods.
func Test_Cover_Iterators(T *testing.T) {
	dg := graph.NewDGraph(ut.Node("A"), ut.Node("B"))
	dg.DFS("A", iterator.PreOrder)
	dg.BFS("A")
	dg.RDFS("A", iterator.PreOrder)
	dg.RBFS("A")

	ug := graph.NewUGraph(ut.Node("A"), ut.Node("B"))
	ug.DFS("A", iterator.PreOrder)
	ug.BFS("A")
	ug.RDFS("A", iterator.PreOrder)
	ug.RBFS("A")
}

func Test_DEdge(T *testing.T) {
	T.Run("Reverse", func(T *testing.T) {
		e := graph.DEdge{
			A: ut.Node("A"),
			B: ut.Node("B"),
		}

		expected := graph.DEdge{
			A: e.B,
			B: e.A,
		}

		assert.Equal(T, expected, e.Reverse())
	})
}

func Test_UEdge(T *testing.T) {
	T.Run("Reverse", func(T *testing.T) {
		e := graph.UEdge{
			A: ut.Node("A"),
			B: ut.Node("B"),
		}

		expected := graph.UEdge{
			A: e.A,
			B: e.B,
		}

		assert.Equal(T, expected, e.Reverse())
	})
}

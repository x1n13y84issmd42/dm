package comp_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/x1n13y84issmd42/gog/graph"
	"github.com/x1n13y84issmd42/gog/graph/collection"
	"github.com/x1n13y84issmd42/gog/graph/comp"
	"github.com/x1n13y84issmd42/gog/graph/ut"
)

func Test_Cycle(T *testing.T) {

	T.Run("Cycle1", func(T *testing.T) {
		g := graph.NewDGraph(
			ut.Node("1"),
			ut.Node("3"),
			ut.Node("2"),
			ut.Node("4"),
			ut.Node("5"),
			ut.Node("7"),
			ut.Node("6"),
			ut.Node("0"),
		)

		g.AddEdge("1", "3")
		g.AddEdge("3", "2")
		g.AddEdge("2", "1")

		g.AddEdge("4", "5")
		g.AddEdge("5", "7")
		g.AddEdge("7", "6")
		g.AddEdge("6", "4")

		g.AddEdge("0", "1")
		g.AddEdge("3", "4")

		expected := &collection.NodeStack{
			g.Node("1"),
			g.Node("3"),
			g.Node("2"),
			g.Node("1"),
		}

		actual := comp.Cycle(g)

		assert.Equal(T, expected, actual)
	})

	T.Run("Cycle2", func(T *testing.T) {
		g := graph.NewDGraph(
			ut.Node("0"),
			ut.Node("1"),
			ut.Node("2"),
			ut.Node("3"),
			ut.Node("4"),
			ut.Node("5"),
			ut.Node("6"),
			ut.Node("7"),
		)

		g.AddEdge("1", "3")
		g.AddEdge("3", "2")

		g.AddEdge("4", "5")
		g.AddEdge("5", "7")
		g.AddEdge("7", "6")
		g.AddEdge("6", "4")

		g.AddEdge("0", "1")
		g.AddEdge("3", "4")

		expected := &collection.NodeStack{
			g.Node("4"),
			g.Node("5"),
			g.Node("7"),
			g.Node("6"),
			g.Node("4"),
		}

		actual := comp.Cycle(g)

		assert.Equal(T, expected, actual)
	})

	T.Run("DAG", func(T *testing.T) {
		g := graph.NewDGraph(
			ut.Node("1"),
			ut.Node("3"),
			ut.Node("2"),
			ut.Node("4"),
			ut.Node("5"),
			ut.Node("7"),
			ut.Node("6"),
			ut.Node("0"),
		)

		g.AddEdge("1", "3")
		g.AddEdge("3", "2")

		g.AddEdge("4", "5")
		g.AddEdge("5", "7")
		g.AddEdge("7", "6")

		g.AddEdge("0", "1")
		g.AddEdge("3", "4")

		expected := &collection.NodeStack{}

		actual := comp.Cycle(g)

		assert.Equal(T, expected, actual)
	})
}

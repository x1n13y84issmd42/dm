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
		g := graph.NewDGraph()

		g.AddEdge(ut.Node("1"), ut.Node("3"))
		g.AddEdge(g.Node("3"), ut.Node("2"))
		g.AddEdge(g.Node("2"), g.Node("1"))

		g.AddEdge(ut.Node("4"), ut.Node("5"))
		g.AddEdge(g.Node("5"), ut.Node("7"))
		g.AddEdge(g.Node("7"), ut.Node("6"))
		g.AddEdge(g.Node("6"), g.Node("4"))

		g.AddEdge(ut.Node("0"), g.Node("1"))
		g.AddEdge(g.Node("3"), g.Node("4"))

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
		g := graph.NewDGraph()

		g.AddEdge(ut.Node("1"), ut.Node("3"))
		g.AddEdge(g.Node("3"), ut.Node("2"))
		// g.AddEdge(g.Node("2"), g.Node("1"))

		g.AddEdge(ut.Node("4"), ut.Node("5"))
		g.AddEdge(g.Node("5"), ut.Node("7"))
		g.AddEdge(g.Node("7"), ut.Node("6"))
		g.AddEdge(g.Node("6"), g.Node("4"))

		g.AddEdge(ut.Node("0"), g.Node("1"))
		g.AddEdge(g.Node("3"), g.Node("4"))

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
		g := graph.NewDGraph()

		g.AddEdge(ut.Node("1"), ut.Node("3"))
		g.AddEdge(g.Node("3"), ut.Node("2"))
		// g.AddEdge(g.Node("2"), g.Node("1"))

		g.AddEdge(ut.Node("4"), ut.Node("5"))
		g.AddEdge(g.Node("5"), ut.Node("7"))
		g.AddEdge(g.Node("7"), ut.Node("6"))
		// g.AddEdge(g.Node("6"), g.Node("4"))

		g.AddEdge(ut.Node("0"), g.Node("1"))
		g.AddEdge(g.Node("3"), g.Node("4"))

		expected := &collection.NodeStack{}

		actual := comp.Cycle(g)

		assert.Equal(T, expected, actual)
	})

}

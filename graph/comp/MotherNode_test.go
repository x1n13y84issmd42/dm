package comp_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/x1n13y84issmd42/gog/graph"
	"github.com/x1n13y84issmd42/gog/graph/comp"
	"github.com/x1n13y84issmd42/gog/graph/ut"
)

func Test_MotherNode(T *testing.T) {
	T.Run("Exists", func(T *testing.T) {
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

		expected := g.Node("0")
		actual := comp.MotherNode(g)

		assert.Equal(T, expected, actual)

	})

	T.Run("Doesn't exist", func(T *testing.T) {
		g := graph.NewDGraph()

		g.AddEdge(ut.Node("1"), ut.Node("3"))
		g.AddEdge(ut.Node("2"), g.Node("3"))

		assert.Nil(T, comp.MotherNode(g))

	})
}

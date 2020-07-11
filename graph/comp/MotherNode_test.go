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

		expected := g.Node("0")
		actual := comp.MotherNode(g)

		assert.Equal(T, expected, actual)

	})

	T.Run("Doesn't exist", func(T *testing.T) {
		g := graph.NewDGraph(
			ut.Node("1"),
			ut.Node("3"),
			ut.Node("2"),
		)

		g.AddEdge("1", "3")
		g.AddEdge("2", "3")

		assert.Nil(T, comp.MotherNode(g))

	})
}

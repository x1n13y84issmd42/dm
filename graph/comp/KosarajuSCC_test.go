package comp_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/x1n13y84issmd42/gog/graph"
	"github.com/x1n13y84issmd42/gog/graph/collection"
	"github.com/x1n13y84issmd42/gog/graph/comp"
	"github.com/x1n13y84issmd42/gog/graph/ut"
)

func Test_KosarajuSCC(T *testing.T) {
	g := graph.NewDGraph(
		ut.Node("A"),
		ut.Node("B"),
		ut.Node("C"),
		ut.Node("D"),
		ut.Node("E"),
		ut.Node("F"),
		ut.Node("G"),
		ut.Node("H"),
		ut.Node("I"),
		ut.Node("J"),
		ut.Node("K"),
	)

	g.AddEdge("A", "B")
	g.AddEdge("B", "C")
	g.AddEdge("C", "A")

	g.AddEdge("D", "E")
	g.AddEdge("E", "F")
	g.AddEdge("F", "D")

	g.AddEdge("G", "H")
	g.AddEdge("H", "I")
	g.AddEdge("I", "J")
	g.AddEdge("J", "G")

	g.AddEdge("J", "K")

	g.AddEdge("B", "D")
	g.AddEdge("G", "F")

	expected := []*collection.Nodes{
		collection.NewNodes(g.Node("G"), g.Node("H"), g.Node("I"), g.Node("J")),
		collection.NewNodes(g.Node("K")),
		collection.NewNodes(g.Node("A"), g.Node("B"), g.Node("C")),
		collection.NewNodes(g.Node("D"), g.Node("E"), g.Node("F")),
	}

	actual := comp.KosarajuSCC(g)

	assert.Equal(T, expected, actual)
}

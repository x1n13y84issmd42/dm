package comp_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/x1n13y84issmd42/dm/graph"
	"github.com/x1n13y84issmd42/dm/graph/collection"
	"github.com/x1n13y84issmd42/dm/graph/comp"
	"github.com/x1n13y84issmd42/dm/graph/ut"
)

func Test_KosarajuSCC(T *testing.T) {
	g := graph.NewDGraph()

	g.AddEdge(ut.Node("A"), ut.Node("B"))
	g.AddEdge(g.Node("B"), ut.Node("C"))
	g.AddEdge(g.Node("C"), g.Node("A"))

	g.AddEdge(ut.Node("D"), ut.Node("E"))
	g.AddEdge(g.Node("E"), ut.Node("F"))
	g.AddEdge(g.Node("F"), g.Node("D"))

	g.AddEdge(ut.Node("G"), ut.Node("H"))
	g.AddEdge(g.Node("H"), ut.Node("I"))
	g.AddEdge(g.Node("I"), ut.Node("J"))
	g.AddEdge(g.Node("J"), g.Node("G"))

	g.AddEdge(g.Node("J"), ut.Node("K"))

	g.AddEdge(g.Node("B"), g.Node("D"))
	g.AddEdge(g.Node("G"), g.Node("F"))

	expected := []*collection.Nodes{
		collection.NewNodes(g.Node("G"), g.Node("H"), g.Node("I"), g.Node("J")),
		collection.NewNodes(g.Node("K")),
		collection.NewNodes(g.Node("A"), g.Node("B"), g.Node("C")),
		collection.NewNodes(g.Node("D"), g.Node("E"), g.Node("F")),
	}

	actual := comp.KosarajuSCC(g)

	assert.Equal(T, expected, actual)
}

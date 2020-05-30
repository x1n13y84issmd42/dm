package graph

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/x1n13y84issmd42/dm/graph/collection"
	"github.com/x1n13y84issmd42/dm/graph/contract"
	"github.com/x1n13y84issmd42/dm/graph/ut"
)

func CreateTestGraph(graph contract.Graph) {
	root := ut.Node("R007")

	graph.AddEdge(root, ut.Node("A"))
	graph.AddEdge(root, ut.Node("B"))
	graph.AddEdge(ut.Node("C"), root)
	graph.AddEdge(ut.Node("D"), root)
}

func Test_DAGraph(T *testing.T) {
	dg := NewDGraph()

	CreateTestGraph(dg)

	expected := collection.NewNodes()
	expected.Add(dg.Node("A"))
	expected.Add(dg.Node("B"))

	actual := dg.AdjacentNodes("R007")

	assert.Equal(T, expected, actual)
}

func Test_UAGraph(T *testing.T) {
	dg := NewUGraph()

	CreateTestGraph(dg)

	expected := collection.NewNodes()
	expected.Add(dg.Node("A"))
	expected.Add(dg.Node("B"))
	expected.Add(dg.Node("C"))
	expected.Add(dg.Node("D"))

	actual := dg.AdjacentNodes("R007")

	assert.Equal(T, expected, actual)
}

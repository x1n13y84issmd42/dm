package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/x1n13y84issmd42/dm/graphs/nodes"
	"github.com/x1n13y84issmd42/dm/graphs/ut"
)

func CreateTestGraph(graph IGraph) {
	root := ut.Node("R007")

	graph.AddEdge(root, ut.Node("A"))
	graph.AddEdge(root, ut.Node("B"))
	graph.AddEdge(ut.Node("C"), root)
	graph.AddEdge(ut.Node("D"), root)
}

func Test_DAGraph(T *testing.T) {
	dg := DAGraph{}

	CreateTestGraph(&dg)

	expected := nodes.Nodes{}
	expected.Add(dg.Node("A"))
	expected.Add(dg.Node("B"))

	root := dg.Node("R007")
	actual := root.Adjacent()

	assert.Equal(T, expected, actual)
}

func Test_UAGraph(T *testing.T) {
	dg := UAGraph{}

	CreateTestGraph(&dg)

	expected := nodes.Nodes{}
	expected.Add(dg.Node("A"))
	expected.Add(dg.Node("B"))
	expected.Add(dg.Node("C"))
	expected.Add(dg.Node("D"))

	root := dg.Node("R007")
	actual := root.Adjacent()

	assert.Equal(T, expected, actual)
}

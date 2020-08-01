package path_test

import (
	"fmt"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/x1n13y84issmd42/gog/graph"
	"github.com/x1n13y84issmd42/gog/graph/contract"
	"github.com/x1n13y84issmd42/gog/graph/path"
	"github.com/x1n13y84issmd42/gog/graph/ut"
)

func Test_DijkstraDistances(T *testing.T) {
	g := graph.NewDWGraph()

	g.AddNode(ut.Node("0"))
	g.AddNode(ut.Node("1"))
	g.AddNode(ut.Node("2"))
	g.AddNode(ut.Node("3"))
	g.AddNode(ut.Node("4"))

	g.AddEdge("0", 4, "1")
	g.AddEdge("0", 1, "2")
	g.AddEdge("1", 1, "3")
	g.AddEdge("2", 5, "3")
	g.AddEdge("2", 2, "1")
	g.AddEdge("3", 3, "4")

	expected := path.NodeDistanceMap{
		"0": 0,
		"1": 3,
		"2": 1,
		"3": 4,
		"4": 7,
	}

	actual := path.DijkstraDistances(g, "0")

	fmt.Printf("the distances\n")
	fmt.Printf("%#v\n", actual)

	distancesK := []string{}
	for dK := range actual {
		distancesK = append(distancesK, string(dK))
	}
	sort.Strings(distancesK)

	for _, dK := range distancesK {
		assert.Equal(T,
			expected[contract.NodeID(dK)], actual[contract.NodeID(dK)],
			"Node %s should have distance of %f", dK, expected[contract.NodeID(dK)],
		)
	}
}

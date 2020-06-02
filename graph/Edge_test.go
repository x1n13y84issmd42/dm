package graph_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/x1n13y84issmd42/dm/graph"
	"github.com/x1n13y84issmd42/dm/graph/ut"
)

func Test_DEdge_Reverse(T *testing.T) {
	e := graph.DEdge{
		A: ut.Node("A"),
		B: ut.Node("B"),
	}

	expected := graph.DEdge{
		A: e.B,
		B: e.A,
	}

	assert.Equal(T, expected, e.Reverse())
}

func Test_WEdge_Reverse(T *testing.T) {
	e := graph.WEdge{
		DEdge: graph.DEdge{
			A: ut.Node("A"),
			B: ut.Node("B"),
		},
		W: 0.333,
	}

	expected := graph.WEdge{
		DEdge: graph.DEdge{
			A: e.B,
			B: e.A,
		},
		W: 0.333,
	}

	assert.Equal(T, expected, e.Reverse())
}

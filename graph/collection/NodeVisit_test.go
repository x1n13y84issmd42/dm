package collection_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/x1n13y84issmd42/dm/graph/collection"
	"github.com/x1n13y84issmd42/dm/graph/contract"
)

func Test_NodeVisitMap(T *testing.T) {
	m := collection.NodeVisitMap{}
	n := contract.NodeID("piggy")

	T.Run("Visit", func(T *testing.T) {
		assert.False(T, m.Visited(n))

		m.Visit(n)

		assert.True(T, m.Visited(n))
		assert.Equal(T, int64(1), m.Visits(n))
	})

	T.Run("Visit more", func(T *testing.T) {
		m.Visit(n)
		m.Visit(n)
		m.Visit(n)

		assert.Equal(T, int64(4), m.Visits(n))
	})

	T.Run("Unknown", func(T *testing.T) {
		un := contract.NodeID("unknwn")

		assert.False(T, m.Visited(un))
		assert.Equal(T, int64(0), m.Visits(un))
	})
}

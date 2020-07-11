package storage_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/x1n13y84issmd42/gog/graph/collection"
	"github.com/x1n13y84issmd42/gog/graph/storage"
	"github.com/x1n13y84issmd42/gog/graph/ut"
)

func Test_AdjacencyList(T *testing.T) {
	al := storage.NewAdjacencyList()

	al.AddNode(ut.Node("R007"))
	al.AddNode(ut.Node("Down 1"))
	al.AddNode(ut.Node("Down 2"))
	al.AddNode(ut.Node("Up 1"))
	al.AddNode(ut.Node("Up 2"))

	al.AddEdge("R007", "Down 1")
	al.AddEdge("R007", "Down 2")
	al.AddEdge("R007", "Down 2")
	al.AddEdge("Up 1", "R007")
	al.AddEdge("Up 2", "R007")

	T.Run("Node", func(T *testing.T) {
		assert.Equal(T, ut.Node("R007"), al.Node("R007"))
		assert.Nil(T, al.Node("Root_INVALID"))
	})

	T.Run("Nodes", func(T *testing.T) {
		expected := collection.NewNodes(
			al.Node("R007"),
			al.Node("Down 1"),
			al.Node("Down 2"),
			al.Node("Up 1"),
			al.Node("Up 2"),
		)
		assert.Equal(T, expected, al.Nodes())
	})

	T.Run("AdjacentNodes", func(T *testing.T) {
		assert.Equal(T,
			collection.NewNodes(ut.Node("Down 1"), ut.Node("Down 2")),
			al.AdjacentNodes("R007"),
		)

		assert.Equal(T,
			collection.NewNodes(),
			al.AdjacentNodes("Root_INVALID"),
		)
	})

	T.Run("UpstreamNodes", func(T *testing.T) {
		usn := al.UpstreamNodes("R007")
		assert.Equal(T,
			collection.NewNodes(ut.Node("Up 1"), ut.Node("Up 2")),
			usn,
		)
	})
}

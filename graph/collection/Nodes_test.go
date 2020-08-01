package collection_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/x1n13y84issmd42/gog/graph/collection"
	"github.com/x1n13y84issmd42/gog/graph/contract"
)

type testNode struct {
	Name string
}

func (node *testNode) ID() contract.NodeID {
	return contract.NodeID(node.Name)
}

func tnode(name string) *testNode {
	node := testNode{
		Name: name,
	}

	return &node
}

func Test_Set(T *testing.T) {
	set := collection.NewNodes()

	T.Run("Empty", func(T *testing.T) {
		assert.Equal(T, uint(0), set.Count())
	})

	T.Run("Add", func(T *testing.T) {
		ares := set.Add(tnode("a"))
		assert.Equal(T, uint(1), set.Count())
		assert.False(T, ares)

		ares = set.Add(tnode("a"))
		assert.True(T, ares)
		set.Add(tnode("b"))
		assert.Equal(T, uint(2), set.Count())
	})

	T.Run("Get", func(T *testing.T) {
		assert.Equal(T, tnode("a"), set.Get("a"))
	})

	T.Run("Remove", func(T *testing.T) {
		ares := set.Remove("a")
		assert.True(T, ares)
		assert.Equal(T, uint(1), set.Count())
		ares = set.Remove("a")
		assert.False(T, ares)
		assert.Equal(T, uint(1), set.Count())
	})

	T.Run("Values", func(T *testing.T) {
		set := collection.NewNodes(tnode("a"), tnode("z"), tnode("c"))

		set.Add(tnode("b"))
		set.Add(tnode("e"))
		set.Add(tnode("a"))
		set.Add(tnode("z"))
		set.Add(tnode("z"))

		expected := []*testNode{
			tnode("a"),
			tnode("b"),
			tnode("c"),
			tnode("e"),
			tnode("z"),
		}

		actual := []*testNode{}

		for _, n := range set.Values() {
			actual = append(actual, n.(*testNode))
		}

		assert.Equal(T, expected, actual)
	})

	T.Run("Clone", func(T *testing.T) {
		set := collection.NewNodes(
			tnode("a"), tnode("z"),
			tnode("c"), tnode("b"),
			tnode("e"), tnode("a"),
			tnode("z"), tnode("c"))

		expected := collection.NewNodes(
			tnode("a"),
			tnode("b"),
			tnode("e"), tnode("c"),
			tnode("z"),
		)

		assert.Equal(T, expected, set.Clone())
	})
}

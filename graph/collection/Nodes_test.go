package collection_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/x1n13y84issmd42/dm/graph/collection"
	"github.com/x1n13y84issmd42/dm/graph/contract"
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

func Test_Set_Add_Get_Remove(T *testing.T) {
	set := collection.NewNodes()
	assert.Equal(T, 0, set.Count())

	ares := set.Add(tnode("a"))
	assert.Equal(T, 1, set.Count())
	assert.False(T, ares)

	ares = set.Add(tnode("a"))
	assert.True(T, ares)
	set.Add(tnode("b"))
	assert.Equal(T, 2, set.Count())

	assert.Equal(T, tnode("a"), set.Get("a"))

	ares = set.Remove("a")
	assert.True(T, ares)
	assert.Equal(T, 1, set.Count())

	ares = set.Remove("a")
	assert.False(T, ares)
	assert.Equal(T, 1, set.Count())
}

func Test_Set_Values(T *testing.T) {
	set := collection.NewNodes()

	set.Add(tnode("a"))
	set.Add(tnode("z"))
	set.Add(tnode("c"))
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
}

func Test_Set_Clone(T *testing.T) {
	set := collection.NewNodes()

	set.Add(tnode("a"))
	set.Add(tnode("z"))
	set.Add(tnode("c"))
	set.Add(tnode("b"))
	set.Add(tnode("e"))
	set.Add(tnode("a"))
	set.Add(tnode("z"))
	set.Add(tnode("z"))

	expected := collection.NewNodes()
	expected.Add(tnode("a"))
	expected.Add(tnode("b"))
	expected.Add(tnode("c"))
	expected.Add(tnode("e"))
	expected.Add(tnode("z"))

	assert.Equal(T, expected, set.Clone())
}

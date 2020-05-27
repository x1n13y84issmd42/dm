package nodes

import (
	"testing"

	"github.com/stretchr/testify/assert"
	// "github.com/x1n13y84issmd42/dm/graphs/ut"
)

// TestNode ...
type TestNode struct {
	Node
	Name string
}

// ID ...
func (node *TestNode) ID() NodeID {
	return NodeID(node.Name)
}

// Node ...
func TNode(name string) *TestNode {
	node := TestNode{
		Node: Node{},
		Name: name,
	}

	node.Node.INode = &node

	return &node
}

func Test_Set_Add_Get_Remove(T *testing.T) {
	set := NewNodes()
	assert.Equal(T, 0, set.Count())

	ares := set.Add(TNode("a"))
	assert.Equal(T, 1, set.Count())
	assert.False(T, ares)

	ares = set.Add(TNode("a"))
	assert.True(T, ares)
	set.Add(TNode("b"))
	assert.Equal(T, 2, set.Count())

	assert.Equal(T, TNode("a"), set.Get("a"))

	ares = set.Remove("a")
	assert.True(T, ares)
	assert.Equal(T, 1, set.Count())

	ares = set.Remove("a")
	assert.False(T, ares)
	assert.Equal(T, 1, set.Count())
}

func Test_Set_Values(T *testing.T) {
	set := NewNodes()

	set.Add(TNode("a"))
	set.Add(TNode("z"))
	set.Add(TNode("c"))
	set.Add(TNode("b"))
	set.Add(TNode("e"))
	set.Add(TNode("a"))
	set.Add(TNode("z"))
	set.Add(TNode("z"))

	expected := []*TestNode{
		TNode("a"),
		TNode("b"),
		TNode("c"),
		TNode("e"),
		TNode("z"),
	}

	actual := []*TestNode{}

	for _, n := range set.Values() {
		actual = append(actual, n.(*TestNode))
	}

	assert.Equal(T, expected, actual)
}

func Test_Set_Clone(T *testing.T) {
	set := NewNodes()

	set.Add(TNode("a"))
	set.Add(TNode("z"))
	set.Add(TNode("c"))
	set.Add(TNode("b"))
	set.Add(TNode("e"))
	set.Add(TNode("a"))
	set.Add(TNode("z"))
	set.Add(TNode("z"))

	expected := NewNodes()
	expected.Add(TNode("a"))
	expected.Add(TNode("b"))
	expected.Add(TNode("c"))
	expected.Add(TNode("e"))
	expected.Add(TNode("z"))

	assert.Equal(T, expected, set.Clone())
}

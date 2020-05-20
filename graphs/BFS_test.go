package graphs

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestNode struct {
	Node
	Name string
}

func Test_BFS_1(T *testing.T) {
	root := TestNode{
		Name: "L0",
	}

	root.Add(&TestNode{
		Name: "L10",
	}).Add(&TestNode{
		Name: "L20",
	}).Add(&TestNode{
		Name: "L30",
	})

	root.Add(&TestNode{
		Name: "L11",
	}).Add(&TestNode{
		Name: "L21",
	})

	root.Add(&TestNode{
		Name: "L12",
	})

	expected := "L0L10L11L12L20L21L30"
	actual := ""

	for node := range BFS(&root) {
		actual = fmt.Sprintf("%s%s", actual, node.(*TestNode).Name)
	}

	assert.Equal(T, expected, actual)
}

package graphs

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_DFS_1(T *testing.T) {
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

	expected := "L0L10L20L30L11L21L12"
	actual := ""

	for node := range DFS(&root) {
		actual = fmt.Sprintf("%s%s", actual, node.(*TestNode).Name)
	}

	assert.Equal(T, expected, actual)
}

func Test_DFS_Loop(T *testing.T) {
	root := TestNode{
		Name: "L0",
	}

	root.Add(&TestNode{
		Name: "L10",
	}).Add(&root)

	expected := "L0L10"
	actual := ""

	for node := range DFS(&root) {
		actual = fmt.Sprintf("%s%s", actual, node.(*TestNode).Name)
	}

	assert.Equal(T, expected, actual)
}

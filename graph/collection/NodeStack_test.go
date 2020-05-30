package collection

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/x1n13y84issmd42/dm/graph/contract"
	"github.com/x1n13y84issmd42/dm/graph/ut"
)

func Test_NodeStack(T *testing.T) {
	stack := NodeStack{}

	T.Run("Push", func(T *testing.T) {
		stack.Push(ut.Node("a"))
		stack.Push(ut.Node("b"))
		stack.Push(ut.Node("c"))

		assert.Equal(T, 3, len(stack))
	})

	T.Run("Append", func(T *testing.T) {
		stack.Append([]contract.Node{
			ut.Node("1"),
			ut.Node("2"),
			ut.Node("3"),
			ut.Node("4"),
		})

		assert.Equal(T, 7, len(stack))
	})

	T.Run("PopFront", func(T *testing.T) {
		n := stack.PopFront()

		assert.Equal(T, n.(*ut.TestNode).Name, "a")
		assert.Equal(T, 6, len(stack))
	})
}

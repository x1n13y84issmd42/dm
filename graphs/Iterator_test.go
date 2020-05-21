package graphs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NodeStack(T *testing.T) {
	stack := NodeStack{}

	T.Run("Push", func(T *testing.T) {
		stack.Push(&TestNode{Name: "a"})
		stack.Push(&TestNode{Name: "b"})
		stack.Push(&TestNode{Name: "c"})

		assert.Equal(T, 3, len(stack))
	})

	T.Run("Append", func(T *testing.T) {
		stack.Append([]IDNode{
			&TestNode{Name: "1"},
			&TestNode{Name: "2"},
			&TestNode{Name: "3"},
			&TestNode{Name: "4"},
		})

		assert.Equal(T, 7, len(stack))
	})

	T.Run("PopFront", func(T *testing.T) {
		n := stack.PopFront()

		assert.Equal(T, n.(*TestNode).Name, "a")
		assert.Equal(T, 6, len(stack))
	})
}

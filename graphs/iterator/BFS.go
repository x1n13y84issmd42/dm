package iterator

import (
	"github.com/x1n13y84issmd42/dm/graphs/nodes"
)

// BFS creates a breadth-first search iterator to traverse nodes.
func BFS(node nodes.INode) NChannel {
	ch := make(NChannel)
	stack := NodeStack{node}
	go func() {
		visited := NodeVisitedMap{}
		traverseBFS(ch, &stack, &visited)
		close(ch)
	}()
	return ch
}

func traverseBFS(ch NChannel, stack *NodeStack, visited *NodeVisitedMap) {
	for len(*stack) > 0 {
		node := stack.PopFront()

		if node != nil && (*visited)[node] == false {
			ch <- node
			(*visited)[node] = true

			if node.Adjacent().Count() > 0 {
				stack.Append(node.Adjacent().Values())
			}
		}
	}
}

package iterator

import (
	"github.com/x1n13y84issmd42/dm/graph/nodes"
)

// BFS creates a breadth-first search iterator to traverse nodes.
func BFS(graph nodes.NodeAccess, root nodes.NodeID) NChannel {
	ch := make(NChannel)
	stack := NodeStack{}
	go func() {
		stack.Push(graph.Node(root))
		visited := NodeVisitedMap{}
		traverseBFS(graph, ch, &stack, &visited)

		close(ch)
	}()
	return ch
}

func traverseBFS(graph nodes.NodeAccess, ch NChannel, stack *NodeStack, visited *NodeVisitedMap) {
	for len(*stack) > 0 {
		node := stack.PopFront()
		nID := node.ID()

		if node != nil && (*visited)[nID] == false {
			ch <- node
			(*visited)[nID] = true

			adj := graph.AdjacentNodes(nID)
			if adj.Count() > 0 {
				stack.Append(adj.Values())
			}
		}
	}
}

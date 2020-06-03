package iterator

import (
	"github.com/x1n13y84issmd42/dm/graph/collection"
	"github.com/x1n13y84issmd42/dm/graph/contract"
)

// BFS creates a breadth-first search iterator to traverse nodes.
func BFS(graph contract.NodeAccess, root contract.NodeID) contract.NChannel {
	ch := make(contract.NChannel)
	stack := collection.NodeStack{}
	go func() {
		stack.Push(graph.Node(root))
		visited := collection.NodeVisitMap{}
		traverseBFS(graph, ch, &stack, &visited, Forward)

		close(ch)
	}()
	return ch
}

// RBFS creates a reversed breadth-first search iterator to traverse nodes.
func RBFS(graph contract.NodeAccess, root contract.NodeID) contract.NChannel {
	ch := make(contract.NChannel)
	stack := collection.NodeStack{}
	go func() {
		stack.Push(graph.Node(root))
		visited := collection.NodeVisitMap{}
		traverseBFS(graph, ch, &stack, &visited, Backward)

		close(ch)
	}()
	return ch
}

func traverseBFS(
	graph contract.NodeAccess,
	ch contract.NChannel,
	stack *collection.NodeStack,
	visited *collection.NodeVisitMap,
	nextNodes contract.TraversalDirection,
) {
	for len(*stack) > 0 {
		node := stack.PopFront()
		nID := node.ID()

		if node != nil && (*visited)[nID] == false {
			ch <- node
			(*visited)[nID] = true

			next := nextNodes(graph, nID)
			if next.Count() > 0 {
				stack.Append(next.Values())
			}
		}
	}
}

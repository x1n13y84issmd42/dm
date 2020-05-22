package graphs

// BFS creates a breadth-first search iterator.
func BFS(node IDNode) IteratorChannel {
	ch := make(chan IDNode)
	stack := NodeStack{node}
	go func() {
		visited := NodeVisitedMap{}
		traverseBFS(ch, &stack, &visited)
		close(ch)
	}()
	return ch
}

func traverseBFS(ch IteratorChannel, stack *NodeStack, visited *NodeVisitedMap) {
	for len(*stack) > 0 {
		node := stack.PopFront()

		if node != nil && (*visited)[node] == false {
			ch <- node
			(*visited)[node] = true

			if node.NumChildren() > 0 {
				stack.Append(node.Children())
			}
		}
	}
}

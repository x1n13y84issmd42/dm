package graphs

// BFS creates a breadth-first search iterator.
func BFS(node IDNode) IteratorChannel {
	ch := make(chan IDNode)
	stack := NodeStack{node}
	go func() {
		traverseBFS(ch, &stack)
		close(ch)
	}()
	return ch
}

func traverseBFS(ch IteratorChannel, stack *NodeStack) {
	for len(*stack) > 0 {
		node := stack.PopFront()

		if node != nil {
			ch <- node

			if node.NumChildren() > 0 {
				stack.Append(node.Children())
			}
		}
	}
}

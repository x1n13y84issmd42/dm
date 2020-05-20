package graphs

// BFS creates a breadth-first search iterator.
func BFS(node INode) IteratorChannel {
	ch := make(chan INode)
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

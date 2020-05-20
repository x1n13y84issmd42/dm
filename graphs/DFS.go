package graphs

// DFS creates a depth-first search iterator.
func DFS(root INode) IteratorChannel {
	ch := make(chan INode)
	go func() {
		traverseDFS(root, ch)
		close(ch)
	}()
	return ch
}

func traverseDFS(node INode, ch IteratorChannel) {
	ch <- node
	for i := 0; i < node.NumChildren(); i++ {
		traverseDFS(node.Child(i), ch)
	}
}

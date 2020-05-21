package graphs

// DFS creates a depth-first search iterator.
func DFS(root IDNode) IteratorChannel {
	ch := make(chan IDNode)
	go func() {
		traverseDFS(root, ch)
		close(ch)
	}()
	return ch
}

func traverseDFS(node IDNode, ch IteratorChannel) {
	ch <- node
	for i := 0; i < node.NumChildren(); i++ {
		traverseDFS(node.Child(i), ch)
	}
}

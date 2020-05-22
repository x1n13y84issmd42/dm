package graphs

// DFS creates a depth-first search iterator.
func DFS(root IDNode) IteratorChannel {
	ch := make(chan IDNode)
	go func() {
		visited := NodeVisitedMap{}
		traverseDFS(root, ch, &visited)
		close(ch)
	}()
	return ch
}

func traverseDFS(node IDNode, ch IteratorChannel, visited *NodeVisitedMap) {
	if (*visited)[node] {
		return
	}

	ch <- node
	(*visited)[node] = true

	for i := 0; i < node.NumChildren(); i++ {
		traverseDFS(node.Child(i), ch, visited)
	}
}

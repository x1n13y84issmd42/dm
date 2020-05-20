package graphs

// IteratorChannel is a channel to deliver items while iterating.
// Exists to have a natural range syntax.
type IteratorChannel = chan INode

// DFS creates a depth-first search iterator.
func DFS(root INode) IteratorChannel {
	ch := make(chan INode, 1)
	go func() {
		traverse(root, ch)
		close(ch)
	}()
	return ch
}

func traverse(node INode, ch IteratorChannel) {
	ch <- node
	for i := 0; i < node.NumChildren(); i++ {
		traverse(node.Child(i), ch)
	}
}

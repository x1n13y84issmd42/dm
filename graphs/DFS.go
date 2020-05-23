package graphs

import "github.com/x1n13y84issmd42/dm/graphs/nodes"

// DFS creates a depth-first search iterator.
func DFS(root nodes.INode) IteratorChannel {
	ch := make(chan nodes.INode)
	go func() {
		visited := NodeVisitedMap{}
		traverseDFS(root, ch, &visited)
		close(ch)
	}()
	return ch
}

func traverseDFS(node nodes.INode, ch IteratorChannel, visited *NodeVisitedMap) {
	if (*visited)[node] {
		return
	}

	ch <- node
	(*visited)[node] = true

	for _, n := range node.Adjacent() {
		traverseDFS(n, ch, visited)
	}
}

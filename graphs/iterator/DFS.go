package iterator

import (
	"github.com/x1n13y84issmd42/dm/graphs/nodes"
)

// DFS creates a depth-first search iterator to traverse nodes.
func DFS(root nodes.INode) NChannel {
	ch := make(NChannel)
	go func() {
		visited := NodeVisitedMap{}
		traverseDFS(root, ch, &visited, false)
		close(ch)
	}()
	return ch
}

func traverseDFS(node nodes.INode, ch NChannel, visited *NodeVisitedMap, postorder bool) {
	if (*visited)[node] {
		return
	}

	if !postorder {
		ch <- node
		(*visited)[node] = true
	}

	for n := range node.Adjacent().Range() {
		traverseDFS(n, ch, visited, postorder)
	}

	if postorder {
		ch <- node
		(*visited)[node] = true
	}
}

// EDFS creates a depth-first search iterator to traverse edges.
func EDFS(root nodes.INode) EChannel {
	ch := make(EChannel)
	go func() {
		visited := NodeVisitedMap{}
		traverseEDFS(root, ch, &visited)
		close(ch)
	}()
	return ch
}

func traverseEDFS(node nodes.INode, ch EChannel, visited *NodeVisitedMap) {
	if (*visited)[node] {
		return
	}

	(*visited)[node] = true

	for n := range node.Adjacent().Range() {
		// ch <- Edge
		traverseEDFS(n, ch, visited)
	}
}

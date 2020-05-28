package iterator

import (
	"github.com/x1n13y84issmd42/dm/graphs/nodes"
)

// DFS creates a depth-first search iterator to traverse the graph.
func DFS(graph nodes.IAdjacency, root nodes.NodeID) NChannel {
	ch := make(NChannel)
	go func() {
		visited := NodeVisitedMap{}
		traverseDFS(graph, graph.Node(root), ch, &visited, false)
		close(ch)
	}()
	return ch
}

func traverseDFS(graph nodes.IAdjacency, node nodes.Node, ch NChannel, visited *NodeVisitedMap, postorder bool) {
	nID := node.ID()
	if (*visited)[nID] {
		return
	}

	if !postorder {
		ch <- node
		(*visited)[nID] = true
	}

	for n := range graph.AdjacentNodes(nID).Range() {
		traverseDFS(graph, n, ch, visited, postorder)
	}

	if postorder {
		ch <- node
		(*visited)[nID] = true
	}
}

/* // EDFS creates a depth-first search iterator to traverse edges.
func EDFS(root nodes.Node) EChannel {
	ch := make(EChannel)
	go func() {
		visited := NodeVisitedMap{}
		traverseEDFS(root, ch, &visited)
		close(ch)
	}()
	return ch
}

func traverseEDFS(node nodes.Node, ch EChannel, visited *NodeVisitedMap) {
	if (*visited)[node] {
		return
	}

	(*visited)[node] = true

	for n := range node.Adjacent().Range() {
		// ch <- Edge
		traverseEDFS(n, ch, visited)
	}
} */

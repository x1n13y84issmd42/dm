package iterator

import (
	"github.com/x1n13y84issmd42/dm/graph/collection"
	"github.com/x1n13y84issmd42/dm/graph/contract"
)

// DFS creates a depth-first search iterator to traverse the graph.
func DFS(graph contract.NodeAccess, root contract.NodeID) contract.NChannel {
	ch := make(contract.NChannel)
	go func() {
		visited := collection.NodeVisitMap{}
		traverseDFS(graph, graph.Node(root), ch, &visited, false)
		close(ch)
	}()
	return ch
}

func traverseDFS(graph contract.NodeAccess, node contract.Node, ch contract.NChannel, visited *collection.NodeVisitMap, postorder bool) {
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
		visited := collection.NodeVisitMap{}
		traverseEDFS(root, ch, &visited)
		close(ch)
	}()
	return ch
}

func traverseEDFS(node nodes.Node, ch EChannel, visited *collection.NodeVisitMap) {
	if (*visited)[node] {
		return
	}

	(*visited)[node] = true

	for n := range node.Adjacent().Range() {
		// ch <- Edge
		traverseEDFS(n, ch, visited)
	}
} */

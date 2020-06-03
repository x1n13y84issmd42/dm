package iterator

import (
	"github.com/x1n13y84issmd42/dm/graph/collection"
	"github.com/x1n13y84issmd42/dm/graph/contract"
)

// DFS creates a depth-first search iterator to traverse the graph.
func DFS(graph contract.NodeAccess, root contract.NodeID, traverse contract.TraversalOrder) contract.NChannel {
	ch := make(contract.NChannel)
	go func() {
		visited := collection.NodeVisitMap{}
		traverseDFS(graph, graph.Node(root), ch, &visited, traverse, Forward)
		close(ch)
	}()
	return ch
}

// RDFS creates a reversed depth-first search iterator to traverse the graph.
// Reversed search from a node N means walking inbpund edges from parent nodes of N depth-first.
func RDFS(graph contract.NodeAccess, root contract.NodeID, traverse contract.TraversalOrder) contract.NChannel {
	ch := make(contract.NChannel)
	go func() {
		visited := collection.NodeVisitMap{}
		traverseDFS(graph, graph.Node(root), ch, &visited, traverse, Backward)
		close(ch)
	}()
	return ch
}

func traverseDFS(
	graph contract.NodeAccess,
	node contract.Node,
	ch contract.NChannel,
	visited *collection.NodeVisitMap,
	traverse contract.TraversalOrder,
	nextNodes contract.TraversalDirection,
) {
	nID := node.ID()
	if visited.Visited(nID) {
		return
	}

	traverse(func() {
		ch <- node
		visited.Visit(nID)
	}, func() {
		for n := range nextNodes(graph, nID).Range() {
			traverseDFS(graph, n, ch, visited, traverse, nextNodes)
		}
	})
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

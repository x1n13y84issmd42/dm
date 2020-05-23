package graphs

import "github.com/x1n13y84issmd42/dm/graphs/nodes"

// KosarajuSCC finds strongly connected components of a graph using the Sambasiva Rao Kosaraju's algorithm.
func KosarajuSCC(root nodes.INode) {
	S := NodeStack{}

	for n := range DFS(root) {
		S.Push(n)
	}
}

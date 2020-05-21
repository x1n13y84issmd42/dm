package graphs

// KosarajuSCC finds strongly connected components of a graph using the Sambasiva Rao Kosaraju's algorithm.
func KosarajuSCC(root IDNode) {
	S := NodeStack{}

	for n := range DFS(root) {
		S.Push(n)
	}
}

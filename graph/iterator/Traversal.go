package iterator

// PreOrder implements a pre-order traversal, i.e. visits a node first, then it's adjacent nodes.
// Used in iterators.
func PreOrder(traverseNode func(), traverseAdjacent func()) {
	traverseNode()
	traverseAdjacent()
}

// PostOrder implements a post-order traversal, i.e. visits a node's adjacent nodes first, then the node itself.
// Used in iterators.
func PostOrder(traverseNode func(), traverseAdjacent func()) {
	traverseAdjacent()
	traverseNode()
}

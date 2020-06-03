package contract

// Traversal is a traversal order function for regular graphs.
type Traversal func(func(), func())

// BinaryTraversal is a traversal order function for binary trees.
type BinaryTraversal func(func(), func(), func())

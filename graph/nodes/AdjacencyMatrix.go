package nodes

// AdjacencyMatrix is a representation of graphs as a matrix of adjacent nodes.
type AdjacencyMatrix struct {
	Nodes Nodes
	M     [][]Node
}

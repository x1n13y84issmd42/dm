package contract

// Graph is an interface for unweighted graph.
type Graph interface {
	NodeAccess
	AddEdge(v1 Node, v2 Node)
}

// WGraph is an interface for weighted graph.
type WGraph interface {
	NodeAccess
	AddEdge(v1 Node, w float64, v2 Node)
}

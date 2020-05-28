package nodes

// IEdge is an interface for edges.
type IEdge interface {
	Reverse() IEdge
}

// Edge is a pair of nodes.
type Edge struct {
	A Node
	B Node
}

// WEdge is a weighted edge.
type WEdge struct {
	Edge
	W float64
}

// Reverse creates a new edge by swapping the receiver's nodes.
func (e Edge) Reverse() IEdge {
	return Edge{
		A: e.A,
		B: e.B,
	}
}

// Reverse creates a new edge by swapping the receiver's nodes.
func (e WEdge) Reverse() IEdge {
	return WEdge{
		Edge: Edge{
			A: e.A,
			B: e.B,
		},
		W: e.W,
	}
}

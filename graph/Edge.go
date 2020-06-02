package graph

import "github.com/x1n13y84issmd42/dm/graph/contract"

// DEdge is a pair of nodes making a directed edge.
type DEdge struct {
	A contract.Node
	B contract.Node
}

// WEdge is a weighted edge.
type WEdge struct {
	DEdge
	W float64
}

// Reverse creates a new edge by swapping the receiver's nodes.
func (e DEdge) Reverse() contract.Edge {
	return DEdge{
		A: e.B,
		B: e.A,
	}
}

// Reverse creates a new edge by swapping the receiver's nodes.
func (e WEdge) Reverse() contract.Edge {
	return WEdge{
		DEdge: DEdge{
			A: e.B,
			B: e.A,
		},
		W: e.W,
	}
}

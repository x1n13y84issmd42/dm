package path

import (
	"container/heap"
	"math"

	"github.com/x1n13y84issmd42/gog/graph/collection"
	"github.com/x1n13y84issmd42/gog/graph/contract"
)

// DijkstraSPT builds a shortest path tree from s to every other
// node in the weighted graph g using a Dijkstra's algorithm.
func DijkstraSPT(s contract.NodeID) {

}

// NodeDistanceMap is a map of NodeID -> distance.
type NodeDistanceMap map[contract.NodeID]float64

// Get returns a stored distance value for nID.
// If there is none, +infinity is returned.
func (ndm NodeDistanceMap) Get(nID contract.NodeID, add float64) float64 {
	if d, ok := ndm[nID]; ok {
		return d + add
	}

	return math.Inf(1) + add
}

// DijkstraDistances computes shortests distances
// from s to every other node in the weighted graph g.
func DijkstraDistances(g contract.WGraph, s contract.NodeID) NodeDistanceMap {
	distances := make(NodeDistanceMap)
	dheap := make(collection.NodeDistanceHeap, 0, g.Len())
	visited := collection.NodeVisitMap{}

	// Putting the starting node to the heap.
	heap.Push(&dheap, collection.NewNodeDistance(g.Node(s).ID(), 0))
	// and to the distance map
	distances[s] = 0

	for dheap.Len() > 0 {
		nd := heap.Pop(&dheap).(collection.NodeDistance)
		visited.Visit(nd.ID)

		for _, e := range g.OutEdges(nd.ID) {
			if visited.Visited(e.B.ID()) {
				continue
			}

			d := distances.Get(nd.ID, e.W)

			if d < distances.Get(e.B.ID(), 0) {
				distances[e.B.ID()] = d
				// e.A is always the id.ID node, so adding e.B to the heap.
				heap.Push(&dheap, collection.NewNodeDistance(e.B.ID(), e.W))
			}

		}
	}

	return distances
}

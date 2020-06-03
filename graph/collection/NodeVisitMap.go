package collection

import "github.com/x1n13y84issmd42/dm/graph/contract"

// NodeVisitMap is a map of visited nodes.
// Used in iterators.
type NodeVisitMap map[contract.NodeID]bool

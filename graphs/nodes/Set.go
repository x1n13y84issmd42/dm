package nodes

// Nodes is a set of INode instances.
type Nodes map[NodeID]INode

// Add adds a node to the set and returns true if it had been inserted for the first time.
func (set *Nodes) Add(n INode) bool {
	had := set.Has(n)
	(*set)[n.ID()] = n
	return had
}

// Remove removes a node from the set.
func (set *Nodes) Remove(n INode) {}

// Has tells whether a node is present in the set.
func (set *Nodes) Has(n INode) bool {
	_, ok := (*set)[n.ID()]
	return ok
}

// Count tells how many nodes are currently in the set.
func (set Nodes) Count() int {
	return len(set)
}

// Clone creates a new set by copying the receiver set.
func (set Nodes) Clone() Nodes {
	res := make(Nodes, set.Count())
	for _, n := range set {
		res.Add(n)
	}
	return res
}

// Values creates a slice of values taken from the set.
func (set Nodes) Values() []INode {
	res := []INode{}
	for _, n := range set {
		res = append(res, n)
	}
	return res
}

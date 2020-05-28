package nodes

import (
	"sort"
)

// Nodes is a set of INode instances.
type Nodes struct {
	Map map[NodeID]Node
}

// NewNodes creates a new node set instance.
func NewNodes() *Nodes {
	return &Nodes{
		Map: make(map[NodeID]Node),
	}
}

// Add adds a node to the set and returns true if it had been inserted for the first time.
func (set *Nodes) Add(n Node) bool {
	had := set.Has(n.ID())
	set.Map[n.ID()] = n
	return had
}

// Remove removes a node from the set. Returns true if node was present in the set before removal.
func (set *Nodes) Remove(nID NodeID) bool {
	had := set.Has(nID)
	delete(set.Map, nID)
	return had
}

// Has tells whether a node is present in the set.
func (set *Nodes) Has(nID NodeID) bool {
	_, ok := set.Map[nID]
	return ok
}

// Get returns a node from the set.
func (set *Nodes) Get(nID NodeID) Node {
	return set.Map[nID]
}

// Count tells how many nodes are currently in the set.
func (set Nodes) Count() int {
	return len(set.Map)
}

// Clone creates a new set by copying the receiver set.
func (set Nodes) Clone() *Nodes {
	res := NewNodes()
	for _, n := range set.Map {
		res.Add(n)
	}
	return res
}

// Values creates a slice of values taken from the set.
func (set Nodes) Values() []Node {
	res := []Node{}
	for n := range set.Range() {
		res = append(res, n)
	}
	return res
}

// Channel is a channel to deliver items while iterating.
// Exists to have a natural range syntax.
type Channel chan Node

// Range is an attempt to make iteration over a map-based set stable in terms of order.
func (set Nodes) Range() Channel {
	ch := make(chan Node)
	go func() {
		// Collecting keys.
		keys := []string{}
		for nID := range set.Map {
			keys = append(keys, string(nID))
		}
		// Sorting them.
		sort.Strings(keys)

		// Sending values to the channel.
		for _, k := range keys {
			ch <- set.Map[NodeID(k)]
		}

		close(ch)
	}()
	return ch
}

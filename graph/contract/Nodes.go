package contract

// Channel is a channel to deliver items while iterating over a node set.
// Exists to have a natural range syntax.
type Channel chan Node

// Nodes is a set of nodes.
type Nodes interface {
	Add(n Node) bool
	Remove(nID NodeID) bool
	Has(nID NodeID) bool
	Get(nID NodeID) Node
	Count() int
	Clone() Nodes
	Values() []Node
	Range() Channel
}

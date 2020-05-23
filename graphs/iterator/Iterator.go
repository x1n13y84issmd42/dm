package iterator

import "github.com/x1n13y84issmd42/dm/graphs/nodes"

// Channel is a channel to deliver items while iterating.
// Exists to have a natural range syntax.
type Channel chan nodes.INode

// NodeStack is a slice of nodes.
type NodeStack []nodes.INode

// Push adds a node to the end of the stack.
func (stack *NodeStack) Push(node nodes.INode) {
	*stack = append(*stack, node)
}

// Append appends a list of nodes to the end of the stack.
func (stack *NodeStack) Append(nodes []nodes.INode) {
	*stack = append(*stack, nodes...)
}

// PopFront removes the first node from the stack and returns it.
func (stack *NodeStack) PopFront() nodes.INode {
	s := *stack
	if len(s) >= 1 {
		res := s[0]
		*stack = s[1:]
		return res
	}

	return nil
}

// NodeVisitedMap is a map of visited nodes.
// Used in iterators.
type NodeVisitedMap map[nodes.INode]bool

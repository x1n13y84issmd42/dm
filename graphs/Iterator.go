package graphs

// IteratorChannel is a channel to deliver items while iterating.
// Exists to have a natural range syntax.
type IteratorChannel chan IDNode

// NodeStack is a slice of nodes.
type NodeStack []IDNode

// Push adds a node to the end of the stack.
func (stack *NodeStack) Push(node IDNode) {
	*stack = append(*stack, node)
}

// Append appends a list of nodes to the end of the stack.
func (stack *NodeStack) Append(nodes []IDNode) {
	*stack = append(*stack, nodes...)
}

// PopFront removes the first node from the stack and returns it.
func (stack *NodeStack) PopFront() IDNode {
	if len(*stack) >= 1 {
		res := (*stack)[0]
		*stack = (*stack)[1:]
		return res
	}

	return nil
}

// NodeVisitedMap is a map of visited nodes.
// Used in iterators.
type NodeVisitedMap map[IDNode]bool

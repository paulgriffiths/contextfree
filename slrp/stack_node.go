package slrp

import (
	"github.com/paulgriffiths/contextfree/tree"
	"github.com/paulgriffiths/gods/stacks"
)

// StackNode implements a stack of node elements.
type StackNode struct {
	stack stacks.StackInterface
}

// NewStackNode creates a new stack of node elements.
func NewStackNode() StackNode {
	return StackNode{stacks.NewStackInterface()}
}

// Push pushes a new node element onto the stack.
func (s *StackNode) Push(n *tree.Node) {
	s.stack.Push(n)
}

// Pop pops the top node element from the stack.
func (s *StackNode) Pop() *tree.Node {
	return s.stack.Pop().(*tree.Node)
}

// Peek returns the top node element from the stack without
// removing it.
func (s *StackNode) Peek() *tree.Node {
	return s.stack.Peek().(*tree.Node)
}

// IsEmpty returns true if the stack is empty, otherwise false.
func (s *StackNode) IsEmpty() bool {
	return s.stack.IsEmpty()
}

package evm

import (
    "github.com/holiman/uint256"
)

// Stack is a Stack datastructure with big.Int as datatype
type Stack []*uint256.Int

// Push appends a *uint256.Int to the top of the stack
// and returns the stack.
//
// After the value has been appended to the stack
// the function should return the new slice to the
// stack in context.
func (s Stack) Push(c *uint256.Int) Stack {
    return append(s, c)
}

// Pop removes a *uint256.Int from the stack and returns
// the Stack and the value that was removed from the top.
func (s Stack) Pop() (Stack, *uint256.Int) {
    if s.IsEmpty() {
        return s, uint256.NewInt()
    }
    return s[:len(s)-1], s[len(s)-1]
}

// IsEmpty Checks if Stack is empty and returns a bool
// value depending on the outcome.
func (s Stack) IsEmpty() bool {
    return len(s) == 0
}

// Peek pulls the head value of the stack
// without removing it.
func (s Stack) Peek() *uint256.Int {
    return s[len(s)-1]
}
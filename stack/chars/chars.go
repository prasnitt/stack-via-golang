// Package chars implement stack of characters (rune) of a string
package chars

import "github.com/prasnitt/stack-via-golang/stack"

const invalidItem = -1

// Considering string is as list of `rune`
type Stack []rune

// IsEmpty checks if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// String convert the current stack data to a string
func (s *Stack) String() string {
	return string(*s)
}

// Push will add new a character `c` on the top of stack
func (s *Stack) Push(c rune) {
	*s = append(*s, c)
}

// Pop will remove and return top character of stack. Return error if stack is empty.
func (s *Stack) Pop() (rune, error) {
	if s.IsEmpty() {
		return invalidItem, stack.ErrorEmptyStack
	} else {

		item := (*s)[len(*s)-1] // Index into the slice and obtain the element.
		*s = (*s)[:len(*s)-1]   // Remove it from the stack by slicing it off.
		return item, nil
	}
}

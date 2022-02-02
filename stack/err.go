package stack

import (
	"fmt"
)

// ErrorEmptyStack will be raised during Pop operation from empty stack
var ErrorEmptyStack = fmt.Errorf("empty stack")

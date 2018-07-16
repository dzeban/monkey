package object

import (
	"fmt"
)

// Error represent an error object returned by evaluator
type Error struct {
	Message string
}

// Inspect implements Object interface
func (e *Error) Inspect() string {
	return fmt.Sprintf("ERROR: %s", e.Message)
}

// Type implements Object interface
func (e *Error) Type() Type {
	return TypeError
}

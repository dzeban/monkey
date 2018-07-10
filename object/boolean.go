package object

import (
	"fmt"
)

// Boolean represent a boolean object
type Boolean struct {
	Value bool
}

// Inspect implements Object interface
func (b *Boolean) Inspect() string {
	return fmt.Sprintf("%t", b.Value)
}

// Type implements Object interface
func (b *Boolean) Type() Type {
	return TypeBoolean
}

package object

import (
	"fmt"
)

// Integer represent an integer object
type Integer struct {
	Value int64
}

// Inspect implements Object interface
func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

// Type implements Object interface
func (i *Integer) Type() Type {
	return TypeInteger
}

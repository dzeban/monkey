package object

// ReturnValue wraps another object for return statements
type ReturnValue struct {
	Value Object
}

// Inspect implements Object interface
func (rv *ReturnValue) Inspect() string {
	return rv.Value.Inspect()
}

// Type implements Object interface
func (rv *ReturnValue) Type() Type {
	return TypeReturnValue
}

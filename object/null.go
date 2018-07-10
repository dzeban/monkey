package object

// Null represent a null object
type Null struct{}

// Inspect implements Object interface
func (n *Null) Inspect() string {
	return "null"
}

// Type implements Object interface
func (n *Null) Type() Type {
	return TypeNull
}

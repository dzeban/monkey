package object

// Type is a string describing type of the object in Monkey language
type Type string

// Object types
const (
	TypeInteger = "INTEGER"
	TypeBoolean = "BOOLEAN"
	TypeNull    = "NULL"
)

// Object defines contract that should be implemented by all objects like integers
type Object interface {
	Type() Type
	Inspect() string
}

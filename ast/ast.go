package ast

// Node is a genenric AST node, that can print it's token literal
type Node interface {
	TokenLiteral() string
}

// Statement is a type for the language statements
type Statement interface {
	Node
	statementNode()
}

// Expression holds an expression of the language
type Expression interface {
	Node
	expressionNode()
}

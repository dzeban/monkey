package ast

import "github.com/dzeban/monkey/token"

// ExpressionStatement is a type for the statements consisting of a
// single expression, e.g.
//
// let x = 5; // this is a LetStatement
// x + 3;     // this is an ExpressionStatement
//
// We use it as a wrapper for parsing expressions
type ExpressionStatement struct {
	Token      token.Token // The first token of the expression
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}

// TokenLiteral returns the first expression's token.
// TokenLiteral implements ast.Node interface.
func (es *ExpressionStatement) TokenLiteral() string { return es.Token.Literal }

// String prints text representation of ExpressionStatement
func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}

	return ""
}

package ast

import "github.com/dzeban/monkey/token"

// ReturnStatement is a type for "return" statements.
// It consists of returned expression.
type ReturnStatement struct {
	Token token.Token
	Expr  Expression
}

func (rs *ReturnStatement) statementNode() {}

// TokenLiteral is an implementation of node interface for ReturnStatement
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

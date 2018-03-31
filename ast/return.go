package ast

import (
	"bytes"

	"github.com/dzeban/monkey/token"
)

// ReturnStatement is a type for "return" statements.
// It consists of returned expression.
type ReturnStatement struct {
	Token token.Token
	Expr  Expression
}

func (rs *ReturnStatement) statementNode() {}

// TokenLiteral is an implementation of node interface for ReturnStatement
func (rs *ReturnStatement) TokenLiteral() string { return rs.Token.Literal }

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.Expr != nil {
		out.WriteString(rs.Expr.String())
	}

	out.WriteString(";")
	return out.String()
}

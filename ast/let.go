package ast

import (
	"bytes"

	"github.com/dzeban/monkey/token"
)

// LetStatement is a type for "let" statements.
// It consists of identifier and expression that it binds to the identifier
type LetStatement struct {
	Token token.Token
	Ident *Identifier
	Expr  Expression
}

func (ls *LetStatement) statementNode() {}

// TokenLiteral is an implementation of node interface for LetStatement
func (ls *LetStatement) TokenLiteral() string { return ls.Token.Literal }

func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Ident.Value)
	out.WriteString(" = ")

	if ls.Expr != nil {
		out.WriteString(ls.Expr.String())
	}

	out.WriteString(";")
	return out.String()
}

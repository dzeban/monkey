package ast

import "github.com/dzeban/monkey/token"

// Identifier keeps a string name of identifier
type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode() {}

// TokenLiteral is an implementation of node interface for Identifier
func (i *Identifier) TokenLiteral() string { return i.Token.Literal }

// String returns text representation of Identifier
func (i *Identifier) String() string { return i.Value }

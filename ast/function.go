package ast

import (
	"bytes"
	"strings"

	"github.com/dzeban/monkey/token"
)

// FunctionExpression represents function expressions like this:
// fn (x, y) { return x+y; }
type FunctionExpression struct {
	Token  token.Token
	Params []*Identifier
	Body   *BlockStatement
}

func (fe *FunctionExpression) expressionNode()      {}
func (fe *FunctionExpression) TokenLiteral() string { return fe.Token.Literal }
func (fe *FunctionExpression) String() string {
	var params []string
	for _, p := range fe.Params {
		params = append(params, p.String())
	}

	var out bytes.Buffer
	out.WriteString(fe.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ","))
	out.WriteString(")")
	out.WriteString(fe.Body.String())

	return out.String()
}

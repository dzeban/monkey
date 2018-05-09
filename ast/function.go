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

func (fl *FunctionExpression) expressionNode()      {}
func (fl *FunctionExpression) TokenLiteral() string { return fl.Token.Literal }
func (fl *FunctionExpression) String() string {
	var params []string
	for _, p := range fl.Params {
		params = append(params, p.String())
	}

	var out bytes.Buffer
	out.WriteString(fl.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ","))
	out.WriteString(")")
	out.WriteString(fl.Body.String())

	return out.String()
}

package ast

import (
	"bytes"
	"strings"

	"github.com/dzeban/monkey/token"
)

// CallExpression represents the expression that calls a function by identifier
// or a function literal.
// Examples:
//    add(2, 3)             // call by identifier
//    fn(x, y){x + y}(2, 3) // call by function expression
type CallExpression struct {
	Token     token.Token // The '(' token
	Function  Expression  // Identifier or FunctionExpression
	Arguments []Expression
}

func (ce *CallExpression) expressionNode()      {}
func (ce *CallExpression) TokenLiteral() string { return ce.Token.Literal }
func (ce *CallExpression) String() string {
	var out bytes.Buffer

	args := []string{}
	for _, a := range ce.Arguments {
		args = append(args, a.String())
	}

	out.WriteString(ce.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")

	return out.String()
}

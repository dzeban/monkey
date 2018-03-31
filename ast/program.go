package ast

import (
	"bytes"
)

// Program is a collection of statements
type Program struct {
	Statements []Statement
}

// TokenLiteral is an impleentation of Node interface for Program
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}

	return ""
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

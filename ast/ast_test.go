package ast

import (
	"testing"

	"github.com/dzeban/monkey/token"
)

func TestString(t *testing.T) {
	programString := "let myVar = anotherVar;"
	parsedProgram := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"},
				Ident: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Expr: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}

	if parsedProgram.String() != programString {
		t.Errorf("program.String() wrong. got=%q", parsedProgram.String())
	}
}

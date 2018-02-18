package parser

import (
	"testing"

	"github.com/dzeban/monkey/ast"
	"github.com/dzeban/monkey/lexer"
)

func TestLetStatements(t *testing.T) {
	input := `
let x = 5;
let y = 10;
let foobar = 182182;
`
	l := lexer.New(input)
	p := New(l)

	program := p.ParseProgram()
	if program == nil {
		t.Fatal("ParseProgram() returned nil")
	}

	if len(program.Statements) != 3 {
		t.Fatalf("program.Statemnts must contain 3 statements, got %d",
			len(program.Statements))
	}

	tests := []struct {
		expectedIdentifier string
	}{
		{"x"},
		{"y"},
		{"foobar"},
	}

	for i, tt := range tests {
		stmt := program.Statements[i]
		if !testLetStatement(t, stmt, tt.expectedIdentifier) {
			return
		}
	}
}

func testLetStatement(t *testing.T, s ast.Statement, name string) bool {
	if s.TokenLiteral() != "let" {
		t.Errorf("TokenLiter must be 'let', got=%q", s.TokenLiteral())
		return false
	}

	letStmt, ok := s.(*ast.LetStatement)
	if !ok {
		t.Errorf("s is not *ast.LetStatment, it's %T", s)
		return false
	}

	if letStmt.Ident.Value != name {
		t.Errorf("letStmt.Name.Value is not '%s', it's '%s'", name, letStmt.Ident.Value)
		return false
	}

	if letStmt.Ident.TokenLiteral() != name {
		t.Errorf("s.Name is not '%s', it's '%s'", name, letStmt.Ident)
		return false
	}

	return true
}

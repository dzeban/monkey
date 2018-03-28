package ast

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

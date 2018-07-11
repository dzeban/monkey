package eval

import (
	"github.com/dzeban/monkey/ast"
	"github.com/dzeban/monkey/object"
)

// Single instances for trivial eval values
var (
	NULL  = &object.Null{}
	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
)

// Eval evaluates parsed AST
func Eval(node ast.Node) object.Object {
	switch node.(type) {
	case *ast.Program:
		p := node.(*ast.Program)
		return evalStatements(p.Statements)

	case *ast.ExpressionStatement:
		s := node.(*ast.ExpressionStatement)
		return Eval(s.Expression)

	case *ast.IntegerLiteral:
		il := node.(*ast.IntegerLiteral)
		return &object.Integer{Value: il.Value}

	case *ast.Boolean:
		b := node.(*ast.Boolean)
		if b.Value {
			return TRUE
		}
		return FALSE
	}

	return nil
}

func evalStatements(stmts []ast.Statement) object.Object {
	var result object.Object

	for _, stmt := range stmts {
		result = Eval(stmt)
	}

	return result
}

package eval

import (
	"fmt"

	"github.com/dzeban/monkey/ast"
	"github.com/dzeban/monkey/object"
)

// Single instances for trivial eval values
var (
	NULL  = &object.Null{}
	TRUE  = &object.Boolean{Value: true}
	FALSE = &object.Boolean{Value: false}
)

func isError(obj object.Object) bool {
	if obj != nil {
		return obj.Type() == object.TypeError
	}

	return false
}

// Eval evaluates parsed AST
func Eval(node ast.Node, env *object.Environment) object.Object {
	switch node.(type) {
	case *ast.Program:
		p := node.(*ast.Program)
		return evalProgram(p, env)

	case *ast.ExpressionStatement:
		s := node.(*ast.ExpressionStatement)
		return Eval(s.Expression, env)

	case *ast.BlockStatement:
		bs := node.(*ast.BlockStatement)
		return evalBlockStatement(bs, env)

	case *ast.IfExpression:
		ie := node.(*ast.IfExpression)
		return evalIfExpression(ie, env)

	case *ast.ReturnStatement:
		rs := node.(*ast.ReturnStatement)

		val := Eval(rs.Expr, env)
		if isError(val) {
			return val
		}

		return &object.ReturnValue{Value: val}

	case *ast.FunctionExpression:
		fe := node.(*ast.FunctionExpression)

		return &object.Function{Params: fe.Params, Body: fe.Body, Env: env}

	case *ast.CallExpression:
		ce := node.(*ast.CallExpression)

		fn := Eval(ce.Function, env)
		if isError(fn) {
			return fn
		}

		args := evalExpressions(ce.Arguments, env)
		if len(args) == 1 && isError(args[0]) {
			return args[0]
		}

		return applyFunction(fn, args)

	case *ast.LetStatement:
		ls := node.(*ast.LetStatement)

		val := Eval(ls.Expr, env)
		if isError(val) {
			return val
		}

		env.Set(ls.Ident.Value, val)

	case *ast.Identifier:
		i := node.(*ast.Identifier)
		return evalIdentifier(i, env)

	case *ast.IntegerLiteral:
		il := node.(*ast.IntegerLiteral)
		return &object.Integer{Value: il.Value}

	case *ast.Boolean:
		b := node.(*ast.Boolean)
		return boolToBoolean(b.Value)

	case *ast.PrefixExpression:
		e := node.(*ast.PrefixExpression)

		right := Eval(e.Right, env)
		if isError(right) {
			return right
		}

		return evalPrefixExpression(e.Operator, right)

	case *ast.InfixExpression:
		e := node.(*ast.InfixExpression)

		left := Eval(e.Left, env)
		if isError(left) {
			return left
		}

		right := Eval(e.Right, env)
		if isError(right) {
			return right
		}

		return evalInfixExpression(e.Operator, left, right)

	}

	return nil
}

func evalProgram(p *ast.Program, env *object.Environment) object.Object {
	var result object.Object

	for _, stmt := range p.Statements {
		result = Eval(stmt, env)

		switch result := result.(type) {
		case *object.ReturnValue:
			return result.Value
		case *object.Error:
			return result
		}
	}

	return result
}

func evalBlockStatement(bs *ast.BlockStatement, env *object.Environment) object.Object {
	var result object.Object

	for _, statement := range bs.Statements {
		result = Eval(statement, env)

		if result != nil {
			rt := result.Type()
			if rt == object.TypeReturnValue || rt == object.TypeError {
				return result
			}
		}
	}

	return result
}

func evalPrefixExpression(op string, right object.Object) object.Object {
	switch op {
	case "!":
		return evalBangOperator(right)
	case "-":
		return evalMinusPrefixOperator(right)
	default:
		return newError("unknown operator: %s%s:", op, right.Type())
	}
}

func evalBangOperator(right object.Object) object.Object {
	switch right {
	case TRUE:
		return FALSE
	case FALSE:
		return TRUE
	case NULL:
		return TRUE
	default:
		return FALSE
	}
}

func evalMinusPrefixOperator(right object.Object) object.Object {
	if right.Type() != object.TypeInteger {
		return newError("unknown operator: -%s", right.Type())
	}

	value := right.(*object.Integer).Value
	return &object.Integer{Value: -value}
}

func evalInfixExpression(op string, left, right object.Object) object.Object {
	if left.Type() == object.TypeInteger && right.Type() == object.TypeInteger {
		return evalIntegerInfixExpression(op, left, right)
	}

	if left.Type() != right.Type() {
		return newError("type mismatch: %s %s %s", left.Type(), op, right.Type())
	}

	if op == "==" {
		// This is a pointers comparison because booleans are single
		// instance values declared as package vars TRUE and FALSE.
		return boolToBoolean(left == right)
	}

	if op == "!=" {
		// This is a pointers comparison because booleans are single
		// instance values declared as package vars TRUE and FALSE.
		return boolToBoolean(left != right)
	}

	return newError("unknown operator: %s %s %s", left.Type(), op, right.Type())
}

func evalIntegerInfixExpression(op string, left, right object.Object) object.Object {
	leftVal := left.(*object.Integer).Value
	rightVal := right.(*object.Integer).Value

	switch op {
	case "+":
		return &object.Integer{Value: leftVal + rightVal}
	case "-":
		return &object.Integer{Value: leftVal - rightVal}
	case "*":
		return &object.Integer{Value: leftVal * rightVal}
	case "/":
		return &object.Integer{Value: leftVal / rightVal}
	case "<":
		return boolToBoolean(leftVal < rightVal)
	case ">":
		return boolToBoolean(leftVal > rightVal)
	case "==":
		return boolToBoolean(leftVal == rightVal)
	case "!=":
		return boolToBoolean(leftVal != rightVal)
	default:
		return newError("unknown operator: %s %s %s", left.Type(), op, right.Type())
	}
}

func boolToBoolean(b bool) *object.Boolean {
	if b {
		return TRUE
	}
	return FALSE
}

func evalIfExpression(ie *ast.IfExpression, env *object.Environment) object.Object {
	condition := Eval(ie.Condition, env)

	if isTruthy(condition) {
		return Eval(ie.Consequence, env)
	} else if ie.Alternative != nil {
		return Eval(ie.Alternative, env)
	} else {
		return NULL
	}
}

func isTruthy(obj object.Object) bool {
	switch obj {
	case NULL:
		return false
	case FALSE:
		return false
	case TRUE:
		return true
	default:
		return true
	}
}

func newError(format string, a ...interface{}) *object.Error {
	return &object.Error{Message: fmt.Sprintf(format, a...)}
}

func evalIdentifier(ident *ast.Identifier, env *object.Environment) object.Object {
	val, ok := env.Get(ident.Value)
	if !ok {
		return newError("identifier not found: " + ident.Value)
	}

	return val
}

func evalExpressions(exps []ast.Expression, env *object.Environment) []object.Object {
	var result []object.Object

	for _, e := range exps {
		evaluated := Eval(e, env)
		if isError(evaluated) {
			return []object.Object{evaluated}
		}

		result = append(result, evaluated)
	}

	return result
}

func applyFunction(obj object.Object, args []object.Object) object.Object {
	fn, ok := obj.(*object.Function)
	if !ok {
		return newError("not a function: %s", obj.Type())
	}

	extendedEnv := extendFunctionEnv(fn, args)
	evaluated := Eval(fn.Body, extendedEnv)
	return unwrapReturnValue(evaluated)
}

func extendFunctionEnv(fn *object.Function, args []object.Object) *object.Environment {
	env := object.NewEnclosedEnvironment(fn.Env)

	for i, param := range fn.Params {
		env.Set(param.Value, args[i])
	}

	return env
}

func unwrapReturnValue(obj object.Object) object.Object {
	if returnValue, ok := obj.(*object.ReturnValue); ok {
		return returnValue.Value
	}

	return obj
}

package object

import (
	"bytes"
	"strings"

	"github.com/dzeban/monkey/ast"
)

// Function represents function objects in evaluator
type Function struct {
	Params []*ast.Identifier
	Body   *ast.BlockStatement
	Env    *Environment
}

// Inspect implements Object interface
func (f *Function) Inspect() string {
	var out bytes.Buffer
	var params []string

	for _, p := range f.Params {
		params = append(params, p.String())
	}

	out.WriteString("fn(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")

	return out.String()
}

// Type implements Object interface
func (f *Function) Type() Type {
	return TypeFunction
}

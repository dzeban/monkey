package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/dzeban/monkey/object"

	"github.com/dzeban/monkey/eval"
	"github.com/dzeban/monkey/lexer"
	"github.com/dzeban/monkey/parser"
)

// PROMPT is a REPL prompt
const PROMPT = ">> "

// Start creates new REPL
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

	for {
		fmt.Print(PROMPT)

		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()

		if len(p.Errors()) > 0 {
			printParserErrors(p.Errors(), out)
			continue
		}

		e := eval.Eval(program, env)
		if e != nil {
			io.WriteString(out, e.Inspect())
			io.WriteString(out, "\n")
		} else {
			io.WriteString(out, "failed to eval program")
			io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(errors []string, out io.Writer) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}

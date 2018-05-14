package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/dzeban/monkey/lexer"
	"github.com/dzeban/monkey/parser"
)

// PROMPT is a REPL prompt
const PROMPT = ">> "

// Start creates new REPL
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

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

		io.WriteString(out, program.String())
		io.WriteString(out, "\n")
	}
}

func printParserErrors(errors []string, out io.Writer) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}

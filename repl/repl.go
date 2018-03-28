package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/dzeban/monkey/lexer"
	"github.com/dzeban/monkey/token"
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

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}

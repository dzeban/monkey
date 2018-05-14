package main

import (
	"os"

	"github.com/dzeban/monkey/repl"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}

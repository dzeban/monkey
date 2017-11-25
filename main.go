package main

import (
	"github.com/dzeban/monkey/repl"

	"os"
)

func main() {
	repl.Start(os.Stdin, os.Stdout)
}

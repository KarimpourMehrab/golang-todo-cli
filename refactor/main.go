package main

import (
	"bufio"
	"flag"
	"os"
	"todo/refactor/helper"
	"todo/refactor/input"
)

func main() {
	command := flag.String("command", "no command", "TODO app")
	scanner := bufio.NewScanner(os.Stdin)
	for {
		input.Handler(helper.RemoveSpace(*command), scanner)
	}
}

package main

import (
	"bufio"
	"fmt"
	"github.com/paulgriffiths/contextfree/pp"
	"os"
	"strings"
)

func main() {
	parser, err := pp.FromReader(strings.NewReader(grammarString))
	if err != nil {
		fmt.Fprintf(os.Stderr, "couldn't get parser: %v", err)
		os.Exit(1)
	}

	fmt.Printf("Enter simple arithmetical expressions ('q' to quit).\n")
	sc := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("> ")

		if !sc.Scan() || sc.Text() == "Q" || sc.Text() == "q" {
			break
		}

		t := parser.Parse(sc.Text())
		if t == nil {
			fmt.Printf("Illegal expression.\n")
			continue
		}

		fmt.Printf(" = %g\n", eval(t))
	}
}

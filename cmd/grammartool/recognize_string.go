package main

import (
	"bytes"
	"fmt"
	"github.com/paulgriffiths/contextfree/grammar"
	"github.com/paulgriffiths/contextfree/rdp"
	"os"
)

func recognizeRdpString(g *grammar.Grammar, input string, verbose bool) {
	parser, err := rdp.New(g)
	if err != nil {
		fmt.Fprintf(os.Stderr,
			"grammartool: ouldn't create parser: %v\n", err)
		os.Exit(1)
	}

	tree := parser.Parse(input)
	if tree == nil {
		fmt.Printf("Grammar does not recognize string '%s'.\n", input)
	} else if verbose {
		fmt.Printf("Grammar recognizes string '%s'.\n", input)
	}
}

func parseRdpString(g *grammar.Grammar, input string) {
	parser, err := rdp.New(g)
	if err != nil {
		fmt.Fprintf(os.Stderr,
			"grammartool: couldn't create parser: %v\n", err)
		os.Exit(1)
	}

	tree := parser.Parse(input)
	if tree == nil {
		fmt.Printf("Grammar does not recognize string '%s'.\n", input)
		return
	}

	outBuffer := bytes.NewBuffer(nil)
	tree.WriteBracketed(outBuffer, "`")
	output := string(outBuffer.Bytes())

	fmt.Printf("%s\n", output)
}

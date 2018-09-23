package main

import (
	"bytes"
	"fmt"
	"github.com/paulgriffiths/contextfree/grammar"
	"github.com/paulgriffiths/contextfree/rdp"
	"os"
)

func recognizeRdpString(g *grammar.Grammar, input string) {
	parser, err := rdp.New(g)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't create parser: %v\n", err)
		return
	}

	tree := parser.Parse(input)
	if tree != nil {
		fmt.Printf("Grammar recognizes string '%s'.\n", input)
	} else {
		fmt.Printf("Grammar does not recognize string '%s'.\n", input)
	}
}

func parseRdpString(g *grammar.Grammar, input string) {
	parser, err := rdp.New(g)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Couldn't create parser: %v\n", err)
		return
	}

	tree := parser.Parse(input)
	if tree == nil {
		fmt.Printf("Grammar does not recognize string '%s'.\n", input)
		return
	}

	outBuffer := bytes.NewBuffer(nil)
	tree.WriteBracketed(outBuffer, "`")
	output := string(outBuffer.Bytes())

	fmt.Printf("Parse tree for string '%s': %s\n", input, output)
}

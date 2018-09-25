package main

import (
	"fmt"
	"github.com/paulgriffiths/contextfree/grammar"
	"github.com/paulgriffiths/lexer"
	"math"
	"os"
	"strings"
)

func showLexerOutput(g *grammar.Grammar, input string) {
	l, err := lexer.New(g.Terminals)
	if err != nil {
		fmt.Fprintf(os.Stderr, "grammartool: couldn't create "+
			"lexer: %v\n", err)
		os.Exit(1)
	}

	tokens, err := l.Lex(strings.NewReader(input))
	if err != nil {
		fmt.Fprintf(os.Stderr,
			"grammartool: couldn't lex tokens: %v\n", err)
		os.Exit(1)
	}

	tLen := int(math.Ceil(math.Log10(float64(len(g.Terminals)))))
	nLen := int(math.Ceil(math.Log10(float64(len(tokens)))))

	fmt.Printf("%[1]*[3]s %[2]*[4]s %[5]s\n", tLen, nLen, "T", "n",
		"Lexeme")
	fmt.Printf("%[1]*[3]s %[2]*[4]s %[5]s\n", tLen, nLen, "-", "-",
		"------")

	for _, t := range tokens {
		fmt.Printf("%[1]*[3]d %[2]*[4]d '%[5]s'\n",
			tLen, nLen, t.ID, t.Index, t.Value)
	}
}

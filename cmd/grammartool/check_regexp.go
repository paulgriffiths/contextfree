package main

import (
	"fmt"
	"github.com/paulgriffiths/contextfree/grammar"
	"regexp"
)

func checkRegexp(g *grammar.Grammar, verbose bool) {
	for _, t := range g.Terminals {
		_, err := regexp.Compile(t)
		if err != nil {
			fmt.Printf("Regexp `%s` did not compile: %v\n", t, err)
		} else if verbose {
			fmt.Printf("Regexp `%s` compiled successfully.\n", t)
		}
	}
}

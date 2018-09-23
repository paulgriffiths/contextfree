package main

import (
	"fmt"
	"github.com/paulgriffiths/contextfree/grammar"
	"github.com/paulgriffiths/contextfree/types/symbols"
	"sort"
)

func outputAttribs(g *grammar.Grammar) {
	if g.IsLeftRecursive() {
		fmt.Printf("The grammar is left-recursive.\n")
	} else {
		fmt.Printf("The grammar is not left-recursive.\n")
	}

	fmt.Printf("The grammar has %s, %s, and %s\n",
		plural(len(g.NonTerminals), "nonterminal", "nonterminals"),
		plural(len(g.Terminals), "terminal", "terminals"),
		plural(g.NumProductions(), "production", "productions"),
	)
}

func outputTerminalsAndNonTerminals(g *grammar.Grammar) {
	nnt := len(g.NonTerminals)
	nt := len(g.Terminals)

	fmt.Printf("The %s ",
		plural(nnt, "nonterminal is", "nonterminals are"))
	printCommaList(intRange(nnt), g.NonTerminals, "")
	fmt.Printf(".\n")

	fmt.Printf("The %s ", plural(nt, "terminal is", "terminals are"))
	printCommaList(intRange(nt), g.Terminals, "`")
	fmt.Printf(".\n")
}

func outputCycles(g *grammar.Grammar) {
	outputNonTerminalList(g, g.NonTerminalsWithCycles(),
		"has a cycle", "have cycles")
}

func outputEProductions(g *grammar.Grammar) {
	outputNonTerminalList(g, g.NonTerminalsWithEProductions(),
		"has an e-production", "have e-productions")
}

func outputNullable(g *grammar.Grammar) {
	outputNonTerminalList(g, g.NonTerminalsNullable(),
		"is nullable", "are nullable")
}

func outputUseless(g *grammar.Grammar) {
	outputNonTerminalList(g, g.Unreachable(),
		"is unreachable", "are unreachable")
	outputNonTerminalList(g, g.Unproductive(),
		"is unproductive", "are unproductive")
}

func outputNonTerminalList(g *grammar.Grammar, list []int,
	singular, plural string) {
	if len(list) == 0 {
		fmt.Printf("No nonterminals %s.\n", plural)
		return
	}

	printCommaList(list, g.NonTerminals, "")
	if len(list) == 1 {
		fmt.Printf(" %s.\n", singular)
	} else {
		fmt.Printf(" %s.\n", plural)
	}
}

func outputFirst(g *grammar.Grammar) {
	ml := g.MaxNonTerminalNameLength()
	fmt.Printf("First sets:\n")
	for n, nt := range g.NonTerminals {
		f := g.First(symbols.NewNonTerminal(n)).Elements()

		terminals := []string{}
		hasEmpty := false
		for _, terminal := range f {
			if terminal.IsEmpty() {
				hasEmpty = true
			} else {
				t := fmt.Sprintf("`%s`", g.Terminals[terminal.I])
				terminals = append(terminals, t)
			}
		}
		sort.Sort(sort.StringSlice(terminals))
		if hasEmpty {
			terminals = append(terminals, "e")
		}

		fmt.Printf("First(%s)", nt)
		for i := 0; i <= (ml - len(nt) + 1); i++ {
			fmt.Printf(" ")
		}
		fmt.Printf(": { ")
		for n, terminal := range terminals {
			if n != 0 {
				fmt.Printf(", ")
			}
			fmt.Printf("%s", terminal)
		}
		fmt.Printf(" }\n")
	}
}

func outputFollows(g *grammar.Grammar) {
	fmt.Printf("Follow sets:\n")
	ml := g.MaxNonTerminalNameLength()
	for nt := range g.NonTerminals {
		set := g.Follow(nt)
		f := set.Elements()

		terminals := []string{}
		hasEmpty := false
		hasEnd := false
		for _, terminal := range f {
			if terminal.IsEmpty() {
				hasEmpty = true
			} else if terminal.IsInputEnd() {
				hasEnd = true
			} else {
				t := fmt.Sprintf("`%s`", g.Terminals[terminal.I])
				terminals = append(terminals, t)
			}
		}
		sort.Sort(sort.StringSlice(terminals))
		if hasEmpty {
			terminals = append(terminals, "e")
		}
		if hasEnd {
			terminals = append(terminals, "$")
		}

		fmt.Printf("Follow(%s)", g.NonTerminals[nt])
		for i := 0; i <= (ml - len(g.NonTerminals[nt])); i++ {
			fmt.Printf(" ")
		}
		fmt.Printf(": { ")
		for n, terminal := range terminals {
			if n != 0 {
				fmt.Printf(", ")
			}
			fmt.Printf("%s", terminal)
		}
		fmt.Printf(" }\n")
	}
}

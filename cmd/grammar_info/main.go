package main

import (
	"flag"
	"fmt"
	"github.com/paulgriffiths/contextfree/grammar"
	"os"
)

func main() {
	fileName := flag.String("f", "", "grammar file name")
	recog := flag.String("r", "", "show if the grammar recognizes "+
		"the provided string")
	ptree := flag.String("p", "", "show a parse tree for the provided "+
		"string, if the grammar recognizes it")
	showFile := flag.Bool("g", true, "show grammar representation")
	listAttribs := flag.Bool("s", true, "show grammar statistics")
	listTerms := flag.Bool("t", false, "list terminals and nonterminals")
	listFirst := flag.Bool("w", false, "list First and Follow for all "+
		"nonterminals")
	listCycles := flag.Bool("c", false, "list nonterminals with cycles")
	listE := flag.Bool("e", false, "list nonterminals with e-productions")
	listNull := flag.Bool("n", false, "list nonterminals which are nullable")
	listUseless := flag.Bool("u", false, "list useless rules")
	listAll := flag.Bool("a", false, "list all grammar information "+
		"(equivalent to -stcen)")
	flag.Parse()

	if *fileName == "" {
		fmt.Fprintf(os.Stderr, "grammar_info: no filename provided.\n")
		os.Exit(1)
	}

	file, err := os.Open(*fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "couldn't open file %q: %v\n",
			os.Args[1], err)
		os.Exit(1)
	}

	defer file.Close()

	g, err := grammar.FromFile(*fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "couldn't create grammar: %v\n", err)
		os.Exit(1)
	}

	if *showFile {
		outputFile(file)
	}
	if *listAttribs || *listAll {
		outputAttribs(g)
	}
	if *listTerms || *listAll {
		outputTerminalsAndNonTerminals(g)
	}
	if *listCycles || *listAll {
		outputCycles(g)
	}
	if *listE || *listAll {
		outputEProductions(g)
	}
	if *listNull || *listAll {
		outputNullable(g)
	}
	if *listUseless || *listAll {
		outputUseless(g)
	}
	if *listFirst || *listAll {
		outputFirst(g)
		outputFollows(g)
	}
	if *recog != "" {
		if g.IsLeftRecursive() {
			fmt.Printf("Parsing currently only implemented for non-" +
				"left-recursive grammars.\n")
		} else {
			recognizeRdpString(g, *recog)
		}
	}
	if *ptree != "" {
		if g.IsLeftRecursive() {
			fmt.Printf("Parsing currently only implemented for non-" +
				"left-recursive grammars.\n")
		} else {
			parseRdpString(g, *ptree)
		}
	}
}

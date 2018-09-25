package main

import (
	"flag"
	"fmt"
	"github.com/paulgriffiths/contextfree/grammar"
	"os"
)

func main() {

	// General flags.

	fileName := flag.String("f", "", "grammar file name")
	verbose := flag.Bool("v", false, "verbose mode")

	// Recognition and parsing flags.

	input := flag.String("i", "", "input string")
	recog := flag.Bool("r", false, "show if the grammar recognizes "+
		"the provided string")
	ptree := flag.Bool("p", false, "show a parse tree for the provided "+
		"string, if the grammar recognizes it")
	lex := flag.Bool("lex", false, "show lexer output")

	// Analysis flags.

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
	checkReg := flag.Bool("checkRegexp", false, "check validity "+
		"of all terminal regular expressions")

	// Code generation flags.

	genNonTerms := flag.Bool("generateNonTerminals", false,
		"generate constant declarations for the grammar nonterminals "+
			"suitable for inclusion in a Go program")
	genTerms := flag.Bool("generateTerminals", false, "generate "+
		"constant declarations for the grammar terminals suitable "+
		"for inclusion in a Go program")
	genSyms := flag.Bool("generateSymbols", false, "generate "+
		"constant declarations for the terminals and nonterminals "+
		"suitable for inclusion in a Go program")
	genString := flag.Bool("generateGrammarString", false, "generate "+
		"a string containing the grammar text suitable for inclusion "+
		"in a Go program")
	genAll := flag.Bool("generate", false, "equivalent to "+
		"-generateSymbols -generateGrammarString")
	ntPrefix := flag.String("nonTerminalPrefix", "nt", "prefix for "+
		"nonterminal identifiers in generated code")
	tPrefix := flag.String("terminalPrefix", "t", "prefix for "+
		"terminal identifiers in generated code")
	pkg := flag.String("pkg", "main", "package name for generated code")

	flag.Parse()

	if *fileName == "" {
		fmt.Fprintf(os.Stderr, "grammartool: no filename provided.\n")
		os.Exit(1)
	}

	file, err := os.Open(*fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr, "grammartool: couldn't open file %q: %v\n",
			*fileName, err)
		os.Exit(1)
	}

	defer file.Close()

	g, err := grammar.FromFile(*fileName)
	if err != nil {
		fmt.Fprintf(os.Stderr,
			"grammartool: couldn't create grammar: %v\n", err)
		os.Exit(1)
	}

	// Generated code options.

	if *genSyms || *genTerms || *genNonTerms || *genString || *genAll {
		fmt.Printf("package %s\n\n", *pkg)

		if *genNonTerms || *genSyms || *genAll {
			generateNonTerminals(g, *ntPrefix)
		}

		if *genTerms || *genSyms || *genAll {
			if *genNonTerms || *genSyms || *genAll {
				fmt.Printf("\n")
			}
			generateTerminals(g, *tPrefix)
		}
		if *genString || *genAll {
			if *genTerms || *genNonTerms || *genSyms || *genAll {
				fmt.Printf("\n")
			}
			generateGrammarString(g, file)
		}
		return
	}

	// String recognition and parsing options.

	if *recog {
		if *input == "" {
			fmt.Fprintf(os.Stderr, "grammartool: you must provide "+
				"an input string with -i when specifying -r.\n")
			os.Exit(1)
		}

		if g.IsLeftRecursive() {
			fmt.Fprintf(os.Stderr, "grammartool: parsing currently "+
				"only implemented for non-left-recursive grammars.\n")
			os.Exit(1)
		} else {
			recognizeRdpString(g, *input, *verbose)
		}

		return
	}

	if *ptree {
		if *input == "" {
			fmt.Fprintf(os.Stderr, "grammartool: you must provide "+
				"an input string with -i when specifying -p.\n")
			os.Exit(1)
		}

		if g.IsLeftRecursive() {
			fmt.Fprintf(os.Stderr, "grammartool: parsing currently "+
				"only implemented for non-left-recursive grammars.\n")
			os.Exit(1)
		} else {
			parseRdpString(g, *input)
		}

		return
	}

	if *lex {
		if *input == "" {
			fmt.Fprintf(os.Stderr, "grammartool: you must provide "+
				"an input string with -i when specifying -lex.\n")
			os.Exit(1)
		}

		showLexerOutput(g, *input)
		return
	}

	// Grammar analysis options.

	if *checkReg {
		checkRegexp(g, *verbose)
		return
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
}

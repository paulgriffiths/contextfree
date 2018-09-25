package main

import (
	"fmt"
	"github.com/paulgriffiths/contextfree/grammar"
	"math"
	"os"
	"strings"
	"unicode"
)

// generateNonTerminals outputs a Go constant declaration group
// representing the indices of nonterminals in the grammar.
// If the nonterminal prefix string begins with an upper case letter,
// each identifier is accompanied by a Go documentation comment
// on the preceding line. If the nonterminal prefix string begins with
// a lower case letter, a simpler comment appears at the end of
// each line.
func generateNonTerminals(g *grammar.Grammar, ntPrefix string) {
	maxNL := g.MaxNonTerminalNameLength() + len(ntPrefix)
	fieldWidth := len(g.NonTerminals[0]) + len(ntPrefix) + 7
	if maxNL > fieldWidth {
		fieldWidth = maxNL
	}

	fmt.Printf("// Nonterminal identifier constants.\nconst (\n")
	for n, nt := range g.NonTerminals {
		ntID := ntPrefix + strings.Replace(nt, "'", "p", -1)
		endComment := ""
		if unicode.IsUpper(rune(ntID[0])) {
			fmt.Printf("\t// %s represents nonterminal %s\n", ntID, nt)
		} else {
			endComment = fmt.Sprintf(" // Nonterminal %s", nt)
		}
		if n == 0 {
			ntID += " = iota"
		}
		fmt.Printf("\t%-[1]*s%s\n", fieldWidth, ntID, endComment)
	}
	fmt.Printf(")\n")
}

// generateTerminals outputs a Go constant declaration group
// representing the indices of terminals in the grammar.
// If the terminal prefix string begins with an upper case letter,
// each identifier is accompanied by a Go documentation comment
// on the preceding line. If the terminal prefix string begins with
// a lower case letter, a simpler comment appears at the end of
// each line.
func generateTerminals(g *grammar.Grammar, tPrefix string) {
	numDigits := int(math.Ceil(math.Log10(float64(len(g.Terminals)))))
	if numDigits < 1 {
		numDigits = 1
	}
	fieldWidth := numDigits + len(tPrefix) + 7

	fmt.Printf("// Terminal identifier constants.\nconst (\n")
	for n, t := range g.Terminals {
		tID := fmt.Sprintf("%s%d", tPrefix, n)
		endComment := ""
		if unicode.IsUpper(rune(tID[0])) {
			fmt.Printf("\t// %s represents terminal `%s`\n", tID, t)
		} else {
			endComment = fmt.Sprintf(" // Terminal `%s`", t)
		}
		if n == 0 {
			tID += " = iota"
		}
		fmt.Printf("\t%-[1]*s%s\n", fieldWidth, tID, endComment)
	}
	fmt.Printf(")\n")
}

// generateGrammarString outputs the declaration of a Go string
// containing the text of the grammar, preceded by a comment
// showing the original grammar text.
func generateGrammarString(g *grammar.Grammar, file *os.File) {
	gs := fileToString(file)
	fmt.Printf("// grammarString is a string representation " +
		"of the following grammar:\n")

	// Precede each line with single line comment characters
	// and output.

	comm := "// " + strings.Replace(gs, "\n", "\n// ", -1)
	if comm[len(comm)-3:] == "// " {
		comm = comm[:len(comm)-3]
	}
	os.Stdout.Write([]byte(comm))

	// Output the grammar string.

	fmt.Printf("var grammarString = `")
	os.Stdout.Write([]byte(strings.Replace(gs, "`", "` + \"`\" + `", -1)))
	os.Stdout.Write([]byte("`\n"))
}

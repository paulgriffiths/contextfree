package grammar

import (
	"github.com/paulgriffiths/contextfree/types/symbols"
	"github.com/paulgriffiths/contextfree/utils"
	"io"
)

// Grammar represents a context-free grammar.
type Grammar struct {
	NonTerminals []string
	Terminals    []string
	NtTable      map[string]int
	TTable       map[string]int
	Prods        []symbols.StringList
	firsts       []symbols.SetSymbol
	follows      []symbols.SetSymbol
}

// New returns a new context-free grammer from the provided
// description.
func New(reader io.Reader) (*Grammar, error) {
	return parse(reader)
}

// Terminal returns a grammar symbol for the nonterminal
// named in the provided string.
func (g *Grammar) Terminal(nt string) symbols.Symbol {
	return symbols.NewNonTerminal(g.NtTable[nt])
}

// TerminalComp returns a grammar symbol for the terminal
// named in the provided string.
func (g *Grammar) TerminalComp(nt string) symbols.Symbol {
	return symbols.NewTerminal(g.TTable[nt])
}

// NonTerminalsSet returns an integer set containing the elements
// 0...𝑛-1, where 𝑛 is the number of nonterminals.
func (g *Grammar) NonTerminalsSet() utils.SetInt {
	return utils.NewSetInt(utils.IntRange(len(g.NonTerminals))...)
}

// TerminalsSet returns an integer set containing the elements
// 0...𝑛-1, where 𝑛 is the number of terminals.
func (g *Grammar) TerminalsSet() utils.SetInt {
	return utils.NewSetInt(utils.IntRange(len(g.Terminals))...)
}

// Symbols returns a slice containing all the nonterminal and
// terminal symbols in the grammar.
func (g *Grammar) Symbols() []symbols.Symbol {
	list := g.NonTerminalSymbols()
	return append(list, g.TerminalSymbols()...)
}

// NonTerminalSymbols returns a slice containing all the nonterminal
// symbols in the grammar.
func (g *Grammar) NonTerminalSymbols() []symbols.Symbol {
	list := []symbols.Symbol{}
	for n := range g.NonTerminals {
		list = append(list, symbols.NewNonTerminal(n))
	}
	return list
}

// TerminalSymbols returns a slice containing all the terminal
// symbols in the grammar.
func (g *Grammar) TerminalSymbols() []symbols.Symbol {
	list := []symbols.Symbol{}
	for n := range g.Terminals {
		list = append(list, symbols.NewTerminal(n))
	}
	return list
}

// ProductionNumber returns an index for the specified production
// for the specified nonterminal, the index representing the position
// of the production in the list of all productions in the grammar in
// the order they appear in the grammar.
func (g *Grammar) ProductionNumber(nt, prod int) int {
	n := 0
	for i := 0; i < nt; i++ {
		n += len(g.Prods[i])
	}
	return n + prod
}

// ProductionFromNumber takes a production index (such as one
// returned from ProductionNumber) and converts it to a nonterminal
// and production index suitable for indexing Grammar.Prods.
func (g *Grammar) ProductionFromNumber(n int) (int, int) {
	nt := 0
	for n >= len(g.Prods[nt]) {
		n -= len(g.Prods[nt])
		nt++
	}
	return nt, n
}

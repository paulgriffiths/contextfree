package grammar

import (
	"github.com/paulgriffiths/contextfree/datastruct"
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
func (g *Grammar) NonTerminalsSet() datastruct.SetInt {
	return datastruct.NewSetInt(utils.IntRange(len(g.NonTerminals))...)
}

// TerminalsSet returns an integer set containing the elements
// 0...𝑛-1, where 𝑛 is the number of terminals.
func (g *Grammar) TerminalsSet() datastruct.SetInt {
	return datastruct.NewSetInt(utils.IntRange(len(g.Terminals))...)
}

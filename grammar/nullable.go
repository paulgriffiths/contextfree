package grammar

import (
	"github.com/paulgriffiths/contextfree/types/symbols"
)

// NonTerminalsNullable returns a list of nonterminals which
// are nullable.
func (g *Grammar) NonTerminalsNullable() []int {
	list := []int{}
	for n := range g.NonTerminals {
		if g.IsNullable(n) {
			list = append(list, n)
		}
	}
	return list
}

// IsNullable checks if a nonterminal is nullable.
func (g *Grammar) IsNullable(nt int) bool {
	return g.First(symbols.NewNonTerminal(nt)).ContainsEmpty()
}

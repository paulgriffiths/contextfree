package grammar

// BUG: A non-terminal consisting solely of the empty symbol
// should not be classified as unproductive.

import (
	"github.com/paulgriffiths/contextfree/types/symbols"
)

// Unproductive returns a list of unproductive nonterminals.
func (g *Grammar) Unproductive() []int {
	list := []int{}

	// A nonterminal 𝐴 is unproductive if First(𝐴) yields the empty set.

	for i := range g.NonTerminals {
		if g.First(symbols.NewNonTerminal(i)).IsEmpty() {
			list = append(list, i)
		}
	}

	return list
}

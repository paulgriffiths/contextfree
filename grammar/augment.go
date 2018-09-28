package grammar

import "github.com/paulgriffiths/contextfree/types/symbols"

// Augment returns a augmented copy of the grammar. If 𝐺 is a grammar
// with start symbol 𝑆, then the augmented grammar for 𝐺 is 𝐺 with a
// new start symbol 𝑆' and production 𝑆' ➝  𝑆.
func (g *Grammar) Augment() *Grammar {

	// Create augmented nonterminals name list.

	nonTerms := make([]string, len(g.NonTerminals))
	copy(nonTerms, g.NonTerminals)
	name := "S"
	for {
		if _, ok := g.NtTable[name]; ok {
			name += "'"
			continue
		}
		break
	}
	nonTerms = append([]string{name}, nonTerms...)

	// Create augmented nonterminals table.

	ntmap := make(map[string]int)
	for key, value := range g.NtTable {
		ntmap[key] = value + 1
	}
	ntmap[name] = 0

	// Create augmented productions.

	prods := make([]symbols.StringList, len(g.Prods))
	for n := range prods {
		prods[n] = g.Prods[n].Copy()
	}

	for p, prod := range prods {
		for s, str := range prod {
			for x, sym := range str {
				if sym.IsNonTerminal() {
					prods[p][s][x].I++
				}
			}
		}
	}
	newProd := symbols.StringList{
		{symbols.NewNonTerminal(1)},
	}
	prods = append([]symbols.StringList{newProd}, prods...)

	// Build and return new grammar.

	newGrammar := Grammar{
		NonTerminals: nonTerms,
		Terminals:    g.Terminals,
		NtTable:      ntmap,
		TTable:       g.TTable,
		Prods:        prods,
		firsts:       nil,
		follows:      nil,
	}
	return &newGrammar
}

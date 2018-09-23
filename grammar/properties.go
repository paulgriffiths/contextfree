package grammar

// MaxNonTerminalNameLength returns the length of the longest
// nonterminal name.
func (g *Grammar) MaxNonTerminalNameLength() int {
	maxNL := -1
	for _, nt := range g.NonTerminals {
		if len(nt) > maxNL {
			maxNL = len(nt)
		}
	}
	return maxNL
}

// NumProductions returns the number of productions in the grammar.
func (g *Grammar) NumProductions() int {
	n := 0
	for _, str := range g.Prods {
		n += len(str)
	}
	return n
}

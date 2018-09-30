package grammar

// HasEProduction checks if the grammar has an e-production.
func (g *Grammar) HasEProduction() bool {
	return len(g.NonTerminalsWithEProductions()) != 0
}

// NonTerminalsWithEProductions returns a list of nonterminals which
// have e-productions.
func (g *Grammar) NonTerminalsWithEProductions() []int {
	list := []int{}
	for nt, prods := range g.Prods {
		for _, prod := range prods {
			if prod.IsEmpty() {
				list = append(list, nt)
			}
		}
	}
	return list
}

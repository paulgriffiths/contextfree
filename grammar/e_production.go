package grammar

// HasEProduction checks if the grammar has an e-production.
func (g *Grammar) HasEProduction() bool {
	return len(g.NonTerminalsWithEProductions()) != 0
}

// NonTerminalsWithEProductions returns a list of nonterminals which
// have e-productions.
func (g *Grammar) NonTerminalsWithEProductions() []int {
	list := []int{}
	for n, prod := range g.Prods {
		if prod.HasEmpty() {
			list = append(list, n)
		}
	}
	return list
}

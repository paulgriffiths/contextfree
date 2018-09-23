package grammar

// NonTerminalsWithCycles returns a list of nonterminals which
// have cycles.
func (g *Grammar) NonTerminalsWithCycles() []int {
	list := []int{}
	for n := range g.NonTerminals {
		if g.hcInternal(n, n, make(map[int]bool)) {
			list = append(list, n)
		}
	}
	return list
}

// HasCycle checks if the grammar contains a cycle.
func (g *Grammar) HasCycle() bool {
	for n := range g.NonTerminals {
		if g.hcInternal(n, n, make(map[int]bool)) {
			return true
		}
	}
	return false
}

func (g *Grammar) hcInternal(nt, interNt int, checked map[int]bool) bool {
	if checked[interNt] {
		return false
	}
	checked[interNt] = true

	for _, body := range g.Prods[interNt] {
		if body.IsNonTerminal() {
			if body[0].I == nt {
				return true
			} else if g.hcInternal(nt, body[0].I, checked) {
				return true
			}
		}
	}
	return false
}

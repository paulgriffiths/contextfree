package grammar

// NonTerminalsImmediatelyLeftRecursive returns a list of nonterminals
// which are immediately left-recursive.
func (g *Grammar) NonTerminalsImmediatelyLeftRecursive() []int {
	list := []int{}
	for n := range g.NonTerminals {
		for _, str := range g.Prods[n] {
			if !str.IsEmpty() && str[0].IsNonTerminal() && str[0].I == n {
				list = append(list, n)
				break
			}
		}
	}
	return list
}

// NonTerminalsLeftRecursive returns a list of nonterminals which
// are left-recursive.
func (g *Grammar) NonTerminalsLeftRecursive() []int {
	list := []int{}
	for n := range g.NonTerminals {
		if g.lrInternal(n, n, make(map[int]bool)) {
			list = append(list, n)
		}
	}
	return list
}

// IsLeftRecursive checks if the grammar is left-recursive.
func (g *Grammar) IsLeftRecursive() bool {
	for n := range g.NonTerminals {
		if g.lrInternal(n, n, make(map[int]bool)) {
			return true
		}
	}
	return false
}

func (g *Grammar) lrInternal(nt, interNt int, checked map[int]bool) bool {
	if checked[interNt] {
		return false
	}
	checked[interNt] = true

	for _, str := range g.Prods[interNt] {
		if !str.IsEmpty() && str[0].IsNonTerminal() {
			if str[0].I == nt {
				return true
			} else if g.lrInternal(nt, str[0].I, checked) {
				return true
			}
		}
	}
	return false
}

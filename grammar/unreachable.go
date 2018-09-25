package grammar

import (
	"github.com/paulgriffiths/contextfree/utils"
	"sort"
)

// Unreachable returns a list all nonterminals which are not reachable
// from the start symbol.
func (g *Grammar) Unreachable() []int {

	// Begin search from start symbol 𝑆, which is always reachable.

	reachable := utils.NewSetInt(0)
	visitNext := utils.NewSetInt(0)

	for !visitNext.IsEmpty() {

		// Loop through all productions of all nonterminals reached
		// on the last loop iteration, and aggregate any nonterminals
		// reachable via those productions which have not previously
		// been reached.

		reached := utils.NewSetInt()
		for _, nt := range visitNext.Elements() {
			for _, str := range g.Prods[nt] {
				for _, sym := range str {
					if sym.IsNonTerminal() && !reachable.Contains(sym.I) {
						reached.Insert(sym.I)
					}
				}
			}
		}

		// Add any newly-reached nonterminals to the reachable set
		// and assign the set of nonterminals reached on this loop
		// iteration to visitNext. If we didn't reach any nonterminals
		// which had not already been reached, then no additional loop
		// iterations will reach any more. In this case, visitSet
		// will be the empty set and the loop will terminate.

		reachable.Merge(reached)
		visitNext = reached
	}

	// The set of unreachable nonterminals is the set difference
	// between the set of all nonterminals and the set of reachable
	// nonterminals.

	list := g.NonTerminalsSet().Difference(reachable).Elements()
	sort.Sort(sort.IntSlice(list))
	return list
}

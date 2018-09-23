package grammar

import (
	"github.com/paulgriffiths/contextfree/datastruct"
	"github.com/paulgriffiths/contextfree/types/symbols"
)

// First returns the set of terminals that begin strings derived
// from the provided string of grammar symbols.
func (g *Grammar) First(syms ...symbols.Symbol) symbols.SetSymbol {
	if g.firsts == nil {
		g.calcFirsts()
	}

	// First(𝒶𝛽) is simply 𝒶, and 𝜀 obviously has no content so
	// return an empty set. For a single nonterminal, return the
	// precomputed set.

	if syms[0].IsTerminal() {
		return symbols.NewSetSymbol(syms[0])
	}

	if len(syms) == 1 {
		if syms[0].IsEmpty() {
			return symbols.NewSetSymbol()
		} else if syms[0].IsNonTerminal() {
			return g.firsts[syms[0].I]
		}
		panic("unexpected symbol passed to First")
	}

	// For a string of 𝑋1𝑋2...𝑋n, start with the non-𝜀 symbols of
	// 𝑋1. If 𝜀 is in 𝑋1, then repeat with 𝑋2, and so on. If we
	// reach 𝑋n and 𝜀 is in 𝑋n, then 𝜀 is also in First(𝑋1𝑋2...𝑋n).

	set := symbols.NewSetSymbol()
	for _, symbol := range syms {
		f := g.First(symbol)
		set.Merge(f)
		if !f.ContainsEmpty() {
			set.DeleteEmpty()
			break
		}
	}
	return set
}

// calcFirsts calculates the First sets for each nonterminal.
func (g *Grammar) calcFirsts() {
	nullables := g.calcNullables()
	g.firsts = make([]symbols.SetSymbol, len(g.NonTerminals))
	lengths := make([]int, len(g.NonTerminals))
	for i := range g.NonTerminals {
		g.firsts[i] = symbols.NewSetSymbol()
		lengths[i] = -1
	}

	setsChanged := true

	for setsChanged {

		// Complete one First set calculation cycle for each nonterminal.

		for n := range g.NonTerminals {
			symbol := symbols.NewNonTerminal(n)
			f := g.firstInternal(symbol, nullables,
				make(map[symbols.Symbol]bool))
			g.firsts[n].Merge(f)
		}

		// We need to apply the rules until nothing can be added to
		// any First set, which will be the case if we've applied
		// the rules to every production and none of the First sets
		// have changed since we started.

		setsChanged = false
		for i, set := range g.firsts {
			if lengths[i] != set.Length() {
				setsChanged = true
			}
			lengths[i] = set.Length()
		}
	}
}

// firstInternal performs one complete cycle of First set
// computation rules for a given symbol.
func (g *Grammar) firstInternal(sym symbols.Symbol,
	nullables datastruct.SetInt,
	checked map[symbols.Symbol]bool) symbols.SetSymbol {

	set := symbols.NewSetSymbol()

	// First(𝒶) is simply 𝒶, and 𝜀 obviously has no content.

	if sym.IsTerminal() {
		set.Insert(sym)
		return set
	} else if sym.IsEmpty() {
		return set
	} else if !sym.IsNonTerminal() {
		panic("unexpected symbol passed to First")
	}

	if checked[sym] {

		// We already calculated First for this nonterminal,
		// so return the empty set and avoid an infinite loop.

		return set
	}
	checked[sym] = true

	for _, body := range g.Prods[sym.I] {
		if body.IsEmptyString() {
			set.InsertEmpty()
			continue
		}

		for _, sym := range body {
			set.Merge(g.firstInternal(sym, nullables, checked))
			if !(sym.IsNonTerminal() && nullables.Contains(sym.I)) {
				set.DeleteEmpty()
				break
			}
		}
	}

	return set
}

// calcNullables returns the set of nonterminals which can derive 𝜀.
func (g *Grammar) calcNullables() datastruct.SetInt {
	nullable := datastruct.NewSetInt()
	newNulls := datastruct.NewSetInt()

	// Add to set any nonterminal 𝐴 where 𝐴 → 𝜀 is a production.

	for n, prod := range g.Prods {
		if prod.HasEmpty() {
			nullable.Insert(n)
		}
	}

	// Identify any remaining indirectly nullable nonterminals.

	for !nullable.Equals(newNulls) {
		newNulls.Merge(nullable)
		nullable.Merge(newNulls)

		for n, prod := range g.Prods {

			// If this nonterminal is already in the set, don't
			// waste time checking it again.

			if newNulls.Contains(n) {
				continue
			}

			for _, body := range prod {

				// If the production body contains a terminal, it
				// can't be nullable, so continue to the next. We
				// already identified any 𝐴 → 𝜀 productions.

				if !body.HasOnlyNonTerminals() {
					continue
				}

				// The production derives 𝜀 if and only if each
				// nonterminal in the production derives 𝜀. If the
				// production derives 𝜀, the whole nonterminal can
				// derive 𝜀 and there's no need to check further.

				derivesEmpty := true
				for _, sym := range body {
					if !newNulls.Contains(sym.I) {
						derivesEmpty = false
						break
					}
				}

				if derivesEmpty {
					newNulls.Insert(n)
					break
				}
			}
		}
	}

	return nullable
}

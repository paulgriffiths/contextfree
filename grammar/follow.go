package grammar

import (
	"github.com/paulgriffiths/contextfree/types/symbols"
)

// Follow calculates the Follow set for the given nonterminal, where
// the Follow set contains the set of terminals, or the end-of-input
// marker, which can follow that nonterminal.
func (g *Grammar) Follow(n int) symbols.SetSymbol {
	if g.follows == nil {
		g.calcFollows()
	}
	return g.follows[n]
}

// calcFollows calculates the Follow set for each nonterminal.
func (g *Grammar) calcFollows() {
	g.follows = make([]symbols.SetSymbol, len(g.NonTerminals))
	lengths := make([]int, len(g.NonTerminals))
	for i := 0; i < len(g.NonTerminals); i++ {
		g.follows[i] = symbols.NewSetSymbol()
		lengths[i] = -1
	}

	setsChanged := true

	// End of input can always follow the start symbol.

	g.follows[0].InsertEndOfInput()

	for setsChanged {
		for head, prod := range g.Prods {
			for _, str := range prod {
				for i, sym := range str {

					if !sym.IsNonTerminal() {

						// We're only calculating Follow for
						// nonterminals, so skip anything else.

						continue
					}

					if !str.IsLast(i) {

						// If 𝛢→𝛼𝛣𝛽, then everything in First(𝛽)
						// is in Follow(𝛣) except 𝜀, since it's not a
						// terminal.

						first := g.First(str[i+1:]...).Copy()

						if first.ContainsEmpty() {

							// If First(𝛽) derives 𝜀, then 𝛣 can appear
							// at the end of an 𝛢 production, therefore
							// anything that follows 𝛢 can also follow 𝛣.

							g.follows[sym.I].Merge(g.follows[head])

							// 𝜀 itself can't follow 𝛣, since it's not a
							// terminal, so remove it if it's present.

							first.DeleteEmpty()
						}

						g.follows[sym.I].Merge(first)

					} else if str.IsLast(i) {

						// If 𝛢→𝛼𝛣, then 𝛣 can appear at the end of an
						// 𝛢 production, therefore anything that follows
						// 𝛢 can also follow 𝛣.

						g.follows[sym.I].Merge(g.follows[head])
					}
				}
			}
		}

		// We need to apply the rules until nothing can be added to
		// any Follow set, which will be the case if we've applied
		// the rules to every production and none of the Follow sets
		// have changed since we started.

		setsChanged = false
		for i, set := range g.follows {
			if lengths[i] != set.Length() {
				setsChanged = true
			}
			lengths[i] = set.Length()
		}
	}
}

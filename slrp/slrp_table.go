package slrp

import (
	"github.com/paulgriffiths/contextfree/grammar"
)

// Table implements an SLR-parsing table for a context-free grammar.
type Table struct {
	C []SetItem
	A [][][]Action
	G [][]int
}

// NewTable constructs an SLR-parsing table for a grammar.
func NewTable(g *grammar.Grammar) Table {
	a := g.Augment()
	c := canonical(a)
	actions := actionTable(c, a)
	gotos := goToTable(c, a)
	return Table{c, actions, gotos}
}

func actionTable(c []SetItem, g *grammar.Grammar) [][][]Action {
	t := make([][][]Action, len(c))
	for row := range t {
		t[row] = make([][]Action, len(g.Terminals)+1) // +1 for EOI
		for col := range t[row] {
			t[row][col] = []Action{}
		}
	}

	nt := len(g.Terminals)

	for i, set := range c {
		for _, item := range set.Elements() {
			prod := g.Prods[item.Nt][item.Prod]
			dot := item.Dot

			if prod.NotLast(dot) {
				sym := prod[dot]
				if !sym.IsTerminal() {
					continue
				}
				gset := goTo(set, sym, g)
				n := canonicalIndex(c, gset)
				t[i][sym.I] = append(t[i][sym.I], NewShift(n))
			} else if item == NewItem(0, 0, 1) {
				t[i][nt] = append(t[i][nt], NewAccept())
			} else if prod.AfterLast(dot) {
				f := g.Follow(item.Nt)
				for _, s := range f.Elements() {
					pn := g.ProductionNumber(item.Nt, item.Prod)
					if s.IsInputEnd() {
						t[i][nt] = append(t[i][nt], NewReduce(pn))
					} else {
						t[i][s.I] = append(t[i][s.I], NewReduce(pn))
					}
				}
			}
		}
	}

	return t
}

func goToTable(c []SetItem, g *grammar.Grammar) [][]int {
	table := make([][]int, len(c))
	for row := range table {
		table[row] = make([]int, len(g.NonTerminals))
		for col := range table[row] {
			table[row][col] = -1
		}
	}

	for setIndex, set := range c {
		for symIndex, sym := range g.NonTerminalSymbols() {
			if symIndex == 0 {
				continue
			}
			gset := goTo(set, sym, g)
			if !gset.IsEmpty() {
				table[setIndex][symIndex] = canonicalIndex(c, gset)
			}
		}
	}

	return table
}

func canonicalIndex(sets []SetItem, set SetItem) int {
	for n, s := range sets {
		if set.Equals(s) {
			return n
		}
	}
	panic("set not found in canonical set!")
}

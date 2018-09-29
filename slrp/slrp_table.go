package slrp

import (
	"github.com/paulgriffiths/contextfree/grammar"
)

type SlrpTable struct {
	C []SetItem
	G [][]int
}

func NewSlrpTable(g *grammar.Grammar) SlrpTable {
	a := g.Augment()
	c := canonical(a)
	t := goToTable(c, a)
	return SlrpTable{c, t}
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

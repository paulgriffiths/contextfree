package slrp

import (
	"github.com/paulgriffiths/contextfree/grammar"
)

func canonical(g *grammar.Grammar) []SetItem {
	a := g.Augment()
	c := []SetItem{closure(NewSetItem(Item{0, 0, 0}), a)}
	oldLen := -1

	for len(c) != oldLen {
		oldLen = len(c)

		for _, set := range c {
			for _, sym := range a.Symbols() {
				if r := goTo(set, sym, a); !r.IsEmpty() && !contains(c, r) {
					c = append(c, r)
				}
			}
		}
	}

	return c
}

func contains(sets []SetItem, s SetItem) bool {
	for _, set := range sets {
		if set.Equals(s) {
			return true
		}
	}
	return false
}

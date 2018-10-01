package slrp

import (
	"github.com/paulgriffiths/contextfree/grammar"
)

func closure(s SetItem, g *grammar.Grammar) SetItem {
	c := s.Copy()
	next := NewSetItem()

	for !next.Equals(c) {
		c.Merge(next)
		next.Merge(c)

		for _, item := range c.Elements() {
			itemProd := g.Prods[item.Nt][item.Prod]
			if itemProd.AfterLast(item.Dot) {
				continue
			}

			if itemProd[item.Dot].IsNonTerminal() {
				nt := itemProd[item.Dot].I
				for prod := range g.Prods[nt] {
					newItem := Item{nt, prod, 0}
					if !next.Contains(newItem) {
						next.Insert(newItem)
					}
				}
			}
		}
	}

	return c
}

package slrp

import (
	"github.com/paulgriffiths/contextfree/grammar"
	"github.com/paulgriffiths/contextfree/types/symbols"
)

func goTo(i SetItem, sym symbols.Symbol, g *grammar.Grammar) SetItem {
	c := NewSetItem()

	for _, item := range i.Elements() {
		itemProd := g.Prods[item.Nt][item.Prod]
		if itemProd.AfterLast(item.Dot) {
			continue
		}

		if itemProd[item.Dot] == sym {
			startSet := NewSetItem(Item{item.Nt, item.Prod, item.Dot + 1})
			c.Merge(closure(startSet, g))
		}
	}

	return c
}

package slrp

import (
	"github.com/paulgriffiths/contextfree/grammar"
	"github.com/paulgriffiths/contextfree/types/symbols"
	"testing"
)

var goToTestCases = []struct {
	i, g SetItem
	sym  symbols.Symbol
}{
	{
		NewSetItem(
			Item{0, 0, 1}, // 𝑆' ⟶  𝐸⋅
			Item{1, 0, 1}, // 𝐸 ⟶  𝐸 ⋅+ 𝑇
		),
		NewSetItem(
			Item{1, 0, 2}, // 𝐸 ⟶  𝐸 + ⋅𝑇
			Item{2, 0, 0}, // 𝑇 ⟶  ⋅𝑇 * 𝐹
			Item{2, 1, 0}, // 𝑇 ⟶  ⋅𝐹
			Item{3, 0, 0}, // 𝐹 ⟶  ⋅( 𝐸 )
			Item{3, 1, 0}, // 𝐹 ⟶  ⋅𝐷
		),
		symbols.NewTerminal(0), // +
	},
	{
		NewSetItem(
			Item{1, 0, 2}, // 𝐸 ⟶  𝐸 + ⋅𝑇
			Item{2, 0, 0}, // 𝑇 ⟶  ⋅𝑇 * 𝐹
			Item{2, 1, 0}, // 𝑇 ⟶  ⋅𝐹
			Item{3, 0, 0}, // 𝐹 ⟶  ⋅( 𝐸 )
			Item{3, 1, 0}, // 𝐹 ⟶  ⋅𝐷
		),
		NewSetItem(
			Item{1, 0, 3}, // 𝐸 ⟶  𝐸 + 𝑇⋅
			Item{2, 0, 1}, // 𝑇 ⟶  𝑇 ⋅* 𝐹
		),
		symbols.NewNonTerminal(2), // 𝑇
	},
}

func TestGoTo(t *testing.T) {
	g, err := grammar.FromFile(tgArithLr)
	if err != nil {
		t.Errorf("couldn't get grammar: %v", err)
		return
	}

	for n, tc := range goToTestCases {
		if o := goTo(tc.i, tc.sym, g.Augment()); !o.Equals(tc.g) {
			t.Errorf("case %d, got %v, want %v", n+1, o, tc.g)
		}
	}
}

package slrp

import (
	"github.com/paulgriffiths/contextfree/grammar"
	"testing"
)

var closureTestCases = []struct {
	i, c SetItem
}{
	{
		NewSetItem(
			Item{0, 0, 0}, // 𝑆' ⟶  ⋅𝐸
		),
		NewSetItem(
			Item{0, 0, 0}, // 𝑆' ⟶  ⋅𝐸
			Item{1, 0, 0}, // 𝐸 ⟶  ⋅𝐸 + 𝑇
			Item{1, 1, 0}, // 𝐸 ⟶  ⋅𝑇
			Item{2, 0, 0}, // 𝑇 ⟶  ⋅𝑇 * 𝐹
			Item{2, 1, 0}, // 𝑇 ⟶  ⋅𝐹
			Item{3, 0, 0}, // 𝐹 ⟶  ⋅( 𝐸 )
			Item{3, 1, 0}, // 𝐹 ⟶  ⋅𝐷
		),
	},
	{
		NewSetItem(
			Item{0, 0, 1}, // 𝑆' ⟶  𝐸⋅
			Item{1, 0, 1}, // 𝑆' ⟶  𝐸⋅ + 𝑇
		),
		NewSetItem(
			Item{0, 0, 1}, // 𝑆' ⟶  𝐸⋅
			Item{1, 0, 1}, // 𝑆' ⟶  𝐸⋅ + 𝑇
		),
	},
	{
		NewSetItem(
			Item{1, 1, 1}, // 𝐸 ⟶  𝑇⋅
			Item{2, 0, 1}, // 𝐸 ⟶  𝑇⋅ * 𝐹
		),
		NewSetItem(
			Item{1, 1, 1}, // 𝐸 ⟶  𝑇⋅
			Item{2, 0, 1}, // 𝐸 ⟶  𝑇⋅ * 𝐹
		),
	},
	{
		NewSetItem(
			Item{2, 1, 1}, // 𝑇 ⟶  𝐹⋅
		),
		NewSetItem(
			Item{2, 1, 1}, // 𝑇 ⟶  𝐹⋅
		),
	},
	{
		NewSetItem(
			Item{3, 0, 1}, // 𝐹 ⟶  ( ⋅𝐸 )
		),
		NewSetItem(
			Item{3, 0, 1}, // 𝐹 ⟶  ( ⋅𝐸 )
			Item{1, 0, 0}, // 𝐸 ⟶  ⋅𝐸 + 𝑇
			Item{1, 1, 0}, // 𝐸 ⟶  ⋅𝑇
			Item{2, 0, 0}, // 𝑇 ⟶  ⋅𝑇 * 𝐹
			Item{2, 1, 0}, // 𝑇 ⟶  ⋅𝐹
			Item{3, 0, 0}, // 𝐹 ⟶  ⋅( 𝐸 )
			Item{3, 1, 0}, // 𝐹 ⟶  ⋅𝐷
		),
	},
	{
		NewSetItem(
			Item{1, 0, 2}, // 𝐸 ⟶  𝐸 + ⋅𝑇
		),
		NewSetItem(
			Item{1, 0, 2}, // 𝐸 ⟶  𝐸 + ⋅𝑇
			Item{2, 0, 0}, // 𝑇 ⟶  ⋅𝑇 * 𝐹
			Item{2, 1, 0}, // 𝑇 ⟶  ⋅𝐹
			Item{3, 0, 0}, // 𝐹 ⟶  ⋅( 𝐸 )
			Item{3, 1, 0}, // 𝐹 ⟶  ⋅𝐷
		),
	},
	{
		NewSetItem(
			Item{2, 0, 2}, // 𝑇 ⟶  𝑇 * ⋅𝐹
		),
		NewSetItem(
			Item{2, 0, 2}, // 𝑇 ⟶  𝑇 * ⋅𝐹
			Item{3, 0, 0}, // 𝐹 ⟶  ⋅( 𝐸 )
			Item{3, 1, 0}, // 𝐹 ⟶  ⋅𝐷
		),
	},
}

func TestClosure(t *testing.T) {
	g, err := grammar.FromFile(tgArithLr)
	if err != nil {
		t.Errorf("couldn't get grammar: %v", err)
		return
	}

	for n, tc := range closureTestCases {
		if c := closure(tc.i, g.Augment()); !tc.c.Equals(c) {
			t.Errorf("case %d, got %v, want %v", n+1, c, tc.c)
		}
	}
}

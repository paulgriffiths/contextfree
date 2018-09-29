package slrp

import (
	"github.com/paulgriffiths/contextfree/grammar"
	"testing"
)

var canonicalTestCase = []SetItem{
	NewSetItem(
		Item{0, 0, 0},
		Item{1, 0, 0},
		Item{1, 1, 0},
		Item{2, 0, 0},
		Item{2, 1, 0},
		Item{3, 0, 0},
		Item{3, 1, 0},
	),
	NewSetItem(
		Item{0, 0, 1},
		Item{1, 0, 1},
	),
	NewSetItem(
		Item{1, 1, 1},
		Item{2, 0, 1},
	),
	NewSetItem(
		Item{2, 1, 1},
	),
	NewSetItem(
		Item{3, 0, 1},
		Item{1, 0, 0},
		Item{1, 1, 0},
		Item{2, 0, 0},
		Item{2, 1, 0},
		Item{3, 0, 0},
		Item{3, 1, 0},
	),
	NewSetItem(
		Item{3, 1, 1},
	),
	NewSetItem(
		Item{1, 0, 2},
		Item{2, 0, 0},
		Item{2, 1, 0},
		Item{3, 0, 0},
		Item{3, 1, 0},
	),
	NewSetItem(
		Item{2, 0, 2},
		Item{3, 0, 0},
		Item{3, 1, 0},
	),
	NewSetItem(
		Item{1, 0, 1},
		Item{3, 0, 2},
	),
	NewSetItem(
		Item{1, 0, 3},
		Item{2, 0, 1},
	),
	NewSetItem(
		Item{2, 0, 3},
	),
	NewSetItem(
		Item{3, 0, 3},
	),
}

func TestCanonical(t *testing.T) {
	g, err := grammar.FromFile(tgArithLr)
	if err != nil {
		t.Errorf("couldn't get grammar: %v", err)
		return
	}

	c := canonical(g)
	if len(c) != len(canonicalTestCase) {
		t.Errorf("length not same, got %d, want %d", len(c),
			len(canonicalTestCase))
		return
	}

	for n, tc := range canonicalTestCase {
		if !tc.Equals(c[n]) {
			t.Errorf("index %d, got %v, want %v", n, tc, c[n])
		}
	}
}

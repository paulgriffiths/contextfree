package grammar_test

import (
	"github.com/paulgriffiths/contextfree/grammar"
	"github.com/paulgriffiths/contextfree/types/symbols"
	"testing"
)

var tc = grammar.Grammar{
	NonTerminals: []string{"S", "E", "T", "F"},
	Terminals:    []string{"\\+", "\\*", "\\(", "\\)", "[[:digit:]]+"},
	NtTable: map[string]int{
		"S": 0, "E": 1, "T": 2, "F": 3,
	},
	TTable: map[string]int{
		"\\+": 0, "\\*": 1, "\\(": 2, "\\)": 3, "[[:digit:]]+": 4,
	},
	Prods: []symbols.StringList{
		{
			symbols.String{
				symbols.NewNonTerminal(1),
			},
		},
		{
			symbols.String{
				symbols.NewNonTerminal(1),
				symbols.NewTerminal(0),
				symbols.NewNonTerminal(2),
			},
			symbols.String{
				symbols.NewNonTerminal(2),
			},
		},
		{
			symbols.String{
				symbols.NewNonTerminal(2),
				symbols.NewTerminal(1),
				symbols.NewNonTerminal(3),
			},
			symbols.String{
				symbols.NewNonTerminal(3),
			},
		},
		{
			symbols.String{
				symbols.NewTerminal(2),
				symbols.NewNonTerminal(1),
				symbols.NewTerminal(3),
			},
			symbols.String{
				symbols.NewTerminal(4),
			},
		},
	},
}

func TestAugmentNonTerminalNames(t *testing.T) {
	old, err := grammar.FromFile(tgArithLr)
	if err != nil {
		t.Errorf("couldn't create grammar: %v", err)
		return
	}

	aug := old.Augment()

	if len(aug.NonTerminals) != len(tc.NonTerminals) {
		t.Errorf("length doesn't match, got %d, want %d",
			len(aug.NonTerminals), len(tc.NonTerminals))
		return
	}

	for n, name := range aug.NonTerminals {
		if name != tc.NonTerminals[n] {
			t.Errorf("name doesn't match, got %s, want %s", name,
				tc.NonTerminals[n])
		}
	}
}

func TestAugmentTerminalNames(t *testing.T) {
	old, err := grammar.FromFile(tgArithLr)
	if err != nil {
		t.Errorf("couldn't create grammar: %v", err)
		return
	}

	aug := old.Augment()

	if len(aug.Terminals) != len(tc.Terminals) {
		t.Errorf("length doesn't match, got %d, want %d",
			len(aug.Terminals), len(tc.Terminals))
		return
	}

	for n, name := range aug.Terminals {
		if name != tc.Terminals[n] {
			t.Errorf("name doesn't match, got %s, want %s", name,
				tc.Terminals[n])
		}
	}
}

func TestAugmentNonTerminalTable(t *testing.T) {
	old, err := grammar.FromFile(tgArithLr)
	if err != nil {
		t.Errorf("couldn't create grammar: %v", err)
		return
	}

	aug := old.Augment()

	if len(aug.NtTable) != len(tc.NtTable) {
		t.Errorf("length doesn't match, got %d, want %d",
			len(aug.NtTable), len(tc.NtTable))
		return
	}

	for key, value := range aug.NtTable {
		if s, _ := tc.NtTable[key]; s != value {
			t.Errorf("value doesn't match, got %d, want %d", s, value)
		}
	}
}

func TestAugmentTerminalTable(t *testing.T) {
	old, err := grammar.FromFile(tgArithLr)
	if err != nil {
		t.Errorf("couldn't create grammar: %v", err)
		return
	}

	aug := old.Augment()

	if len(aug.TTable) != len(tc.TTable) {
		t.Errorf("length doesn't match, got %d, want %d",
			len(aug.TTable), len(tc.TTable))
		return
	}

	for key, value := range aug.TTable {
		if s, _ := tc.TTable[key]; s != value {
			t.Errorf("value doesn't match, got %d, want %d", s, value)
		}
	}
}

func TestAugmentProductions(t *testing.T) {
	old, err := grammar.FromFile(tgArithLr)
	if err != nil {
		t.Errorf("couldn't create grammar: %v", err)
		return
	}

	aug := old.Augment()

	for p, prod := range aug.Prods {
		for s, str := range prod {
			for x, sym := range str {
				if sym != tc.Prods[p][s][x] {
					t.Errorf("symbols not equal, got %v, want %v",
						sym, tc.Prods[p][s][x])
				}
			}
		}
	}
}

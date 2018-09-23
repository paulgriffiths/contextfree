package grammar_test

import (
	"github.com/paulgriffiths/contextfree/grammar"
	"testing"
)

func TestGrammarCycles(t *testing.T) {
	for _, tc := range grammarTestCases {
		g, err := grammar.FromFile(tc.filename)
		if err != nil {
			t.Errorf("couldn't parse grammar file %q: %v", tc.filename, err)
			continue
		}

		if len(g.NonTerminalsWithCycles()) != len(tc.haveCycles) {
			t.Errorf("case %s, got %d, want %d", tc.filename,
				len(g.NonTerminalsWithCycles()), len(tc.haveCycles))
			continue
		}

		for n, nt := range g.NonTerminalsWithCycles() {
			if r := g.NonTerminals[nt]; r != tc.haveCycles[n] {
				t.Errorf("case %s, number %d, got %s, want %s",
					tc.filename, n+1, r, tc.haveCycles[n])
			}
		}
	}
}

func TestGrammarHasCycles(t *testing.T) {
	for _, tc := range grammarTestCases {
		g, err := grammar.FromFile(tc.filename)
		if err != nil {
			t.Errorf("couldn't parse grammar file %q: %v", tc.filename, err)
			continue
		}

		if r := g.HasCycle(); r != (len(tc.haveCycles) > 0) {
			t.Errorf("case %s, got %t, want %t", tc.filename,
				r, len(tc.haveCycles) > 0)
		}
	}
}

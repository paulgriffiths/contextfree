package grammar_test

import (
	"github.com/paulgriffiths/contextfree/grammar"
	"testing"
)

func TestGrammarEProduction(t *testing.T) {
	for _, tc := range grammarTestCases {
		g, err := grammar.FromFile(tc.filename)
		if err != nil {
			t.Errorf("couldn't parse grammar file %q: %v", tc.filename, err)
			continue
		}

		if len(g.NonTerminalsWithEProductions()) != len(tc.haveEProds) {
			t.Errorf("case %s, got %d, want %d", tc.filename,
				len(g.NonTerminalsWithEProductions()), len(tc.haveEProds))
			continue
		}

		for n, nt := range g.NonTerminalsWithEProductions() {
			if r := g.NonTerminals[nt]; r != tc.haveEProds[n] {
				t.Errorf("case %s, number %d, got %s, want %s",
					tc.filename, n+1, r, tc.haveEProds[n])
			}
		}
	}
}

func TestGrammarHasEProduction(t *testing.T) {
	for _, tc := range grammarTestCases {
		g, err := grammar.FromFile(tc.filename)
		if err != nil {
			t.Errorf("couldn't parse grammar file %q: %v", tc.filename, err)
			continue
		}

		if r := g.HasEProduction(); r != (len(tc.haveEProds) > 0) {
			t.Errorf("case %s, got %t, want %t", tc.filename,
				r, len(tc.haveEProds) > 0)
		}
	}
}

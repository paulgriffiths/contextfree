package grammar_test

import (
	"github.com/paulgriffiths/contextfree/grammar"
	"testing"
)

func TestGrammarUnproductive(t *testing.T) {
	for _, tc := range grammarTestCases {
		g, err := grammar.FromFile(tc.filename)
		if err != nil {
			t.Errorf("couldn't parse grammar file %q: %v", tc.filename, err)
			continue
		}

		if len(g.Unproductive()) != len(tc.unproductive) {
			t.Errorf("case %s, got %d, want %d", tc.filename,
				len(g.Unproductive()), len(tc.unproductive))
			continue
		}

		for n, nt := range g.Unproductive() {
			if r := g.NonTerminals[nt]; r != tc.unproductive[n] {
				t.Errorf("case %s, number %d, got %s, want %s",
					tc.filename, n+1, r, tc.unproductive[n])
			}
		}
	}
}

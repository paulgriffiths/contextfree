package grammar_test

import (
	"github.com/paulgriffiths/contextfree/grammar"
	"testing"
)

func TestGrammarUnreachable(t *testing.T) {
	for _, tc := range grammarTestCases {
		g, err := grammar.FromFile(tc.filename)
		if err != nil {
			t.Errorf("couldn't parse grammar file %q: %v", tc.filename, err)
			continue
		}

		if len(g.Unreachable()) != len(tc.unreachable) {
			t.Errorf("case %s, got %d, want %d", tc.filename,
				len(g.Unreachable()), len(tc.unreachable))
			continue
		}

		for n, nt := range g.Unreachable() {
			if r := g.NonTerminals[nt]; r != tc.unreachable[n] {
				t.Errorf("case %s, number %d, got %s, want %s",
					tc.filename, n+1, r, tc.unreachable[n])
			}
		}
	}
}

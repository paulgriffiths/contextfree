package grammar_test

import (
	"github.com/paulgriffiths/contextfree/grammar"
	"testing"
)

func TestGrammarLeftRecursive(t *testing.T) {
	for _, tc := range grammarTestCases {
		g, err := grammar.FromFile(tc.filename)
		if err != nil {
			t.Errorf("couldn't parse grammar file %q: %v", tc.filename, err)
			continue
		}

		if len(g.NonTerminalsLeftRecursive()) != len(tc.leftRecursive) {
			t.Errorf("case %s, got %d, want %d", tc.filename,
				len(g.NonTerminalsLeftRecursive()), len(tc.leftRecursive))
			continue
		}

		for n, nt := range g.NonTerminalsLeftRecursive() {
			if r := g.NonTerminals[nt]; r != tc.leftRecursive[n] {
				t.Errorf("case %s, number %d, got %s, want %s",
					tc.filename, n+1, r, tc.leftRecursive[n])
			}
		}
	}
}

func TestGrammarImmediatelyLeftRecursive(t *testing.T) {
	for _, tc := range grammarTestCases {
		g, err := grammar.FromFile(tc.filename)
		if err != nil {
			t.Errorf("couldn't parse grammar file %q: %v", tc.filename, err)
			continue
		}

		nilr := len(g.NonTerminalsImmediatelyLeftRecursive())
		if nilr != len(tc.immediatelyLeftRecursive) {
			t.Errorf("case %s, got %d, want %d", tc.filename,
				nilr, len(tc.immediatelyLeftRecursive))
			continue
		}

		for n, nt := range g.NonTerminalsImmediatelyLeftRecursive() {
			if r := g.NonTerminals[nt]; r != tc.immediatelyLeftRecursive[n] {
				t.Errorf("case %s, number %d, got %s, want %s",
					tc.filename, n+1, r, tc.immediatelyLeftRecursive[n])
			}
		}
	}
}

func TestGrammarIsLeftRecursive(t *testing.T) {
	for _, tc := range grammarTestCases {
		g, err := grammar.FromFile(tc.filename)
		if err != nil {
			t.Errorf("couldn't parse grammar file %q: %v", tc.filename, err)
			continue
		}

		if r := g.IsLeftRecursive(); r != tc.isLeftRecursive {
			t.Errorf("case %s, got %t, want %t", tc.filename,
				r, tc.isLeftRecursive)
		}
	}
}

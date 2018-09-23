package grammar_test

import (
	"github.com/paulgriffiths/contextfree/grammar"
	"testing"
)

func TestGrammarParseNumNonTerminals(t *testing.T) {
	for _, tc := range grammarTestCases {
		g, err := grammar.FromFile(tc.filename)
		if err != nil {
			t.Errorf("couldn't parse grammar file %q: %v", tc.filename, err)
			continue
		}

		if len(g.NonTerminals) != tc.numNonTerminals {
			t.Errorf("case %s, got %d, want %d", tc.filename,
				len(g.NonTerminals), tc.numNonTerminals)
		}
	}
}

func TestGrammarParseNumTerminals(t *testing.T) {
	for _, tc := range grammarTestCases {
		g, err := grammar.FromFile(tc.filename)
		if err != nil {
			t.Errorf("couldn't parse grammar file %q: %v", tc.filename, err)
			continue
		}

		if len(g.Terminals) != tc.numTerminals {
			t.Errorf("case %s, got %d, want %d", tc.filename,
				len(g.Terminals), tc.numTerminals)
		}
	}
}

func TestGrammarParseNumProductions(t *testing.T) {
	for _, tc := range grammarTestCases {
		g, err := grammar.FromFile(tc.filename)
		if err != nil {
			t.Errorf("couldn't parse grammar file %q: %v", tc.filename, err)
			continue
		}

		if g.NumProductions() != tc.numProductions {
			t.Errorf("case %s, got %d, want %d", tc.filename,
				g.NumProductions(), tc.numProductions)
		}
	}
}

func TestGrammarParseNonTerminalNames(t *testing.T) {
	for _, tc := range grammarTestCases {
		g, err := grammar.FromFile(tc.filename)
		if err != nil {
			t.Errorf("couldn't parse grammar file %q: %v", tc.filename, err)
			continue
		}

		if len(g.NonTerminals) != len(tc.nonTerminalNames) {
			t.Errorf("case %s, got %d, want %d", tc.filename,
				len(tc.nonTerminalNames), len(g.NonTerminals))
			continue
		}

		for n, ntName := range g.NonTerminals {
			if ntName != tc.nonTerminalNames[n] {
				t.Errorf("case %s, nonterminal %d, got %s, want %s",
					tc.filename, n, tc.nonTerminalNames[n], ntName)
			}
		}
	}
}

func TestGrammarParseTerminalNames(t *testing.T) {
	for _, tc := range grammarTestCases {
		g, err := grammar.FromFile(tc.filename)
		if err != nil {
			t.Errorf("couldn't parse grammar file %q: %v", tc.filename, err)
			continue
		}

		if len(g.Terminals) != len(tc.terminalNames) {
			t.Errorf("case %s, got %d, want %d", tc.filename,
				len(tc.terminalNames), len(g.Terminals))
			continue
		}

		for n, tName := range g.Terminals {
			if tName != tc.terminalNames[n] {
				t.Errorf("case %s, terminal %d, got %s, want %s",
					tc.filename, n, tc.terminalNames[n], tName)
			}
		}
	}
}

func TestGrammarParseNonTerminalTable(t *testing.T) {
	for _, tc := range grammarTestCases {
		g, err := grammar.FromFile(tc.filename)
		if err != nil {
			t.Errorf("couldn't parse grammar file %q: %v", tc.filename, err)
			continue
		}

		for n, ntName := range g.NonTerminals {
			if r := g.NtTable[ntName]; r != n {
				t.Errorf("case %s, nonterminal %s, got %d, want %d",
					tc.filename, ntName, r, n)
			}
		}
	}
}

func TestGrammarParseTerminalTable(t *testing.T) {
	for _, tc := range grammarTestCases {
		g, err := grammar.FromFile(tc.filename)
		if err != nil {
			t.Errorf("couldn't parse grammar file %q: %v", tc.filename, err)
			continue
		}

		for n, tName := range g.Terminals {
			if r := g.TTable[tName]; r != n {
				t.Errorf("case %s, terminal %s, got %d, want %d",
					tc.filename, tName, r, n)
			}
		}
	}
}

package grammar_test

import (
	"github.com/paulgriffiths/contextfree/grammar"
	"github.com/paulgriffiths/contextfree/types/symbols"
	"testing"
)

func TestGrammarFollows(t *testing.T) {
	for n, tc := range grammarTestCases {
		g, err := grammar.FromFile(tc.filename)
		if err != nil {
			t.Errorf("couldn't parse file %q: %v", tc.filename, err)
			continue
		}

		symSet := make([]symbols.SetSymbol, len(g.NonTerminals))
		for i := 0; i < len(g.NonTerminals); i++ {
			symSet[i] = symbols.NewSetSymbol()
		}

		for nonTerm, terminals := range tc.follows {
			n := g.NtTable[nonTerm]
			for _, term := range terminals {
				if term == "$" {
					symSet[n].Insert(symbols.NewSymbolEndOfInput())
					continue
				}
				symSet[n].Insert(g.TerminalComp(term))
			}
		}

		for i := 0; i < len(g.NonTerminals); i++ {
			if !g.Follow(i).Equals(symSet[i]) {
				t.Errorf("case %d, nonterminal %s, got %v, want %v",
					n+1, g.NonTerminals[i], g.Follow(i).Elements(),
					symSet[i].Elements())
			}
		}
	}
}

package grammar_test

import (
	"github.com/paulgriffiths/contextfree/grammar"
	"github.com/paulgriffiths/contextfree/types/symbols"
	"testing"
)

func TestGrammarFirstStrings(t *testing.T) {
	testCases := []struct {
		filename   string
		nt, result []string
	}{
		{
			tgArithNlr,
			[]string{"F", "T"},
			[]string{"\\(", "[[:digit:]]+"},
		},
		{
			tgArithNlr,
			[]string{"T", "E"},
			[]string{"\\(", "[[:digit:]]+"},
		},
		{
			tgArithNlr,
			[]string{"E", "F"},
			[]string{"\\(", "[[:digit:]]+"},
		},
		{
			tgArithNlr,
			[]string{"E'", "T'"},
			[]string{"\\*", "\\+", ""},
		},
		{
			tgArithNlr,
			[]string{"E'", "F"},
			[]string{"\\+", "\\(", "[[:digit:]]+"},
		},
		{
			tgArithNlr,
			[]string{"F", "E'"},
			[]string{"\\(", "[[:digit:]]+"},
		},
		{
			tgArithNlr,
			[]string{"E'", "T'", "F"},
			[]string{"\\*", "\\+", "\\(", "[[:digit:]]+"},
		},
	}

	for n, tc := range testCases {
		g, err := grammar.FromFile(tc.filename)
		if err != nil {
			t.Errorf("couldn't parse file %q: %v", tc.filename, err)
			continue
		}

		components := symbols.String{}
		for _, nt := range tc.nt {
			components = append(components, g.Terminal(nt))
		}
		resultSet := g.First(components...)
		cmpSet := symbols.NewSetSymbol()

		for _, s := range tc.result {
			if s == "" {
				cmpSet.Insert(symbols.NewSymbolEmpty())
				continue
			}
			cmpSet.Insert(g.TerminalComp(s))
		}

		if !resultSet.Equals(cmpSet) {
			t.Errorf("case %d, got %v, want %v", n+1, resultSet, cmpSet)
		}
	}
}

func TestGrammarFirst(t *testing.T) {
	for n, tc := range grammarTestCases {
		g, err := grammar.FromFile(tc.filename)
		if err != nil {
			t.Errorf("couldn't parse file %q: %v", tc.filename, err)
			continue
		}

		cmpSet := make([]symbols.SetSymbol, len(g.NonTerminals))
		for i := 0; i < len(g.NonTerminals); i++ {
			cmpSet[i] = symbols.NewSetSymbol()
		}

		for nonTerm, terminals := range tc.firsts {
			n := g.NtTable[nonTerm]
			for _, term := range terminals {
				if term == "" {
					cmpSet[n].Insert(symbols.NewSymbolEmpty())
					continue
				}
				cmpSet[n].Insert(g.TerminalComp(term))
			}
		}

		for i := 0; i < len(g.NonTerminals); i++ {
			if !g.First(symbols.NewNonTerminal(i)).Equals(cmpSet[i]) {
				t.Errorf("case %d, nonterminal %s, got %v, want %v",
					n+1, g.NonTerminals[i],
					g.First(symbols.NewNonTerminal(i)).Elements(),
					cmpSet[i].Elements())
			}
		}
	}
}

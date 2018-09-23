package lexer

import (
	"github.com/paulgriffiths/contextfree/grammar"
	"testing"
)

func TestLexerMatch(t *testing.T) {
	testCases := []struct {
		filename  string
		input     string
		terminals []string
	}{
		{tgArithLr, "(3+4)*5",
			[]string{"(", "3", "+", "4", ")", "*", "5"}},
		{tgArithLr, "   ( 3  +4)  * 5      ",
			[]string{"(", "3", "+", "4", ")", "*", "5"}},
		{tgArithNlr, "(3+4)*5",
			[]string{"(", "3", "+", "4", ")", "*", "5"}},
		{tgArithNlr, "   ( 3  +4)  * 5      ",
			[]string{"(", "3", "+", "4", ")", "*", "5"}},
		{tgArithAmbig, "(3+4)*5",
			[]string{"(", "3", "+", "4", ")", "*", "5"}},
		{tgArithAmbig, "   ( 3  +4)  * 5      ",
			[]string{"(", "3", "+", "4", ")", "*", "5"}},
		{tgBalParens1, "((()))",
			[]string{"(", "(", "(", ")", ")", ")"}},
		{tgBalParens2, "((()))",
			[]string{"(", "(", "(", ")", ")", ")"}},
		{tgZeroOne, "000111",
			[]string{"0", "0", "01", "1", "1"}},
		{tgIndirectLr1, "abcd",
			[]string{"a", "b", "c", "d"}},
		{tgCycle1, "abaababa",
			[]string{"a", "b", "a", "a", "b", "a", "b", "a"}},
	}

	for n, tc := range testCases {
		g, err := grammar.FromFile(tc.filename)
		if err != nil {
			t.Errorf("couldn't parse file: %v", err)
			continue
		}

		lexer, err := New(g)
		if err != nil {
			t.Errorf("case %d, failed to lex: %v", n+1, err)
			continue
		}

		terminals, err := lexer.Lex(tc.input)
		if err != nil {
			t.Errorf("case %d, failed to lex: %v", n+1, err)
			continue
		}

		if len(terminals) != len(tc.terminals) {
			t.Errorf("case %d, lengths not same, got %d [%v], want %d [%v]",
				n+1, len(terminals), terminals,
				len(tc.terminals), tc.terminals)
			continue
		}

		for m, term := range terminals {
			if term.S != tc.terminals[m] {
				t.Errorf("case %d, %d, strings not same, got %s, want %s",
					n+1, m, term.S, tc.terminals[m])
			}
		}
	}
}

func TestLexerNonMatch(t *testing.T) {
	testCases := []struct {
		filename string
		input    string
	}{
		{tgArithLr, "  frisbee  "},
		{tgArithLr, "  (3+4) * frisbee + 8  "},
		{tgBalParens1, "((()))p"},
		{tgZeroOne, "000a111"},
	}

	for n, tc := range testCases {
		g, perr := grammar.FromFile(tc.filename)
		if perr != nil {
			t.Errorf("couldn't parse file: %v", perr)
			continue
		}

		lexer, err := New(g)
		if err != nil {
			t.Errorf("case %d, failed to lex: %v", n+1, err)
			continue
		}

		if _, lerr := lexer.Lex(tc.input); lerr == nil {
			t.Errorf("case %d, lex unexpectedly succeeded", n+1)
		}
	}
}

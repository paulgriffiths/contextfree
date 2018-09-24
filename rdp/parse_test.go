package rdp_test

import (
	"bytes"
	"github.com/paulgriffiths/contextfree/rdp"
	"testing"
)

func TestParserAccepts(t *testing.T) {
	for _, tc := range parserTestCases {
		parser, err := rdp.FromFile(tc.filename)
		if err != nil {
			t.Errorf("couldn't get parser tree for file %q: %v",
				tc.filename, err)
			continue
		}

		for _, input := range tc.accepts {
			if parser.Parse(input) == nil {
				t.Errorf("grammar %s, string %s, improperly rejected",
					tc.filename, input)
			}
		}
	}
}

func TestParserRejects(t *testing.T) {
	for _, tc := range parserTestCases {
		parser, err := rdp.FromFile(tc.filename)
		if err != nil {
			t.Errorf("couldn't get parser tree for file %q: %v",
				tc.filename, err)
			continue
		}

		for _, input := range tc.rejects {
			if parser.Parse(input) != nil {
				t.Errorf("grammar %s, string %s, improperly accepted",
					tc.filename, input)
			}
		}
	}
}

func TestParseWriteBracketed(t *testing.T) {
	testCases := []struct {
		filename, input string
		opts            []string
		output          string
	}{
		{
			tgArithNlr,
			"(3+4)*5",
			[]string{"", "[", "]"},
			"[E [T [F ( [E [T [F [Digits 3]] [T' e]] [E' + " +
				"[T [F [Digits 4]] [T' e]] [E' e]]] )] [T' * [F " +
				"[Digits 5]] [T' e]]] [E' e]]",
		},
		{
			tgBalParens2,
			"((()))",
			[]string{"", "[", "]"},
			"[S ( [S ( [S ( [S e] ) [S e]] ) [S e]] ) [S e]]",
		},
		{
			tgZeroOne,
			"00001111",
			[]string{},
			"(S 0 (S 0 (S 0 (S 01) 1) 1) 1)",
		},
	}

	for n, tc := range testCases {
		parser, err := rdp.FromFile(tc.filename)
		if err != nil {
			t.Errorf("couldn't get parser tree for file %q: %v",
				tc.filename, err)
			continue
		}

		tree := parser.Parse(tc.input)
		if tree == nil {
			t.Errorf("couldn't get parse tree for file %q", tc.filename)
			continue
		}

		outBuffer := bytes.NewBuffer(nil)
		tree.WriteBracketed(outBuffer, tc.opts...)
		output := string(outBuffer.Bytes())

		if output != tc.output {
			t.Errorf("case %d, got %q, want %q", n+1, output, tc.output)
		}
	}
}

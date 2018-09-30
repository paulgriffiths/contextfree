package slrp_test

import (
	"bytes"
	"github.com/paulgriffiths/contextfree/slrp"
	"testing"
)

func TestParserAccepts(t *testing.T) {
	for _, tc := range parserTestCases {
		parser, err := slrp.FromFile(tc.filename)
		if err != nil {
			t.Errorf("couldn't get parser for file %q: %v",
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
		parser, err := slrp.FromFile(tc.filename)
		if err != nil {
			t.Errorf("couldn't get parser for file %q: %v",
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
			tgArithLr,
			"3+4",
			[]string{"", "[", "]"},
			"[E [E [T [F 3]]] + [T [F 4]]]",
		},
		{
			tgArithLr,
			"(3 + 4) * 5",
			[]string{"", "[", "]"},
			"[E [T [T [F ( [E [E [T [F 3]]] + [T [F 4]]] )]] * [F 5]]]",
		},
	}

	for n, tc := range testCases {
		parser, err := slrp.FromFile(tc.filename)
		if err != nil {
			t.Errorf("couldn't get parser for file %q: %v",
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

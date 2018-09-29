package slrp_test

import (
	"github.com/paulgriffiths/contextfree/slrp"
	"testing"
)

/*
func TestParse(t *testing.T) {
    parser, err := slrp.FromFile(tgArithLr)
    if err != nil {
        t.Errorf("couldn't get parser for file %q: %v",
            tgArithLr, err)
    }

    tree := parser.Parse("(3 + 4) * 5")
    tree = parser.Parse("(3 + 4) *")
    if tree == nil {}
}
*/
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

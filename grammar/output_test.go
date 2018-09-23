package grammar_test

import (
	"bufio"
	"bytes"
	"github.com/paulgriffiths/contextfree/grammar"
	"os"
	"testing"
)

func TestGrammarOutput(t *testing.T) {
	testCases := []struct {
		infile, cmpfile string
	}{
		{tgArithLr, tgOutArithLrRaw},
		{tgArithNlr, tgOutArithNlrRaw},
		{tgArithAmbig, tgOutArithAmbigRaw},
		{tgBalParens1, tgOutBalParens1Raw},
		{tgBalParens2, tgOutBalParens2Raw},
		{tgZeroOne, tgOutZeroOneRaw},
	}

	for _, tc := range testCases {
		g, err := grammar.FromFile(tc.infile)
		if err != nil {
			t.Errorf("couldn't parse file %q: %v", tc.infile, err)
			continue
		}

		outBuffer := bytes.NewBuffer(nil)
		g.Output(outBuffer)
		outScanner := bufio.NewScanner(outBuffer)

		infile, fileErr := os.Open(tc.cmpfile)
		if fileErr != nil {
			t.Errorf("couldn't open file %q: %v", tc.infile, fileErr)
			continue
		}

		cmpScanner := bufio.NewScanner(infile)

		for cmpScanner.Scan() {
			if !outScanner.Scan() {
				t.Errorf("fewer lines than %q", tc.infile)
				break
			}
			if cmpScanner.Text() != outScanner.Text() {
				t.Errorf("%q: got %q, want %q", tc.infile,
					outScanner.Text(), cmpScanner.Text())
			}
		}

		if outScanner.Scan() {
			t.Errorf("more lines than %q", tc.infile)
		}

		infile.Close()
	}
}

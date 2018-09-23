package lexer_test

import (
	"github.com/paulgriffiths/contextfree/grammar/internal/lar"
	"github.com/paulgriffiths/contextfree/grammar/internal/lexer"
	"os"
	"testing"
)

func TestLexerSuccess(t *testing.T) {
	infiles := []string{
		tgArithLr,
		tgArithNlr,
		tgArithAmbig,
		tgBalParens1,
		tgBalParens2,
		tgZeroOne,
	}

	for _, f := range infiles {
		infile, fileErr := os.Open(f)
		if fileErr != nil {
			t.Errorf("couldn't open file %q: %v", f, fileErr)
			continue
		}

		_, err := lexer.Lex(infile)
		if err != nil {
			t.Errorf("couldn't get list of tokens for file %q", f)
		}

		infile.Close()
	}
}

func TestLexerTokenListArithLr(t *testing.T) {
	expected := lexer.TokenList{
		lexer.Token{lexer.TokenNonTerminal, "E", lar.FilePos{0, 8}},
		lexer.Token{lexer.TokenArrow, ":", lar.FilePos{7, 8}},
		lexer.Token{lexer.TokenNonTerminal, "E", lar.FilePos{9, 8}},
		lexer.Token{lexer.TokenTerminal, "\\+", lar.FilePos{11, 8}},
		lexer.Token{lexer.TokenNonTerminal, "T", lar.FilePos{16, 8}},
		lexer.Token{lexer.TokenAlt, "|", lar.FilePos{18, 8}},
		lexer.Token{lexer.TokenNonTerminal, "T", lar.FilePos{20, 8}},
		lexer.Token{lexer.TokenEndOfLine, "", lar.FilePos{21, 8}},
		lexer.Token{lexer.TokenNonTerminal, "T", lar.FilePos{0, 9}},
		lexer.Token{lexer.TokenArrow, ":", lar.FilePos{7, 9}},
		lexer.Token{lexer.TokenNonTerminal, "T", lar.FilePos{9, 9}},
		lexer.Token{lexer.TokenTerminal, "\\*", lar.FilePos{11, 9}},
		lexer.Token{lexer.TokenNonTerminal, "F", lar.FilePos{16, 9}},
		lexer.Token{lexer.TokenAlt, "|", lar.FilePos{18, 9}},
		lexer.Token{lexer.TokenNonTerminal, "F", lar.FilePos{20, 9}},
		lexer.Token{lexer.TokenEndOfLine, "", lar.FilePos{21, 9}},
		lexer.Token{lexer.TokenNonTerminal, "F", lar.FilePos{0, 10}},
		lexer.Token{lexer.TokenArrow, ":", lar.FilePos{7, 10}},
		lexer.Token{lexer.TokenTerminal, "\\(", lar.FilePos{9, 10}},
		lexer.Token{lexer.TokenNonTerminal, "E", lar.FilePos{14, 10}},
		lexer.Token{lexer.TokenTerminal, "\\)", lar.FilePos{16, 10}},
		lexer.Token{lexer.TokenAlt, "|", lar.FilePos{21, 10}},
		lexer.Token{lexer.TokenNonTerminal, "Digits", lar.FilePos{23, 10}},
		lexer.Token{lexer.TokenEndOfLine, "", lar.FilePos{29, 10}},
		lexer.Token{lexer.TokenNonTerminal, "Digits", lar.FilePos{0, 11}},
		lexer.Token{lexer.TokenArrow, ":", lar.FilePos{7, 11}},
		lexer.Token{lexer.TokenTerminal, "[[:digit:]]+",
			lar.FilePos{9, 11}},
		lexer.Token{lexer.TokenEndOfLine, "", lar.FilePos{23, 11}},
	}

	infileName := tgArithLr
	infile, err := os.Open(tgArithLr)
	if err != nil {
		t.Errorf("couldn't open file %q: %v", infileName, err)
		return
	}

	tokens, err := lexer.Lex(infile)
	if err != nil {
		t.Errorf("couldn't get list of tokens")
		return
	}

	if len(tokens) != len(expected) {
		t.Errorf("Got %d tokens, want %d", len(tokens), len(expected))
		return
	}

	for n, token := range expected {
		if token != tokens[n] {
			t.Errorf("token %d, got %v, want %v", n+1, tokens[n], token)
		}
	}
}

func TestLexerErrors(t *testing.T) {
	testCases := []struct {
		filename string
		err      lexer.LexError
	}{
		{
			tgBadUnterminatedTerminal1,
			lexer.LexError{lexer.LexErrUnterminatedTerminal, "", lar.FilePos{17, 3}},
		},
		{
			tgBadUnterminatedTerminal2,
			lexer.LexError{lexer.LexErrUnterminatedTerminal, "", lar.FilePos{19, 5}},
		},
		{
			tgBadIllegalCharacter1,
			lexer.LexError{lexer.LexErrIllegalCharacter, "%", lar.FilePos{16, 3}},
		},
		{
			tgBadIllegalCharacter2,
			lexer.LexError{lexer.LexErrIllegalCharacter, "$", lar.FilePos{1, 3}},
		},
	}

	for n, tc := range testCases {
		infile, fileErr := os.Open(tc.filename)
		if fileErr != nil {
			t.Errorf("couldn't open file %q: %v", tc.filename, fileErr)
			continue
		}

		if _, err := lexer.Lex(infile); err != tc.err {
			t.Errorf("case %d, got %v, want %v", n+1, err, tc.err)
		}

		infile.Close()
	}
}

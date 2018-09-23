package grammar_test

import (
	"github.com/paulgriffiths/contextfree/grammar"
	"github.com/paulgriffiths/contextfree/grammar/internal/lar"
	"testing"
)

func TestParserErrors(t *testing.T) {
	testCases := []struct {
		filename string
		err      grammar.ParseError
	}{
		{tgBadMissingHead1,
			grammar.ParseError{grammar.ParseErrMissingHead, lar.FilePos{0, 4}}},
		{tgBadMissingBody1,
			grammar.ParseError{grammar.ParseErrEmptyBody, lar.FilePos{8, 4}}},
		{tgBadMissingBody2,
			grammar.ParseError{grammar.ParseErrEmptyBody, lar.FilePos{18, 4}}},
		{tgBadMissingBody3,
			grammar.ParseError{grammar.ParseErrEmptyBody, lar.FilePos{8, 4}}},
		{tgBadMissingBody4,
			grammar.ParseError{grammar.ParseErrEmptyBody, lar.FilePos{8, 5}}},
		{tgBadENotAlone1,
			grammar.ParseError{grammar.ParseErrEmptyNotAlone, lar.FilePos{24, 4}}},
		{tgBadENotAlone2,
			grammar.ParseError{grammar.ParseErrEmptyNotAlone, lar.FilePos{26, 4}}},
		{tgBadMissingArrow1,
			grammar.ParseError{grammar.ParseErrMissingArrow, lar.FilePos{1, 4}}},
	}

	for n, tc := range testCases {
		if _, err := grammar.FromFile(tc.filename); err != tc.err {
			t.Errorf("case %d, got %v, want %v", n+1, err, tc.err)
		}
	}
}

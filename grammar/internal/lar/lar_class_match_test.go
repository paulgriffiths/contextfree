package lar

import (
	"strings"
	"testing"
)

func TestSuccessfulClassMatch(t *testing.T) {
	for i, testCase := range classMatchGoodCases {
		lar, err := NewLookaheadReader(strings.NewReader(testCase.input))
		if err != nil {
			t.Errorf("couldn't create lookahead reader: %v", err)
			continue
		}

		for j, match := range testCase.matches {
			if !match.matchFunc(&lar) {
				t.Errorf("case %d, %d, matching method failed", i, j)
			}
		}
	}
}

func TestSuccessfulClassMatchResultValue(t *testing.T) {
	for i, testCase := range classMatchGoodCases {
		lar, err := NewLookaheadReader(strings.NewReader(testCase.input))
		if err != nil {
			t.Errorf("couldn't create lookahead reader: %v", err)
			continue
		}

		for j, match := range testCase.matches {
			match.matchFunc(&lar)
			if result := string(lar.Result.Value); result != match.result {
				t.Errorf("case %d, %d, got %s, want %s", i, j,
					result, match.result)
			}
		}
	}
}

func TestSuccessfulClassMatchResultPosition(t *testing.T) {
	for i, testCase := range classMatchGoodCases {
		lar, err := NewLookaheadReader(strings.NewReader(testCase.input))
		if err != nil {
			t.Errorf("couldn't create lookahead reader: %v", err)
			continue
		}

		for j, match := range testCase.matches {
			match.matchFunc(&lar)
			if result := lar.Result.Pos; result != match.pos {
				t.Errorf("case %d, %d, got %v, want %v", i, j,
					result, match.pos)
			}
		}
	}
}

func TestSuccessfulClassMatchEndOfInputReached(t *testing.T) {
	for i, testCase := range classMatchGoodCases {
		lar, err := NewLookaheadReader(strings.NewReader(testCase.input))
		if err != nil {
			t.Errorf("couldn't create lookahead reader: %v", err)
			continue
		}

		for _, match := range testCase.matches {
			match.matchFunc(&lar)
		}

		if result := lar.EndOfInput(); !result {
			t.Errorf("case %d, end of input not found when expected", i)
		}
	}
}

func TestUnsuccessfulClassMatch(t *testing.T) {
	for i, testCase := range classMatchBadCases {
		lar, err := NewLookaheadReader(strings.NewReader(testCase.input))
		if err != nil {
			t.Errorf("couldn't create lookahead reader: %v", err)
			continue
		}

		for j, match := range testCase.terms {
			if match(&lar) {
				t.Errorf("case %d, %d, matching method succeeded", i, j)
			}
		}
	}
}

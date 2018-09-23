package lar

import (
	"strings"
	"testing"
)

func TestSuccessfulRuneMatch(t *testing.T) {
	for i, testCase := range runeMatchGoodCases {
		lar, err := NewLookaheadReader(strings.NewReader(testCase.input))
		if err != nil {
			t.Errorf("couldn't create lookahead reader: %v", err)
			continue
		}

		for j, match := range testCase.matches {
			if !lar.MatchOneOf(match.args...) {
				t.Errorf("case %d, %d, matching method failed", i, j)
			}
		}
	}
}

func TestSuccessfulRuneMatchResultValue(t *testing.T) {
	for i, testCase := range runeMatchGoodCases {
		lar, err := NewLookaheadReader(strings.NewReader(testCase.input))
		if err != nil {
			t.Errorf("couldn't create lookahead reader: %v", err)
			continue
		}

		for j, match := range testCase.matches {
			lar.MatchOneOf(match.args...)
			if result := string(lar.Result.Value); result != match.result {
				t.Errorf("case %d, %d, got %s, want %s", i, j,
					result, match.result)
			}
		}
	}
}

func TestSuccessfulRuneMatchResultPosition(t *testing.T) {
	for i, testCase := range runeMatchGoodCases {
		lar, err := NewLookaheadReader(strings.NewReader(testCase.input))
		if err != nil {
			t.Errorf("couldn't create lookahead reader: %v", err)
			continue
		}

		for j, match := range testCase.matches {
			lar.MatchOneOf(match.args...)
			if result := lar.Result.Pos; result != match.pos {
				t.Errorf("case %d, %d, got %v, want %v", i, j,
					result, match.pos)
			}
		}
	}
}

func TestSuccessfulRuneMatchEndOfInputReached(t *testing.T) {
	for i, testCase := range runeMatchGoodCases {
		lar, err := NewLookaheadReader(strings.NewReader(testCase.input))
		if err != nil {
			t.Errorf("couldn't create lookahead reader: %v", err)
			continue
		}

		for _, match := range testCase.matches {
			lar.MatchOneOf(match.args...)
		}

		if result := lar.EndOfInput(); !result {
			t.Errorf("case %d, end of input not found when expected", i)
		}
	}
}

func TestUnsuccessfulRuneMatch(t *testing.T) {
	for i, testCase := range runeMatchBadCases {
		lar, err := NewLookaheadReader(strings.NewReader(testCase.input))
		if err != nil {
			t.Errorf("couldn't create lookahead reader: %v", err)
			continue
		}

		if lar.MatchOneOf(testCase.args...) {
			t.Errorf("case %d, matching method succeeded", i)
		}
	}
}

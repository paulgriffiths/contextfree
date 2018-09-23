package lar

import (
	"strings"
	"testing"
)

var nextTestInput = "ab\n123\n?!*:\n"
var nextTestCases = []struct {
	value string
	pos   FilePos
}{
	{"a", FilePos{0, 1}},
	{"b", FilePos{1, 1}},
	{"\n", FilePos{2, 1}},
	{"1", FilePos{0, 2}},
	{"2", FilePos{1, 2}},
	{"3", FilePos{2, 2}},
	{"\n", FilePos{3, 2}},
	{"?", FilePos{0, 3}},
	{"!", FilePos{1, 3}},
	{"*", FilePos{2, 3}},
	{":", FilePos{3, 3}},
	{"\n", FilePos{4, 3}},
}

func TestNextValue(t *testing.T) {
	lar, err := NewLookaheadReader(strings.NewReader(nextTestInput))
	if err != nil {
		t.Errorf("couldn't create lookahead reader: %v", err)
		return
	}

	for n, testCase := range nextTestCases {
		result, err := lar.Next()
		if err != nil {
			t.Errorf("index %d, got %v, want %v", n, err, nil)
			continue
		}

		if result := string(result); result != testCase.value {
			t.Errorf("got %s, want %s", result, testCase.value)
		}
	}
}

func TestNextPos(t *testing.T) {
	lar, err := NewLookaheadReader(strings.NewReader(nextTestInput))
	if err != nil {
		t.Errorf("couldn't create lookahead reader: %v", err)
		return
	}

	for n, r := range nextTestCases {
		if _, err := lar.Next(); err != nil {
			t.Errorf("case %d, couldn't get next byte: %v", n, err)
			continue
		}

		if pos := lar.pos; pos != r.pos {
			t.Errorf("case %d, unexpected line, got %v, want %v",
				n, pos, r.pos)
		}
	}
}

func TestNextEndOfInput(t *testing.T) {
	lar, err := NewLookaheadReader(strings.NewReader(nextTestInput))
	if err != nil {
		t.Errorf("couldn't create lookahead reader: %v", err)
		return
	}

	for n := range nextTestCases {
		if _, err := lar.Next(); err != nil {
			t.Errorf("case %d, couldn't get next byte: %v", n, err)
		}
	}

	if !lar.EndOfInput() {
		t.Errorf("end of input not found when expected")
	}
}

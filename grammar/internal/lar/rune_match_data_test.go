package lar

type runeMatch struct {
	result string
	args   []rune
	pos    FilePos
}

var runeMatchGoodCases = []struct {
	input   string
	matches []runeMatch
}{
	{"!", []runeMatch{
		{"!", []rune{'!'}, FilePos{0, 1}},
	}},
	{"?", []runeMatch{
		{"?", []rune{'?'}, FilePos{0, 1}},
	}},
	{"#", []runeMatch{
		{"#", []rune{'#'}, FilePos{0, 1}},
	}},
	{"!", []runeMatch{
		{"!", []rune{'!', '?', '#'}, FilePos{0, 1}},
	}},
	{"?", []runeMatch{
		{"?", []rune{'!', '?', '#'}, FilePos{0, 1}},
	}},
	{"#", []runeMatch{
		{"#", []rune{'!', '?', '#'}, FilePos{0, 1}},
	}},
	{"!?#", []runeMatch{
		{"!", []rune{'!'}, FilePos{0, 1}},
		{"?", []rune{'?'}, FilePos{1, 1}},
		{"#", []rune{'#'}, FilePos{2, 1}},
	}},
	{"!?#", []runeMatch{
		{"!", []rune{'!', '?', '#', '$', '%'}, FilePos{0, 1}},
		{"?", []rune{'!', '?', '#', '$', '%'}, FilePos{1, 1}},
		{"#", []rune{'!', '?', '#', '$', '%'}, FilePos{2, 1}},
	}},
	{"!?#\n%^&\n@()", []runeMatch{
		{"!", []rune{'!'}, FilePos{0, 1}},
		{"?", []rune{'?'}, FilePos{1, 1}},
		{"#", []rune{'#'}, FilePos{2, 1}},
		{"\n", []rune{'\n'}, FilePos{3, 1}},
		{"%", []rune{'%'}, FilePos{0, 2}},
		{"^", []rune{'^'}, FilePos{1, 2}},
		{"&", []rune{'&'}, FilePos{2, 2}},
		{"\n", []rune{'\n'}, FilePos{3, 2}},
		{"@", []rune{'@'}, FilePos{0, 3}},
		{"(", []rune{'('}, FilePos{1, 3}},
		{")", []rune{')'}, FilePos{2, 3}},
	}},
	{"!?#\n%^&\n@()", []runeMatch{
		{"!", []rune{'!', '?', '#', '\n', '%', '^', '&'}, FilePos{0, 1}},
		{"?", []rune{'!', '?', '#', '\n', '%', '^', '&'}, FilePos{1, 1}},
		{"#", []rune{'!', '?', '#', '\n', '%', '^', '&'}, FilePos{2, 1}},
		{"\n", []rune{'!', '?', '#', '\n', '%', '^', '&'}, FilePos{3, 1}},
		{"%", []rune{'\n', '%', '^', '&'}, FilePos{0, 2}},
		{"^", []rune{'\n', '%', '^', '&'}, FilePos{1, 2}},
		{"&", []rune{'\n', '%', '^', '&'}, FilePos{2, 2}},
		{"\n", []rune{'\n', '%', '^', '&'}, FilePos{3, 2}},
		{"@", []rune{'\n', '@', '(', ')'}, FilePos{0, 3}},
		{"(", []rune{'\n', '@', '(', ')'}, FilePos{1, 3}},
		{")", []rune{'\n', '@', '(', ')'}, FilePos{2, 3}},
	}},
}

var runeMatchBadCases = []struct {
	input string
	args  []rune
}{
	{"!", []rune{'@', '#', '$', '%', '^', '&', '\n'}},
}

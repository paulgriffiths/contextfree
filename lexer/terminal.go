package lexer

// Terminal represents a terminal found by the lexer.
type Terminal struct {
	N int    // ID number in grammar
	S string // Actual string found in input
}

package lexer

import (
	"fmt"
	"github.com/paulgriffiths/contextfree/grammar/internal/lar"
)

// TokenType represents the type of a context-free grammar lexer token.
type TokenType int

const (
	// TokenTerminal represents a terminal grammar symbol.
	TokenTerminal TokenType = iota
	// TokenNonTerminal represents a nonterminal grammar symbol.
	TokenNonTerminal
	// TokenAlt represents an alternative production separator.
	TokenAlt
	// TokenArrow represents a production → symbol separating a
	// production head from its body or bodies.
	TokenArrow
	// TokenEndOfLine represents an end-of-line marker.
	TokenEndOfLine
	// TokenEmpty represents an empty token.
	TokenEmpty
)

// Token represents a context-free grammar lexer token.
type Token struct {
	T   TokenType
	S   string
	Pos lar.FilePos
}

// TokenList represents a list of context-free grammar lexer tokens.
type TokenList []Token

// typeNames associates token type values with descriptive strings.
var typeNames = []string{
	TokenTerminal:    "Terminal",
	TokenNonTerminal: "Non-terminal",
	TokenAlt:         "Alternative",
	TokenArrow:       "Arrow",
	TokenEndOfLine:   "End-of-line",
	TokenEmpty:       "Empty",
}

// String returns a string representation of a context-free grammar
// lexer token.
func (t Token) String() string {
	return fmt.Sprintf("%s: %q (line %d, ch %d)",
		typeNames[t.T], t.S, t.Pos.Line, t.Pos.Ch)
}

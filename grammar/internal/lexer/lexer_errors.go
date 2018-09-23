package lexer

import (
	"fmt"
	"github.com/paulgriffiths/contextfree/grammar/internal/lar"
)

// LexErrorType represents the type of a lexer error.
type LexErrorType int

// lexer error type values.
const (
	LexErrIllegalCharacter LexErrorType = iota
	LexErrUnterminatedTerminal
	LexErrBadInput
	LexErrUnknownErr
)

// LexErr is an interface for lexer errors. It is provided so that
// lexer functions can return nil errors.
type LexErr interface {
	error
	implementsLexError()
}

// LexError is a concrete lexer error type.
type LexError struct {
	T   LexErrorType
	I   string
	Pos lar.FilePos
}

// lexErrorNames associate lexer error type values with descriptive
// strings.
var lexErrorNames = []string{
	LexErrIllegalCharacter:     "illegal character",
	LexErrUnterminatedTerminal: "unterminated terminal",
	LexErrBadInput:             "bad input",
	LexErrUnknownErr:           "unknown error",
}

// implementsLexError is a dummy method to satisfy the interface.
func (e LexError) implementsLexError() {}

// Error returns a string representation of a lexer error.
func (e LexError) Error() string {
	return fmt.Sprintf("%s at line %d, char %d",
		lexErrorNames[e.T], e.Pos.Line, e.Pos.Ch)
}

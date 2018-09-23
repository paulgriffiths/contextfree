package lexer

import "fmt"

// lexErrorType represents the type of a lexer error.
type lexErrorType int

// lexer error type values.
const (
	lexErrMatchFailed lexErrorType = iota
	lexErrEndOfInput
	lexErrBadRegexp
)

// lexErr is an interface for lexer errors. It is provided so that
// lexer functions can return nil errors.
type lexErr interface {
	error
	implementsLexError()
}

// lexError is a concrete lexer error type.
type lexError struct {
	t lexErrorType
}

// lexErrorNames associate lexer error type values with descriptive
// strings.
var lexErrorNames = []string{
	lexErrMatchFailed: "failed to match terminal",
	lexErrEndOfInput:  "end of input",
	lexErrBadRegexp:   "could not compile regular expression",
}

// implementsLexError is a dummy method to satisfy the interface.
func (e lexError) implementsLexError() {}

// Error returns a string representation of a lexer error.
func (e lexError) Error() string {
	return fmt.Sprintf("%s", lexErrorNames[e.t])
}

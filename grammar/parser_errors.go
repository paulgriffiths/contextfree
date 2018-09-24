package grammar

import (
	"fmt"
)

// ParseErrorType represents the type of a parser error.
type ParseErrorType int

// parser error type values.
const (
	ParseErrMissingArrow ParseErrorType = iota
	ParseErrEmptyBody
	ParseErrEmptyNotAlone
	ParseErrMissingNonTerminal
	ParseErrMissingHead
)

// ParseErr is an interface for parser errors. It is provided so that
// parser functions can return nil errors.
type ParseErr interface {
	error
	implementsParseError()
}

// ParseError is a concrete parser error type.
type ParseError struct {
	T   ParseErrorType
	Pos int
}

// parseErrorNames associate parser error type values with descriptive
// strings.
var parseErrorNames = []string{
	ParseErrMissingArrow:       "missing arrow",
	ParseErrEmptyBody:          "empty body",
	ParseErrEmptyNotAlone:      "empty body not alone",
	ParseErrMissingNonTerminal: "missing nonterminal",
	ParseErrMissingHead:        "missing head",
}

// implementsParseError is a dummy method to satisfy the interface.
func (e ParseError) implementsParseError() {}

// Error returns a string representation of a parser error.
func (e ParseError) Error() string {
	return fmt.Sprintf("%s at position %d",
		parseErrorNames[e.T], e.Pos)
}

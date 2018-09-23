package lexer

import (
	"fmt"
	"github.com/paulgriffiths/contextfree/grammar"
	"regexp"
	"sort"
	"unicode"
)

// Lexer implements a lexer for a context-free grammar.
type Lexer struct {
	terminals TerminalList   // List of terminals to search for
	regexps   *regexp.Regexp // Compiled regular expression
	input     string         // The input string
	n         int            // Current position in the input
}

// New creates and returns a new lexer from the provided grammar.
func New(g *grammar.Grammar) (*Lexer, error) {

	// Retrieve terminals from grammar and sort in reverse order.
	// Sorting in reverse order is done in an attempt to avoid
	// conflicts and match longest strings. It's not guaranteed
	// to avoid problems with amgibuous terminals, and differs
	// from the usual Lex approach of matching the earliest-defined
	// pattern first in the event of multiple matches, but this is
	// not a general-purpose library and it's likely sufficient
	// for our current purposes.

	terminals := TerminalList{}
	for n, t := range g.Terminals {
		terminals = append(terminals, Terminal{n, t})
	}
	sort.Sort(sort.Reverse(terminals))

	// Build a single regular expression of the form
	// (^term1)|(^term2)|...|(^termn). Go's regular expression package
	// will prefer leftmost submatches, and this is why we sorted them
	// in reverse order, so that the longest matches which we ought to
	// check first are at the left. The parenthesized expressions will
	// enable us to identify which terminal was matched.

	regexpString := ""
	for i, t := range terminals {
		if i != 0 {
			regexpString += "|"
		}
		regexpString += fmt.Sprintf("(^%s)", t.S)
	}

	r, err := regexp.Compile(regexpString)
	if err != nil {
		return nil, lexError{lexErrBadRegexp}
	}

	// Build lexer and return it.

	l := Lexer{terminals, r, "", 0}
	return &l, nil
}

// Lex returns a list of terminals of the provided grammar extracted
// from the provided input.
func (l Lexer) Lex(input string) (TerminalList, error) {
	list := TerminalList{}
	l.input = input
	l.n = 0

	for {
		terminal, err := l.getNextTerminal()
		if err != nil {
			if err.(lexError).t == lexErrEndOfInput {
				break
			}
			return nil, err
		}
		list = append(list, terminal)
	}

	return list, nil
}

// getNextTerminal attempts to match the input from the current
// input index against one of the terminal regular expressions.
func (l *Lexer) getNextTerminal() (Terminal, lexErr) {

	// Skip leading whitespace and test for end of input.

	l.skipWhitespace()
	if l.endOfInput() {
		return Terminal{}, lexError{lexErrEndOfInput}
	}

	// Test the input at the current position for a match with
	// any of our terminals, returning an error if there's no match.

	result := l.regexps.FindAllStringSubmatchIndex(l.input[l.n:], 1)
	if len(result) == 0 {
		return Terminal{}, lexError{lexErrMatchFailed}
	}
	matches := result[0]
	if matches[0] == -1 {
		return Terminal{}, lexError{lexErrMatchFailed}
	}

	// Find out which subexpression matched, and build and
	// return a terminal accordingly, and advance the input index.

	for i, t := range l.terminals {
		beg, end := matches[2*(i+1)], matches[2*(i+1)+1]
		if beg == -1 {
			continue
		}
		term := Terminal{t.N, l.input[l.n : l.n+end-beg]}
		l.n += end - beg
		return term, nil
	}

	// If we get here we found a match, but then didn't find it.

	panic("failed to find regexp match index")
}

// skipWhitespace advances the current input index past any
// whitespace characters.
func (l *Lexer) skipWhitespace() {
	for l.n < len(l.input) && unicode.IsSpace(rune(l.input[l.n])) {
		l.n++
	}
}

// endOfInput checks if the input index is at the end of the input.
func (l *Lexer) endOfInput() bool {
	return l.n >= len(l.input)
}

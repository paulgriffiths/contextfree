package lar

import (
	"bufio"
	"fmt"
	"io"
	"unicode"
)

// LookaheadReader implements a single character lookahead reader.
type LookaheadReader struct {
	reader    *bufio.Reader
	lookahead rune
	current   rune
	pos       FilePos
	Result    ReaderResult
}

// NewLookaheadReader returns a single character lookahead reader from
// an io.Reader
func NewLookaheadReader(reader io.Reader) (LookaheadReader, error) {
	r := LookaheadReader{bufio.NewReader(reader), 0, 0,
		FilePos{-1, 1}, ReaderResult{[]rune{}, FilePos{0, 0}}}
	peek, _, err := r.reader.ReadRune()
	if err != nil && err != io.EOF {
		return r, fmt.Errorf("couldn't create lookahead reader: %v", err)
	}
	r.lookahead = peek
	return r, nil
}

// Next returns the next character from a lookahead reader.
// If there are no more characters, the function returns 0 and io.EOF.
// On any other error, the function returns 0 and that error.
func (r *LookaheadReader) Next() (rune, error) {
	if r.lookahead == 0 {
		return 0, io.EOF
	}

	if r.current == '\n' {
		r.pos.incLine()
	} else {
		r.pos.inc()
	}

	r.current = r.lookahead

	if lookahead, _, err := r.reader.ReadRune(); err != nil {
		r.lookahead = 0
		if err != io.EOF {
			return 0, err
		}
	} else {
		r.lookahead = lookahead
	}

	return r.current, nil
}

// MatchOneOf returns true if the next character to be read is among
// the characters passed to the function and stores that character in
// the result, otherwise it returns false and clears the result.
func (r *LookaheadReader) MatchOneOf(vals ...rune) bool {
	r.Result.clear()
	for _, b := range vals {
		if r.lookahead == b {
			r.Next()
			r.Result.setPos(r.pos)
			r.Result.appendRune(b)
			return true
		}
	}
	return false
}

// MatchAnyExcept returns true if the next character to be read is not
// among the characters passed to the function and stores that character
// in the result, otherwise it returns false and clears the result.
func (r *LookaheadReader) MatchAnyExcept(vals ...rune) bool {

	// TODO: Write tests for this function

	r.Result.clear()
	if r.EndOfInput() {
		return false
	}
	for _, b := range vals {
		if r.lookahead == b {
			return false
		}
	}
	r.Result.appendRune(r.lookahead)
	r.Next()
	r.Result.setPos(r.pos)
	return true
}

// MatchIdentifier returns true if the next character to be read is among
// the characters passed to the function and stores that character in
// the result, otherwise it returns false and clears the result.
func (r *LookaheadReader) MatchIdentifier() bool {
	r.Result.clear()
	found := false
	for isIdentifierCharacter(r.lookahead) {
		if !found && unicode.IsDigit(r.lookahead) {
			return false
		}
		current, _ := r.Next()
		if !found {
			r.Result.setPos(r.pos)
			found = true
		}
		r.Result.appendRune(current)
	}
	return found
}

// isIdentifierCharacter returns true if the specifier rune is
// a letter, or a digit or an underscore character, i.e. if it is
// a legal character for a C identifier (excluding the first character,
// which may not be a digit)
func isIdentifierCharacter(r rune) bool {
	return unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_'
}

// MatchNewline returns true if the next character to be read is
// a newline character and stores that character in the result,
// otherwise it returns false and clears the result
func (r *LookaheadReader) MatchNewline() bool {
	return r.MatchOneOf('\n')
}

// MatchLetter returns true if the next character to be read is a letter
// and stores that character in the result, otherwise it returns false
// and clears the result.
func (r *LookaheadReader) MatchLetter() bool {
	return r.matchSingleIsFunc(unicode.IsLetter)
}

// MatchSpace returns true if the next character to be read is whitespace
// and stores that character in the result, otherwise it returns false
// and clears the result.
func (r *LookaheadReader) MatchSpace() bool {
	return r.matchSingleIsFunc(unicode.IsSpace)
}

// MatchDigit returns true if the next character to be read is a digit
// and stores that character in the result, otherwise it returns false
// and clears the result.
func (r *LookaheadReader) MatchDigit() bool {
	return r.matchSingleIsFunc(unicode.IsDigit)
}

// MatchLetters returns true if the next character to be read is a letter
// and stores that and all immediately following letter characters in
// the result, otherwise it returns false and clears the result.
func (r *LookaheadReader) MatchLetters() bool {
	return r.matchMultipleIsFunc(unicode.IsLetter)
}

// MatchSpaces returns true if the next character to be read is whitespace
// and stores that and all immediately following whitespace characters in
// the result, otherwise it returns false and clears the result.
func (r *LookaheadReader) MatchSpaces() bool {
	return r.matchMultipleIsFunc(unicode.IsSpace)
}

// MatchDigits returns true if the next character to be read is a digit
// and stores that and all immediately following digit characters in
// the result, otherwise it returns false and clears the result.
func (r *LookaheadReader) MatchDigits() bool {
	return r.matchMultipleIsFunc(unicode.IsDigit)
}

// EndOfInput returns true if end of input has been reached,
// otherwise it returns false.
func (r LookaheadReader) EndOfInput() bool {
	return r.lookahead == 0
}

// matchSingleIsFunc packages up the matching logic for MatchLetter,
// MatchDigit, and MatchSpace, which otherwise differ only in the
// function used to test the rune.
func (r *LookaheadReader) matchSingleIsFunc(isFunc func(rune) bool) bool {
	r.Result.clear()
	if r.lookahead != '\n' && isFunc(r.lookahead) {
		current, _ := r.Next()
		r.Result.setPos(r.pos)
		r.Result.appendRune(current)
		return true
	}
	return false
}

// matchMultipleIsFunc packages up the matching logic for MatchLetters,
// MatchDigits, and MatchSpaces, which otherwise differ only in the
// function used to test the rune.
func (r *LookaheadReader) matchMultipleIsFunc(isFunc func(rune) bool) bool {
	r.Result.clear()
	found := false
	for r.lookahead != '\n' && isFunc(r.lookahead) {
		current, _ := r.Next()
		if !found {
			r.Result.setPos(r.pos)
			found = true
		}
		r.Result.appendRune(current)
	}
	return found
}

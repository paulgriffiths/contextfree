package lexer

import (
	"github.com/paulgriffiths/contextfree/grammar/internal/lar"
	"io"
)

// Lex extracts a list of tokens from a context-free grammar.
func Lex(input io.Reader) (TokenList, LexErr) {
	reader, err := lar.NewLookaheadReader(input)
	if err != nil {
		return nil, LexError{LexErrBadInput, "", lar.FilePos{0, 0}}
	}

	tokens := []Token{}
	startOfLine := true // True if no tokens yet on current line

	for {

		// Ignore leading whitespace

		reader.MatchSpaces()
		if reader.EndOfInput() {
			break
		}

		// Only return an end-of-line token for lines which
		// have some other token on them (i.e. don't return
		// end-of-line tokens for blank lines and comment-only
		// lines).

		if startOfLine && reader.MatchOneOf('#') {
			for reader.MatchAnyExcept('\n') {
			}
		}

		if startOfLine && reader.MatchNewline() {
			continue
		} else {
			startOfLine = false
		}

		// Extract the next Token

		switch {
		case reader.MatchOneOf('#'):
			for reader.MatchAnyExcept('\n') {
			}
		case reader.MatchNewline():
			tokens = append(tokens,
				Token{TokenEndOfLine, "", reader.Result.Pos})
			startOfLine = true
		case reader.MatchOneOf(':'):
			tokens = append(tokens,
				Token{TokenArrow, ":", reader.Result.Pos})
		case reader.MatchOneOf('|'):
			tokens = append(tokens,
				Token{TokenAlt, "|", reader.Result.Pos})
		case reader.MatchLetter():
			t := string(reader.Result.Value[0])
			pos := reader.Result.Pos
			for {
				if reader.MatchLetters() {
					t = t + string(reader.Result.Value)
				} else if reader.MatchDigits() {
					t = t + string(reader.Result.Value)
				} else if reader.MatchOneOf('\'') {
					t = t + string(reader.Result.Value)
				} else {
					break
				}
			}
			if t == "e" {
				tokens = append(tokens, Token{TokenEmpty, t, pos})
			} else {
				tokens = append(tokens, Token{TokenNonTerminal, t, pos})
			}
		case reader.MatchOneOf('`'):
			t := ""
			pos := lar.FilePos{}
			startPos := reader.Result.Pos
			for reader.MatchAnyExcept('`', '\n') {
				t += string(reader.Result.Value)
				pos = reader.Result.Pos
			}
			if !reader.MatchOneOf('`') {
				return nil, LexError{LexErrUnterminatedTerminal, "", pos}
			}
			tokens = append(tokens, Token{TokenTerminal, t, startPos})
		default:
			reader.MatchAnyExcept()
			return nil, LexError{LexErrIllegalCharacter,
				string(reader.Result.Value[0]), reader.Result.Pos}
		}
	}
	return tokens, nil
}

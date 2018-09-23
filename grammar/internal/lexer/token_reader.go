package lexer

// TokenReader provides lookahead and matching functionality
// for a list of tokens.
type TokenReader struct {
	tokens                       []Token
	currentIndex, lookaheadIndex int
}

// NewTokenReader creates a new tokenReader.
func NewTokenReader(tokens []Token) TokenReader {
	return TokenReader{tokens, -1, 0}
}

// Reset resets the current and lookahead indices.
func (r *TokenReader) Reset() {
	r.currentIndex = -1
	r.lookaheadIndex = 0
}

// AtEnd checks if if we've matched the last token.
func (r TokenReader) AtEnd() bool {
	return r.lookaheadIndex == len(r.tokens)
}

// Match returns true and advances the indices if the next token
// to be read matches the provided type.
func (r *TokenReader) Match(t TokenType) bool {
	if r.AtEnd() {
		return false
	}
	if r.tokens[r.lookaheadIndex].T == t {
		r.currentIndex++
		r.lookaheadIndex++
		return true
	}
	return false
}

// Peek returns true if the next to be read matches the provided type.
func (r TokenReader) Peek(t TokenType) bool {
	if r.AtEnd() {
		return false
	}
	if r.tokens[r.lookaheadIndex].T == t {
		return true
	}
	return false
}

// Current returns the most recently read token. This should only
// be called after a successful match, as it will panic if no tokens
// have yet been read.
func (r TokenReader) Current() Token {
	if r.currentIndex < 0 || r.currentIndex >= len(r.tokens) {
		panic("no current token")
	}
	return r.tokens[r.currentIndex]
}

// Lookahead returns the token that would be read next. Either atEnd,
// peek or match should be called and checked prior to calling
// lookahead, as it will panic if all the tokens have been read.
func (r TokenReader) Lookahead() Token {
	if r.lookaheadIndex < 0 || r.lookaheadIndex >= len(r.tokens) {
		panic("no current token")
	}
	return r.tokens[r.lookaheadIndex]
}

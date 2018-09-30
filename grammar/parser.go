package grammar

// BUG: Doesn't throw error if undefined nonterminal is present.

import (
	"github.com/paulgriffiths/contextfree/types/symbols"
	"github.com/paulgriffiths/lexer"
	"io"
)

var lexemePatterns = []string{
	"#[^\n]*\n",
	"[[:alpha:]][[:alnum:]']*",
	"`[^`]+`",
	"\\|",
	":",
	"\n",
}

const (
	tComment = iota
	tNonT
	tT
	tAlt
	tArrow
	tNewline
	tEmpty
)

// parse parses a context-free grammar and creates a corresponding
// data structure.
func parse(input io.Reader) (*Grammar, error) {
	l, lerr := lexer.New(lexemePatterns)
	if lerr != nil {
		return nil, lerr
	}

	tokens, err := l.Lex(input)
	if err != nil {
		return nil, lerr
	}

	g := firstPass(tokens)
	if perr := secondPass(g, tokens); perr != nil {
		return nil, perr
	}

	return g, nil
}

// secondPass takes a second pass through the grammar and extracts
// the productions.
func secondPass(g *Grammar, t lexer.TokenList) ParseErr {
	tRead := 0
	for tRead < t.Len() {

		// Skip comments and empty lines

		for tokenIndexIs(t, tRead, tComment, tNewline) {
			tRead++
			continue
		}

		// Get the next grammar production

		numRead, err := getNextProduction(g, t[tRead:])
		if err != nil {
			return err
		}
		tRead += numRead
	}
	return nil
}

// getNextProduction extracts the next production.
func getNextProduction(g *Grammar, t lexer.TokenList) (int, ParseErr) {
	if !tokenIndexIs(t, 0, tNonT) {
		return 0, ParseError{ParseErrMissingHead, t[0].Index}
	}
	if !tokenIndexIs(t, 1, tArrow) {
		return 1, ParseError{ParseErrMissingArrow, t[1].Index}
	}

	head := g.NtTable[t[0].Value]
	tRead := 2

	for {
		str, n, perr := getNextBody(g, t[tRead:])
		tRead += n
		if perr != nil {
			return tRead, perr
		}

		g.Prods[head] = append(g.Prods[head], str)

		if tokenIndexIs(t, tRead, tNewline) {
			tRead++
		}
		if !tokenIndexIs(t, tRead, tAlt) {
			break
		}
		tRead++
	}

	return tRead, nil
}

// getNextBody extracts the next production body.
func getNextBody(g *Grammar, t lexer.TokenList) (symbols.String, int, ParseErr) {
	if tokenIndexIs(t, 0, tEmpty) {
		if tokenIndexIs(t, 1, tNonT, tT) {
			return nil, 0, ParseError{ParseErrEmptyNotAlone, t[1].Index}
		}
		return symbols.String{}, 1, nil
	}

	syms := symbols.String{}
	n := 0

loop:
	for {
		switch {
		case tokenIndexIs(t, n, tNonT):
			ntID := g.NtTable[t[n].Value]
			syms = append(syms, symbols.NewNonTerminal(ntID))
		case tokenIndexIs(t, n, tT):
			tID := g.TTable[tTrim(t[n].Value)]
			syms = append(syms, symbols.NewTerminal(tID))
		case tokenIndexIs(t, n, tEmpty):
			return nil, n, ParseError{ParseErrEmptyNotAlone, t[n].Index}
		default:
			break loop
		}
		n++
	}

	if len(syms) == 0 {
		return nil, n, ParseError{ParseErrEmptyBody, t[n].Index}
	}

	return syms, n, nil
}

// firstPass makes a first pass through the grammar to identify
// the terminals and non-terminals, and to set up the symbol tables.
func firstPass(tokens lexer.TokenList) *Grammar {
	nonTerminals := []string{}
	terminals := []string{}
	ntTable := make(map[string]int)
	tTable := make(map[string]int)

	for _, token := range tokens {
		switch {
		case tokenIs(token, tNonT):
			if _, ok := ntTable[token.Value]; !ok {
				ntTable[token.Value] = len(nonTerminals)
				nonTerminals = append(nonTerminals, token.Value)
			}
		case tokenIs(token, tT):
			tvalue := tTrim(token.Value)
			if _, ok := tTable[tvalue]; !ok {
				tTable[tvalue] = len(terminals)
				terminals = append(terminals, tvalue)
			}
		}
	}

	g := Grammar{
		NonTerminals: nonTerminals,
		Terminals:    terminals,
		NtTable:      ntTable,
		TTable:       tTable,
		Prods:        make([]symbols.StringList, len(nonTerminals)),
		firsts:       nil,
		follows:      nil,
	}
	return &g
}

// tTrim removes the leading and trailing backquotes from a
// terminal token value.
func tTrim(term string) string {
	return term[1 : len(term)-1]
}

// tokenIs checks if the type of the token matches any of the
// provided types. This function distinguishes between nonterminals
// and the empty symbol, even though the patterns passed to the
// lexical analyzer do not.
func tokenIs(t lexer.Token, tokenID ...int) bool {
	for _, id := range tokenID {
		if t.ID == tNonT {
			if id == tEmpty && t.Value == "e" {
				return true
			}
			if id == tNonT && t.Value != "e" {
				return true
			}
		} else if t.ID == id {
			return true
		}
	}
	return false
}

// tokenIndexIs checks if the token at index n of the provided list
// matches any of the provided types. If index n is past the end
// of the list, the function returns false, so false means the
// token at the position in question either doesn't match any of the
// provided types, or it doesn't exist. Put another way, this function
// may safely be used to check if the token at the specified index
// exists.
func tokenIndexIs(tokens lexer.TokenList, n int, tokenID ...int) bool {
	if n >= len(tokens) {
		return false
	}
	return tokenIs(tokens[n], tokenID...)
}

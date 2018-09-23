package grammar

// BUG: Doesn't throw error if undefined nonterminal is present.

import (
	"github.com/paulgriffiths/contextfree/grammar/internal/lexer"
	"github.com/paulgriffiths/contextfree/types/symbols"
	"io"
)

// parse parses a context-free grammar and creates a corresponding
// data structure.
func parse(input io.Reader) (*Grammar, error) {
	tokens, lerr := lexer.Lex(input)
	if lerr != nil {
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
func secondPass(g *Grammar, tokens []lexer.Token) ParseErr {
	reader := lexer.NewTokenReader(tokens)

	for !reader.AtEnd() {
		if perr := getNextProduction(g, &reader); perr != nil {
			return perr
		}
	}

	return nil
}

// getNextProduction extracts the next production.
func getNextProduction(g *Grammar, reader *lexer.TokenReader) ParseErr {
	if !reader.Match(lexer.TokenNonTerminal) {
		token := reader.Lookahead()
		return ParseError{ParseErrMissingHead, token.Pos.LineOnly()}
	}

	head := g.NtTable[reader.Current().S]

	if !reader.Match(lexer.TokenArrow) {
		token := reader.Current()
		return ParseError{ParseErrMissingArrow,
			token.Pos.Advance(len(reader.Current().S))}
	}

	for {
		cmp, perr := getNextBody(g, reader)
		if perr != nil {
			return perr
		}

		g.Prods[head] = append(g.Prods[head], cmp)

		reader.Match(lexer.TokenEndOfLine)
		if !reader.Match(lexer.TokenAlt) {
			break
		}
	}

	return nil
}

// getNextBody extracts the next production body.
func getNextBody(g *Grammar, reader *lexer.TokenReader) (symbols.String, ParseErr) {
	if reader.Match(lexer.TokenEmpty) {
		token := reader.Current()
		if reader.Peek(lexer.TokenNonTerminal) ||
			reader.Peek(lexer.TokenTerminal) ||
			reader.Peek(lexer.TokenEmpty) {
			return nil, ParseError{ParseErrEmptyNotAlone,
				token.Pos.Advance(1)}
		}
		return symbols.String{symbols.NewSymbolEmpty()}, nil
	}

	cmps := symbols.String{}

	for {
		if reader.Match(lexer.TokenNonTerminal) {
			token := reader.Current()
			cmps = append(cmps,
				symbols.NewNonTerminal(g.NtTable[token.S]))
		} else if reader.Match(lexer.TokenTerminal) {
			token := reader.Current()
			cmps = append(cmps,
				symbols.NewTerminal(g.TTable[token.S]))
		} else if reader.Match(lexer.TokenEmpty) {
			token := reader.Current()
			return nil, ParseError{ParseErrEmptyNotAlone,
				token.Pos.Advance(1)}
		} else {
			break
		}
	}

	if len(cmps) == 0 {
		token := reader.Current()
		return nil, ParseError{ParseErrEmptyBody,
			token.Pos.Advance(1)}
	}

	return cmps, nil
}

// firstPass makes a first pass through the grammar to identify
// the terminals and non-terminals, and to set up the symbol tables.
func firstPass(tokens []lexer.Token) *Grammar {
	nonTerminals := []string{}
	terminals := []string{}
	ntTable := make(map[string]int)
	tTable := make(map[string]int)

	for _, token := range tokens {
		switch token.T {
		case lexer.TokenNonTerminal:
			if _, ok := ntTable[token.S]; !ok {
				ntTable[token.S] = len(nonTerminals)
				nonTerminals = append(nonTerminals, token.S)
			}
		case lexer.TokenTerminal:
			if _, ok := tTable[token.S]; !ok {
				tTable[token.S] = len(terminals)
				terminals = append(terminals, token.S)
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

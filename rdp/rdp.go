package rdp

import (
	"github.com/paulgriffiths/contextfree/grammar"
	"github.com/paulgriffiths/contextfree/tree"
	"github.com/paulgriffiths/contextfree/types/symbols"
	"github.com/paulgriffiths/lexer"
	"io"
	"strings"
)

// Rdp implements a recursive descent parser.
type Rdp struct {
	g     *grammar.Grammar
	lexer *lexer.Lexer
}

// New creates a new recursive descent parser.
func New(g *grammar.Grammar) (*Rdp, error) {
	l, err := lexer.New(g.Terminals)
	if err != nil {
		return nil, err
	}
	r := Rdp{g, l}
	return &r, nil
}

// FromFile constructs a recursive descent parser from a
// context-free grammar representation in a text file.
func FromFile(filename string) (*Rdp, error) {
	g, gerr := grammar.FromFile(filename)
	if gerr != nil {
		return nil, gerr
	}
	return New(g)
}

// FromReader constructs a recursive descent parser from a
// context-free grammar representation in an io.Reader
func FromReader(reader io.Reader) (*Rdp, error) {
	g, gerr := grammar.New(reader)
	if gerr != nil {
		return nil, gerr
	}
	return New(g)
}

// Parse parses input against a grammar and returns a parse tree,
// or nil on failure.
func (r Rdp) Parse(input string) *tree.Node {
	tokens, err := r.lexer.Lex(strings.NewReader(input))
	if err != nil {
		return nil
	}

	node, n := r.parseNT(tokens, 0)
	if n == len(tokens) {
		return node
	}
	return nil
}

// parseSym parses a grammar symbol.
func (r Rdp) parseSym(t lexer.TokenList, sym symbols.Symbol) (*tree.Node, int) {
	var node *tree.Node
	numTerms := 0

	switch sym.T {
	case symbols.SymbolNonTerminal:
		node, numTerms = r.parseNT(t, sym.I)
	case symbols.SymbolTerminal:
		if len(t) > 0 && sym.I == t[0].ID {
			node, numTerms = tree.New(sym, t[0].Value, nil), 1
		}
	}

	return node, numTerms
}

// parseNT parses a non-terminal.
func (r Rdp) parseNT(t lexer.TokenList, nt int) (*tree.Node, int) {
	for _, body := range r.g.Prods[nt] {
		if children, numTerms := r.parseBody(t, body); children != nil {
			return tree.New(
				symbols.Symbol{symbols.SymbolNonTerminal, nt},
				r.g.NonTerminals[nt],
				children,
			), numTerms
		}
	}

	return nil, 0
}

// parseBody parses a production body.
func (r Rdp) parseBody(t lexer.TokenList, body symbols.String) ([]*tree.Node, int) {
	var children []*tree.Node
	matchLength := 0

	if body.IsEmpty() {
		return []*tree.Node{}, 0
	}

	for _, symbol := range body {
		node, numTerms := r.parseSym(t[matchLength:], symbol)
		if node == nil {
			return nil, 0
		}
		children = append(children, node)
		matchLength += numTerms
	}

	return children, matchLength
}

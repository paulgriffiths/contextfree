package rdp

import (
	"github.com/paulgriffiths/contextfree/grammar"
	"github.com/paulgriffiths/contextfree/tree"
	"github.com/paulgriffiths/contextfree/types/symbols"
	"github.com/paulgriffiths/lexer"
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

	l, lerr := lexer.New(g.Terminals)
	if lerr != nil {
		return nil, lerr
	}

	newParser := Rdp{g, l}
	return &newParser, nil
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

// parseComp parses a grammar symbol.
func (r Rdp) parseComp(t lexer.TokenList, sym symbols.Symbol) (*tree.Node, int) {
	var node *tree.Node
	numTerms := 0

	switch sym.T {
	case symbols.SymbolNonTerminal:
		node, numTerms = r.parseNT(t, sym.I)
	case symbols.SymbolTerminal:
		if len(t) > 0 && sym.I == t[0].ID {
			node, numTerms = tree.NewNode(sym, t[0].Value, nil), 1
		}
	case symbols.SymbolEmpty:
		node = tree.NewNode(sym, "e", nil)
	}

	return node, numTerms
}

// parseNT parses a non-terminal.
func (r Rdp) parseNT(t lexer.TokenList, nt int) (*tree.Node, int) {
	for _, body := range r.g.Prods[nt] {
		if children, numTerms := r.parseBody(t, body); children != nil {
			return tree.NewNode(
				symbols.Symbol{symbols.SymbolNonTerminal, nt},
				r.g.NonTerminals[nt],
				children,
			), numTerms
		}
	}

	return nil, 0
}

// parseBody parses a production body.
func (r Rdp) parseBody(t lexer.TokenList, body []symbols.Symbol) ([]*tree.Node, int) {
	var children []*tree.Node
	matchLength := 0

	for _, symbol := range body {
		node, numTerms := r.parseComp(t[matchLength:], symbol)
		if node == nil {
			return nil, 0
		}
		children = append(children, node)
		matchLength += numTerms
	}

	return children, matchLength
}

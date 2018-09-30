package pp

import (
	"github.com/paulgriffiths/contextfree/grammar"
	"github.com/paulgriffiths/contextfree/tree"
	"github.com/paulgriffiths/contextfree/types/symbols"
	"github.com/paulgriffiths/lexer"
	"io"
	"strings"
)

// Pp represents a predictive parser.
type Pp struct {
	g     *grammar.Grammar
	Table ppTable
	lexer *lexer.Lexer
}

// New constructs a new predictive parser for a context-free grammar.
func New(g *grammar.Grammar) (*Pp, error) {
	table := makePPTable(g)
	l, err := lexer.New(g.Terminals)
	if err != nil {
		return nil, err
	}
	newParser := Pp{g, table, l}
	return &newParser, nil
}

// FromFile constructs a predictive parser from a context-free grammar
// representation in a text file.
func FromFile(filename string) (*Pp, error) {
	g, gerr := grammar.FromFile(filename)
	if gerr != nil {
		return nil, gerr
	}
	return New(g)
}

// FromReader constructs a predictive parser from a context-free grammar
// representation in an io.Reader
func FromReader(reader io.Reader) (*Pp, error) {
	g, gerr := grammar.New(reader)
	if gerr != nil {
		return nil, gerr
	}
	return New(g)
}

// Parse parses input against a grammar and returns a parse tree,
// or nil on failure.
func (p Pp) Parse(input string) *tree.Node {
	tokens, err := p.lexer.Lex(strings.NewReader(input))
	if err != nil {
		return nil
	}

	if node, n := p.parseNT(tokens, 0); n == tokens.Len() {
		return node
	}
	return nil
}

// parseSym parses a production str sym.
func (p Pp) parseSym(t lexer.TokenList, sym symbols.Symbol) (*tree.Node, int) {
	var node *tree.Node
	numTerms := 0

	switch sym.T {
	case symbols.SymbolNonTerminal:
		node, numTerms = p.parseNT(t, sym.I)
	case symbols.SymbolTerminal:
		if !t.IsEmpty() && t[0].ID == sym.I {
			node, numTerms = tree.New(sym, t[0].Value, nil), 1
		}
	}

	return node, numTerms
}

// parseNT parses a non-terminal.
func (p Pp) parseNT(t lexer.TokenList, nt int) (*tree.Node, int) {

	// If there are no more terminals in the list, check whether
	// the current nonterminal can be followed by end-of-input.

	if t.IsEmpty() {
		str := p.Table[nt][len(p.g.Terminals)]
		if str.IsEmpty() {
			return nil, 0
		}
		if str[0].IsEmpty() {
			term := tree.New(symbols.NewNonTerminal(nt),
				p.g.NonTerminals[nt], []*tree.Node{})
			return term, 0
		}
		panic("unexpected terminal condition")
	}

	// Get the str for this nonterminal with the next terminal,
	// returning an error if the predictive parsing table doesn't
	// contain an entry.

	str := p.Table[nt][t[0].ID]
	if str.IsEmpty() {
		return nil, 0
	}

	// No need to parse the body if it's empty.

	if str[0].IsEmpty() {
		return tree.New(
			symbols.Symbol{symbols.SymbolNonTerminal, nt},
			p.g.NonTerminals[nt],
			[]*tree.Node{},
		), 0
	}

	// Production body isn't empty, so parse it.

	if children, numTerms := p.parseString(t, str[0]); children != nil {
		return tree.New(
			symbols.Symbol{symbols.SymbolNonTerminal, nt},
			p.g.NonTerminals[nt],
			children,
		), numTerms
	}

	return nil, 0
}

// parseString parses a production str.
func (p Pp) parseString(t lexer.TokenList, str symbols.String) ([]*tree.Node, int) {
	var children []*tree.Node
	matchLength := 0

	for _, sym := range str {
		node, numTerms := p.parseSym(t[matchLength:], sym)
		if node == nil {
			return nil, 0
		}
		children = append(children, node)
		matchLength += numTerms
	}

	return children, matchLength
}

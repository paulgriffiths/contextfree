package pp

import (
	"github.com/paulgriffiths/contextfree/grammar"
	"github.com/paulgriffiths/contextfree/lexer"
	"github.com/paulgriffiths/contextfree/tree"
	"github.com/paulgriffiths/contextfree/types/symbols"
)

// Pp represents a predictive parser.
type Pp struct {
	g     *grammar.Grammar
	Table ppTable
	lexer *lexer.Lexer
}

// NewPp constructs a new predictive parser for a context-free grammar.
func New(grammar *grammar.Grammar) *Pp {
	table := makePPTable(grammar)
	l, err := lexer.New(grammar)
	if err != nil {
		return nil
	}
	newParser := Pp{grammar, table, l}
	return &newParser
}

// Parse parses input against a grammar and returns a parse tree,
// or nil on failure.
func (p Pp) Parse(input string) *tree.Node {
	terminals, err := p.lexer.Lex(input)
	if err != nil {
		return nil
	}

	node, n := p.parseNT(terminals, 0)
	if n == terminals.Len() {
		return node
	}
	return nil
}

// parseSymbol parses a production str sym.
func (p Pp) parseSymbol(t lexer.TerminalList, comp symbols.Symbol) (*tree.Node, int) {
	var node *tree.Node
	numTerms := 0

	switch comp.T {
	case symbols.SymbolNonTerminal:
		node, numTerms = p.parseNT(t, comp.I)
	case symbols.SymbolTerminal:
		if !t.IsEmpty() && t[0].N == comp.I {
			node, numTerms = tree.NewNode(comp, t[0].S, nil), 1
		}
	case symbols.SymbolEmpty:
		node = tree.NewNode(comp, "e", nil)
	}

	return node, numTerms
}

// parseNT parses a non-terminal.
func (p Pp) parseNT(t lexer.TerminalList, nt int) (*tree.Node, int) {

	// If there are no more terminals in the list, check whether
	// the current nonterminal can be followed by end-of-input.

	if t.IsEmpty() {
		str := p.Table[nt][len(p.g.Terminals)]
		if str.IsEmpty() {
			return nil, 0
		}
		if str[0].IsEmptyString() {
			em := tree.NewNode(symbols.NewSymbolEmpty(), "e", nil)
			term := tree.NewNode(symbols.NewNonTerminal(nt),
				p.g.NonTerminals[nt], []*tree.Node{em})
			return term, 0
		}
		panic("unexpected terminal condition")
	}

	// Get the str for this nonterminal with the next terminal,
	// returning an error if the predictive parsing table doesn't
	// contain an entry.

	str := p.Table[nt][t[0].N]
	if str.IsEmpty() {
		return nil, 0
	}

	if children, numTerms := p.parseString(t, str[0]); children != nil {
		return tree.NewNode(
			symbols.Symbol{symbols.SymbolNonTerminal, nt},
			p.g.NonTerminals[nt],
			children,
		), numTerms
	}

	return nil, 0
}

// parseString parses a production str.
func (p Pp) parseString(t lexer.TerminalList, str symbols.String) ([]*tree.Node, int) {
	var children []*tree.Node
	matchLength := 0

	for _, sym := range str {
		node, numTerms := p.parseSymbol(t[matchLength:], sym)
		if node == nil {
			return nil, 0
		}
		children = append(children, node)
		matchLength += numTerms
	}

	return children, matchLength
}

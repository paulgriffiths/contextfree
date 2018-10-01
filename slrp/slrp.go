package slrp

import (
	"github.com/paulgriffiths/contextfree/grammar"
	"github.com/paulgriffiths/contextfree/tree"
	"github.com/paulgriffiths/contextfree/types/symbols"
	"github.com/paulgriffiths/gods/stacks"
	"github.com/paulgriffiths/lexer"
	"io"
	"strings"
)

// Slrp implements a simple LR parser for a context-free grammar.
type Slrp struct {
	g     *grammar.Grammar
	t     Table
	lexer *lexer.Lexer
}

// New constructs a new simple LR parser for a context-free grammar.
func New(g *grammar.Grammar) (*Slrp, error) {
	table := NewTable(g)
	l, err := lexer.New(g.Terminals)
	if err != nil {
		return nil, err
	}
	newParser := Slrp{g, table, l}
	return &newParser, nil
}

// FromFile constructs a simple LR parser from a context-free grammar
// representation in a text file.
func FromFile(filename string) (*Slrp, error) {
	g, gerr := grammar.FromFile(filename)
	if gerr != nil {
		return nil, gerr
	}

	return New(g)
}

// FromReader constructs a simple LR parser from a context-free grammar
// representation in an io.Reader
func FromReader(reader io.Reader) (*Slrp, error) {
	g, gerr := grammar.New(reader)
	if gerr != nil {
		return nil, gerr
	}
	return New(g)
}

// Parse parses input against a grammar and returns a parse tree,
// or nil on failure.
func (p Slrp) Parse(input string) *tree.Node {
	tokens, err := p.lexer.Lex(strings.NewReader(input))
	if err != nil {
		return nil
	}

	n := 0

	stack := stacks.NewStackInt()
	stack.Push(0)
	tstack := NewStackNode()

	for {
		var actionList []Action
		var sym symbols.Symbol

		// Check if we have a terminal, or if we're at end-of-input.

		if n < len(tokens) {
			sym = symbols.NewTerminal(tokens[n].ID)
			actionList = p.t.A[stack.Peek()][sym.I]
		} else {
			actionList = p.t.A[stack.Peek()][len(p.t.M.Terminals)]
		}

		// Return nil if we don't have a valid action.

		if actionList == nil || len(actionList) < 1 {
			return nil
		}

		// Perform the action.

		if action := actionList[0]; action.IsShift() {
			stack.Push(action.S)
			tstack.Push(tree.New(sym, tokens[n].Value, nil))
			n++
		} else if action.IsReduce() {
			nt, n := p.t.M.ProductionFromNumber(action.S)
			prod := p.t.M.Prods[nt][n]
			children := []*tree.Node{}
			for i := 0; i < len(prod); i++ {
				stack.Pop()
				children = append([]*tree.Node{tstack.Pop()}, children...)
			}
			stack.Push(p.t.G[stack.Peek()][nt])
			tstack.Push(tree.New(symbols.NewNonTerminal(nt-1),
				p.t.M.NonTerminals[nt], children))
		} else if action.IsAccept() {
			return tstack.Pop()
		}
	}
}

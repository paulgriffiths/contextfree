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

	table := NewTable(g)
	l, lerr := lexer.New(g.Terminals)
	if lerr != nil {
		return nil, lerr
	}

	newParser := Slrp{g, table, l}
	return &newParser, nil
}

// FromReader constructs a simple LR parser from a context-free grammar
// representation in an io.Reader
func FromReader(reader io.Reader) (*Slrp, error) {
	g, gerr := grammar.New(reader)
	if gerr != nil {
		return nil, gerr
	}

	table := NewTable(g)
	l, lerr := lexer.New(g.Terminals)
	if lerr != nil {
		return nil, lerr
	}

	newParser := Slrp{g, table, l}
	return &newParser, nil
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

		if n < len(tokens) {
			sym := symbols.NewTerminal(tokens[n].ID)
			actionList = p.t.A[stack.Peek()][sym.I]
		} else {
			actionList = p.t.A[stack.Peek()][len(p.t.M.Terminals)]
		}

		if actionList == nil || len(actionList) < 1 {
			return nil
		}

		action := actionList[0]

		if action.IsShift() {
			stack.Push(action.S)
			sym := symbols.NewTerminal(tokens[n].ID)
			newNode := tree.Node{sym, tokens[n].Value, nil}
			tstack.Push(&newNode)
			n++
		} else if action.IsReduce() {
			nt, n := p.t.M.ProductionFromNumber(action.S)
			prod := p.t.M.Prods[nt][n]
			children := []*tree.Node{}
			for i := 0; i < len(prod); i++ {
				stack.Pop()
				node := tstack.Pop()
				children = append([]*tree.Node{node}, children...)
			}
			node := tree.Node{symbols.NewNonTerminal(nt - 1),
				p.t.M.NonTerminals[nt], children}
			tstack.Push(&node)
			stack.Push(p.t.G[stack.Peek()][nt])
		} else if action.IsAccept() {
			node := tstack.Pop()
			return node
		}
	}

	return nil
}

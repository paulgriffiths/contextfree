package tree

import (
	"fmt"
	"github.com/paulgriffiths/contextfree/types/symbols"
	"io"
)

// Node represents a parse tree node.
type Node struct {
	Sym      symbols.Symbol
	Value    string
	Children []*Node
}

// NewNode creates a new parse tree node.
func NewNode(sym symbols.Symbol, value string, children []*Node) *Node {
	node := Node{sym, value, children}
	return &node
}

// WriteTerminals writes the terminals in the parse tree to
// the provided io.Writer.
func (t *Node) WriteTerminals(writer io.Writer) {
	if t.Sym.T == symbols.SymbolTerminal {
		writer.Write([]byte(fmt.Sprintf("%s", t.Value)))
	}
	for _, child := range t.Children {
		child.WriteTerminals(writer)
	}
}

// WriteBracketed outputs the parse tree in a bracketed-parse format.
func (t *Node) WriteBracketed(writer io.Writer, opts ...string) {
	qc := ""
	ob := "("
	cb := ")"
	if len(opts) > 0 {
		qc = opts[0]
	}
	if len(opts) > 2 {
		ob = opts[1]
		cb = opts[2]
	}

	switch t.Sym.T {
	case symbols.SymbolTerminal:
		writer.Write([]byte(fmt.Sprintf("%s%s%s", qc, t.Value, qc)))
	case symbols.SymbolNonTerminal:
		writer.Write([]byte(fmt.Sprintf("%s%s ", ob, t.Value)))
		if len(t.Children) == 0 {
			writer.Write([]byte(fmt.Sprintf("e")))
		} else {
			for n, child := range t.Children {
				child.WriteBracketed(writer, qc, ob, cb)
				if n < len(t.Children)-1 {
					writer.Write([]byte(fmt.Sprintf(" ")))
				}
			}
		}
		writer.Write([]byte(fmt.Sprintf("%s", cb)))
	}
}

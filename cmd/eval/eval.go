package main

import (
	"github.com/paulgriffiths/contextfree/tree"
	"strconv"
)

// Nonterminal ID constants, in the order they appear in the grammar.
const (
	ntE = iota
	ntT
	ntEp
	ntF
	ntTp
	ntN
)

// Terminal ID constants, in the order they appear in the grammar.
const (
	tPlus = iota
	tMinus
	tMultiply
	tDivide
	tLeftParen
	tRightParen
	tNumber
)

// eval evaluates the expression represented by the provided tree.
// In some cases, a float64 argument needs to be passed to evaluate
// the expression. As written, if this function is called from main
// with the tree returned from the parser, then the tree root passed
// to this function will always represent a nonterminal.
func eval(t *tree.Node, n ...float64) float64 {
	ch := t.Children

	if len(ch) == 0 {
		return n[0]
	}

	sym := ch[0].Sym

	switch t.Sym.I {
	case ntE:
		return eval(ch[1], eval(ch[0]))

	case ntT:
		return eval(ch[1], eval(ch[0]))

	case ntEp:
		switch {
		case sym.IsTerminal() && sym.I == tPlus:
			return eval(ch[2], n[0]+eval(ch[1]))
		case sym.IsTerminal() && sym.I == tMinus:
			return eval(ch[2], n[0]-eval(ch[1]))
		}

	case ntTp:
		switch {
		case sym.IsTerminal() && sym.I == tMultiply:
			return eval(ch[2], n[0]*eval(ch[1]))
		case sym.IsTerminal() && sym.I == tDivide:
			return eval(ch[2], n[0]/eval(ch[1]))
		}

	case ntF:
		switch {
		case sym.IsNonTerminal() && sym.I == ntN:
			return eval(ch[0])
		case sym.IsTerminal() && sym.I == tLeftParen:
			return eval(ch[1])
		}

	case ntN:
		val, err := strconv.ParseFloat(ch[0].Value, 64)
		if err != nil {
			panic("unexpectedly failed to convert float")
		}
		return val
	}

	panic("unexpectedly reached end of eval function")
}

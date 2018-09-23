package pp_test

import (
	"github.com/paulgriffiths/contextfree/grammar"
	"github.com/paulgriffiths/contextfree/pp"
	"github.com/paulgriffiths/contextfree/tree"
)

func getParserFromFile(filename string) (*pp.Pp, error) {
	g, err := grammar.FromFile(filename)
	if err != nil {
		return nil, err
	}

	parser := pp.New(g)
	if parser == nil {
		return nil, err
	}

	return parser, nil
}

func getParseTreeFromFile(filename, input string) (*tree.Node, error) {
	g, err := grammar.FromFile(filename)
	if err != nil {
		return nil, err
	}

	parser := pp.New(g)
	if parser == nil {
		return nil, err
	}

	tree := parser.Parse(input)
	if tree == nil {
		return nil, nil
	}

	return tree, nil
}

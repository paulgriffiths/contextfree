package rdp_test

import (
	"github.com/paulgriffiths/contextfree/grammar"
	"github.com/paulgriffiths/contextfree/rdp"
	"github.com/paulgriffiths/contextfree/tree"
	"testing"
)

func getParserFromFile(t *testing.T, filename string) (*rdp.Rdp, error) {
	g, err := grammar.FromFile(filename)
	if err != nil {
		return nil, err
	}

	parser, err := rdp.New(g)
	if err != nil {
		return nil, err
	}

	return parser, nil
}

func getParseTreeFromFile(t *testing.T, filename,
	input string) (*tree.Node, error) {
	g, err := grammar.FromFile(filename)
	if err != nil {
		return nil, err
	}

	parser, err := rdp.New(g)
	if err != nil {
		return nil, err
	}

	tree := parser.Parse(input)
	if tree == nil {
		return nil, nil
	}

	return tree, nil
}

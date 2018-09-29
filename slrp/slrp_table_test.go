package slrp

import (
	"github.com/paulgriffiths/contextfree/grammar"
	"testing"
)

var goToTableTestCase = [][]int{
	[]int{-1, 1, 2, 3},
	[]int{-1, -1, -1, -1},
	[]int{-1, -1, -1, -1},
	[]int{-1, -1, -1, -1},
	[]int{-1, 8, 2, 3},
	[]int{-1, -1, -1, -1},
	[]int{-1, -1, 9, 3},
	[]int{-1, -1, -1, 10},
	[]int{-1, -1, -1, -1},
	[]int{-1, -1, -1, -1},
	[]int{-1, -1, -1, -1},
	[]int{-1, -1, -1, -1},
}

var actionTableTestCase = [][][]Action{
	[][]Action{
		[]Action{},
		[]Action{},
		[]Action{NewShift(4)},
		[]Action{},
		[]Action{NewShift(5)},
		[]Action{},
	},
	[][]Action{
		[]Action{NewShift(6)},
		[]Action{},
		[]Action{},
		[]Action{},
		[]Action{},
		[]Action{NewAccept()},
	},
	[][]Action{
		[]Action{NewReduce(2)},
		[]Action{NewShift(7)},
		[]Action{},
		[]Action{NewReduce(2)},
		[]Action{},
		[]Action{NewReduce(2)},
	},
	[][]Action{
		[]Action{NewReduce(4)},
		[]Action{NewReduce(4)},
		[]Action{},
		[]Action{NewReduce(4)},
		[]Action{},
		[]Action{NewReduce(4)},
	},
	[][]Action{
		[]Action{},
		[]Action{},
		[]Action{NewShift(4)},
		[]Action{},
		[]Action{NewShift(5)},
		[]Action{},
	},
	[][]Action{
		[]Action{NewReduce(6)},
		[]Action{NewReduce(6)},
		[]Action{},
		[]Action{NewReduce(6)},
		[]Action{},
		[]Action{NewReduce(6)},
	},
	[][]Action{
		[]Action{},
		[]Action{},
		[]Action{NewShift(4)},
		[]Action{},
		[]Action{NewShift(5)},
		[]Action{},
	},
	[][]Action{
		[]Action{},
		[]Action{},
		[]Action{NewShift(4)},
		[]Action{},
		[]Action{NewShift(5)},
		[]Action{},
	},
	[][]Action{
		[]Action{NewShift(6)},
		[]Action{},
		[]Action{},
		[]Action{NewShift(11)},
		[]Action{},
		[]Action{},
	},
	[][]Action{
		[]Action{NewReduce(1)},
		[]Action{NewShift(7)},
		[]Action{},
		[]Action{NewReduce(1)},
		[]Action{},
		[]Action{NewReduce(1)},
	},
	[][]Action{
		[]Action{NewReduce(3)},
		[]Action{NewReduce(3)},
		[]Action{},
		[]Action{NewReduce(3)},
		[]Action{},
		[]Action{NewReduce(3)},
	},
	[][]Action{
		[]Action{NewReduce(5)},
		[]Action{NewReduce(5)},
		[]Action{},
		[]Action{NewReduce(5)},
		[]Action{},
		[]Action{NewReduce(5)},
	},
}

func TestSlrpActionTable(t *testing.T) {
	g, err := grammar.FromFile(tgArithLr)
	if err != nil {
		t.Errorf("couldn't get grammar: %v", err)
		return
	}

	table := NewTable(g)

	if len(table.A) != len(actionTableTestCase) {
		t.Errorf("lengths not equal, got %d, want %d", len(table.A),
			len(actionTableTestCase))
		return
	}

	for i, a := range actionTableTestCase {
		for j := range a {
			if table.A[i][j] == nil && actionTableTestCase[i][i] != nil {
				t.Errorf("(%d, %d), not both equal to nil", i, j)
				continue
			}

			if len(table.A[i][j]) != len(actionTableTestCase[i][j]) {
				t.Errorf("(%d, %d), for length got %d, want %d", i, j,
					len(table.A[i][j]), len(actionTableTestCase[i][j]))
				continue
			}

			if len(actionTableTestCase[i][j]) > 0 {
				if table.A[i][j][0] != actionTableTestCase[i][j][0] {
					t.Errorf("(%d, %d) got %d, want %d", i, j,
						table.A[i][j], actionTableTestCase[i][j])
				}
			}
		}
	}
}

func TestSlrpGotoTable(t *testing.T) {
	g, err := grammar.FromFile(tgArithLr)
	if err != nil {
		t.Errorf("couldn't get grammar: %v", err)
		return
	}

	table := NewTable(g)

	if len(table.G) != len(goToTableTestCase) {
		t.Errorf("lengths not equal, got %d, want %d", len(table.G),
			len(goToTableTestCase))
	}

	for i, a := range goToTableTestCase {
		for j := range a {
			if table.G[i][j] != goToTableTestCase[i][j] {
				t.Errorf("(%d, %d) got %d, want %d", i, j,
					table.G[i][j], goToTableTestCase[i][j])
			}
		}
	}
}

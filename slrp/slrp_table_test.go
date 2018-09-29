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

func TestSlrpGotoTable(t *testing.T) {
	g, err := grammar.FromFile(tgArithLr)
	if err != nil {
		t.Errorf("couldn't get grammar: %v", err)
		return
	}

	table := NewSlrpTable(g)

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

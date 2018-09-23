package pp_test

import (
	"github.com/paulgriffiths/contextfree/grammar"
	"github.com/paulgriffiths/contextfree/pp"
	"testing"
)

func TestTable(t *testing.T) {
	g, err := grammar.FromFile(tgArithNlr)
	if err != nil {
		t.Errorf("couldn't get grammar from file %s: %v", tgArithNlr, err)
		return
	}

	pp := pp.New(g)

	matrix := [][]int{
		//    +  *  (  )  n  $
		[]int{0, 0, 1, 0, 1, 0}, // E
		[]int{0, 0, 1, 0, 1, 0}, // T
		[]int{1, 0, 0, 1, 0, 1}, // E'
		[]int{0, 0, 1, 0, 1, 0}, // F
		[]int{1, 1, 0, 1, 0, 1}, // T'
		[]int{0, 0, 0, 0, 1, 0}, // Digits
	}

	for i, row := range matrix {
		for j, l := range row {
			v := len(pp.Table[i][j])
			if v != l {
				t.Errorf("For (%d,%d), got %d, want %d", i, j, v, l)
			}
		}
	}
}

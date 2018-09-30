package grammar

import (
	"fmt"
	"io"
)

// Output outputs a representation of the grammar.
func (g *Grammar) Output(writer io.Writer) {
	maxNL := g.MaxNonTerminalNameLength()

	for i, prod := range g.Prods {
		for n, str := range prod {
			var s string
			if n == 0 {
				s = fmt.Sprintf("%-[1]*s :", maxNL, g.NonTerminals[i])
			} else {
				s = fmt.Sprintf("%-[1]*s |", maxNL, "")
			}
			writer.Write([]byte(s))

			for _, sym := range str {
				switch {
				case sym.IsNonTerminal():
					s = fmt.Sprintf(" %s", g.NonTerminals[sym.I])
				case sym.IsTerminal():
					s = fmt.Sprintf(" `%s`", g.Terminals[sym.I])
				default:
					panic("unexpected str component")
				}
				writer.Write([]byte(s))
			}

			if str.IsEmpty() {
				writer.Write([]byte(" e"))
			}

			writer.Write([]byte("\n"))
		}
	}
}

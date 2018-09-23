package symbols

// SymbolType represents the type of a grammar symbol.
type SymbolType int

const (
	// SymbolNonTerminal represents a nonterminal grammar symbol.
	SymbolNonTerminal SymbolType = iota
	// SymbolTerminal represents a terminal grammar symbol.
	SymbolTerminal
	// SymbolEmpty represents the empty grammar symbol.
	SymbolEmpty
	// SymbolInputEnd represents the end-of-input marker.
	SymbolInputEnd
)

// Symbol represents a grammar symbol.
type Symbol struct {
	T SymbolType
	I int
}

// IsNonTerminal checks if a grammar symbol is a nonterminal.
func (s Symbol) IsNonTerminal() bool {
	return s.T == SymbolNonTerminal
}

// IsTerminal checks if a grammar symbol is a terminal.
func (s Symbol) IsTerminal() bool {
	return s.T == SymbolTerminal
}

// IsEmpty checks if a grammar symbol is the empty symbol.
func (s Symbol) IsEmpty() bool {
	return s.T == SymbolEmpty
}

// IsInputEnd checks if a grammar symbol is the end-of-input marker.
func (s Symbol) IsInputEnd() bool {
	return s.T == SymbolInputEnd
}

// NewNonTerminal returns a new nonterminal grammar symbol with
// a specified identifier.
func NewNonTerminal(n int) Symbol {
	return Symbol{SymbolNonTerminal, n}
}

// NewTerminal returns a new terminal grammar symbol with
// a specified identifier.
func NewTerminal(n int) Symbol {
	return Symbol{SymbolTerminal, n}
}

// NewSymbolEmpty returns a new empty grammar symbol.
func NewSymbolEmpty() Symbol {
	return Symbol{SymbolEmpty, 0}
}

// NewSymbolEndOfInput returns a new end-of-input marker.
func NewSymbolEndOfInput() Symbol {
	return Symbol{SymbolInputEnd, -1}
}

package symbols

// String represents a string of grammar symbols.
type String []Symbol

// IsEmpty checks if a string of grammar symbols is empty.
func (s String) IsEmpty() bool {
	return len(s) == 0
}

// IsEmptyString checks if a string of grammar symbols contains
// only a single empty symbol.
func (s String) IsEmptyString() bool {
	return len(s) == 1 && s[0].T == SymbolEmpty
}

// IsNonTerminal checks if a string of grammar symbols contains
// only a single nonterminal.
func (s String) IsNonTerminal() bool {
	return len(s) == 1 && s[0].T == SymbolNonTerminal
}

// IsTerminal checks if a string of grammar symbols contains only
// a single terminal.
func (s String) IsTerminal() bool {
	return len(s) == 1 && s[0].T == SymbolTerminal
}

// HasOnlyNonTerminals checks if a string of grammar symbols
// contains only nonterminals.
func (s String) HasOnlyNonTerminals() bool {
	for _, symbol := range s {
		if !symbol.IsNonTerminal() {
			return false
		}
	}
	return true
}

// HasOnlyTerminals checks if a string of grammar symbols
// contains only terminals.
func (s String) HasOnlyTerminals() bool {
	for _, symbol := range s {
		if !symbol.IsTerminal() {
			return false
		}
	}
	return true
}

// IsLast returns true if the provided index refers to the last
// symbol of the string.
func (s String) IsLast(n int) bool {
	return n == len(s)-1
}

// Copy makes a copy of a string.
func (s String) Copy() String {
	newString := make(String, len(s))
	copy(newString, s)
	return newString
}

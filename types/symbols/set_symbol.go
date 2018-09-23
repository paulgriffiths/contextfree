package symbols

// SetSymbol implements a set of grammar symbols.
type SetSymbol map[Symbol]bool

// NewSetSymbol creates a new set of grammar symbols with
// optional initial elements.
func NewSetSymbol(values ...Symbol) SetSymbol {
	newSet := make(map[Symbol]bool)
	for _, value := range values {
		newSet[value] = true
	}
	return newSet
}

// IsEmpty returns true if a set is the empty set.
func (s SetSymbol) IsEmpty() bool {
	return len(s) == 0
}

// Length returns the number of grammar symbols in the set.
func (s SetSymbol) Length() int {
	return len(s)
}

// Elements returns an array of the grammar symbols in the set.
func (s SetSymbol) Elements() String {
	list := make([]Symbol, 0, len(s))
	for key := range s {
		list = append(list, key)
	}
	return list
}

// Equals tests if two sets contain the same grammar symbols.
func (s SetSymbol) Equals(other SetSymbol) bool {
	if len(s) != len(other) || len(s) != len(s.Union(other)) {
		return false
	}
	return true
}

// Contains returns true if the set contains the specified grammar symbol.
func (s SetSymbol) Contains(n Symbol) bool {
	return s[n]
}

// ContainsEmpty returns true if the set contains the empty grammar symbol.
func (s SetSymbol) ContainsEmpty() bool {
	return s[Symbol{SymbolEmpty, 0}]
}

// ContainsEndOfInput returns true if the set contains an end-of-input
// marker.
func (s SetSymbol) ContainsEndOfInput() bool {
	return s[Symbol{SymbolInputEnd, -1}]
}

// Insert inserts a grammar symbol into a set if it isn't already
// in the set.
func (s *SetSymbol) Insert(n Symbol) {
	(*s)[n] = true
}

// InsertEmpty inserts an empty grammar symbol into a set if it
// isn't already in the set.
func (s *SetSymbol) InsertEmpty() {
	(*s)[Symbol{SymbolEmpty, 0}] = true
}

// InsertEndOfInput inserts an end-of-input marker into a set if it
// isn't already in the set.
func (s *SetSymbol) InsertEndOfInput() {
	(*s)[Symbol{SymbolInputEnd, -1}] = true
}

// Merge inserts into a set the grammar symbols from another set.
func (s *SetSymbol) Merge(other SetSymbol) {
	for key, value := range other {
		if value {
			(*s)[key] = true
		}
	}
}

// Copy returns a copy of the set.
func (s SetSymbol) Copy() SetSymbol {
	c := NewSetSymbol()
	for key := range s {
		c[key] = true
	}
	return c
}

// Delete deletes a grammar symbol from a set.
func (s *SetSymbol) Delete(n Symbol) {
	delete(*s, n)
}

// DeleteEmpty deletes the empty grammar symbol from a set.
func (s *SetSymbol) DeleteEmpty() {
	delete(*s, Symbol{SymbolEmpty, 0})
}

// Intersection returns the intersection of two sets.
func (s SetSymbol) Intersection(other SetSymbol) SetSymbol {
	inter := NewSetSymbol()
	for key := range s {
		if other[key] {
			inter[key] = true
		}
	}
	return inter
}

// Union returns the union of two sets.
func (s SetSymbol) Union(other SetSymbol) SetSymbol {
	return NewSetSymbol(append(s.Elements(), other.Elements()...)...)
}

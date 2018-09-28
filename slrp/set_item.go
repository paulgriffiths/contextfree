package slrp

// SetItem implements a set of integers.
type SetItem map[Item]bool

// NewSetItem creates a new set of integers with optional initial elements.
func NewSetItem(values ...Item) SetItem {
	newSet := make(map[Item]bool)
	for _, value := range values {
		newSet[value] = true
	}
	return newSet
}

// IsEmpty returns true if a set is the empty set.
func (s SetItem) IsEmpty() bool {
	return len(s) == 0
}

// Length returns the number of elements in the set.
func (s SetItem) Length() int {
	return len(s)
}

// Elements returns an array of the elements in the set.
func (s SetItem) Elements() []Item {
	list := make([]Item, 0, len(s))
	for key := range s {
		list = append(list, key)
	}
	return list
}

// Equals tests if two sets contain the same members
func (s SetItem) Equals(other SetItem) bool {
	if len(s) != len(other) || len(s) != len(s.Union(other)) {
		return false
	}
	return true
}

// Contains returns true if the set contains the specified integer.
func (s SetItem) Contains(n Item) bool {
	return s[n]
}

// Insert inserts an integer into a set if it isn't already in the set.
func (s *SetItem) Insert(n Item) {
	(*s)[n] = true
}

// Merge inserts into a set the elements from another set.
func (s *SetItem) Merge(other SetItem) {
	for key, value := range other {
		if value {
			(*s)[key] = true
		}
	}
}

// Copy returns a copy of the set.
func (s SetItem) Copy() SetItem {
	c := NewSetItem()
	for key := range s {
		c[key] = true
	}
	return c
}

// Intersection returns the intersection of two sets.
func (s SetItem) Intersection(other SetItem) SetItem {
	inter := NewSetItem()
	for key := range s {
		if other[key] {
			inter[key] = true
		}
	}
	return inter
}

// Difference returns the elements which are in set s, but not
// in set other.
func (s SetItem) Difference(other SetItem) SetItem {
	diff := NewSetItem()
	for key := range s {
		if !other[key] {
			diff[key] = true
		}
	}
	return diff
}

// Union returns the union of two sets.
func (s SetItem) Union(other SetItem) SetItem {
	return NewSetItem(append(s.Elements(), other.Elements()...)...)
}

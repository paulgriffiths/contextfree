package datastruct

// SetInt implements a set of integers.
type SetInt map[int]bool

// NewSetInt creates a new set of integers with optional initial elements.
func NewSetInt(values ...int) SetInt {
	newSet := make(map[int]bool)
	for _, value := range values {
		newSet[value] = true
	}
	return newSet
}

// IsEmpty returns true if a set is the empty set.
func (s SetInt) IsEmpty() bool {
	return len(s) == 0
}

// Length returns the number of elements in the set.
func (s SetInt) Length() int {
	return len(s)
}

// Elements returns an array of the elements in the set.
func (s SetInt) Elements() []int {
	list := make([]int, 0, len(s))
	for key := range s {
		list = append(list, key)
	}
	return list
}

// Equals tests if two sets contain the same members
func (s SetInt) Equals(other SetInt) bool {
	if len(s) != len(other) || len(s) != len(s.Union(other)) {
		return false
	}
	return true
}

// Contains returns true if the set contains the specified integer.
func (s SetInt) Contains(n int) bool {
	return s[n]
}

// Insert inserts an integer into a set if it isn't already in the set.
func (s *SetInt) Insert(n int) {
	(*s)[n] = true
}

// Merge inserts into a set the elements from another set.
func (s *SetInt) Merge(other SetInt) {
	for key, value := range other {
		if value {
			(*s)[key] = true
		}
	}
}

// Copy returns a copy of the set.
func (s SetInt) Copy() SetInt {
	c := NewSetInt()
	for key := range s {
		c[key] = true
	}
	return c
}

// Intersection returns the intersection of two sets.
func (s SetInt) Intersection(other SetInt) SetInt {
	inter := NewSetInt()
	for key := range s {
		if other[key] {
			inter[key] = true
		}
	}
	return inter
}

// Difference returns the elements which are in set s, but not
// in set other.
func (s SetInt) Difference(other SetInt) SetInt {
	diff := NewSetInt()
	for key := range s {
		if !other[key] {
			diff[key] = true
		}
	}
	return diff
}

// Union returns the union of two sets.
func (s SetInt) Union(other SetInt) SetInt {
	return NewSetInt(append(s.Elements(), other.Elements()...)...)
}

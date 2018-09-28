package symbols

// StringList represents a list of symbol strings.
type StringList []String

// IsEmpty checks is a body list is empty.
func (l StringList) IsEmpty() bool {
	return len(l) == 0
}

// HasEmpty checks if a list of symbol string contains a string
// consisting solely of the empty symbol.
func (l StringList) HasEmpty() bool {
	for _, str := range l {
		if str.IsEmptyString() {
			return true
		}
	}
	return false
}

// Copy makes a copy of a StringList.
func (l StringList) Copy() StringList {
	newList := make(StringList, len(l))
	for n := range l {
		newList[n] = l[n].Copy()
	}
	return newList
}

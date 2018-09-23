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

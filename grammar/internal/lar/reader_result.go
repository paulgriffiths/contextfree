package lar

// ReaderResult encapsulates a result buffer and its position
// and line in the input.
type ReaderResult struct {
	Value []rune
	Pos   FilePos
}

// clear clears all the fields in a ReaderResult.
func (r *ReaderResult) clear() {
	r.Value = []rune{}
	r.Pos.clear()
}

// appendByte appends a byte to a ReaderResult's value buffer
func (r *ReaderResult) appendRune(rn rune) {
	r.Value = append(r.Value, rn)
}

// setPos sets the ReaderResult's position and line
func (r *ReaderResult) setPos(pos FilePos) {
	r.Pos.setPos(pos)
}

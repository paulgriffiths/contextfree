package lar

// FilePos represents a line number and character position in a file.
type FilePos struct {
	Ch, Line int
}

// clear clears all the fields in a FilePos
func (f *FilePos) clear() {
	f.Ch, f.Line = 0, 0
}

// setPos sets the FilePos's character position and line
func (f *FilePos) setPos(pos FilePos) {
	f.Ch, f.Line = pos.Ch, pos.Line
}

// inc increments the FilePos's character position
func (f *FilePos) inc() {
	f.Ch++
}

// inc increments the FilePos's character position
func (f *FilePos) incLine() {
	f.Ch = 0
	f.Line++
}

// Advance returns a new FilePos with the character index incremented
// by the provided amount.
func (f FilePos) Advance(n int) FilePos {
	return FilePos{f.Ch + n, f.Line}
}

// LineOnly returns a new FilePos with the character index set to zero,
// but retaining the original line number.
func (f FilePos) LineOnly() FilePos {
	return FilePos{0, f.Line}
}

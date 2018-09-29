package slrp

// Item represents an LR(0) item.
type Item struct {
	Nt   int
	Prod int
	Dot  int
}

// NewItem creates a new item.
func NewItem(n, p, d int) Item {
	return Item{n, p, d}
}

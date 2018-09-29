package slrp

type actionType int

const (
	actionShift actionType = iota
	actionReduce
	actionAccept
)

// Action represents an action in an SLR parsing table.
type Action struct {
	T actionType
	S int
}

// NewShift creates a new shift action to the specified state.
func NewShift(s int) Action {
	return Action{actionShift, s}
}

// NewReduce creates a new reduce action to the specified state.
func NewReduce(s int) Action {
	return Action{actionReduce, s}
}

// NewAccept creates a new accept action.
func NewAccept() Action {
	return Action{actionAccept, -1}
}

// IsShift tests if an action is a shift action.
func (a Action) IsShift() bool {
	return a.T == actionShift
}

// IsReduce tests if an action is a reduce action.
func (a Action) IsReduce() bool {
	return a.T == actionReduce
}

// IsAccept tests if an action is an accept action.
func (a Action) IsAccept() bool {
	return a.T == actionAccept
}

// Equals tests if an action is equal to another.
func (a Action) Equals(other Action) bool {
	return a.T == other.T && a.S == other.S
}

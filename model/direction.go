package model

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

// Direction is either Across or Down (according to the enumerated
// constants of those names).
type Direction string

const (
	ACROSS Direction = "A"
	DOWN   Direction = "D"
)

// ---------------------------------------------------------------------
// Methods
// ---------------------------------------------------------------------

// String returns a string representation of this object
func (dir Direction) String() string {
	var s string
	switch dir {
	case ACROSS:
		s = "across"
	case DOWN:
		s = "down"
	}
	return s
}

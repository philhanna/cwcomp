package model

import (
	"fmt"
	"strings"
)

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
// Functions
// ---------------------------------------------------------------------

// DirectionFromString parses a string for a direction. It will accept
// anything that starts with the direction letter value.
// Panics if the parameter is not a valid direction.
func DirectionFromString(s string) Direction {
	switch {
	case strings.HasPrefix(s, "A"), strings.HasPrefix(s, "a"):
		return ACROSS
	case strings.HasPrefix(s, "D"), strings.HasPrefix(s, "d"):
		return DOWN
	default:
		errmsg := fmt.Sprintf("%q is not a valid direction string", s)
		panic(errmsg)
	}
}

// ---------------------------------------------------------------------
// Methods
// ---------------------------------------------------------------------

// Other returns the other direction.
func (dir Direction) Other() Direction {
	var other Direction
	switch dir {
	case ACROSS:
		other = DOWN
	case DOWN:
		other = ACROSS
	}
	return other
}

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

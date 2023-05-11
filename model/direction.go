package model

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

// Direction is either Across or Down (according to the enumerated
// constants of those names).
type Direction byte

const (
	ACROSS Direction = iota
	DOWN
)

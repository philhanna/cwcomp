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

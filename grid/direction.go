package grid

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

// Direction is either Across or Down (according to the enumerated
// constants of those names).
type Direction string

const (
	Across Direction = "across"
	Down   Direction = "down"
)

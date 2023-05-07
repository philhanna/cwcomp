package grid

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

// Direction is either Across or Down (according to the enumerated
// constants of those names).
type Direction byte

const (
	Across Direction = iota
	Down
)

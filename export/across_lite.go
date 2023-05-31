package export

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

// AcrossLite is a representation of a puzzle in a standard interchange
// format.  This is proprietary format is defined and maintained by
// https://www.litsoft.com/.  It is described in
// https://www.litsoft.com/across/docs/AcrossTextFormat.pdf#31
type AcrossLite struct {
	title       string
	author      string
	copyright   string
	size        int
	grid        []string
	acrossClues []string
	downClues   []string
	notepad     []string
}

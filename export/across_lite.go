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

// ---------------------------------------------------------------------
// Constructor
// ---------------------------------------------------------------------

// NewAcrossLite creates a new, empty AcrossLite structure and returns a
// pointer to it.
func NewAcrossLite() *AcrossLite {
	p := new(AcrossLite)

	p.grid = make([]string, 0)
	p.acrossClues = make([]string, 0)
	p.downClues = make([]string, 0)
	p.notepad = make([]string, 0)

	return p
}

// ---------------------------------------------------------------------
// Methods
// ---------------------------------------------------------------------

// Import parses the text in an AcrossLite puzzle
func Import(filename string) (*AcrossLite, error) {
	p := NewAcrossLite()

	type ParsingState byte
	const (
		INIT ParsingState = iota
	)

	return p, nil
}

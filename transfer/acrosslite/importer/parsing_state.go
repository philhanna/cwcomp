package importer

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

// The parsing states of the finite state machine. These are used in
// determining the next state after each line is handled. There is a map
// of parsing states to handler functions in AcrossLiteImporter.
type ParsingState byte

const (
	UNKNOWN ParsingState = iota
	INIT
	LOOKING_FOR_TITLE
	READING_TITLE
	LOOKING_FOR_AUTHOR
	READING_AUTHOR
	LOOKING_FOR_COPYRIGHT
	READING_COPYRIGHT
	LOOKING_FOR_SIZE
	READING_SIZE
	LOOKING_FOR_GRID
	READING_GRID
	READING_ACROSS
	READING_DOWN
	READING_NOTEPAD
	DONE
)

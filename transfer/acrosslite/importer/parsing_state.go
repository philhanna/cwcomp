package importer

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

// The parsing states of the finite state machine
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

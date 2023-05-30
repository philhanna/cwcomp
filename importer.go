package cwcomp

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

// Importer is an interface that must be implemented by any source of
// puzzle data that can be imported (e.g., AcrossLite)
type Importer interface {
	Read() error
}
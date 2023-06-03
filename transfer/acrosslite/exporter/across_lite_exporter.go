package exporter

import (
	al "github.com/philhanna/cwcomp/transfer/acrosslite"
	"io"
)

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

// ---------------------------------------------------------------------
// Functions
// ---------------------------------------------------------------------

// Write will create an AcrossLite text format object and write it to
// the specified writer.
func Write(pal *al.AcrossLite, writer io.Writer) error {
	return nil // TODO implement me
}

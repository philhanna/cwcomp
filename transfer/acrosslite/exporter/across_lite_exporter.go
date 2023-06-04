package exporter

import (
	"io"
	"log"
	"strings"

	al "github.com/philhanna/cwcomp/transfer/acrosslite"
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
	log.Println("Entering Write")

	// Use a string builder to collect the output
	sb := strings.Builder{}

	// Write each section
	sb.WriteString(WriteAcrossPuzzle(pal) + "\n")
	sb.WriteString(WriteTitle(pal) + "\n")
	sb.WriteString(WriteAuthor(pal) + "\n")
	sb.WriteString(WriteCopyright(pal) + "\n")
	sb.WriteString(WriteSize(pal) + "\n")
	sb.WriteString(WriteGrid(pal) + "\n")
	sb.WriteString(WriteAcrossClues(pal) + "\n")
	sb.WriteString(WriteDownClues(pal) + "\n")
	if pal.GetNotepad() != "" {
		sb.WriteString(WriteNotepad(pal) + "\n")
	}

	// Write the results
	result := sb.String()
	blob := []byte(result)
	nBytes, err := writer.Write(blob)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Printf("Wrote %d bytes to output\n", nBytes)
	log.Println("Leaving Write")
	return nil
}

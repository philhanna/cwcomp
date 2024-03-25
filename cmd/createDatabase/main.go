package main

import (
	"io"
	"log"
	"os"

	"github.com/philhanna/cwcomp"
	"github.com/philhanna/cwcomp/model"
)

func main() {
	// Get the database file name
	config := cwcomp.Configuration
	filename := config.DATABASE.NAME
	log.Printf("Database file name = %v\n", filename)

	// Create a backup
	backup := filename + ".bak"
	os.Remove(backup)

	// Copy current to backup
	CopyFile(backup, filename)
	log.Printf("Created backup as %v\n", backup)

	// Delete current
	os.Remove(filename)

	// Create new
	model.CreateDatabase()
}

// CopyFile copies src into dst (Note the order of the arguments!)
func CopyFile(dst, src string) int64 {
	source, _ := os.Open(src)
	defer source.Close()

	destination, _ := os.Create(dst)
	defer destination.Close()

	nBytes, _ := io.Copy(destination, source)
	return nBytes
}

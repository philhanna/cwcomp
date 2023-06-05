package main

import (
	"io"
	"log"
	"os"

	"github.com/philhanna/cwcomp"
	"github.com/philhanna/cwcomp/model"
)

func main() {
	var (
		err error
	)

	// Get the database file name
	config := cwcomp.Configuration
	filename := config.DATABASE.NAME
	log.Printf("Database file name = %v\n", filename)

	// Create a backup
	backup := filename + ".bak"
	os.Remove(backup)

	// Copy current to backup
	_, err = CopyFile(backup, filename)
	if err != nil {
		log.Fatalf("Could not copy file: %v\n", err)
	}
	log.Printf("Created backup as %v\n", backup)

	// Delete current
	os.Remove(filename)

	// Create new
	model.CreateDatabase()
}

// CopyFile copies src into dst (Note the order of the arguments!)
func CopyFile(dst, src string) (int64, error) {
	source, err := os.Open(src)
	if os.IsNotExist(err) {
		return 0, nil
	}
	if err != nil {
		return 0, err
	}
	defer source.Close()

	destination, err := os.Create(dst)
	if err != nil {
		return 0, err
	}
	defer destination.Close()
	nBytes, err := io.Copy(destination, source)
	return nBytes, err
}

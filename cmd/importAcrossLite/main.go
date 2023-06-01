package main

import (
	"log"
	"path/filepath"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/philhanna/cwcomp"
	"github.com/philhanna/cwcomp/transfer/acrosslite/importer"
)

// This program will read a file in AcrossLite text format, create a
// puzzle from it, and save the puzzle in a database. The database will
// be created if it does not already exist.

const USERNAME = "Phil Hanna"

func main() {

	var err error

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Add the test user
	userid, err := createTestUser()
	if err != nil {
		log.Fatal(err)
	}

	// Load easy.txt into an AcrossLite structure
	filename := getFileName()
	acrossLite, err := importer.Parse(filename)
	if err != nil {
		log.Fatal(err)
	}

	// Import the AcrossLite structure into a Puzzle object
	puzzle, err := cwcomp.ImportPuzzle(acrossLite)
	if err != nil {
		log.Fatal(err)
	}

	// Save the puzzle in the database
	err = puzzle.SavePuzzle(userid)
	if err != nil {
		log.Fatal(err)
	}

	// Reload the puzzle and print it
	grid, err := cwcomp.LoadPuzzle(userid, "TODO")
	if err != nil {
		log.Fatal(err)
	}
	cwcomp.DumpPuzzle(grid)
}

// Returns the test file name abs(easy.txt)
func getFileName() string {
	dirname := filepath.Join("..", "..", "transfer", "acrosslite", "testdata")
	dirname, _ = filepath.Abs(dirname)
	filename := filepath.Join(dirname, "easy.txt")
	return filename
}

// Creates this user in the database
func createTestUser() (int, error) {
	var userid int
	var err error

	// Connect to the database
	con, err := cwcomp.Connect()
	defer con.Close()
	if err != nil {
		log.Println(err)
		return 0, err
	}

	// Create the test user
	sql := `INSERT INTO users
			(username, password, created)
			VALUES (?, ?, ?)`
	_, err = con.Exec(sql, USERNAME, getPassword(), getTime())
	if err != nil {
		log.Println(err)
		return 0, err
	}

	// Get the userid just added
	rows, _ := con.Query(`SELECT last_insert_rowid()`)
	defer rows.Close()

	rows.Next()
	err = rows.Scan(&userid)
	if err != nil {
		log.Println(err)
		return 0, err
	}

	log.Printf("Added test user\n")
	return userid, nil
}

// Returns the encrypted password
func getPassword() []byte {
	return cwcomp.Hash256(USERNAME)
}

// Returns the current time as a string
func getTime() string {
	return time.Now().Format(time.RFC3339)
}

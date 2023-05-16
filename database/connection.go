package database

import (
	"crypto/sha256"
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
	"github.com/philhanna/cwcomp"
)

// ---------------------------------------------------------------------
// Methods
// ---------------------------------------------------------------------

// Connect opens a connection to the cwcomp database.
func Connect() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", cwcomp.Configuration.DATABASE.NAME)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

// Hash256 returns the sha256 of the specified string
func Hash256(s string) []byte {
	blob := []byte(s)
	blob32 := sha256.Sum256(blob)
	newBlob := blob32[:]
	return newBlob
}

package database

import (
	"crypto/sha256"
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
	"github.com/philhanna/cwcomp"
)

// ---------------------------------------------------------------------
// Methods
// ---------------------------------------------------------------------

// Connect opens a connection to the cwcomp database.
func Connect(filename ...string) (*sql.DB, error) {
	var dbName string
	if len(filename) > 0 {
		dbName = filename[0]
	} else {
		dbName = cwcomp.Configuration.DATABASE.NAME
	}
	dataSourceName := fmt.Sprintf("file:%s?_foreign_keys=on", dbName)
	con, err := sql.Open("sqlite3", dataSourceName)
	if err != nil {
		return nil, err
	}
	err = con.Ping()
	if err != nil {
		return nil, err
	}
	return con, nil
}

// Hash256 returns the sha256 of the specified string
func Hash256(s string) []byte {
	blob := []byte(s)
	blob32 := sha256.Sum256(blob)
	newBlob := blob32[:]
	return newBlob
}

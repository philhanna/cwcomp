package model

import (
	"crypto/sha256"
	"database/sql"
	_ "embed"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
	"github.com/philhanna/cwcomp"
)

// The DDL to create the tables
//
//go:embed ddl.sql
var ddl string

// ---------------------------------------------------------------------
// Methods
// ---------------------------------------------------------------------

// Connect opens a connection to the cwcomp database.
func Connect() (*sql.DB, error) {
	dbName := cwcomp.Configuration.DATABASE.NAME
	dataSourceName := fmt.Sprintf("file:%s?_foreign_keys=on", dbName)
	con, _ := sql.Open("sqlite3", dataSourceName)
	return con, nil
}

// CreateDatabase creates the database, either the production one
// or the test one, depending on Configuration.DATABASE.NAME
func CreateDatabase() {

	// Connect to the database
	con, _ := Connect()

	// Run the DDL
	ddl := GetDDL()
	con.Exec(ddl)
	log.Printf("Created %v\n", cwcomp.Configuration.DATABASE.NAME)

}

// GetDDL returns a string containing the contents of the tables.sql file.
func GetDDL() string {
	return ddl
}

// Hash256 returns the sha256 of the specified string
func Hash256(s string) []byte {
	blob := []byte(s)
	blob32 := sha256.Sum256(blob)
	newBlob := blob32[:]
	return newBlob
}

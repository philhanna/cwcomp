package cwcomp

import (
	"crypto/sha256"
	"database/sql"
	_ "embed"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

//go:embed ddl.sql
var _contents string

// ---------------------------------------------------------------------
// Methods
// ---------------------------------------------------------------------

// Connect opens a connection to the cwcomp database.
func Connect() (*sql.DB, error) {
	dbName := Configuration.DATABASE.NAME
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

// CreateDatabase creates the database, either the production one
// or the test one, depending on Configuration.DATABASE.NAME
func CreateDatabase() {

	// Connect to the database
	con, err := Connect()
	if err != nil {
		log.Fatalf("Could not connect to db: %v\n", err)
	}

	// Run the DDL
	ddl := GetDDL()
	_, err = con.Exec(ddl)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Created %v\n", Configuration.DATABASE.NAME)

}

// GetDDL returns a string containing the contents of the tables.sql file.
func GetDDL() string {
	return _contents
}

// Hash256 returns the sha256 of the specified string
func Hash256(s string) []byte {
	blob := []byte(s)
	blob32 := sha256.Sum256(blob)
	newBlob := blob32[:]
	return newBlob
}

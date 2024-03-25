package model

import (
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
	dbName := cwcomp.GetConfiguration().DATABASE.NAME
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
	log.Printf("Created %v\n", cwcomp.GetConfiguration().DATABASE.NAME)

}

// GetDDL returns a string containing the contents of the tables.sql file.
func GetDDL() string {
	return ddl
}

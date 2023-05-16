package main

import (
	"database/sql"
	_ "embed"
	"io"
	"log"
	"os"
	"time"

	"github.com/philhanna/cwcomp"
	db "github.com/philhanna/cwcomp/database"
)

//go:embed tables.sql
var _contents string

func main() {
	var (
		err error
		con *sql.DB
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

	// Connect to the database
	con, err = db.Connect()
	if err != nil {
		log.Fatalf("Could not connect to db: %v\n", err)
	}

	// Run the DDL
	ddl := GetDDL()
	_, err = con.Exec(ddl)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Created %v\n", filename)

	// Create the admin user
	sql := `INSERT INTO users (username, password, created) values(?, ?, ?);`
	username := "admin"
	password := db.Hash256(username)
	created := time.Now().Format(time.RFC3339)
	_, err = con.Exec(sql, username, password, created)
	if err != nil {
		log.Fatalf("Could not add admin user: %v\n", err)
	}
	log.Printf("Added admin user\n")
}

// CopyFile copies src into dst (Note the order of the arguments!)
func CopyFile(dst, src string) (int64, error) {
	source, err := os.Open(src)
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

// GetDDL returns a string containing the contents of the tables.sql file.
func GetDDL() string {
	return _contents
}

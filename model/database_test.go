package model

import (
	"log"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/philhanna/cwcomp"
	"github.com/stretchr/testify/assert"
)

const TEST_USERID = 1

func createTestDatabase() {

	// Set configuration to use the test database name
	tmp := os.TempDir()
	config := cwcomp.Configuration
	dbName := filepath.Join(tmp, "cwcomp_test.db")
	if fileExists(dbName) {
		os.Remove(dbName)
	}
	config.DATABASE.NAME = dbName
	cwcomp.Configuration = config

	// Create the test database
	cwcomp.CreateDatabase()

	// Connect to the test database
	con, err := cwcomp.Connect()
	defer con.Close()

	// Create the test user
	sql := `INSERT INTO users (username, password, created) values(?, ?, ?);`
	username := "test"
	password := cwcomp.Hash256(username)
	created := time.Now().Format(time.RFC3339)
	_, err = con.Exec(sql, username, password, created)
	if err != nil {
		log.Fatalf("Could not add test user: %v\n", err)
	}
	log.Printf("Added test user\n")

}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
	   return false
	}
	return !info.IsDir()
 }

func setUp() {
	createTestDatabase()
}

func tearDown() {
	return
}

func TestGrid_SaveGrid(t *testing.T) {

	setUp()
	defer tearDown()

	// Create a new grid and populate it with words
	grid := getGoodGrid()
	type test struct {
		seq  int
		dir  Direction
		text string
	}
	testWords := []test{
		{1, ACROSS, "NOW"},
		{7, DOWN, "COW"},
		{8, ACROSS, "BLUE"},
		{20, DOWN, "HOW"},
	}
	for _, test := range testWords {
		word := grid.LookupWordByNumber(test.seq, test.dir)
		grid.SetText(word, test.text)
	}
	grid.SetGridName("Rhyme")

	_, err := grid.SaveGrid(TEST_USERID)
	assert.Nil(t, err)

	// Done with the grid
	grid.DeleteGrid(TEST_USERID, "Rhyme")
}

func TestGrid_GetGridList(t *testing.T) {
	setUp()
	defer tearDown()
	
	grid := getGoodGrid()
	gridNames, err := grid.GetGridList(TEST_USERID)
	assert.Nil(t, err)
	assert.Equal(t, 0, len(gridNames))
}

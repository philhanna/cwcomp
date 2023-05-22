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

// ---------------------------------------------------------------------
// Internal functions
// ---------------------------------------------------------------------

const test_userid = 1

// Creates a disposable database with a schema identical to the
// production database
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
	con, _ := cwcomp.Connect()
	defer con.Close()

	// Create the test user
	sql := `INSERT INTO users (username, password, created) values(?, ?, ?);`
	username := "test"
	password := cwcomp.Hash256(username)
	created := time.Now().Format(time.RFC3339)
	_, err := con.Exec(sql, username, password, created)
	if err != nil {
		log.Fatalf("Could not add test user: %v\n", err)
	}
	log.Printf("Added test user\n")

}

// Creates a test grid, populates some words and clues, and saves it.
func saveGrid(grid *Grid, gridName string) error {
	var err error

	// Create a new grid and populate it with words
	type test struct {
		seq  int
		dir  Direction
		text string
		clue string
	}
	testWords := []test{
		{1, ACROSS, "NOW", "At this time"},
		{7, DOWN, "COW", "Bovine"},
		{8, ACROSS, "BLUE", "Azure"},
		{20, DOWN, "HOW", "In what manner"},
	}
	for _, test := range testWords {
		word := grid.LookupWordByNumber(test.seq, test.dir)
		if err = grid.SetText(word, test.text); err != nil {
			return err
		}
		if err = grid.SetClue(word, test.clue); err != nil {
			return err
		}
	}

	grid.SetGridName(gridName)
	if err = grid.SaveGrid(test_userid); err != nil {
		return err
	}

	return nil
}

// ---------------------------------------------------------------------
// Test fixtures
// ---------------------------------------------------------------------

// Run at the beginning of every test function
func setUp() {
	createTestDatabase()
}

// Run at the end of every test function
func tearDown() {
	tmp := os.TempDir()
	dbName := filepath.Join(tmp, "cwcomp_test.db")
	if fileExists(dbName) {
		os.Remove(dbName)
	}
}

// ---------------------------------------------------------------------
// Unit tests
// ---------------------------------------------------------------------

// Tests whether the list of grid names obtained from the
// grid.GetGridList method is expected.
func TestGrid_GetGridList(t *testing.T) {
	setUp()
	defer tearDown()

	grid := getGoodGrid()
	gridNames := grid.GetGridList(test_userid)
	assert.Equal(t, 0, len(gridNames))
}

// Tests whether the specified grid name is already used.
func TestGrid_GridNameUsed(t *testing.T) {
	setUp()
	defer tearDown()

	var (
		err       error
		gridNames []string
		used      bool
	)
	grid := getGoodGrid()

	gridNames = grid.GetGridList(test_userid)
	assert.Equal(t, 0, len(gridNames))

	err = grid.SaveGrid(test_userid)
	assert.NotNilf(t, err, "save grid")

	used = grid.GridNameUsed(test_userid, "good9")
	assert.False(t, used)

	grid.SetGridName("good9")
	grid.SaveGrid(test_userid)
	gridNames = grid.GetGridList(test_userid)
	assert.Equal(t, 1, len(gridNames))

	used = grid.GridNameUsed(test_userid, "good9")
	assert.True(t, used)
}

// Tests whether a grid can be loaded correctly.
func TestGrid_LoadGrid(t *testing.T) {

	setUp()
	defer tearDown()

	var (
		err          error
		grid         *Grid
		reloadedGrid *Grid
	)

	_, err = LoadGrid(test_userid, "BOGUS")
	assert.NotNil(t, err)

	const gridName = "Rhyme"

	grid = getGoodGrid()
	err = saveGrid(grid, gridName)
	assert.Nil(t, err)

	// Reload the grid from the database
	reloadedGrid, err = LoadGrid(test_userid, gridName)
	assert.Nil(t, err)

	// Compare to the original grid

	assert.True(t, grid.Equal(reloadedGrid))
}

// Tests whether a grid can be saved correctly.
func TestGrid_SaveGrid(t *testing.T) {

	setUp()
	defer tearDown()

	const gridName = "Rhyme"
	grid := getGoodGrid()
	err := saveGrid(grid, gridName)
	assert.Nil(t, err)
	grid.DeleteGrid(test_userid, gridName)
}

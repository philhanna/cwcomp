package model

import (
	"log"
	"os"
	"path/filepath"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/philhanna/cwcomp"
	"github.com/philhanna/cwcomp/util"
	"github.com/stretchr/testify/assert"
)

// ---------------------------------------------------------------------
// Internal functions
// ---------------------------------------------------------------------

const TEST_USERID = 1

// Creates a disposable database with a schema identical to the
// production database
func createTestDatabase() {

	// Set configuration to use the test database name
	tmp := os.TempDir()
	config := cwcomp.GetConfiguration()
	dbName := filepath.Join(tmp, "cwcomp_test.db")
	if cwcomp.FileExists(dbName) {
		os.Remove(dbName)
	}
	config.DATABASE.NAME = dbName
	cwcomp.GetConfiguration = func() *cwcomp.Configuration{
		return config
	}

	// Create the test database
	CreateDatabase()

	// Connect to the test database
	con, _ := Connect()
	defer con.Close()

	// Create the test user
	sql := `INSERT INTO users (username, password, created) values(?, ?, ?);`
	username := "test"
	password := util.Hash256(username)
	created := time.Now().Format(time.RFC3339)
	_, err := con.Exec(sql, username, password, created)
	if err != nil {
		log.Fatalf("Could not add test user: %v\n", err)
	}
	log.Printf("Added test user\n")

}

// Creates a test puzzle, populates some words and clues, and saves it.
func savePuzzle(puzzle *Puzzle, puzzleName string) error {
	var err error

	// Create a new puzzle and populate it with words
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
		word := puzzle.LookupWordByNumber(test.seq, test.dir)
		if err = puzzle.SetText(word, test.text); err != nil {
			return err
		}
		if err = puzzle.SetClue(word, test.clue); err != nil {
			return err
		}
	}

	puzzle.SetPuzzleName(puzzleName)
	if err = puzzle.SavePuzzle(TEST_USERID); err != nil {
		return err
	}

	return nil
}

// ---------------------------------------------------------------------
// Test fixtures
// ---------------------------------------------------------------------

func runtest(f func(t *testing.T)) func(t *testing.T) {
	return func(t *testing.T) {
		setUp()
		f(t)
		tearDown()
	}
}

// Run at the beginning of every test function
func setUp() {
	createTestDatabase()
}

// Run at the end of every test function
func tearDown() {
	tmp := os.TempDir()
	dbName := filepath.Join(tmp, "cwcomp_test.db")
	if cwcomp.FileExists(dbName) {
		os.Remove(dbName)
	}
}

// ---------------------------------------------------------------------
// Unit tests
// ---------------------------------------------------------------------

// Tests whether the list of puzzle names obtained from the
// puzzle.GetPuzzleList method is expected.
func TestPuzzle_GetPuzzleList(t *testing.T) {
	runtest(func(*testing.T) {
		puzzle := getGoodPuzzle()
		puzzleNames := puzzle.GetPuzzleList(TEST_USERID)
		assert.Equal(t, 0, len(puzzleNames))
	})(t)
}

// Tests whether the specified puzzle name is already used.
func TestPuzzle_PuzzleNameUsed(t *testing.T) {
	runtest(func(*testing.T) {
		var (
			err         error
			puzzleNames []string
			used        bool
		)
		puzzle := getGoodPuzzle()

		puzzleNames = puzzle.GetPuzzleList(TEST_USERID)
		assert.Equal(t, 0, len(puzzleNames))

		err = puzzle.SavePuzzle(TEST_USERID)
		assert.NotNilf(t, err, "save puzzle")

		used = puzzle.PuzzleNameUsed(TEST_USERID, "good9")
		assert.False(t, used)

		puzzle.SetPuzzleName("good9")
		puzzle.SavePuzzle(TEST_USERID)
		puzzleNames = puzzle.GetPuzzleList(TEST_USERID)
		assert.Equal(t, 1, len(puzzleNames))

		used = puzzle.PuzzleNameUsed(TEST_USERID, "good9")
		assert.True(t, used)

	})(t)
}

// Tests whether a puzzle can be loaded correctly.
func TestPuzzle_LoadPuzzle(t *testing.T) {
	runtest(func(*testing.T) {
		var (
			err            error
			puzzle         *Puzzle
			reloadedPuzzle *Puzzle
		)

		_, err = LoadPuzzle(TEST_USERID, "BOGUS")
		assert.NotNil(t, err)

		const puzzleName = "Rhyme"

		puzzle = getGoodPuzzle()
		err = savePuzzle(puzzle, puzzleName)
		assert.Nil(t, err)

		// Reload the puzzle from the database
		reloadedPuzzle, err = LoadPuzzle(TEST_USERID, puzzleName)
		assert.Nil(t, err)

		// Compare to the original puzzle

		assert.True(t, puzzle.Equal(reloadedPuzzle))
	})(t)
}

func TestPuzzle_RenamePuzzle(t *testing.T) {
	runtest(func(*testing.T) {
		puzzle := getGoodPuzzle()
		puzzle.SetPuzzleName("foo")
		puzzle.SavePuzzle(TEST_USERID)
		err := puzzle.RenamePuzzle(TEST_USERID, "foo", "bar")
		assert.Nil(t, err)
		err = puzzle.RenamePuzzle(TEST_USERID, "baz", "bam")
		assert.NotNil(t, err)
	})(t)
}

// Tests whether a puzzle can be saved correctly.
func TestPuzzle_SavePuzzle(t *testing.T) {
	runtest(func(*testing.T) {
		const puzzleName = "Rhyme"
		puzzle := getGoodPuzzle()
		err := savePuzzle(puzzle, puzzleName)
		assert.Nil(t, err)
		puzzle.DeletePuzzle(TEST_USERID, puzzleName)
	})(t)
}

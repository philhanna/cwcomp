package model

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/philhanna/cwcomp"
)

// Letter value of a black cell in the cells table
const BLACK_CELL = "\x00"

// ---------------------------------------------------------------------
// Functions
// ---------------------------------------------------------------------

// DeletePuzzle deletes the specified puzzle
func (puzzle *Puzzle) DeletePuzzle(userid int, puzzlename string) error {
	con, _ := cwcomp.Connect()
	sql := `DELETE FROM puzzles WHERE userid=? AND puzzlename=?`
	_, err := con.Exec(sql, userid, puzzlename)
	return err
}

// GetPuzzleList returns a list of puzzles for the specified user.
func (puzzle *Puzzle) GetPuzzleList(userid int) []string {
	puzzlenames := make([]string, 0)

	// Get a database connection
	con, _ := cwcomp.Connect()
	defer con.Close()

	// Make the query
	sql := `
		SELECT		puzzlename
		FROM		puzzles
		WHERE		userid = ?
		ORDER BY	modified`
	rows, _ := con.Query(sql, userid)
	defer rows.Close()

	// Copy the names into the slice
	for {
		more := rows.Next()
		if !more {
			break
		}
		var puzzlename string
		rows.Scan(&puzzlename)
		puzzlenames = append(puzzlenames, puzzlename)
	}

	// Return the names of the puzzles

	return puzzlenames

}

// PuzzleNameUsed returns true if the specified puzzle name for this user is
// already saved in the database
func (puzzle *Puzzle) PuzzleNameUsed(userid int, puzzlename string) bool {
	// Open a connection
	con, _ := cwcomp.Connect()
	defer con.Close()

	// Query for this user/puzzlename
	sql := `SELECT COUNT(*) FROM puzzles WHERE userid=? AND puzzlename=?`
	rows, _ := con.Query(sql, userid, puzzlename)
	defer rows.Close()

	// Get the count of saved puzzles with that name
	rows.Next()
	count := 0
	rows.Scan(&count)

	// Set the return value
	used := (count > 0)

	return used
}

// LoadPuzzle reads puzzle data from the database and creates a Puzzle object from it.
func LoadPuzzle(userid int, puzzlename string) (*Puzzle, error) {

	// Connect to the database
	con, _ := cwcomp.Connect()
	defer con.Close()

	// Get the puzzle record from the database
	puzzleRows, _ := con.Query(`
		SELECT COUNT(*), id, n FROM puzzles WHERE userid=? AND puzzlename=?`,
		userid, puzzlename)
	defer puzzleRows.Close()

	var count, id, n int

	for puzzleRows.Next() {
		puzzleRows.Scan(&count, &id, &n)
		if count == 0 {
			return nil, fmt.Errorf("no puzzle named %q found", puzzlename)
		}
	}

	// Create an empty puzzle and begin populating it from the database
	puzzle := NewPuzzle(n)
	puzzle.SetPuzzleName(puzzlename)

	// Populate the cells (black cells and other)
	cellRows, _ := con.Query(`
		SELECT r, c, letter FROM cells WHERE id=?`,
		id)
	defer cellRows.Close()

	var r, c int
	var letter string

	for cellRows.Next() {
		cellRows.Scan(&r, &c, &letter)
		point := NewPoint(r, c)
		switch letter {
		case BLACK_CELL:
			puzzle.SetCell(point, NewBlackCell(point))
		default:
			puzzle.SetLetter(point, letter)
		}
	}

	// Renumber the puzzle to create the word and word number arrays
	puzzle.RenumberCells()

	// Populate the words
	wordRows, _ := con.Query(`
		SELECT r, c, dir, clue FROM words WHERE id=?`,
		id)
	defer wordRows.Close()

	for wordRows.Next() {

		var r, c int
		var dir, clue string

		wordRows.Scan(&r, &c, &dir, &clue)
		point := NewPoint(r, c)
		direction := DirectionFromString(dir)
		word := puzzle.LookupWord(point, direction)
		puzzle.SetClue(word, clue)
	}

	// Return the newly reconstituted puzzle with no error
	return puzzle, nil
}

// SavePuzzle adds or updates a record for this puzzle in the database
func (puzzle *Puzzle) SavePuzzle(userid int) error {
	var (
		err  error
		id   int
		rows *sql.Rows
		sql  string
	)

	// Ensure the puzzle has been named
	puzzlename := puzzle.GetPuzzleName()
	if puzzlename == "" {
		err = fmt.Errorf("cannot save a puzzle without a name")
		return err
	}

	// Open a connection
	con, _ := cwcomp.Connect()
	defer con.Close()

	// Delete any previous records for this puzzle
	// (should do cascading delete to other tables)
	sql = `DELETE FROM puzzles WHERE puzzlename = ?`
	con.Exec(sql, puzzlename)

	// Save the data in the puzzles table
	// and get the generated puzzle ID
	sql = `
		INSERT INTO puzzles(userid, puzzlename, created, modified, n)
		VALUES(?, ?, ?, ?, ?)
		`
	timenow := time.Now()
	created := timenow.Format(time.RFC3339)
	modified := created
	con.Exec(sql, userid, puzzlename, created, modified, puzzle.n)
	rows, _ = con.Query("SELECT last_insert_rowid()")
	rows.Next()
	rows.Scan(&id) // Return this later

	// Save the cell data in the cells table
	sql = `
		INSERT INTO cells(id, r, c, letter)
		VALUES(?, ?, ?, ?)
		`
	for cell := range puzzle.CellIterator() {
		var (
			r      int
			c      int
			letter string
		)
		r, c = cell.GetPoint().r, cell.GetPoint().c
		switch typedCell := cell.(type) {
		case LetterCell:
			letter = typedCell.letter
		case BlackCell:
			letter = BLACK_CELL
		}
		con.Exec(sql, id, r, c, letter)
	}

	// Save the word data in the words table
	sql = `
		INSERT INTO words(id, r, c, dir, length, clue)
		VALUES(?, ?, ?, ?, ?, ?)
		`
	for _, word := range puzzle.words {
		point := word.point
		con.Exec(sql, id, point.r, point.c, word.direction, word.length, word.clue)
	}

	// Successful completion
	return nil
}

// RenamePuzzle renames a puzzle in the database
func (puzzle *Puzzle) RenamePuzzle(userid int, oldPuzzleName, newPuzzleName string) error {

	// See if there is a puzzle by the old name

	if !puzzle.PuzzleNameUsed(userid, oldPuzzleName) {
		return fmt.Errorf("no puzzle found for name=%q", oldPuzzleName)
	}

	// Connect to the database

	con, _ := cwcomp.Connect()
	defer con.Close()

	// Do the update and return no error
	con.Exec(`
		UPDATE	puzzles
		SET		puzzlename=?
		WHERE	userid=?
		AND		puzzlename=?`,
		newPuzzleName, userid, oldPuzzleName)
	return nil
}

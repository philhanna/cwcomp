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

// DeleteGrid deletes the specified grid
func (grid *Grid) DeleteGrid(userid int, gridname string) error {
	con, _ := cwcomp.Connect()
	sql := `DELETE FROM grids WHERE userid = ? AND gridname = ?`
	_, err := con.Exec(sql, userid, gridname)
	return err
}

// GetGridList returns a list of grids for the specified user.
func (grid *Grid) GetGridList(userid int) []string {
	gridnames := make([]string, 0)

	// Get a database connection
	con, _ := cwcomp.Connect()
	defer con.Close()

	// Make the query
	sql := `
		SELECT		gridname
		FROM		grids
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
		var gridname string
		rows.Scan(&gridname)
		gridnames = append(gridnames, gridname)
	}

	// Return the names of the grids

	return gridnames

}

// GridNameUsed returns true if the specified grid name for this user is
// already saved in the database
func (grid *Grid) GridNameUsed(userid int, gridname string) bool {
	// Open a connection
	con, _ := cwcomp.Connect()
	defer con.Close()

	// Query for this user/gridname
	sql := `SELECT COUNT(*) FROM grids WHERE userid=? AND gridname=?`
	rows, _ := con.Query(sql, userid, gridname)
	defer rows.Close()

	// Get the count of saved grids with that name
	rows.Next()
	count := 0
	rows.Scan(&count)

	// Set the return value
	used := (count > 0)

	return used
}

// LoadGrid reads grid data from the database and creates a Grid object from it.
func LoadGrid(userid int, gridname string) (*Grid, error) {

	// Connect to the database
	con, _ := cwcomp.Connect()
	defer con.Close()

	// Get the grid record from the database
	gridRows, _ := con.Query(`
		SELECT COUNT(*), gridid, n FROM grids WHERE userid=? AND gridname=?`,
		userid, gridname)
	defer gridRows.Close()

	var count, gridid, n int

	for gridRows.Next() {
		gridRows.Scan(&count, &gridid, &n)
		if count == 0 {
			return nil, fmt.Errorf("no grid named %q found", gridname)
		}
	}

	// Create an empty grid and begin populating it from the database
	grid := NewGrid(n)
	grid.SetGridName(gridname)

	// Populate the cells (black cells and other)
	cellRows, _ := con.Query(`
		SELECT r, c, letter FROM cells WHERE gridid=?`,
		gridid)
	defer cellRows.Close()

	var r, c int
	var letter string

	for cellRows.Next() {
		cellRows.Scan(&r, &c, &letter)
		point := NewPoint(r, c)
		switch letter {
		case BLACK_CELL:
			grid.SetCell(point, NewBlackCell(point))
		default:
			grid.SetLetter(point, letter)
		}
	}

	// Renumber the grid to create the word and word number arrays
	grid.RenumberCells()

	// Populate the words
	wordRows, _ := con.Query(`
		SELECT r, c, dir, clue FROM words WHERE gridid=?`,
		gridid)
	defer wordRows.Close()

	for wordRows.Next() {

		var r, c int
		var dir, clue string

		wordRows.Scan(&r, &c, &dir, &clue)
		point := NewPoint(r, c)
		direction := DirectionFromString(dir)
		word := grid.LookupWord(point, direction)
		grid.SetClue(word, clue)
	}

	// Return the newly reconstituted grid with no error
	return grid, nil
}

// SaveGrid adds or updates a record for this grid in the database
func (grid *Grid) SaveGrid(userid int) error {
	var (
		err    error
		gridid int
		rows   *sql.Rows
		sql    string
	)

	// Ensure the grid has been named
	gridname := grid.GetGridName()
	if gridname == "" {
		err = fmt.Errorf("cannot save a grid without a name")
		return err
	}

	// Open a connection
	con, _ := cwcomp.Connect()
	defer con.Close()

	// Delete any previous records for this grid
	// (should do cascading delete to other tables)
	sql = `DELETE FROM grids WHERE gridname = ?`
	con.Exec(sql, gridname)

	// Save the data in the grids table
	// and get the generated grid ID
	sql = `
		INSERT INTO grids(userid, gridname, created, modified, n)
		VALUES(?, ?, ?, ?, ?)
		`
	timenow := time.Now()
	created := timenow.Format(time.RFC3339)
	modified := created
	con.Exec(sql, userid, gridname, created, modified, grid.n)
	rows, _ = con.Query("SELECT last_insert_rowid()")
	rows.Next()
	rows.Scan(&gridid) // Return this later

	// Save the cell data in the cells table
	sql = `
		INSERT INTO cells(gridid, r, c, letter)
		VALUES(?, ?, ?, ?)
		`
	for cell := range grid.CellIterator() {
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
		con.Exec(sql, gridid, r, c, letter)
	}

	// Save the word data in the words table
	sql = `
		INSERT INTO words(gridid, r, c, dir, length, clue)
		VALUES(?, ?, ?, ?, ?, ?)
		`
	for _, word := range grid.words {
		point := word.point
		con.Exec(sql, gridid, point.r, point.c, word.direction, word.length, word.clue)
	}

	// Successful completion
	return nil
}

// RenameGrid renames a grid in the database
func (grid *Grid) RenameGrid(userid int, oldGridName, newGridName string) error {
	
	// See if there is a grid by the old name

	if !grid.GridNameUsed(userid, oldGridName) {
		return fmt.Errorf("no grid found for name=%q", oldGridName)
	}
	
	// Connect to the database

	con, _ := cwcomp.Connect()
	defer con.Close()

	// Do the update and return no error
	con.Exec(`
		UPDATE	grids
		SET		gridname=?
		WHERE	userid=?
		AND		gridname=?`,
			newGridName, userid, oldGridName)
	return nil
}
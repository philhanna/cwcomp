package model

import (
	"errors"
	"fmt"

	db "github.com/philhanna/cwcomp/database"
)

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

type GridTable struct {
}

// ---------------------------------------------------------------------
// Functions
// ---------------------------------------------------------------------

// GetGridList returns a list of grids for the specified user.
func (grid *Grid) GetGridList(userid int) ([]string, error) {

	gridnames := make([]string, 0)

	// Get a database connection
	con, err := db.Connect()
	if err != nil {
		errmsg := fmt.Sprintf("Unable to connect to database: %v\n", err)
		err := errors.New(errmsg)
		return gridnames, err
	}
	defer con.Close()

	// Make the query
	sql := `
		SELECT		gridname
		FROM		grids
		WHERE		userid = ?
		ORDER BY	modified`
	rows, err := con.Query(sql, userid)
	if err != nil {
		errmsg := fmt.Sprintf("Unable to query grids table: %v\n", err)
		err := errors.New(errmsg)
		return gridnames, err
	}
	defer rows.Close()

	// Copy the names into the slice
	for {
		more := rows.Next()
		if !more {
			break
		}
		var gridname string
		err = rows.Scan(&gridname)
		if err != nil {
			errmsg := fmt.Sprintf("Unable to read gridname from grids table: %v\n", err)
			err := errors.New(errmsg)
			return gridnames, err
		}
		gridnames = append(gridnames, gridname)
	}

	// Return the names of the grids

	return gridnames, nil

}

// Save adds or updates a record for this grid in the database
func (grid *Grid) Save(userid int) error {
	
	// Ensure the grid has been named
	if grid.GetGridName() == "" {
		errmsg := "Cannot save a grid without a name"
		err := errors.New(errmsg)
		return err
	}
	
	// Open a connection
	con, _ := db.Connect()
	defer con.Close()

	// Delete any previous records for this grid
	gridnames, err := grid.GetGridList(userid)
	if err != nil {
		return err
	}
	if len(gridnames) > 0 {
		for _, gridname := range gridnames {
			var sql string
			sql = `DELETE FROM CELLS WHERE gridid = ?`
			_, err := con.Exec(sql, userid, gridname)
			if err != nil {
				errmsg := fmt.Sprintf("Could not delete grid %s: %v\n", gridname, err)
				err := errors.New(errmsg)
				return err
			}
		}
	}


	// TODO finish me
	return nil
}
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

	gridNames := make([]string, 0)

	// Get a database connection
	con, err := db.Connect()
	if err != nil {
		errmsg := fmt.Sprintf("Unable to connect to database: %v\n", err)
		err := errors.New(errmsg)
		return gridNames, err
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
		return gridNames, err
	}
	defer rows.Close()

	// Copy the names into the slice
	for {
		more := rows.Next()
		if !more {
			break
		}
		var gridName string
		err = rows.Scan(&gridName)
		if err != nil {
			errmsg := fmt.Sprintf("Unable to read gridname from grids table: %v\n", err)
			err := errors.New(errmsg)
			return gridNames, err
		}
		gridNames = append(gridNames, gridName)
	}

	// Return the names of the grids

	return gridNames, nil

}

package rest

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/philhanna/cwcomp/model"
)

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

type PuzzleEntries struct {
	Entries []PuzzleEntry `json:"entries"`
}

type PuzzleEntry struct {
	ID         int    `json:"id"`
	Puzzlename string `json:"puzzlename"`
	Modified   string `json:"modified"`
}

// ---------------------------------------------------------------------
// Functions
// ---------------------------------------------------------------------

// PuzzlesHandler serves REST requests for:
//
//   - GET:  Returns a list of puzzles for this user
//   - POST: Adds a new puzzle
func PuzzlesHandler(w http.ResponseWriter, r *http.Request) {

	log.Println("Entering PuzzlesHandler")

	// Get the session
	session, err := GetSession(w, r)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	log.Printf("Found session %v for userid %d\n", session.ID, session.USERID)
	userid := session.USERID

	// Send back the list of puzzle names for this user
	con, _ := model.Connect()
	defer con.Close()

	rows, err := con.Query(`
		SELECT		id, puzzlename, modified
		FROM		puzzles
		WHERE		userid=?
		ORDER BY	3 DESC, 2`, userid)
	defer rows.Close()
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	entries := new(PuzzleEntries)
	for rows.Next() {
		entry := PuzzleEntry{}
		rows.Scan(&entry.ID, &entry.Puzzlename, &entry.Modified)
		entries.Entries = append(entries.Entries, entry)
	}

	// Convert the slice to JSON
	jsonBlob, err := json.MarshalIndent(entries, "", "  ")
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBlob)

	log.Println("Leaving PuzzlesHandler")

}

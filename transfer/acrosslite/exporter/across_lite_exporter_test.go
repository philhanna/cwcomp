package exporter

import (
	"bufio"
	"encoding/json"
	"os"
	"testing"
	"time"

	al "github.com/philhanna/cwcomp/transfer/acrosslite"
	"github.com/stretchr/testify/assert"
)

func getTestStructure() (*al.AcrossLite, error) {

	pal := al.NewAcrossLite()
	n := 9
	pal.SetTitle("My title")
	// pal.SetAuthor("Phil Hanna")
	pal.SetCopyright("2023")
	pal.SetSize(n)
	cellStrings := []string{
		".NOW.   C",
		"BLUE.   O",
		"    .   W",
		"        .",
		"...   ...",
		".        ",
		"H   .    ",
		"O   .    ",
		"W   .    .",
	}
	for r := 1; r <= n; r++ {
		s := cellStrings[r-1]
		for c := 1; c <= n; c++ {
			letter := s[c-1]
			pal.SetCell(r, c, letter)
		}
	}
	clueMap := map[int]string{
		1: "Not then but",
		4: "Ends in C",
		8: "Not green but",
	}
	pal.SetAcrossClues(clueMap)

	clueMap = map[int]string{
		1:  "Friend of labor?",
		2:  "Pain expression",
		3:  "Starts with WE",
		7:  "Sacred animal",
		20: "Not why but",
	}
	pal.SetDownClues(clueMap)

	if false {
		type Extra struct {
			Created  string
			Modified string
		}
		const ISO8601 = "2006-01-02T15:04:05.999999"
		created := time.Now()
		modified := time.Now().Add(3 * 24 * time.Hour)
		extra := Extra{
			Created:  created.Format(ISO8601),
			Modified: modified.Format(ISO8601),
		}
		jsonblob, err := json.Marshal(extra)
		if err != nil {
			return pal, err
		}
		jsonstr := string(jsonblob)
		pal.SetNotepad(jsonstr)
	}
	return pal, nil
}

func TestExport(t *testing.T) {
	pal, err := getTestStructure()
	assert.Nil(t, err)

	w := bufio.NewWriter(os.Stdout)
	Write(pal, w)
	w.Flush()
}

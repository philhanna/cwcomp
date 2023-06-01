package main

import (
	"bufio"
	"log"
	"os"
	"path/filepath"
	"strings"

	_ "github.com/mattn/go-sqlite3"
	"github.com/philhanna/cwcomp"
	"github.com/philhanna/cwcomp/transfer/acrosslite/importer"
)

// This program will read a file in AcrossLite text format, create a
// puzzle from it, and dump the grid in an SVG in /tmp
func main() {

	var err error

	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Load easy.txt into an AcrossLite structure

	reader := strings.NewReader(SAMPLE_DATA)
	bufferedReader := bufio.NewReader(reader)
	acrossLite, err := importer.Parse(bufferedReader)
	if err != nil {
		log.Fatal(err)
	}

	// Import the AcrossLite structure into a Puzzle object
	puzzle, err := cwcomp.ImportPuzzle(acrossLite)
	if err != nil {
		log.Fatal(err)
	}

	// Dump the puzzle and clues
	cwcomp.DumpPuzzle(puzzle)

	// Save the SVG image of the puzzle
	filename := filepath.Join(os.TempDir(), "across_lite.svg")
	svg := cwcomp.NewSVGFromPuzzle(puzzle)
	svgString := svg.GenerateSVG()
	svgBytes := []byte(svgString)
	os.WriteFile(filename, svgBytes, 0644)

	log.Printf("SVG written to %v\n", filename)

}

// Sample file from https://www.litsoft.com/across/docs/AcrossTextFormat.pdf#31
var SAMPLE_DATA = `<ACROSS PUZZLE>
<TITLE> 
	Politics: Who, what, where and why
<AUTHOR> 
	Created by Avalonian
<COPYRIGHT> 
	Literate Software Systems
<SIZE> 
	15x15
<GRID>
	FATE.AWASH.AWOL
	LIES.CURIO.SHOE
	ELECTORATE.SIZE
	ASS.ERST.DIETED
	...CENT.HOSTESS
	REFITS.JEWISH..
	ARITH.KERNS.OAF
	NILE.ANNES.DUPE
	DEI.OVENS.LOSER
	..BODILY.RACERS
	GLUTEAL.PEPS...
	RESIST.SLUE.SKI
	OTTO.REPUBLICAN
	OMES.IRATE.RAMS
	MERE.XENON.ABET
<ACROSS>    			 
	Destiny
	Above water, barely
	Deserter
	True ------ : Arnold S. movie
	Novel or rare item
	If the ---- fits, ...
	Favorite group of 53 across
	EEE with 16 across
	Midsummer's Night Dream character
	Formerly
	Ate sparingly
	Monetary unit
	She entertains guests
	Make ready for use again
	Yiddish, informally
	Math. subject
	Parts of typeset characters
	Clod
	Egyptian river
	Bancroft and Archer
	Fool
	God in latin
	Pizza place fixtures
	Sometimes a dieter is one
	Physical
	Some kinds of snakes
	Of posterior muscles
	Raises one's spirits
	Block
	Swing about
	--- slopes
	A dog's name
	GOP member
	Of masses
	Angry
	Strikes violently
	Just
	Inert gas
	Encourage
<DOWN>
	Biting insect
	Troubles
	Golf equipments
	Computer keyboard key
	Oak seed or fruit
	Sausage
	I smell ----
	Squat
	Community dances
	Owned property
	House on Pennsylvania Ave.
	Seeps 
	City NE of Manchester
	Subject of dentistry
	Wife of Asiris
	Use reference
	"------, Johnny!"
	South African currency unit
	Canal or Lake
	Delaying tactic of 53 across
	Spinning ------
	Toll
	One who mimics
	Bearers : comb. form
	Female pilot
	Medics
	Homages
	Coat part
	Idle
	Type of sandwich
	Clean and care for
	"----- through!"
	Smallest planet of the Sun
	Cover
	One who works despite a strike
	Glacial ridge
	Org.
	Before
	Tax break savings account
`

package main

import (
	"archive/zip"
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/philhanna/cwcomp/model"
	"github.com/philhanna/cwcomp/svg"

	imp "github.com/philhanna/cwcomp/acrosslite/importer"
)

const USERID = 1

var (
	OPTION_SVG      bool
	OPTION_DATABASE bool
	zipFileName     string
)

// This program will extract AcrossLite text files from a zip file,
// create puzzles from them, and save them in the database.  It can
// also generate an SVG with the puzzle image in /tmp.
func main() {

	const (
		usage = `usage: importPythonDatabase [OPTIONS] [FILENAME]

Imports puzzles from the old Python database in a zip file. The puzzles
are in the AcrossLite text format. The format is specified at
https://www.litsoft.com/across/docs/AcrossTextFormat.pdf

positional arguments:
  filename                 path to zipfile containing AcrossLite text files

options:
  -h, --help               display this help text and exit
  -s, --svg                create SVG images for each puzzle
  -d, --database           save puzzles in the database
`
	)

	// Parse the command line arguments
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, usage)
	}
	flag.BoolVar(&OPTION_SVG, "s", false, "generate SVG")
	flag.BoolVar(&OPTION_SVG, "svg", false, "generate SVG")
	flag.BoolVar(&OPTION_DATABASE, "d", false, "Save in database")
	flag.BoolVar(&OPTION_DATABASE, "database", false, "Save in database")
	flag.Parse()

	if flag.NArg() == 0 {
		log.Fatal("No zip file name specified")
	}

	// Check for the existence of the zip file
	zipFileName = flag.Arg(0)
	_, err := os.Stat(zipFileName)
	if err != nil {
		log.Fatalf("Zip file error = %#v\n", err)
	}

	// Open the zip file
	log.Printf("Reading entries from %s...\n", zipFileName)
	zipFile, err := zip.OpenReader(zipFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer zipFile.Close()

	// Read each entry in the zip file
	for _, entry := range zipFile.File {
		log.Printf("Importing %s\n", entry.Name)

		// Open the zip entry
		file, err := entry.Open()
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		// Read its contents into a byte slice
		data, err := ioutil.ReadAll(file)
		if err != nil {
			log.Fatal(err)
		}

		// Parse it as AcrossLite text and create the structure
		reader := bytes.NewReader(data)
		pal, err := imp.Parse(reader)
		if err != nil {
			log.Fatal(err)
		}

		// Now create the puzzle itself
		puzzle, err := model.ImportPuzzle(pal)
		if err != nil {
			log.Fatal(err)
		}

		// If --svg was requested, write the SVG
		if OPTION_SVG {
			svgFileName := strings.ReplaceAll(entry.Name, ".txt", ".svg")
			filename := filepath.Join(os.TempDir(), svgFileName)

			log.Printf("Creating SVG in %s\n", filename)
			cells := model.PuzzleToSimpleMatrix(puzzle)
			image := svg.NewSVG(cells)
			svgString := image.GenerateSVG()
			svgBytes := []byte(svgString)
			os.WriteFile(filename, svgBytes, 0644)
		}

		// If --database was requested, save puzzle in database
		if OPTION_DATABASE {
			re := regexp.MustCompile(`\d+_([^.]+).txt`)
			filename := entry.Name
			group := re.FindStringSubmatch(filename)
			puzzleName := group[1]

			puzzle.SetPuzzleName(puzzleName)
			err = puzzle.SavePuzzle(USERID)
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("Saved %s puzzle in database\n", puzzleName)
		}
	}
}

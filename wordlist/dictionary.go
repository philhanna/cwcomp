package wordlist

import (
	"bufio"
	_ "embed"
	"log"
	"strings"
	"time"
)

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

// Dictionary is the in-memory word list.  There is an entry in this map
// for each word encountered in the words.txt file
type Dictionary []string

// ---------------------------------------------------------------------
// Global variables
// ---------------------------------------------------------------------

var dictionary = make(Dictionary, 0)

// words is the embedded contents of the raw words.txt file
//
//go:embed words.txt
var words string

// ---------------------------------------------------------------------
// Functions
// ---------------------------------------------------------------------

// init loads the dictionary at startup.  I was going to get fancy and
// use a goroutine to do this, but it takes only less than a second on
// my Linux machine (as opposed to 5-6 seconds in the Python version!)
func init() {
	stime := time.Now()
	count := 0
	reader := strings.NewReader(words)
	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		line := scanner.Text()
		dictionary = append(dictionary, line)
		count++
	}
	duration := time.Since(stime)
	log.Printf("Dictionary loaded with %d words in %v\n", count, duration)
}
package cwcomp

import (
	"encoding/json"
	"regexp"
)

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

// Constraint is a structure that describes constraints imposed on this
// word by its crossing words.
type Constraint struct {
	//
	// Index within the main word (1, 2, ..., length)
	//
	Pos int `json:"pos"`
	//
	// Index within the crossing word (1, 2, ..., length)
	//
	Index int `json:"index"`
	//
	// Letter at index
	//
	Letter string `json:"letter"`
	//
	// Text of crossing word
	//
	Text string `json:"text"`
	//
	// Word number of crossing word
	//
	Seq int `json:"seq"`
	//
	// Direction of crossing word
	//
	Dir Direction `json:"dir"`
	//
	// Regular expression for possibilities for this cell
	//
	Pattern string `json:"pattern"`
	//
	// Number of words that match that pattern
	//
	NChoices int `json:"nChoices"`
}

// ---------------------------------------------------------------------
// Methods
// ---------------------------------------------------------------------

// GetConstraints finds the constraints imposed on this word by its
// crossing words.
func (puzzle *Puzzle) GetConstraints(word *Word) []*Constraint {

	// Create a slice to return the constraints we find for each
	// crossing word.
	constraints := make([]*Constraint, word.length)

	// Iterate through the main word point by point
	for index := 0; index < word.length; index++ {

		// Figure out the point in the word at this index
		var point Point
		switch word.direction {
		case ACROSS:
			point = NewPoint(word.point.r, word.point.c+index)
		case DOWN:
			point = NewPoint(word.point.r+index, word.point.c)
		}

		// Lookup the crossing word
		crosser := puzzle.LookupWord(point, word.direction.Other())
		crosserWordNumber := puzzle.GetWordNumber(crosser)

		// Create an empty constraint object pointer
		cst := new(Constraint)

		// Set the index (1, 2, ..., ) within the main word at which
		// this crossing occurs.
		cst.Pos = index + 1

		// Set the index (1, 2, ..., ) within the crossing word at which
		// this crossing occurs.
		crossIndex := 0
		for crossPoint := range puzzle.WordIterator(crosserWordNumber.point, crosser.direction) {
			crossIndex++
			if crossPoint == point {
				cst.Index = crossIndex
			}
		}

		// Get the letter at that point
		cst.Letter = puzzle.GetLetter(point)

		// Get the text of the crossing word
		cst.Text = puzzle.GetText(crosser)

		// Get the sequence number of the crossing word
		cst.Seq = crosserWordNumber.seq

		// Get the direction of the crossing word
		cst.Dir = crosser.direction

		// Get the pattern for a regular expression for the possible
		// choices of the crossing word.
		pattern := cst.Text
		re := regexp.MustCompile(` `)
		pattern = re.ReplaceAllLiteralString(pattern, ".")

		// Examine all the words that match this pattern. Count them and
		// store the number in cst.NChoices. Store the letters of the
		// matching words at the crossing point in a set (from which we
		// will figure out a regular expression)
		cst.NChoices = 0
		letterSet := make(map[byte]bool)
		for matcher := range GetMatchingWords(pattern, make(chan struct{})) {
			cst.NChoices++
			letter := matcher[cst.Index-1]
			letterSet[letter] = true
		}

		// Now take all the letters in the set and make a regular expression
		// that describes every one.
		letterList := make([]byte, 0)
		for letter := range letterSet {
			letterList = append(letterList, letter)
		}
		letterString := string(letterList)
		cst.Pattern = Regexp(letterString)

		// Special case - crossing word is not in the dictionary
		if cst.Pattern == "" {
			cst.Pattern = cst.Letter
			cst.NChoices = 1
		}

		// Add the constraint to the array
		constraints[index] = cst
	}

	return constraints
}

// ToJSON returns the JSON representation of a constraints object
func (cst *Constraint) ToJSON() string {
	jsonBlob, _ := json.Marshal(*cst)
	jsonstr := string(jsonBlob)
	return jsonstr
}

package cwcomp

/***********************************************************************
 * LetterList is a utility class that forms a regular expression
 * from a list of letters
 **********************************************************************/

import (
	"sort"
	"strings"

	"golang.org/x/exp/slices"
)

const ALPHABET = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

// Used by GetBlocks
type Block struct {
	First int
	Last  int
}

// ---------------------------------------------------------------------
// Functions
// ---------------------------------------------------------------------

// Blocks is a generator that will take a list of indices
// and yield pairs of (first, last) indices of sublists that
// are consecutive integers. For example:
//
// ilist = [3, 2, 3, 4, 5, 1, 1, 2, 3] yields 4 pairs:
// (3, 3)
// (2, 5) - 2 through 5
// (1, 1)
// (1, 3) - 1 through 3
//
// Call it like this:
//
//	 c := Blocks(ilist) 	// c is the channel
//	 for pair := range c {	// but you only need to use range
//
//		   first := pair.First
//		   last := pair.Last
//		   ...
//		}
func Blocks(ilist []int) <-chan Block {
	c := make(chan Block)
	go func() {
		defer close(c)
		if len(ilist) < 1 {
			return
		}
		first := ilist[0]
		last := first
		for i := 1; i < len(ilist); i++ {
			x := ilist[i]
			if x == last+1 { // Consecutive
				last = x
			} else {
				c <- Block{first, last}
				first = x
				last = first
			}
		}
		c <- Block{first, last} // Last one
	}()
	return c
}

// Complement is a function that will return a list of indices
// into the ALPHABET string that are not in the list passed
// as a parameter
func Complement(ilist []int) []int {
	xa := []int{}
	for i := 0; i < len(ALPHABET); i++ {
		p := slices.Index(ilist, i)
		if p == -1 {
			xa = append(xa, i)
		}
	}
	return xa
}

// Pattern creates part of a regular expression string
// that will go between brackets. For example, if the input
// list is []int{2, 3, 4, 5, 25}, it will return "C-DF"
func Pattern(ilist []int) string {
	pattern := ""
	c := Blocks(ilist)
	for pair := range c {
		f, l := pair.First, pair.Last
		pattern += string(ALPHABET[f])
		if f < l {
			pattern += "-"
			pattern += string(ALPHABET[l])
		}
	}
	return pattern
}

// Regexp will return a regular expression representing
// the list of letters
func Regexp(letters string) string {

	// Make a list of the distinct integer indices of the letters
	// pointing to ALPHABET. Here we emulate a Python set().
	makeiset := func() []int {
		m := make(map[int]bool)
		iset := make([]int, 0)
		for _, ch := range letters {
			x := strings.Index(ALPHABET, string(ch))
			if x > -1 {
				iset = append(iset, x)
			}
		}
		for _, r := range iset {
			m[r] = true
		}
		iset = make([]int, 0)
		for k := range m {
			iset = append(iset, k)
		}
		sort.Ints(iset)
		return iset
	}
	ilist := makeiset()

	// Easy cases: empty, single letter, or all letters
	if len(ilist) == 0 {
		return ""
	}
	if len(ilist) == 1 {
		x := ilist[0]
		return string(ALPHABET[x])
	}
	if len(ilist) == len(ALPHABET) {
		return "."
	}

	// Not easy cases...
	pattern1 := "[" + Pattern(ilist) + "]"
	pattern2 := "[^" + Pattern(Complement(ilist)) + "]"
	if len(pattern1) <= len(pattern2) {
		return pattern1
	} else {
		return pattern2
	}
}

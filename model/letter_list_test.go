package model

import (
	"slices"
	"testing"

	"github.com/philhanna/cwcomp"
	"github.com/stretchr/testify/assert"
)

func TestGetPackage(t *testing.T) {
	want := "model"
	have := cwcomp.GetPackageName()
	assert.Equal(t, want, have)
}

func TestLetterList_Complement(t *testing.T) {

	// Helper function to make a list of consecutive integers.
	// Operates like a for loop
	straight := func(from int, to int, step int) []int {
		list := make([]int, 0)
		for i := from; i < to; i += step {
			list = append(list, i)
		}
		return list
	}

	// Array of test cases
	testCases := []struct {
		ilist []int
		want  []int
	}{
		{straight(0, 2+1, 1), straight(3, 25+1, 1)},
		{straight(0, 24+1, 2), straight(1, 25+1, 2)},
	}
	for _, tc := range testCases {
		ilist, want := tc.ilist, tc.want
		have := Complement(ilist)
		if ok := slices.Equal(want, have); !ok {
			t.Errorf("have=%v, want=%v", have, want)
		}
	}
}

func TestLetterList_GetBlocks(t *testing.T) {
	testCases := []struct {
		ilist []int   // Just an array of integers
		want  []Block // An array of pairs of integers
	}{
		{[]int{}, []Block{}},
		{[]int{3, 2, 3, 4, 5, 1, 1, 2, 3}, []Block{{3, 3}, {2, 5}, {1, 1}, {1, 3}}},
		{[]int{2, 3, 4, 5, 2}, []Block{{2, 5}, {2, 2}}},
	}
	for _, tc := range testCases {
		ilist := tc.ilist
		c := Blocks(ilist)
		i := 0
		for have := range c {
			if i >= len(tc.want) {
				t.Errorf("More haves than wants, i=%d", i)
				break
			}
			want := tc.want[i]
			i++
			if want.First != have.First {
				t.Errorf("want.First=%d,have.First=%d", want.First, have.First)
			}
			if want.Last != have.Last {
				t.Errorf("want.Last=%d,have.Last=%d", want.Last, have.Last)
			}
		}

		// At this point, i is the number of haves.
		nWants := len(tc.want)
		nHaves := i
		if nWants > nHaves {
			t.Errorf("More wants(%d) than haves(%d)", nWants, nHaves)
		}
	}
}

func TestLetterList_GetPattern(t *testing.T) {
	ilist := []int{2, 3, 4, 5, 25}
	want := "C-FZ"
	have := Pattern(ilist)
	if have != want {
		t.Errorf("want=%v,have=%v", want, have)
	}
}

func TestLetterList_GetRegexp(t *testing.T) {
	testCases := []struct {
		letters string
		want    string
	}{
		{"", ""},                              // empty list
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZ", "."},   // all
		{"ABCD", "[A-D]"},                     // small straight
		{"BCDKLMWXZ", "[^AE-JN-VY]"},          // with gaps
		{"ABCDEFGHIKLMNOPRSTUVWXYZ", "[^JQ]"}, // all but J and Q
		{"ABCDEFGHIJKLMNOPQRSTUVWXY", "[^Z]"}, // all but Z
		{"S", "S"},                            // single letter
		{"", ""},                              // empty pattern
	}
	for i, tc := range testCases {
		letters := tc.letters
		want := tc.want
		have := Regexp(letters)
		if have != want {
			t.Errorf("test case %d: letters=%s, want=%s,have=%s",
				i+1, letters, want, have)
		}
	}
}

package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLetterCell_String(t *testing.T) {
	tests := []struct {
		name     string
		point    Point
		ncAcross *Point
		ncDown   *Point
		letter   string
		want     string
	}{
		{"vanilla", Point{}, &Point{1, 3}, &Point{4, 2}, "A",
			`point:{0,0},ncAcross:(1,3),ncDown:(4,2),letter:"A"`},
		{"no ncAcross", Point{}, nil, &Point{4, 3}, "B",
			`point:{0,0},ncAcross:<nil>,ncDown:(4,3),letter:"B"`},
		{"no ncDown", Point{}, &Point{5, 7}, nil, "C",
			`point:{0,0},ncAcross:(5,7),ncDown:<nil>,letter:"C"`},
		{"no letter", Point{}, &Point{1, 3}, &Point{4, 2}, "",
			`point:{0,0},ncAcross:(1,3),ncDown:(4,2),letter:""`},
		{"no pointers", Point{}, nil, nil, "E",
			`point:{0,0},ncAcross:<nil>,ncDown:<nil>,letter:"E"`},
		{"nothing", Point{}, nil, nil, "",
			`point:{0,0},ncAcross:<nil>,ncDown:<nil>,letter:""`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lc := &LetterCell{
				ncAcross: tt.ncAcross,
				ncDown:   tt.ncDown,
				letter:   tt.letter,
			}
			want := tt.want
			have := lc.String()
			assert.Equal(t, want, have)
		})
	}
}

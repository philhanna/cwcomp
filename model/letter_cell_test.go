package model

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPuzzle_LetterCellIterator(t *testing.T) {
	puzzle := getGoodPuzzle()
	nlc := 0
	for range puzzle.LetterCellIterator() {
		nlc++
	}
	assert.Equal(t, 9*9-16, nlc)
}

func TestLetterCell_GetPoint(t *testing.T) {
	lc := NewLetterCell(NewPoint(1, 2))
	assert.Equal(t, Point{1, 2}, lc.GetPoint())

	lc = *new(LetterCell)
	assert.Equal(t, Point{0, 0}, lc.GetPoint())
}

func TestNewLetterCell(t *testing.T) {
	type args struct {
		point Point
	}
	tests := []struct {
		name string
		args args
		want LetterCell
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewLetterCell(tt.args.point); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewLetterCell() = %v, want %v", got, tt.want)
			}
		})
	}
}

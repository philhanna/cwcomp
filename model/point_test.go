package model

import "testing"
import "github.com/stretchr/testify/assert"

func TestPoint_Compare(t *testing.T) {
	tests := []struct {
		name string
		this Point
		that Point
		want int
	}{
		{"same point", Point{Row: 1, Col: 3}, Point{Row: 1, Col: 3}, 0},
		{"this column is greater", Point{Row: 1, Col: 3}, Point{Row: 1, Col: 2}, 1},
		{"this column is less", Point{Row: 1, Col: 3}, Point{Row: 1, Col: 4}, -1},
		{"this row is greater", Point{Row: 1, Col: 3}, Point{Row: 0, Col: 4}, 1},
		{"this row is less", Point{Row: 0, Col: 3}, Point{Row: 1, Col: 0}, -1},
		{"same row, col less", Point{Row: 0, Col: 0}, Point{Row: 0, Col: 4}, -1},
		{"same row, col more", Point{Row: 0, Col: 0}, Point{Row: 0, Col: -4}, 1},
		{"same col, row less", Point{2, 3}, Point{3, 3}, -1},
		{"same col, row more", Point{4, 3}, Point{3, 3}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := tt.this
			that := tt.that
			assert.Equal(t, tt.want, this.Compare(that))
		})
	}
}

func TestPoint_Equal(t *testing.T) {
	tests := []struct {
		name string
		this Point
		that Point
		want bool
	}{
		{"same point", Point{Row: 1, Col: 3}, Point{Row: 1, Col: 3}, true},
		{"this column is greater", Point{Row: 1, Col: 3}, Point{Row: 1, Col: 2}, false},
		{"this column is less", Point{Row: 1, Col: 3}, Point{Row: 1, Col: 4}, false},
		{"this row is greater", Point{Row: 1, Col: 3}, Point{Row: 0, Col: 4}, false},
		{"this row is less", Point{Row: 0, Col: 3}, Point{Row: 1, Col: 0}, false},
		{"same row, col less", Point{Row: 0, Col: 0}, Point{Row: 0, Col: 4}, false},
		{"same row, col more", Point{Row: 0, Col: 0}, Point{Row: 0, Col: -4}, false},
		{"same col, row less", Point{2, 3}, Point{3, 3}, false},
		{"same col, row more", Point{4, 3}, Point{3, 3}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := tt.this
			that := tt.that
			assert.Equal(t, tt.want, this.Equal(that))
		})
	}
}

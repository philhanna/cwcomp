package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPoint_Compare(t *testing.T) {
	tests := []struct {
		name string
		this Point
		that Point
		want int
	}{
		{"same point", Point{r: 1, c: 3}, Point{r: 1, c: 3}, 0},
		{"this column is greater", Point{r: 1, c: 3}, Point{r: 1, c: 2}, 1},
		{"this column is less", Point{r: 1, c: 3}, Point{r: 1, c: 4}, -1},
		{"this row is greater", Point{r: 1, c: 3}, Point{r: 0, c: 4}, 1},
		{"this row is less", Point{r: 0, c: 3}, Point{r: 1, c: 0}, -1},
		{"same row, col less", Point{r: 0, c: 0}, Point{r: 0, c: 4}, -1},
		{"same row, col more", Point{r: 0, c: 0}, Point{r: 0, c: -4}, 1},
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
	p1 := Point{r: 1, c: 3}
	p2 := p1
	p3 := Point{}
	p3.r++
	p3.c++
	p3.c++
	p3.c++

	tests := []struct {
		name string
		this Point
		that Point
		want bool
	}{
		{"same point", Point{r: 1, c: 3}, Point{r: 1, c: 3}, true},
		{"this column is greater", Point{r: 1, c: 3}, Point{r: 1, c: 2}, false},
		{"identical objects", p1, p2, true},
		{"same values", p1, p3, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			this := tt.this
			that := tt.that
			assert.Equal(t, tt.want, this.Equal(that))
		})
	}
}

func TestPoint_String(t *testing.T) {
	want := `{r:1,c:3}`
	point := Point{r: 1, c: 3}
	have := point.String()
	assert.Equal(t, want, have)
}

func TestPoint_ToXY(t *testing.T) {
	tests := []struct {
		name  string
		point Point
		x     int
		y     int
	}{
		{"simple", Point{1, 7}, 6, 0},
		{"zeros", Point{0, 0}, -1, -1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			point := tt.point
			x, y := point.ToXY()
			assert.Equal(t, tt.x, x)
			assert.Equal(t, tt.y, y)
		})
	}
}

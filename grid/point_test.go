package grid

import (
	"encoding/json"
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
	p1 := Point{Row: 1, Col: 3}
	p2 := p1
	p3 := Point{}
	p3.Row++
	p3.Col++
	p3.Col++
	p3.Col++

	tests := []struct {
		name string
		this Point
		that Point
		want bool
	}{
		{"same point", Point{Row: 1, Col: 3}, Point{Row: 1, Col: 3}, true},
		{"this column is greater", Point{Row: 1, Col: 3}, Point{Row: 1, Col: 2}, false},
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
func TestPoint_FromJSON(t *testing.T) {
	jsonBlob := []byte(`{"r":1,"c":3}`)
	want := Point{Row: 1, Col: 3}
	p := new(Point)
	err := json.Unmarshal(jsonBlob, p)
	assert.Nil(t, err)
	have := *p
	assert.Equal(t, want, have)
}
func TestPoint_String(t *testing.T) {
	want := `(1,3)`
	point := Point{Row: 1, Col: 3}
	have := point.String()
	assert.Equal(t, want, have)
}

func TestPoint_ToJSON(t *testing.T) {
	point := Point{Row: 3, Col: -1}
	jsonBlob, err := json.Marshal(point)
	assert.Nil(t, err)
	have := string(jsonBlob)
	want := `{"r":3,"c":-1}`
	assert.JSONEq(t, want, have)
}

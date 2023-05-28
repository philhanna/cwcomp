package model

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPuzzle_PointIterator(t *testing.T) {
	const n = 3
	puzzle := NewPuzzle(n)

	// Make a list of all points received from the iterator
	list1 := make([]Point, n*n)
	it := puzzle.PointIterator()
	index := 0
	for point := range it {
		list1[index] = point
		index++
	}

	// Make another list of points created from nested loops
	list2 := make([]Point, n*n)
	index = 0
	for r := 1; r <= n; r++ {
		for c := 1; c <= n; c++ {
			list2[index] = NewPoint(r, c)
			index++
		}
	}

	// Should be the same
	assert.Equal(t, list1, list2)
}

func TestPuzzle_SymmetricPoint(t *testing.T) {
	puzzle := NewPuzzle(9)
	tests := []struct {
		p  Point
		sp Point
	}{
		{NewPoint(1, 1), NewPoint(9, 9)},
		{NewPoint(3, 5), NewPoint(7, 5)},
	}
	for _, tt := range tests {
		want := tt.sp
		have := puzzle.SymmetricPoint(tt.p)
		assert.Equal(t, want, have)
	}
}

func TestPoint_Compare(t *testing.T) {
	tests := []struct {
		name string
		this Point
		that Point
		want int
	}{
		{"same point", NewPoint(1, 3), NewPoint(1, 3), 0},
		{"this column is greater", NewPoint(1, 3), NewPoint(1, 2), 1},
		{"this column is less", NewPoint(1, 3), NewPoint(1, 4), -1},
		{"this row is greater", NewPoint(1, 3), NewPoint(0, 4), 1},
		{"this row is less", NewPoint(0, 3), NewPoint(1, 0), -1},
		{"same row, col less", NewPoint(0, 0), NewPoint(0, 4), -1},
		{"same row, col more", NewPoint(0, 0), NewPoint(0, -4), 1},
		{"same col, row less", NewPoint(2, 3), NewPoint(3, 3), -1},
		{"same col, row more", NewPoint(4, 3), NewPoint(3, 3), 1},
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
	p1 := NewPoint(1, 3)
	p2 := p1
	p3 := NewPoint(0, 0)
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
		{"same point", NewPoint(1, 3), NewPoint(1, 3), true},
		{"this column is greater", NewPoint(1, 3), NewPoint(1, 2), false},
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

func TestPoint_NewPoint(t *testing.T) {
	type args struct {
		r int
		c int
	}
	tests := []struct {
		name string
		args args
		want Point
	}{
		{"1, 3", args{1, 3}, NewPoint(1, 3)},
		{"Bad row", args{0, 3}, NewPoint(0, 3)},
		{"Bad column", args{3, 0}, NewPoint(3, 0)},
		{"Bad both", args{-0, -1}, NewPoint(-0, -1)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewPoint(tt.args.r, tt.args.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewPoint() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPoint_String(t *testing.T) {
	want := `{r:1,c:3}`
	point := NewPoint(1, 3)
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
		{"simple", NewPoint(1, 7), 6, 0},
		{"zeros", NewPoint(0, 0), -1, -1},
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

func TestPuzzle_ValidIndex(t *testing.T) {
	tests := []struct {
		name    string
		point   Point
		wantErr bool
	}{
		{"simple point", NewPoint(2, 3), false},
		{"head cell", NewPoint(6, 2), false},
		{"black cell", NewPoint(4, 9), false},
		{"bad row", NewPoint(-10, 3), true},
		{"bad col", NewPoint(1, -10), true},
		{"bad both", NewPoint(0, 0), true},
	}

	puzzle := getGoodPuzzle()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			point := tt.point
			err := puzzle.ValidIndex(point)
			switch tt.wantErr {
			case true:
				assert.NotNil(t, err)
			case false:
				assert.Nil(t, err)
			}
		})
	}
}

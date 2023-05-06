package grid

import "testing"

func TestLetterCell_String(t *testing.T) {
	type fields struct {
		ncAcross Point
		ncDown   Point
		letter   string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{"vanilla", fields{Point{1, 3}, Point{4, 2}, "A"}, `ncAcross:(1,3), ncDown:(4,2), letter:"A"`},
		{"no letter", fields{Point{1, 3}, Point{4, 2}, ""}, `ncAcross:(1,3), ncDown:(4,2), letter:""`},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lc := &LetterCell{
				ncAcross: &tt.fields.ncAcross,
				ncDown:   &tt.fields.ncDown,
				letter:   tt.fields.letter,
			}
			if got := lc.String(); got != tt.want {
				t.Errorf("LetterCell.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

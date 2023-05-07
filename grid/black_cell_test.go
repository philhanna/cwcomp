package grid

import "testing"

func TestBlackCell_String(t *testing.T) {
	tests := []struct {
		name string
		bc   *BlackCell
		want string
	}{
		{"simple", &BlackCell{}, "bc"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bc := &BlackCell{}
			if got := bc.String(); got != tt.want {
				t.Errorf("BlackCell.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

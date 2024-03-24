package importer

import (
	"testing"

	al "github.com/philhanna/cwcomp/acrosslite"
	"github.com/stretchr/testify/assert"
)

func TestHandleReadingSize(t *testing.T) {
	tests := []struct {
		name    string
		line    string
		want    ParsingState
		wantErr bool
	}{
		{
			name: "Good",
			line: "15x15",
			want: LOOKING_FOR_GRID,
		},
		{
			name:    "Not WFF",
			line:    "15xx",
			wantErr: true,
		},
		{
			name:    "Unsquare grid",
			line:    "15x16",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pal := new(al.AcrossLite)
			have, err := HandleReadingSize(pal, tt.line)
			if tt.wantErr {
				assert.NotNil(t, err)
				return
			}
			assert.Equal(t, tt.want, have)
		})
	}
}

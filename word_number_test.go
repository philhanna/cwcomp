package cwcomp

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordNumber_String(t *testing.T) {
	wn := NewWordNumber(1, NewPoint(1, 2))
	want := `seq:1,point:{r:1,c:2}`
	have := wn.String()
	assert.Equal(t, want, have)
}

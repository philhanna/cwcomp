package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWordNumber_String(t *testing.T) {
	wn := NewWordNumber(1, NewPoint(1, 2), 3, 4)
	want := `seq:1,point:{r:1,c:2},aLen:3,dLen:4,aClue:"",dClue:""`
	have := wn.String()
	assert.Equal(t, want, have)
}

package importer

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	dirname, err := filepath.Abs(filepath.Join("..", "testdata"))
	assert.Nil(t, err)

	filename := filepath.Join(dirname, "easy.txt")

	acrossLite, err := Parse(filename)
	assert.Nil(t, err)

	fmt.Println(acrossLite)
}

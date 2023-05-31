package importer

import (
	"fmt"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestImport(t *testing.T) {
	dirname, err := filepath.Abs(filepath.Join("..", "testdata"))
	assert.Nil(t, err)

	filename := filepath.Join(dirname, "easy.txt")

	acrossLite, err := Import(filename)
	assert.Nil(t, err)

	fmt.Println(acrossLite)
}

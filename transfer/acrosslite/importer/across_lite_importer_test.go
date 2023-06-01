package importer

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	dirname, err := filepath.Abs(filepath.Join("..", "testdata"))
	assert.Nil(t, err)
	filename := filepath.Join(dirname, "easy.txt")
	fp, err := os.Open(filename)
	defer fp.Close()
	assert.Nil(t, err)
	reader := bufio.NewReader(fp)
	acrossLite, err := Parse(reader)
	assert.Nil(t, err)

	fmt.Println(acrossLite)
}

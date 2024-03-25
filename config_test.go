package cwcomp

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetConfiguration(t *testing.T) {
	config := GetConfiguration()
	_ = config
}

func TestGetConfigFileName(t *testing.T) {
	filename := GetConfigFileName()
	assert.True(t, strings.HasSuffix(filename, "config.yaml"))
}

func TestNewConfiguration(t *testing.T) {
	// Good case
	config, err := NewConfiguration()
	assert.Nil(t, err)
	assert.NotNil(t, config)

	// Non-existent file
	GetConfigFileName = func() string {
		return "bogus"
	}
	config, err = NewConfiguration()
	assert.NotNil(t, err)
	assert.Nil(t, config)

	// File with invalid yaml
	filename := filepath.Join(os.TempDir(), "invalid.yaml")
	defer os.Remove(filename)
	
	data := []byte(`this is invalid yaml`)
	os.WriteFile(filename, data, 0644)
	GetConfigFileName = func() string {
		return filename
	}
	config, err = NewConfiguration()
	assert.NotNil(t, err)
	assert.Nil(t, config)
}

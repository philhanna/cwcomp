package cwcomp

import (
	"log"
	"os"
	"path/filepath"

	"github.com/ghodss/yaml"
)

var Configuration *configuration

func init() {
	// Load the configuration
	var err error
	Configuration, err = newConfiguration()
	if err != nil {
		log.Fatal(err)
	}
}

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

type configuration struct {
	DATABASE database `json:"database"`
	SERVER   server   `json:"server"`
}

type database struct {
	NAME string `json:"name"`
	DDL  string `json:"ddl"`
}

type server struct {
	HOST string `json:"host"`
	PORT int    `json:"port"`
}

// ---------------------------------------------------------------------
// Constants
// ---------------------------------------------------------------------

const (
	YAML_FILE_NAME = "config.yaml"
)

var (
	PACKAGE_NAME = GetPackageName()
)

// ---------------------------------------------------------------------
// Constructor
// ---------------------------------------------------------------------

// newConfiguration creates a configuration structure from the YAML file
// in the user configuration directory.
func newConfiguration() (*configuration, error) {

	// Get the configuration file name
	configfile, err := configurationFile()
	if err != nil {
		return nil, err
	}

	// Load its data
	yamlBlob, err := os.ReadFile(configfile)
	if err != nil {
		return nil, err
	}

	// Create a configuration structure from the YAML
	p := new(configuration)
	err = yaml.Unmarshal(yamlBlob, p)
	if err != nil {
		return nil, err
	}

	return p, nil
}

// ---------------------------------------------------------------------
// Functions
// ---------------------------------------------------------------------

// configurationFile returns the name of the configuration YAML file in
// .config
func configurationFile() (string, error) {

	// Start with the user configuration directory
	// (on Unix, "$HOME/.config")

	dirname, err := os.UserConfigDir()
	if err != nil {
		return "", err
	}

	// Concatenate the path to the yaml

	configfile := filepath.Join(dirname, PACKAGE_NAME, YAML_FILE_NAME)

	return configfile, nil
}

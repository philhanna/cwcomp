package cwcomp

import (
	"log"
	"os"
	"path/filepath"

	"github.com/ghodss/yaml"
)

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

type Configuration struct {
	DATABASE database `json:"database"`
	SERVER   server   `json:"server"`
}

type database struct {
	NAME string `json:"name"`
}

type server struct {
	HOST string `json:"host"`
	PORT int    `json:"port"`
}

// ---------------------------------------------------------------------
// Constants and variables
// ---------------------------------------------------------------------

// Name of the config.yaml file in the configuration directory
const YAML_FILE_NAME = "config.yaml"

// Name of this package. Used to specify the subdirectory for this
// application in the configuration directory.
var PACKAGE_NAME = GetPackageName()

// A pointer to the loaded instance of the configuration
var configInstance *Configuration

var GetConfiguration = func() *Configuration {
	return configInstance
}

// ---------------------------------------------------------------------
// Constructor
// ---------------------------------------------------------------------

// newConfiguration creates a configuration structure from the YAML file
// in the user configuration directory.
func newConfiguration() (*Configuration, error) {

	// Get the configuration file name
	configfile, err := configurationFile()
	if err != nil {
		return nil, err
	}
	log.Printf("Reading configuration from %v\n", configfile)

	// Load its data
	yamlBlob, err := os.ReadFile(configfile)
	if err != nil {
		return nil, err
	}

	// Create a configuration structure from the YAML
	p := new(Configuration)
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

func init() {
	// Load the configuration
	var err error
	configInstance, err = newConfiguration()
	if err != nil {
		log.Fatal(err)
	}
}

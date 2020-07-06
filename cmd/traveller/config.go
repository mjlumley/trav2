package main

// config.go contains all the code for application configuration.

import (
	"bytes"
	"encoding/json"
	"os"
	"runtime"
)

// windowWidth is the width of the main window in pixels.
const windowWidth = 1280

// windowHeight is the height of the main window in pixels.
const windowHeight = 720

// Constants for mouse button indexes.
const (
	mouseButtonPrimary   = 0
	mouseButtonSecondary = 1
	mouseButtonTertiary  = 2
	mouseButtonCount     = 3
)

// configFile is the (hard-coded) confiuration file.
const configFile string = "./config.json"

// appVersion is the (hard-coded) version
const appVersion string = "0.8.0.20200618151325"

// LogFile is the name of the file to which log output is written.
const LogFile string = "traveller.log"

// configuration contains all configurable fields. Note that configuration items that come from the JSON
// configuration file need to be exported fields
type configuration struct {

	// These are the "built-in" configuration items.
	version string // Program version
	os      string // Operating System, determined at runtime.
	program string // Program name

	// Minor config items
	languageNumItems int   // How many random language words to display
	seedValue        int64 // The value for seeding generation. For testing purposes.

	// These are configuration options from the JSON config file.
	DatabaseFile     string // DatabaseFile is the name of the SQLite database file.
	DataDir          string // The data directory for file input/output
	WorldOutputFile  string // Output tab file for worlds generated
	SectorOutputFile string // Output tab file for sectors generated
	WorldGenNumber   int    // Number of worlds generated in Auto mode
	ForevenFile      string // Output worlds tab file for Foreven sector
}

// config is the global configuration item.
var config configuration

// memLog defines an in-memory logging buffer.
var memLog bytes.Buffer

// Initialisation for this file. In this case, most application initialisation of globals is performed here.
func init() {

}

// initConfig initialises the configuration module. The config items that are needed for the JSON decoder are in uppercase.
func initConfig() {

	// Basic configuration
	config.os = runtime.GOOS
	config.version = appVersion

	// Set some config
	config.languageNumItems = 20
	config.seedValue = 101

	// Read in rest of configuration from file
	file, err := os.Open(configFile)
	if err != nil {
		panic(err)
	}
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&config)
	if err != nil {
		panic(err)
	}

}

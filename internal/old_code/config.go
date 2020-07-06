package main

// config.go contains all the code for application configuration.

import (
	"encoding/json"
	"os"
	"runtime"
)

// configFile is the (hard-coded) confiuration file.
const configFile string = "./config.json"

// appVersion is the (hard-coded) version
const appVersion string = "0.6.24.20200312145752"

// LogFile is the name of the file to which log output is written.
const LogFile string = "traveller.log"

// configuration contains all configurable fields. Note that configuration items that come from the JSON
// configuration file need to be exported fields
type configuration struct {

	// These are the "built-in" configuration items.
	version string // Program version
	os      string // Operating System, determined at runtime.
	program string // Program name

	// These are configuration options from the JSON config file.
	DatabaseFile     string // DatabaseFile is the name of the SQLite database file.
	PageOptionSize   int    // How many lines of page to display
	WebRetrieveWait  int    // Number of seconds to wait when making Web calls to travellermap.com
	WebLineString    string // What we expect to see at the end of every line when retrieved from travellermap.com, CRLF for instance
	WebDelimiter     string // Comma or tab separated data from travellermap.com
	WorldFields      int    // How many fields we expect to see from travellermap.com.
	SeedValue        int64  // The value for seeding generation. For testing purposes.
	DataDir          string // The data directory for file input/output
	WorldOutputFile  string // Output tab file for worlds generated
	SectorOutputFile string // Output tab file for sectors generated
	WorldGenNumber   int    // Number of worlds generated in Auto mode
	ForevenFile      string // Output worlds tab file for Foreven sector
}

// config is the global configuration item.
var config configuration

// initConfig initialises the configuration module. The config items that are needed for the JSON decoder are in uppercase.
func initConfig() {

	// Basic configuration
	config.os = runtime.GOOS
	config.version = appVersion

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

// Package main is the executable for The Traveller's Tool utility.
package main

// main.go contains the main function and associated basic functions for the traveller program.

import (
	"fmt"
	"log"
	"os"
	"strings"
)

// main is the entry point for the traveller application.
func main() {

	// Logging
	logf, err := os.OpenFile(LogFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer logf.Close()
	log.SetOutput(logf)
	log.Print("---------------------")
	log.Print("Traveller started")

	// Get the command-line args and decide what to do.
	args := os.Args

	// Initialise the remainder of the configuration from file
	config.program = args[0]
	initConfig()
	//initDb(config.DatabaseFile)    // This is the old version and should be removed once all functions are moved across.
	log.Printf("Setting database file to %s", config.DatabaseFile)
	SetDbFile(config.DatabaseFile) // This is the new version.

	// TODO: Swap this code out for flags, as command-line argument is likely to get more complicated in the future
	switch len(args) {
	case 1:
		// Single arg would be the command itself, so run interactively.
	case 2:
		// Two args would be either help or verion or error.
		switch strings.ToLower(args[1]) {
		case "--help":
			// Print help information.
			usage(config.program)
		case "--version":
			// Print out program version information.
			displayVersion()
		default:
			// Any case that hasn't been accounted for now is incorrect.
			fmt.Println("Incorrect command-line parameters")
			usage(config.program)
			// Theoretically we should exit with an error number, and not exit with 0.
		}
		// In any case we exit immediately
		os.Exit(0)
	default:
		// Any other
		usage(config.program)
		os.Exit(0)
	}

	// Continue on from here interactively
	fmt.Println("Running The Traveller's Tool interactively...")

	notMain()
	fmt.Println()
	fmt.Println("Thank you for using The Traveller's Tool. Bye")
}

// usage prints out a usage statement, taking the program name as an argument.
func usage(prog string) {
	fmt.Printf("Usage: %s [--help|--version]\n", prog)
}

// displayVersion displays the application version.
func displayVersion() {
	fmt.Printf("%s - Version %s\n", config.program, config.version)
}

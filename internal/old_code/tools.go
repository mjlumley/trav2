package main

// tools.go contains tools for use throughout the entire application.

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"sync"
	"time"
)

// randSeq is a pointer to a RNG
var randSeq *rand.Rand

// Allows seeding only once.
var onlyOnce sync.Once

// init performs initialisation of this module.
func init() {

	// Seed the RNG
	randSeq = rand.New(rand.NewSource(time.Now().UnixNano()))
}

///////// Random Number Generation Tools

// Dice rolls a dice of the specified size. Returns the number rolled.
func Dice(sides int) int {
	if randSeq == nil {
		onlyOnce.Do(func() {
			seed1 := rand.NewSource(time.Now().UnixNano())
			randSeq = rand.New(seed1)
		})
	}

	return randSeq.Intn(sides) + 1
}

// D6 returns the result of a 6-sided dice rolled.
func D6() int {
	return Dice(6)
}

// Flux makes a Traveller "flux" roll, which is 1d6 - 1d6, with possible addition of Dice Modifier.
// It returns the integer result.
func Flux(dm int) int {
	return (D6() - D6() + dm)
}

// SeedForTesting provides an opportunity to seed the RNG for testing purposes.
// You must provide a seed value (hint: perhaps from config?).
func SeedForTesting(s int64) {
	randSeq = rand.New(rand.NewSource(s))
}

////// These have to go.
// clear provides a map for storing clear funcs (although one is only needed!).
var clear map[string]func() //create a map for storing clear funcs

// init performs initialisation specific to these tools.
func init() {
	//Initialize the clear map - for screen clears
	clear = make(map[string]func())
	clear["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

/////////////////
// Tools
/////////////////

// callClear clears the console window. Unbelievably, this actually works!
func callClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

// getChoice gets a users input choice. It returns the string with the users choice, trimmed
// for spaces and carriage returns/line feeds/newlines.
func getChoice(questionString string) (choice string) {

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf(questionString)
	choice, _ = reader.ReadString('\n')

	switch config.os {
	case "windows":
		choice = strings.Replace(choice, "\r\n", "", -1)
	default:
		choice = strings.Replace(choice, "\n", "", -1)
	}

	return
}

// enterToContinue asks for the user to press enter to continue. Actually, anything will do.
func enterToContinue() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Press <Enter> to continue")
	_, _ = reader.ReadString('\n')

}

// getYesNoAnswer gets a yes/no answer. Provide a question in the string, and whether the default is yes (true) or no (false).
// Returns true for yes or false for no.
func getYesNoAnswer(question string, defaultYes bool) bool {
	for {
		if defaultYes {
			question += " [Y/n]?"
		} else {
			question += " [y/N]?"
		}
		r := getChoice(question)
		if r == "" {
			return defaultYes
		}
		r = strings.ToLower(r)
		switch r {
		case "y", "yes":
			return true
		case "n", "no":
			return false
		default:
			fmt.Println("Incorrect choice - please enter 'y' or 'n'")
		}
	}

}

package main

// ehex.go contains code for Ehex characters/values.

import "strings"

// Ehex is an integer value that for values larger than 10 can be displayed in one
// character space, and that these values are represented by the upper-case characters
// starting from 'A'. As such it is similar to a hexadecimal value but go all the
// way up to 33 ('Z').
type Ehex int8

// ehex contains a map of Ehex to integer conversions
var ehexMap map[string]Ehex

// EhexMax is the maximum value for Ehex characters
const EhexMax = 33

// EhexOor is a string indicating the Ehex value is Out Of Range.
const EhexOor = "?"

// init intialises the ehex code of the tools package.
// Specifically, this initialises the
func init() {

	// Init the ehex map
	ehexMap = make(map[string]Ehex)
	ehexMap["0"] = 0
	ehexMap["1"] = 1
	ehexMap["2"] = 2
	ehexMap["3"] = 3
	ehexMap["4"] = 4
	ehexMap["5"] = 5
	ehexMap["6"] = 6
	ehexMap["7"] = 7
	ehexMap["8"] = 8
	ehexMap["9"] = 9
	ehexMap["A"] = 10
	ehexMap["B"] = 11
	ehexMap["C"] = 12
	ehexMap["D"] = 13
	ehexMap["E"] = 14
	ehexMap["F"] = 15
	ehexMap["G"] = 16
	ehexMap["H"] = 17
	ehexMap["J"] = 18
	ehexMap["K"] = 19
	ehexMap["L"] = 20
	ehexMap["M"] = 21
	ehexMap["N"] = 22
	ehexMap["P"] = 23
	ehexMap["Q"] = 24
	ehexMap["R"] = 25
	ehexMap["S"] = 26
	ehexMap["T"] = 27
	ehexMap["U"] = 28
	ehexMap["V"] = 29
	ehexMap["W"] = 30
	ehexMap["X"] = 31
	ehexMap["Y"] = 32
	ehexMap["Z"] = 33
}

// String returns the string value of the Ehex.
func (e Ehex) String() string {
	if e < 0 || e > EhexMax {
		return EhexOor
	}
	vals := [34]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F", "G", "H",
		"J", "K", "L", "M", "N", "P", "Q", "R", "S", "T", "U", "V", "W", "X", "Y", "Z"}
	return vals[e]

}

// Int returns the string value of the Ehex.
func (e Ehex) Int() int {
	return int(e)
}

// EhexVal converts a string representing an ehex to an Ehex type. Error value is -1.
func EhexVal(e string) Ehex {
	e = strings.ToUpper(e)

	if val, found := ehexMap[e]; found {
		return val
	}
	return -1
}

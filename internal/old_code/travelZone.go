package main

import (
	"errors"
	"strings"
)

// travelZone.go contains code for handling Travel Zones

// TravelZone describes the allocated Travel Zone for a world: Green, Amber or Red.
type TravelZone int

// Constants for travel zones.
const (
	// Travel zone Green
	TzGreen TravelZone = iota
	// Travel zone Amber
	TzAmber
	// Travel zone Red
	TzRed
	// Invalid travel zone
	TzUnknown
)

// String returns the (short) string for the travel zone
func (t TravelZone) String() string {
	return [...]string{"", "A", "R", "?"}[t]
}

// Desc returns the longer descriptive string for the Travel Zone.
func (t TravelZone) Desc() string {
	return [...]string{"Green", "Amber", "Red", "Unknown"}[t]
}

// ColouredString returns a coloured string for the Travel Zone.
func (t TravelZone) ColouredString() string {
	return [...]string{"[green]Green[-]", "[yellow]Amber[-]", "[red]Red[-]", "Unknown"}[t]
}

// ZoneFromString returns a TravelZone type and error based on the given string. If the string
// cannot be converted into a TravelZone, then TravelZone will be TzUnknown and error will
// not be nil.
func ZoneFromString(z string) (TravelZone, error) {

	z = strings.ToUpper(z)

	if z == "" || z == "G" || z == "GREEN" {
		return TzGreen, nil
	}
	if z == "A" || z == "AMBER" {
		return TzAmber, nil
	}
	if z == "R" || z == "RED" {
		return TzRed, nil
	}
	return TzUnknown, errors.New("TravelZone: Unable to convert " + z)

}

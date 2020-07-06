package main

// tables.go contains a bunch of the Traveller (T5) tables. For instance,
// the description of Hydrographic C may be here. these tables, ideally, will be moved to the database.

// Constants for starports and spaceports.
const (
	stpUnknown   = "Unknown"            // An unknown Star- or Space-port
	stpExcellent = "Starport excellent" // Excellent quality starport
	stpGood      = "Starport good"
	stpRoutine   = "Starport routine"
	stpPoor      = "Starport poor"
	stpFrontier  = "Starport frontier"
	stpNone      = "Starport none"
	// Spaceports below
	sppGood  = "Spaceport good"
	sppPoor  = "Spaceport poor"
	sppBasic = "Spaceport basic"
	sppNone  = "Spaceport none"
)

// tableStarport takes a single character starport (or spaceport), and returns a description from the T5 table.
// If the starport type is not recognised stpUnknown ("Unknown") is returned.
func tableStarport(s string) string {
	switch s {
	case "A":
		return stpExcellent
	case "B":
		return stpGood
	case "C":
		return stpRoutine
	case "D":
		return stpPoor
	case "E":
		return stpFrontier
	case "X":
		return stpNone
	case "F":
		return sppGood
	case "G":
		return sppPoor
	case "H":
		return sppBasic
	case "Y":
		return sppNone
	default:
		return stpUnknown
	}
}

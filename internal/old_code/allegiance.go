package main

// allegiance.go contains code for allegiances. Predominantly used for Traveller5 allegiances

type allegiance struct {
	code       string // The unique code for the allegiance.
	name       string // A descriptive string for the allegiance.
	legacyCode string // The legacy code for this allegiance. Definitely not unique, and there are clashes.
}

// getAllegianceForCode returns the allegiance matching the unique 4-character code found in the search string sc.
// It returns the allegiance found or a blank allegiance.
func getAllegianceForCode(s []allegiance, sc string) allegiance {
	for _, alleg := range s {
		if alleg.code == sc {
			return alleg
		}
	}
	return allegiance{"", "", ""}
}

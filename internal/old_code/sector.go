package main

import (
	"log"
	"os"
)

// sector.go contains code for sectors and subsectors.

// sectorDTO stores details about a system mapping sector in the database.
type sectorDTO struct {
	id     int    // The Sector's ID from the database.
	name   string // The (official) name of the Sector.
	abbrev string // A four-letter abbreviation for the Sector (usually first 4 chars of the name).
	xLoc   int    // The travellermap.com x offset from Core sector.
	yLoc   int    // The travellermap.com y offset from Core sector.
}

// subsector stores details about a subsector, which can contain up to 80 systems. There are 16 subsectors to  sector in a 4x4 grid.
//
type subsector struct {
	id        int    // The subsector's ID from the database.
	name      string // The Subsector's name.
	remarks   string // Any remarks for the subsector.
	language  string // The majority language and language used to name the Subsector.
	capitalID int    // The ID of the mainworld that is the subsector capital.
}

// subsectorDTO is used for collecting subsector information from the database.
type subsectorDTO struct {
	id             int    // The subsector's ID from the database.
	name           string // The Subsector's name
	sectorID       int    // The database ID of the Sector containing this Subsector.
	subsectorIndex string // The "index" (A through P) of the subsector within the sector. See map.
	remarks        string // Any remarks for the Subsector.s
	langID         int    // The databse ID of the majority language that is used in the Subsector. This will be the language that the Subsector name is in.
	capitalID      int    // The database ID of the mainworld that is the subsector capital if any.
}

// sector contains an ordered collection of worlds in a grid.
type sector struct {
	id         int           // The Sector's ID from the database if it is from the OTU, or -1 if not.
	name       string        // The name of the sector
	abbrev     string        // The four-letter abbreviation for the sector (usually first 4 characters of the name).
	worlds     []world       // The list of worlds
	saved      bool          // If the sector has been saved.
	subsectors [16]subsector // The subsectors (in order from A to P) for this sector if known.
	otu        bool          // Whether the sector belongs to the "Official Traveller Universe"
}

// toTab writes the sector to a tab-delimited string, suitable for displaying on screen or in a file.
func (s sector) toTab() (st string) {
	for _, w := range s.worlds {
		st += w.String() + "\n"
	}
	return
}

// getAbbreviationForSector gets the abbreviation for a Sector. If it finds it in the
// database it uses that, if not, it uses the first four characters of the sector name.
func getAbbreviationForSector(s string) (a string) {

	if len(s) == 0 {
		return ""
	}

	sectorMap, _ := GetAllDetailedSectorAbbrev()
	a = sectorMap[s]

	if len(a) == 0 {
		if len(s) >= 4 {
			a = s[0:4]
		} else {
			a = s
		}
	}
	return
}

/* ssIndex[16] is the definitive map from array to string. */

// // subsectorIndex returns the string for the given index.
// func subsectorIndex(idx int) string {
// 	return [...]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P"}
// }

// toFile writes a sector to the given filename.
func (s *sector) toFile(fn string) error {

	f, err := os.OpenFile(fn, os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("Unable to open file "+fn+". Error: %v", err)
		return err
	}
	defer f.Close()
	// Write the list of worlds to the sector file.
	if _, err := f.WriteString(headerOut[WgtMtBasic] + "\n" + s.toTab()); err != nil {
		log.Printf("Write error : %v", err)
		return err
	}
	logPane.Log("Sector written to file: " + fn)
	s.saved = true
	return nil
}

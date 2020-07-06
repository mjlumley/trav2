package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// hexLoc.go contains code for dealing with star mapping Hex Locations.
//
// Subsectors "Subsector Index" are laid out in each sector like this:
//  /---------------\
//  | A | B | C | D |
//  |---+---+---+---|
//  | E | F | G | H |
//  |---+---+---+---|
//  | I | J | K | L |
//  |---+---+---+---|
//  | M | N | O | P |
//  \---------------/

// ssIndex contains an array of subsector index identifiers.
var ssIndex [16]string

// init initialises the array of subsector index identifiers.
func init() {
	ssIndex = [16]string{"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P"}
}

// HexLoc defines a Hex Location used for Star Mapping. If used for
type HexLoc struct {
	x int // The x location within the sector/subsector
	y int // The y location within the sector/subsector

	sector bool // true if the HexLoc indicates a sector location, false if a subsector location
}

// IsSector returns whether the HexLoc refers to a Sector location or a subsector location.
func (h *HexLoc) IsSector() bool {
	return h.sector
}

// String returns the String representation of the HexLoc
func (h *HexLoc) String() string {
	return fmt.Sprintf("%02d%02d", h.x, h.y)
}

// Coordinates returns the x and y coordinates of the HexLoc.
func (h *HexLoc) Coordinates() (x, y int) {
	return h.x, h.y
}

// IntIndex returns a 0 - 15 index, for instance "A" is zero, "E" is 4. A
// -1 is returned if the HexLoc is invalid.
func (h *HexLoc) IntIndex() int {
	if !h.IsValid() {
		return -1
	}

	str := h.GetIndex()
	for i := 0; i < 16; i++ {
		if str == ssIndex[i] {
			return i
		}
	}
	return -1
}

// GetIndex returns the "Subsector Index" of the HexLoc, that is a string value from "A" to "P".
// If the HexLoc is a subsector location (sector=false), then a blank string is returned.
func (h *HexLoc) GetIndex() string {
	if !h.sector || !h.IsValid() {
		return ""
	}
	// Work out the row and column for the subsector index map, thence the index.
	idx := 4*((h.y-1)/10) + ((h.x - 1) / 8)
	return ssIndex[idx]
}

// IsValid returns whether the HexLoc is valid or not.
func (h *HexLoc) IsValid() bool {
	if h.sector {
		return h.x >= 1 && h.x <= 32 && h.y >= 1 && h.y <= 40
	}
	return h.x >= 1 && h.x <= 8 && h.y >= 1 && h.y <= 10
}

// ConvertToSector converts the (subsector) HexLoc to a sector HexLoc, given
// the Subsector Index as a string. It returns the converted HexLoc, or an
// error if the conversion is invalid.
func (h *HexLoc) ConvertToSector(idxStr string) (*HexLoc, error) {
	if !h.IsValid() {
		return nil, errors.New("HexLoc: invalid HexLoc")
	}
	index := -1
	for i, v := range ssIndex {
		if v == idxStr {
			index = i
		}
	}
	if index == -1 {
		return nil, errors.New("HexLoc: invalid subsector index " + idxStr)
	}
	h.sector = true
	h.x = (index%4)*8 + h.x
	h.y = (index/4)*10 + h.y
	return h, nil
}

// ConvertToSubsector converts the (sector) HexLoc to a subsector HexLoc.
// It returns the converted HexLoc, or an error if the conversion is invalid.
func (h *HexLoc) ConvertToSubsector() (*HexLoc, error) {
	if !h.IsValid() {
		return nil, errors.New("HexLoc: invalid HexLoc")
	}
	h.sector = false
	x := h.x % 8
	if x == 0 {
		x = 8
	}
	y := h.y % 10
	if y == 0 {
		y = 10
	}
	h.x = x
	h.y = y
	return h, nil
}

/////////////////////////////////////////////
// Some tools for working with Hex locations.
//

// NewHexLoc returns a new HexLoc pointer given a string representing the Hex Location, and
// whether the HexLoc is refers to a Sector or Subsector location. If there is any
// error creating the HexLoc (for instance value out of range), then nil is returned.
func NewHexLoc(h string, isSector bool) *HexLoc {
	if len(h) != 4 {
		return nil
	}
	// Grab the individual "numbers" out of this.
	a := []rune(h)
	x, err1 := strconv.Atoi(string(a[0:2]))
	y, err2 := strconv.Atoi(string(a[2:4]))

	if err1 != nil || err2 != nil {
		return nil
	}
	if (x < 1 || x > 32 || y < 1 || y > 40) && isSector {
		return nil
	}
	if (x < 1 || x > 8 || y < 1 || y > 10) && !isSector {
		return nil
	}
	return &HexLoc{x: x, y: y, sector: isSector}

}

// Compare compares two HexLocs and returns -1 if h1 < h2, 0 if h1 = h2, and 1 if h1 > h2.
// If you are trying to compare two different types of hex locations (sector and subsector),
// then an error will be returned.
//
// In order to understand how one location is smaller than another, it is the same order as
// shown in sector listings. The rules for this are as follows:
// - The locations are in subsector order, ie regardless of x,y numbers, subsector "A" is
// lower than subsector "B".
// - Within each subsector, we hold x while advancing y, so 0101 is followed by 0102.
// - When you get to the last row in a column for that subsector, you advance to the next
// column, so the hexloc that follows 0110 is 0201.
// - When you get to the final hex of a subsector, you will advance to the next subsector.
// For instance, going from subsector G to H, would be 2420 -> 2511.
func Compare(h1, h2 HexLoc) (int, error) {
	if !h1.IsValid() {
		return -2, errors.New("HexLoc: Invalid hex Location " + h1.String())
	}
	if !h2.IsValid() {
		return -2, errors.New("HexLoc: Invalid hex Location " + h2.String())
	}
	if h1.IsSector() != h2.IsSector() {
		return -2, errors.New("HexLoc: Cannot compare sector and subsector locations")
	}

	// Should have valid HexLocs now.

	// Check for simple case of different subsectors
	if h1.IsSector() && (strings.Compare(h1.GetIndex(), h2.GetIndex()) != 0) {
		return strings.Compare(h1.GetIndex(), h2.GetIndex()), nil
	}

	// In the same subsector, make a linear value from the x,y values.
	val1 := h1.x*10 + h1.y
	val2 := h2.x*10 + h2.y
	if val1 == val2 {
		return 0, nil
	}
	if val1 < val2 {
		return -1, nil
	}
	return 1, nil
}

// ByLoc implements the sort.Interface for []HexLoc based on the location fields.
// To use this:
//
//    locs := []HexLoc
//    ...
//    sort.Sort(ByLoc(locs))
//
type ByLoc []HexLoc

// Len returns the length of the slice of HexLoc.
func (h ByLoc) Len() int {
	return len(h)
}

// Swap swaps the HexLocs at the indexes around.
func (h ByLoc) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

// Less compares two of the HexLocs at the indexes and returns true if the first
// is less than the second.
func (h ByLoc) Less(i, j int) bool {
	res, err := Compare(h[i], h[j])
	if err != nil {
		panic(err)
	}
	return res == -1
}

package main

// stars.go contains code for stars and star systems

import (
	"log"
	"strconv"
	"strings"
)

// starDetail stores details of a star in a system
type starDetail struct {
	spectralType    string      // The Star Type from : O B A F G K M BD. If BD, size will be zero.
	spectralDecimal int         // The decimal value for spectral 0 to 9. This will be 0 for D (dwarf) size stars
	size            string      // The star size from one of these luminosity classes : Ia Ib II III IV V VI D
	description     string      // The description of the star
	companion       *starDetail // The companion star
	orbit           int         // Orbit number (in the Primary's system) for Non-Primary stars
	//	mass            float64     // The mass of the star (in earth masses)
}

// satelliteOrbit is an array of names for Satellite Orbits.
var satelliteOrbit [26]string

// satelliteOrbitMultiplier is a corresponding array for Orbit multipliers.
var satelliteOrbitMultiplier [26]int

// init intialises the structures needed for this module.
func init() {
	satelliteOrbit[0] = "Ay"
	satelliteOrbit[1] = "Bee"
	satelliteOrbit[2] = "Cee"
	satelliteOrbit[3] = "Dee"
	satelliteOrbit[4] = "Ee"
	satelliteOrbit[5] = "Eff"
	satelliteOrbit[6] = "Gee"
	satelliteOrbit[7] = "Aitch"
	satelliteOrbit[8] = "Eye"
	satelliteOrbit[9] = "Jay"
	satelliteOrbit[10] = "Kay"
	satelliteOrbit[11] = "Ell"
	satelliteOrbit[12] = "Em"
	satelliteOrbit[13] = "En"
	satelliteOrbit[14] = "Oh"
	satelliteOrbit[15] = "Pee"
	satelliteOrbit[16] = "Que"
	satelliteOrbit[17] = "Arr"
	satelliteOrbit[18] = "Ess"
	satelliteOrbit[19] = "Tee"
	satelliteOrbit[20] = "Yu"
	satelliteOrbit[21] = "Vee"
	satelliteOrbit[22] = "Dub"
	satelliteOrbit[23] = "Ex"
	satelliteOrbit[24] = "Wye"
	satelliteOrbit[25] = "Zee"

	satelliteOrbitMultiplier[0] = 1
	satelliteOrbitMultiplier[1] = 2
	satelliteOrbitMultiplier[2] = 3
	satelliteOrbitMultiplier[3] = 4
	satelliteOrbitMultiplier[4] = 5
	satelliteOrbitMultiplier[5] = 6
	satelliteOrbitMultiplier[6] = 8
	satelliteOrbitMultiplier[7] = 10
	satelliteOrbitMultiplier[8] = 20
	satelliteOrbitMultiplier[9] = 30
	satelliteOrbitMultiplier[10] = 40
	satelliteOrbitMultiplier[11] = 50
	satelliteOrbitMultiplier[12] = 60
	satelliteOrbitMultiplier[13] = 70
	satelliteOrbitMultiplier[14] = 80
	satelliteOrbitMultiplier[15] = 100
	satelliteOrbitMultiplier[16] = 150
	satelliteOrbitMultiplier[17] = 200
	satelliteOrbitMultiplier[18] = 250
	satelliteOrbitMultiplier[19] = 300
	satelliteOrbitMultiplier[20] = 400
	satelliteOrbitMultiplier[21] = 500
	satelliteOrbitMultiplier[22] = 600
	satelliteOrbitMultiplier[23] = 700
	satelliteOrbitMultiplier[24] = 800
	satelliteOrbitMultiplier[25] = 1000
}

// getDescription gets the description of the star based on its size.
func (s starDetail) getDescription() string {
	switch s.size {
	case "Ia":
		return "Bright Supergiant"
	case "Ib":
		return "Supergiant"
	case "II":
		return "Bright Giant"
	case "III":
		return "Giant"
	case "IV":
		return "Sub-giant"
	case "V":
		return "Main Sequence"
	case "VI":
		return "Sub-dwarf"
	case "D":
		return "White Dwarf"
	default:
		if s.spectralType == "BD" {
			return "Brown Dwarf"
		}
		return "Unknown"
	}
}

// String shows brief star type (luminosity and size) info for a star.
func (s starDetail) String() string {
	if s.spectralType == "BD" {
		return s.spectralType
	}
	if s.size == "D" {
		return s.size + s.spectralType
	}
	decimal := strconv.Itoa(s.spectralDecimal)
	return s.spectralType + decimal + " " + s.size
}

// StarString shows brief details about all of a system's stars.
func StarString(ss []*starDetail) (ret string) {
	ret = ""
	for i, star := range ss {
		if i != 0 {
			ret = ret + " "
		}
		ret = ret + star.String()
		// Print out the companion(s) -- could potentially be a endless linked list.
		for {
			if star.companion == nil {
				break
			}
			star = star.companion
			ret = ret + " " + star.String()
		}
	}
	return
}

// determineStar generates Homestar spectral class and luminosity (or type and size). A DM (usually -1 to +1) can be added to the rolls, and flux for the
// Primary star is given. Set isHomestar to true if homestar (Primary). Returns a starDetail structure containing the type/spectral and size/luminosity
// details for the star.
func determineStar(dm int, specFlux int, sizeFlux int, isHomestar bool) (s starDetail) {

	// Our "roll" for the Star Spectral Type
	roll := dm + specFlux
	if !isHomestar {
		roll = specFlux + D6() + 1
	}
	if roll < -6 {
		roll = -6
	}
	if roll > 6 {
		roll = 6
	}

	// Our roll for the Star Spectal decimal
	s.spectralDecimal = Dice(10) - 1 // May be ignored for Dwarfs

	// Determine the star spectral type
	if isHomestar {
		switch roll {
		case -6:
			s.spectralType = "O"
		case -5:
			if Dice(2) == 2 {
				s.spectralType = "O"
			} else {
				s.spectralType = "B"
			}
		case -4, -3:
			s.spectralType = "A"
		case -2, -1:
			s.spectralType = "F"
		case 0:
			s.spectralType = "G"
		case 1, 2:
			s.spectralType = "K"
		default:
			s.spectralType = "M"
		}
	} else {
		switch roll {
		case -6:
			if Dice(2) == 2 {
				s.spectralType = "O"
			} else {
				s.spectralType = "B"
			}
		case -5, -4:
			s.spectralType = "A"
		case -3, -2:
			s.spectralType = "F"
		case -1, 0:
			s.spectralType = "G"
		case 1, 2:
			s.spectralType = "K"
		case 3, 4, 5:
			s.spectralType = "M"
		default:
			s.spectralType = "BD"
			// Ignore remaining rolls
			s.size = ""
			s.spectralDecimal = 0
			s.description = "Brown Dwarf"
			//s.mass = s.getMass()
			return
		}
	}

	// Now determine the star size
	roll = sizeFlux
	if !isHomestar {
		roll = sizeFlux + D6() + 2
	}
	if roll > 6 {
		roll = 6
	}
	if roll < -5 {
		roll = -5
	}

	switch roll {
	case -5:
		switch s.spectralType {
		case "O", "B", "A":
			s.size = "Ia"
		default:
			s.size = "II"
		}
	case -4:
		switch s.spectralType {
		case "O", "B", "A":
			s.size = "Ib"
		case "M":
			s.size = "II"
		default:
			s.size = "III"
		}
	case -3:
		if s.spectralType == "F" || s.spectralType == "G" {
			s.size = "IV"
		} else {
			s.size = "II"
			if s.spectralType == "K" {
				if s.spectralDecimal >= 5 {
					s.size = "V"
				} else {
					s.size = "IV"
				}
			}
		}
	case -2:
		if s.spectralType == "F" || s.spectralType == "G" || s.spectralType == "K" {
			s.size = "V"
		} else {
			s.size = "III"
		}
	case -1:
		switch s.spectralType {
		case "O", "B":
			s.size = "III"
		case "A":
			s.size = "IV"
		default:
			s.size = "V"
		}
	case 0:
		if s.spectralType == "O" || s.spectralType == "B" {
			s.size = "III"
		} else {
			s.size = "V"
		}
	case 1:
		if s.spectralType == "B" {
			s.size = "III"
		} else {
			s.size = "V"
		}
	case 2, 3:
		s.size = "V"
	case 4:
		switch s.spectralType {
		case "A":
			s.size = "V"
		case "O", "B":
			s.size = "IV"
		case "F":
			if s.spectralDecimal < 5 {
				s.size = "V"
			} else {
				s.size = "VI"
			}
		default:
			s.size = "VI"
		}
	case 5:
		s.size = "D"
	default:
		if isHomestar {
			s.size = "D"
		} else {
			switch s.spectralType { // Stars other than primary
			case "O", "B":
				s.size = "IV"
			case "A":
				s.size = "V"
			default:
				s.size = "VI"
				if s.spectralDecimal < 5 {
					s.size = "V"
				}
			}
		}
	}
	if s.size == "D" {
		s.spectralDecimal = 0
	}
	s.description = s.getDescription()
	//s.mass = s.getMass()
	if isHomestar {
		s.orbit = -1
	}
	return
}

// parseStars takes a string containing the list of star(s) for a world, and populates a slice of pointers to starDetail structs. This slice is returned.
// The difficulty (or note to be taken) is that the string containing a list of stars for a system contains no other information other than the type and
// size, and the number. Information about how far any companion stars are from the primary, or in fact whether a particular star is a companion star, or
// the orbits occupied by particular stars are not found in the string, and so has to be guessed or determined.
func parseStars(s string) (ss []*starDetail) {

	// TODO: Complete this (https://github.com/mjlumley/traveller/issues/8)

	// A star system can hold up to eight stars, primary + companion, close star + companion, near star + companion, far star + companion.
	// Determining where they are and which they are depends on the statistical likelihood, and the size of the original.

	// Stars will appear only as one of the following formats: "xy z" "D" "Dx" "BD", where
	// - x is spectral type (O,B,A,F,G,K,M)
	// - y is spectral decimal (0 - 9)
	// - z is stellar size (O, Ia, Ib, II, II, IV, V, VI)
	// - D is literal "White Dwarf". This may include spectral decimal or it may be bare.
	// - BD is literal "Brown Dwarf"

	// For this simple system though, (ie INITIALLY) we will simply return a slice of individual starDetails in the ss slice.

	// Simple case.
	if len(s) < 1 {
		return nil
	}

	// The strategy is to split the incoming string on the spaces, and then progressively to "consume" the parts to construct the list of stars.
	parts := strings.Split(s, " ")

	for i := 0; i < len(parts); i++ {

		var star starDetail
		// Examine the string
		str := parts[i]

		switch str[0:1] {
		case "O", "A", "F", "G", "K", "M", "B":
			star.spectralDecimal = 0
			if len(str) < 2 {
				star.spectralType = str[0:1]
			} else {
				// "B" is either going to be a Brown Dwarf or Spectral Type B.
				if str == "BD" {
					star.spectralType = "BD"
					star.size = ""
					break
				}
				star.spectralType = str[0:1]
				if num, err := strconv.Atoi(str[1:2]); err == nil {
					star.spectralDecimal = num
				} else {
					log.Printf("Unable to parse spectral decimal from %s. Using 0.", str)
				}
			}
			// Now we attempt to get the next part of the star descriptor, stellar size.
			if i+1 >= len(parts) {
				// We will be unable to get the next part, this is likely malformed.
				log.Printf("Missing stellar size part for star")
				continue
			}
			size := parts[i+1]
			switch size {
			case "O", "Ia", "Ib", "II", "III", "IV", "V", "VI":
				star.size = size
				i++
			default:
				log.Printf("Unable to determine stellar size for part %s", size)
			}
		case "D":
			// White dwarfs can either be just "D" or "Dx" where x is the spectral type.
			star.size = "D"
			if len(str) == 1 {
				star.spectralType = ""
			} else {
				star.spectralType = str[1:2]
			}
		default:
			// Problems here
			log.Printf("Unable to parse star partial data %s", str)
			continue
		}
		star.description = star.getDescription()
		log.Printf("Star detail: %v", star)
		ss = append(ss, &star)

	}

	return
}

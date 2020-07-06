package main

import (
	"log"
	"strings"
)

// worldGen.go contains code for world generation. The code is ONLY for the world generation process,
// not the use of the world, stars, sector or subsector objects. Basically if the dice needs to be
// rolled for something, it is here (and the function should start with "determineXXX"). If it is
// just retrieving or setting information about a world then it should be setXXX or getXXX and be
// in worlds.go.

// worldGenState contains infomration about the world's generation, it's current stage of generation
// (if multi-stage) and the type of generation used.
// type worldGenState struct {
// 	genType         string
// 	completionState string
// }

// WorldGenType defines the type of world generation used for this world.
type WorldGenType int

// Constants for generation type
const (
	// WgtCt03 is used for Classic Traveller Book 3.
	WgtCt03 WorldGenType = iota
	// WgtCt06 is used for Classic Traveller Book 6.
	WgtCt06
	// WgtMtBasic is used for MegaTraveller Basic.
	WgtMtBasic
	// WgtMtExtended is used for MegaTraveller Extended.
	WgtMtExtended
	// WgtMtWBH is used for World Builders Handbook (MT).
	WgtMtWBH
	// WgtT5ss is used for Traveller5 Second Survey.
	WgtT5ss
	// WgtInvalid is used to indicate a generation that has failed.
	WgtInvalid
)

// Header out prints out a header for the various types of worlds generated.
var headerOut [6]string

func init() {
	headerOut[WgtCt03] = "Sector\tSS\tHex\tName\tUWP\tBases\tRemarks\tZone"
	headerOut[WgtMtBasic] = headerOut[WgtCt03] + "\tPBG\tAllegiance"
	headerOut[WgtT5ss] = headerOut[WgtMtBasic] + "\tStars\t{Ix}\t(Ex)\t[Cx[]\tNobility\tW\tRU"
}

// String displays a string representing the type of world generation process.
func (g WorldGenType) String() string {
	return [...]string{"Classic Traveller Book 3", "Classic Traveller Book 6", "MegaTraveller Basic", "MegaTraveller Extended", "World Builders Handbook", "Traveller5 Second Survey"}[g]
}

// generateCT03World generates a basic Classic Traveller world with the given
// basic information. It returns the word generated.
func generateCT03World(name, hexLoc, sector string) (w world) {

	w.name = name
	hloc := NewHexLoc(hexLoc, true)
	if hloc == nil {
		w.genType = WgtInvalid
		return
	}
	w.hexLoc = *hloc
	w.subsectorIndex = w.hexLoc.GetIndex()
	w.sector = sector
	w.sectorAbbrev = getAbbreviationForSector(sector)
	w.allegiance = basicAllegianceMap["Imperial"]

	ss, err := getSubsectorBySectorNameAndIndex(sector, w.subsectorIndex)
	log.Printf("Subsector is %v", ss)
	if err == nil {
		w.subsector = ss.name
	}

	w.genType = WgtCt03

	// Generate System contents
	// Starport
	w.uwp.starport = determineStarport(ssStandard)
	// Bases
	w.determineBases()
	// Gas Giant
	if D6()+D6() <= 9 {
		w.pbg.gasGiants = 1
	}

	// Generate UWP
	w.uwp.createWorldBasic()
	// Adjustments for Classic Traveller Book 3
	if w.uwp.lawInt > 15 {
		w.uwp.lawInt = 15
	}
	if w.uwp.techInt > 15 {
		w.uwp.techInt = 15
	}

	// Determine Trade Classifications (Basic)
	w.remarks = w.determineTradeClassifications()

	// Finished - return result
	return
}

// generateMTWorld generates a basic MegaTraveller world with the given basic information.
// It returns the word generated.
func generateMTWorld(name, hexLoc, sector, allegiance, traffic string) (w world) {

	w.name = name
	hloc := NewHexLoc(hexLoc, true)
	if hloc == nil {
		w.genType = WgtInvalid
		return
	}
	w.hexLoc = *hloc
	w.subsectorIndex = w.hexLoc.GetIndex()
	w.sector = sector
	w.sectorAbbrev = getAbbreviationForSector(sector)
	w.allegiance = basicAllegianceMap[allegiance]
	w.genType = WgtMtBasic

	if ss, err := getSubsectorBySectorNameAndIndex(sector, w.subsectorIndex); err == nil {
		w.subsector = ss.name
	}

	// Generate System contents

	// Step 3. Starport
	sparcity := ssStandard
	for ss := range mtSubsectorTrafficArr {
		if traffic == mtSubsectorTrafficArr[ss] {
			sparcity = ss
		}
	}
	w.uwp.starport = determineStarport(sparcity)

	// Step 4 - 10. Create UWP for world
	w.uwp.createWorldBasic()
	// Adjustments for MegaTraveller
	if w.uwp.lawInt > 20 {
		w.uwp.lawInt = 20
	}
	if w.uwp.govInt == 14 || w.uwp.govInt == 15 {
		w.uwp.techInt = w.uwp.techInt - 1
	}
	if w.uwp.starport == "F" {
		w.uwp.techInt = w.uwp.techInt + 1
	}
	if w.uwp.techInt > 15 {
		w.uwp.techInt = 15
	}

	// Step 11. Bases
	w.determineBases()

	// Step 12. Determine Trade Classifications (Basic)
	w.remarks = w.determineTradeClassifications()
	//displayObject(w.ObjectBasicString())

	// Step 13. Supplemental Remarks (none)
	// Step 14. Population Multiplier.
	//w.pbg.populationDigit = Dice(10) - 1
	w.pbg.populationDigit = Dice(9)
	// Step 15. Gas Giants
	if D6()+D6() >= 5 {
		roll := D6() + D6()
		switch roll {
		case 2, 3:
			w.pbg.gasGiants = 1
		case 4, 5:
			w.pbg.gasGiants = 2
		case 6, 7:
			w.pbg.gasGiants = 3
		case 8, 9, 10:
			w.pbg.gasGiants = 4
		case 11, 12:
			w.pbg.gasGiants = 5
		}
	}
	// Step 16. Planetoid Belts
	if D6()+D6()+w.pbg.gasGiants >= 8 {
		roll := D6() + D6()
		switch roll {
		case 2, 3, 4, 5, 6, 7:
			w.pbg.planetoids = 1
		case 13:
			w.pbg.planetoids = 3
		default:
			w.pbg.planetoids = 2
		}
	}
	// Step 17. Travel Zone
	w.determineZone()

	// Finished, return the result.
	return

}

// generateT5World generates a single Traveller5 Second Survey mainworld, based on the provided details.
// It returns the world generated.
func generateT5World(name, hexLoc, sector, allegianceCode string) (w world) {

	w.genType = WgtT5ss
	w.name = name
	hloc := NewHexLoc(hexLoc, true)
	if hloc == nil {
		w.genType = WgtInvalid
		return
	}
	w.hexLoc = *hloc
	w.allegiance = allegianceCode
	w.sector = sector
	w.subsectorIndex = w.hexLoc.GetIndex()
	w.sectorAbbrev = getAbbreviationForSector(w.sector)

	ss, err := getSubsectorBySectorNameAndIndex(sector, w.subsectorIndex)
	if err == nil {
		w.subsector = ss.name
	}

	// ---- Step B ---- Basic System features
	starSpectralFlux := Flux(0)
	starSizeFlux := Flux(0)

	//The "Primary" is the main or homeworld star in a star system.
	primary := determineStar(0, starSpectralFlux, starSizeFlux, true)
	if primary.spectralType == "" {
		log.Panic("Error: Primary Star generation has failed.")
		return
	}
	// Add the primary star to the world's information
	w.stars = append(w.stars, &primary)

	// Get some stellar information
	primaryDetail := getStellarDetail(primary.String())

	// Determine the world's habitable zone and orbit
	w.habZoneVar = determineHabitableZoneVariance(primary)
	w.orbit = primaryDetail.habitableZone + w.habZoneVar
	// Possibly adjust for a minimum orbit. Both habitable zone variance and orbit will need to change.
	if w.orbit < primaryDetail.minOrbit {
		w.habZoneVar = w.habZoneVar + primaryDetail.minOrbit
		w.orbit = primaryDetail.minOrbit
	}

	//climate, _ := w.getClimate()
	w.planetOrSat = determineMainworldType()
	w.mwSatGG = false
	if strings.Contains(w.planetOrSat, "Sa") {
		if Flux(0) <= 0 {
			w.mwSatGG = true
		}
		w.satOrbit = determineSatOrbit(w.mwSatGG, strings.Contains(w.planetOrSat, "Close"))
	}
	w.pbg = determinePBG(w.mwSatGG)

	// ---- Step F ---- WorldGen Additional Data (Stellar)
	w.generateSystemStars(starSpectralFlux, starSizeFlux)

	// ---- Step C ---- Generate the UWP
	w.uwp = createWorld(wtMainworld, w.uwp, w.habZoneVar)

	// Adjustments
	if strings.Contains(w.allegiance, "Zh") {
		if w.uwp.popInt < 4 {
			w.uwp.popInt = 0
		}
		if w.uwp.techInt > 14 {
			w.uwp.techInt = 14
		}
	}
	if w.uwp.popInt == 0 {
		w.pbg.populationDigit = 0
	}

	// ---- Step C ---- WorldGen Trade Classes and Zones
	w.remarks = w.determineTradeClassifications()
	w.determineZone()
	w.worlds = 1 + w.pbg.gasGiants + w.pbg.planetoids + D6() + D6()

	// Bases
	w.determineBases()

	// ---- Step E ---- Extensions
	w.determineExtensions()

	return
}

// createWorldBasic create the physical and population details of mainworld using
// CT Book 3 rules.
func (u *worldUwp) createWorldBasic() {

	var dm int // Generic Dice Modifier

	// Size
	u.sizeInt = D6() + D6() - 2
	// Atmosphere
	u.atmInt = D6() + D6() - 7 + u.sizeInt
	if u.atmInt < 0 || u.sizeInt == 0 {
		u.atmInt = 0
	}
	// Hydrographics
	dm = 0
	if u.atmInt < 2 || u.atmInt > 9 {
		dm = -4
	}
	u.hydInt = D6() + D6() - 7 + u.atmInt + dm
	if u.sizeInt == 0 || u.hydInt < 0 {
		u.hydInt = 0
	}
	if u.hydInt > 10 {
		u.hydInt = 10
	}
	// Population
	u.popInt = D6() + D6() - 2
	// Government
	u.govInt = D6() + D6() - 7 + u.popInt
	if u.govInt > 15 {
		u.govInt = 15
	}
	if u.govInt < 0 {
		u.govInt = 0
	}
	// Law Level
	u.lawInt = D6() + D6() - 7 + u.govInt

	// Law level limit varies depending on the generation system. Leave this to the client,
	// as it does not affect anything else below this.
	if u.lawInt < 0 {
		u.lawInt = 0
	}

	// Tech Level
	dm = 0
	switch u.starport {
	case "A":
		dm += 6
	case "B":
		dm += 4
	case "C":
		dm += 2
	case "X":
		dm += -4
	}
	switch u.sizeInt {
	case 0, 1:
		dm += 2
	case 2, 3, 4:
		dm++
	}
	switch u.atmInt {
	case 0, 1, 2, 3:
		dm++
	case 10, 11, 12, 13, 14, 15:
		dm++
	}
	switch u.hydInt {
	case 9:
		dm++
	case 10:
		dm += 2
	}
	switch u.popInt {
	case 1, 2, 3, 4, 5:
		dm++
	case 9:
		dm += 2
	case 10, 11, 12, 13, 14, 15:
		dm += 4
	}
	switch u.govInt {
	case 0, 5:
		dm++
	case 13:
		dm += -2
	}
	u.techInt = D6() + dm
	if u.techInt < 0 {
		u.techInt = 0
	}

	return
}

// createWorld creates a Mainworld or secondary world (out of almost nothing - how about that?) and returns a world UWP structure containing all the cool (but basic) stuff.
// Parameters are the worldType, the mainworld UWP, and the habitable zone variance.
// If you want to create a mainworld, set worldType to "" and habitable zone is ignored. hzVariance should
// be set to negative, postive or zero, a worlds cal. Returns the new world UWP.
func createWorld(worldType string, uwp worldUwp, hzVariance int) (ret worldUwp) {

	var dm int // Generic Dice Modifier

	if worldType != wtMainworld {
		log.Panic("At this time, only able to create Mainworld. Please come back later.")
		return uwp
	}

	ret.starport = determineStarport(ssStandard)

	// Size
	ret.sizeInt = D6() + D6() - 2
	if ret.sizeInt == 10 {
		ret.sizeInt = D6() + 9
	}
	// Atmosphere
	ret.atmInt = ret.sizeInt + Flux(0)
	if ret.atmInt < 0 || ret.sizeInt == 0 {
		ret.atmInt = 0
	}
	if ret.atmInt > 15 {
		ret.atmInt = 15
	}
	// Hydrographics
	dm = 0
	if ret.atmInt < 2 || ret.atmInt > 9 {
		dm = -4
	}
	ret.hydInt = Flux(0) + ret.atmInt + dm
	if ret.sizeInt < 2 || ret.hydInt < 0 {
		ret.hydInt = 0
	}
	if ret.hydInt > 10 {
		ret.hydInt = 10
	}
	// Population
	ret.popInt = D6() + D6() - 2
	if ret.popInt == 10 {
		ret.popInt = D6() + D6() + 3
	}
	// Government
	ret.govInt = Flux(0) + ret.popInt
	if ret.govInt > 15 {
		ret.govInt = 15
	}
	if ret.govInt < 0 {
		ret.govInt = 0
	}
	// Law Level
	ret.lawInt = Flux(0) + ret.govInt
	if ret.lawInt > 18 {
		ret.lawInt = 18
	}
	if ret.lawInt < 0 {
		ret.lawInt = 0
	}
	// Tech Level
	dm = 0
	switch ret.starport {
	case "A":
		dm += 6
	case "B":
		dm += 4
	case "C":
		dm += 2
	case "X":
		dm += -4
	}
	switch ret.sizeInt {
	case 0, 1:
		dm += 2
	case 2, 3, 4:
		dm++
	}
	switch ret.atmInt {
	case 0, 1, 2, 3:
		dm++
	case 10, 11, 12, 13, 14, 15:
		dm++
	}
	switch ret.hydInt {
	case 9:
		dm++
	case 10:
		dm += 2
	}
	switch ret.popInt {
	case 1, 2, 3, 4, 5:
		dm++
	case 9:
		dm += 2
	case 10, 11, 12, 13, 14, 15:
		dm += 4
	}
	switch ret.govInt {
	case 0, 5:
		dm++
	case 13:
		dm += -2
	}
	ret.techInt = D6() + dm
	if ret.techInt < 0 {
		ret.techInt = 0
	}

	return
}

// determineSatOrbit gets the Mainworld satellite orbit name based on mainworld host body (GG or planet) and orbit zone. Returns the orbit name as string.
func determineSatOrbit(gg, close bool) string {
	var dm int

	if gg {
		dm = -2
	} else {
		dm = 2
	}
	roll := Flux(dm)
	if roll < -6 {
		roll = -6
	}
	if roll > 6 {
		roll = 6
	}
	if close {
		roll = roll + 6
	} else {
		roll = roll + 19
	}
	return satelliteOrbit[roll]
}

// determineStarport determines the Starport Type. Returns a random value A thru E or X.
// You should provide an integer indicating how well-travelled this particular subsector
// is (aka the sparcity). Use the "ss" constants. Standard is ssStandard (=1)
func determineStarport(sparcity int) string {
	roll := D6() + D6()
	switch sparcity {
	case ssBackwater:
		switch roll {
		case 2, 3:
			return "A"
		case 4, 5:
			return "B"
		case 6, 7, 8:
			return "C"
		case 9:
			return "D"
		case 10, 11:
			return "E"
		default:
			return "X"
		}
	case ssCluster:
		switch roll {
		case 2, 3, 4, 5:
			return "A"
		case 6, 7:
			return "B"
		case 8, 9:
			return "C"
		case 10:
			return "D"
		case 11:
			return "E"
		default:
			return "X"
		}
	default:
		switch roll {
		case 2, 3, 4:
			return "A"
		case 5, 6:
			return "B"
		case 7, 8:
			return "C"
		case 9:
			return "D"
		case 10, 11:
			return "E"
		default:
			if sparcity == ssMature {
				return "E"
			}
			return "X"
		}
	}
}

// determineMainworldType determines the mainworld type, Planet or Satellite(Close|Far). Returns the type as string.
func determineMainworldType() string {
	switch roll := Flux(0); {
	case roll <= -4:
		return mwTypeFarSatellite
	case roll == -3:
		return mwTypeCloseSatellite
	default:
		return mwTypePlanet
	}
}

// determineHabitableZoneVariance determines mainworld orbit modifier based on Star Spectral type and flux. This affects climate and Trade classifications
// that are based on the climate. Returns value -7 to +7.
func determineHabitableZoneVariance(s starDetail) int {
	var dm int
	switch s.spectralType {
	case "M":
		dm = 2
	case "O", "B":
		dm = -2
	default:
		dm = 0
	}

	x := Flux(dm)
	if x <= -6 {
		return -2
	}
	if x <= -3 {
		return -1
	}
	if x <= 2 {
		return 0
	}
	if x <= 5 {
		return 1
	}
	return 2
}

// determineExtensions determines all Extensions and Nobility for a world. Returns the updated world.
func (w *world) determineExtensions() *world {
	// Importance
	w.determineImportanceExtension()

	// Economic Extensions
	w.determineEconomicExtension()
	w.ru = w.economics.calcRU()

	// Cultural Extension
	w.determineCulturalExtension()

	// Nobility
	w.getNobility()

	return w

}

// determineImportanceExtension gets the Importance for a world, which looks as a string like this "{ +/-x }",
// where x is between -3 and +8.
func (w *world) determineImportanceExtension() (i importanceExt) {

	// Importance
	importInt := 0
	switch w.uwp.starport {
	case "A", "B":
		importInt++
	case "C":
		break
	default:
		importInt--

	}
	if w.uwp.techInt >= 16 {
		importInt++
	}
	if w.uwp.techInt >= 10 {
		importInt++
	}
	if w.uwp.techInt <= 8 {
		importInt--
	}
	if strings.Contains(w.remarks, "Ag") {
		importInt++
	}
	if strings.Contains(w.remarks, "Hi") {
		importInt++
	}
	if strings.Contains(w.remarks, "Ri") {
		importInt++
	}
	if w.uwp.popInt <= 6 {
		importInt--
	}
	if strings.Contains(w.bases, "S") && (strings.ContainsAny(w.bases, "DKN")) {
		importInt++
	}
	if strings.Contains(w.bases, "W") {
		importInt++
	}
	w.importance.Importance = importInt
	return w.importance
}

// determineCulturalExtension determines the Cultural extension for a world, which looks like this "[HASs]" where
// H=homogenity, A=acceptance, S=strangeness, and s=Symbols, all expressed in Extended Hex.
// This is returned as a cultureExt type.
func (w *world) determineCulturalExtension() cultureExt {

	w.culture.Homogenity = Flux(w.uwp.popInt)
	if w.culture.Homogenity < 1 {
		w.culture.Homogenity = 1
	}
	w.culture.Acceptance = w.uwp.popInt + w.importance.Importance
	if w.culture.Acceptance < 1 {
		w.culture.Acceptance = 1
	}
	w.culture.Strangeness = Flux(5)
	if w.culture.Strangeness < 1 {
		w.culture.Strangeness = 1
	}
	w.culture.Symbols = Flux(w.uwp.techInt)
	if w.culture.Symbols < 1 {
		w.culture.Symbols = 1
	}
	if w.uwp.popInt == 0 {
		w.culture.Homogenity = 0
		w.culture.Acceptance = 0
		w.culture.Strangeness = 0
		w.culture.Symbols = 0
	}
	return w.culture
}

// determineEconomicExtension determines the Economic Extension for a world. Economic extension is in the form "(RLI+/-E)", where R=resources,
// L=labour, I=infrastructure, and E=+/- efficiency. Resource units is calculated from this. Function returns a economicExt struct.
func (w *world) determineEconomicExtension() economicExt {

	w.economics.Resource = D6() + D6()
	if w.uwp.techInt >= 8 {
		w.economics.Resource += w.pbg.gasGiants + w.pbg.planetoids
	}
	if w.economics.Resource < 0 {
		w.economics.Resource = 0
	}
	w.economics.Labour = w.uwp.popInt - 1
	if w.economics.Labour < 0 {
		w.economics.Labour = 0
	}
	w.economics.Infrastructure = D6() + D6() + w.importance.Importance
	if strings.Contains(w.remarks, "Ba") && strings.Contains(w.remarks, "Di") && strings.Contains(w.remarks, "Lo") {
		w.economics.Infrastructure = 0
	} else if strings.Contains(w.remarks, "Lo") {
		w.economics.Infrastructure = 1
	}
	if strings.Contains(w.remarks, "Ni") {
		w.economics.Infrastructure = D6() + w.importance.Importance
	}
	if w.economics.Infrastructure < 0 {
		w.economics.Infrastructure = 0
	}
	w.economics.Efficiency = Flux(0)

	return w.economics
}

// determineBases determines the bases for a mainworld/system. It returns the Bases string
// but also sets the bases string in the world object. It bases (heh) the generation method
// on the worldGenState object, and does NOT ask the user for input.
//
// In some cases, percentage for some rolls use equivalents on 2d6, ie 10+ = 4-.
func (w *world) determineBases() {

	w.bases = ""
	if w.uwp.starport == "E" || w.uwp.starport == "X" {
		return
	}
	imperial := false
	if strings.Contains(w.allegiance, "Im") {
		imperial = true
	}

	switch w.genType {
	case WgtCt03:
		if w.uwp.starport == "A" || w.uwp.starport == "B" {
			if D6()+D6() >= 8 {
				w.bases += "N"
			}
		}
		sbDM := 0
		switch w.uwp.starport {
		case "C":
			sbDM = -1
		case "B":
			sbDM = -2
		case "A":
			sbDM = -3
		}
		if D6()+D6()+sbDM >= 7 {
			w.bases += "S"
		}
	case WgtMtBasic:
		if !imperial {
			roll := D6() + D6()
			if (w.uwp.starport == "A" && roll >= 10) || (w.uwp.starport == "B" && roll >= 9) || (w.uwp.starport == "C" && roll >= 8) {
				w.bases = "M"
			}
			return
		}
		switch w.uwp.starport {
		case "A":
			if D6()+D6() >= 8 {
				w.bases += "N"
			}
			if D6()+D6() >= 10 {
				w.bases += "S"
			}
		case "B":
			if D6()+D6() >= 8 {
				w.bases += "N"
			}
			if D6()+D6() >= 9 {
				w.bases += "S"
			}
		case "C":
			if D6()+D6() >= 8 {
				w.bases += "S"
			}
		case "D":
			if D6()+D6() >= 7 {
				w.bases += "S"
			}
		}
	case WgtT5ss:
		switch w.uwp.starport {
		case "A":
			if D6()+D6() <= 6 {
				w.bases += "N"
			}
			if D6()+D6() <= 4 {
				w.bases += "S"
			}
		case "B":
			if D6()+D6() <= 5 {
				w.bases += "N"
			}
			if D6()+D6() <= 5 {
				w.bases += "S"
			}
		case "C":
			if D6()+D6() <= 6 {
				w.bases += "S"
			}
		case "D":
			if D6()+D6() <= 7 {
				w.bases += "S"
			}
		}
	}
	return
}

// determineZone determines the travel zone for a mainworld based on the world characteristics.
func (w *world) determineZone() {
	amberZone := false
	redZone := false

	w.zone = TzGreen

	switch w.genType {
	case WgtMtBasic:
		if w.uwp.starport == "X" {
			redZone = true
		} else {
			switch w.uwp.govInt {
			case 10:
				if w.uwp.lawInt == 20 {
					amberZone = true
				}
			case 11:
				if w.uwp.lawInt >= 19 {
					amberZone = true
				}
			case 12:
				if w.uwp.lawInt >= 18 {
					amberZone = true
				}
			case 13:
				if w.uwp.lawInt >= 17 && w.uwp.lawInt <= 19 {
					amberZone = true
				} else if w.uwp.lawInt == 20 {
					redZone = true
				}
			case 14:
				if w.uwp.lawInt == 17 || w.uwp.lawInt == 18 {
					amberZone = true
				} else if w.uwp.lawInt >= 19 {
					redZone = true
				}
			case 15:
				if w.uwp.lawInt == 16 || w.uwp.lawInt == 17 {
					amberZone = true
				} else if w.uwp.lawInt >= 18 {
					redZone = true
				}
			}
		}
	case WgtT5ss:
		if w.uwp.govInt+w.uwp.lawInt >= 20 {
			amberZone = true
		}
		if w.uwp.govInt+w.uwp.lawInt >= 22 {
			redZone = true
		}
		if w.uwp.starport == "X" {
			redZone = true
		}

		// For Zhodani - assign some amber zones.
		if strings.Contains(w.allegiance, basicAllegianceMap["Zhodani"]) && !redZone && (w.uwp.govInt == 0 || w.uwp.govInt == 7 || w.uwp.govInt >= 13 || w.uwp.techInt <= 7) {
			if Dice(2) == 1 {
				amberZone = true
			}
		}
	}
	if amberZone {
		w.zone = TzAmber
	}
	if redZone {
		w.zone = TzRed
	}
}

// determinePBG generate the PBG fields for a homeworld star system. If mwSatGG is true, then the mainworld is a satellite of a Gas Giant.
func determinePBG(mwSatGG bool) (p worldPBG) {

	p.populationDigit = Dice(9)
	p.planetoids = D6() - 3 // 1d6-3
	if p.planetoids < 0 {
		p.planetoids = 0
	}
	p.gasGiants = (D6()+D6())/2 - 2 // 2d6/2-2
	if p.gasGiants <= 0 {
		if mwSatGG {
			p.gasGiants = 1
		} else {
			p.gasGiants = 0
		}
	}
	return
}

// generateSystemStars generates the stars for a system, given the "flux" for the
// Primary Spectral and Size. Returns the updated world struct.
func (w *world) generateSystemStars(starSpectralFlux, starSizeFlux int) *world {
	// ---- Step F ---- WorldGen Additional Data
	// Determine is there is a primary companion
	var closeStar starDetail
	var nearStar starDetail
	var farStar starDetail

	// Determine if we have a companion to the Primary star
	if Flux(0) >= 3 {
		pCompanion := determineStar(0, starSpectralFlux, starSizeFlux, false)
		w.stars[0].companion = &pCompanion
	}
	// Close star and companion
	if Flux(0) >= 3 {
		closeStar = determineStar(0, starSpectralFlux, starSizeFlux, false)
		closeStar.orbit = D6() - 1 // 1d6-1
		//closeStar.habitableZone = closeStar.getHabitableZone()
		if Flux(0) >= 3 {
			closeCompanion := determineStar(0, starSpectralFlux, starSizeFlux, false)
			closeStar.companion = &closeCompanion
		}
		w.stars = append(w.stars, &closeStar)
	}
	// Near star and companion
	if Flux(0) >= 3 {
		nearStar = determineStar(0, starSpectralFlux, starSizeFlux, false)
		nearStar.orbit = 5 + D6() // 1d6+5
		if Flux(0) >= 3 {
			nearCompanion := determineStar(0, starSpectralFlux, starSizeFlux, false)
			nearStar.companion = &nearCompanion
		}
		w.stars = append(w.stars, &nearStar)
	}
	// Far star and companion
	if Flux(0) >= 3 {
		farStar = determineStar(0, starSpectralFlux, starSizeFlux, false)
		farStar.orbit = 11 + D6() // d6+11
		if Flux(0) >= 3 {
			farCompanion := determineStar(0, starSpectralFlux, starSizeFlux, false)
			farStar.companion = &farCompanion
		}
		w.stars = append(w.stars, &farStar)
	}
	return w
}

// determineTradeClassificationsBasic provides the Trade Classifications (aka remarks) for
// a Basic (Classic Traveller Book 3, MegaTraveller and T5) world. The TCs are returned as a string.
func (w world) determineTradeClassifications() (r string) {

	// Common for all types

	// ---- Planetary
	// Asteroid Belt && Vacuum
	if w.uwp.sizeInt == 0 && w.uwp.atmInt == 0 && w.uwp.hydInt == 0 {
		r += "As "
	} else if w.uwp.atmInt == 0 {
		r += "Va "
	}
	// Ice-capped
	if w.uwp.atmInt <= 1 && w.uwp.hydInt != 0 {
		r += "Ic "
	}
	// Agricultural
	if w.uwp.atmInt >= 4 && w.uwp.atmInt <= 9 && w.uwp.hydInt >= 4 && w.uwp.hydInt <= 8 && w.uwp.popInt >= 5 && w.uwp.popInt <= 7 {
		r += "Ag "
	}
	// Non-agricultural
	if w.uwp.atmInt <= 3 && w.uwp.hydInt <= 3 && w.uwp.popInt >= 6 {
		r += "Na "
	}

	// These are for both MT and T5SS worlds
	if w.genType == WgtMtBasic || w.genType == WgtT5ss {
		// Barren
		if w.uwp.popInt == 0 && w.uwp.govInt == 0 && w.uwp.lawInt == 0 {
			r += "Ba "
		}
		// Fluid
		if w.uwp.atmInt >= 10 && w.uwp.atmInt <= 12 && w.uwp.hydInt >= 1 {
			r += "Fl "
		}
		// High Population
		if w.uwp.popInt >= 9 {
			r += "Hi "
		}
		// Low Population
		if w.uwp.popInt >= 1 && w.uwp.popInt <= 3 {
			r += "Lo "
		}
	}

	// These remainder vary for each type of generation :-/

	switch w.genType {
	case WgtCt03:
		// Desert
		if w.uwp.sizeInt != 0 && w.uwp.hydInt == 0 {
			r += "De "
		}
		// Water world
		if w.uwp.hydInt == 10 {
			r += "Wa "
		}
		// Poor
		if w.uwp.atmInt >= 2 && w.uwp.atmInt <= 5 && w.uwp.hydInt <= 3 {
			r += "Po "
		}
		// Industrial
		if (w.uwp.atmInt <= 2 || w.uwp.atmInt == 4 || w.uwp.atmInt == 7 || w.uwp.atmInt == 9) && w.uwp.popInt >= 9 {
			r += "In "
		}
		// Non-industrial
		if w.uwp.popInt <= 6 {
			r += "Ni "
		}
		// Rich
		if (w.uwp.atmInt == 6 || w.uwp.atmInt == 8) && w.uwp.popInt >= 6 && w.uwp.popInt <= 8 && w.uwp.govInt >= 4 && w.uwp.govInt <= 9 {
			r += "Ri "
		}
	case WgtMtBasic:
		// Desert
		if w.uwp.sizeInt != 0 && w.uwp.hydInt == 0 && w.uwp.sizeInt >= 2 {
			r += "De "
		}
		// Water world
		if w.uwp.hydInt == 10 {
			r += "Wa "
		}
		// Industrial
		if (w.uwp.atmInt <= 2 || w.uwp.atmInt == 4 || w.uwp.atmInt == 7 || w.uwp.atmInt == 9) && w.uwp.popInt >= 9 {
			r += "In "
		}
		// Non-industrial
		if w.uwp.popInt <= 6 && w.uwp.popInt >= 1 {
			r += "Ni "
		}
		// Poor
		if w.uwp.atmInt >= 2 && w.uwp.atmInt <= 5 && w.uwp.hydInt <= 3 && w.uwp.popInt > 0 {
			r += "Po "
		}
		// Rich
		if (w.uwp.atmInt == 6 || w.uwp.atmInt == 8) && w.uwp.popInt >= 6 && w.uwp.popInt <= 8 && strings.Contains(w.allegiance, basicAllegianceMap["Aslan"]) {
			r += "Ri "
		}

	case WgtT5ss:
		// Desert
		if w.uwp.atmInt >= 2 && w.uwp.atmInt <= 9 && w.uwp.hydInt == 0 {
			r += "De "
		}
		// Water world
		if w.uwp.sizeInt >= 3 && w.uwp.sizeInt <= 9 && w.uwp.atmInt >= 3 && w.uwp.atmInt <= 12 && w.uwp.hydInt == 10 {
			r += "Wa "
		}
		// Poor
		if w.uwp.atmInt >= 2 && w.uwp.atmInt <= 5 && w.uwp.hydInt <= 3 {
			r += "Po "
		}
		// Industrial
		if (w.uwp.atmInt <= 2 || w.uwp.atmInt == 4 || w.uwp.atmInt == 7 || (w.uwp.atmInt >= 9 && w.uwp.atmInt <= 12)) && w.uwp.popInt >= 9 {
			r += "In "
		}
		// Non-industrial
		if w.uwp.popInt >= 4 && w.uwp.popInt <= 6 {
			r += "Ni "
		}
		// Garden World
		if w.uwp.sizeInt >= 6 && w.uwp.sizeInt <= 8 && (w.uwp.atmInt == 5 || w.uwp.atmInt == 6 || w.uwp.atmInt == 8) && w.uwp.hydInt >= 5 && w.uwp.hydInt <= 7 {
			r += "Ga "
		}
		// Hell World
		if w.uwp.sizeInt >= 3 && w.uwp.sizeInt <= 12 && (w.uwp.atmInt == 2 || w.uwp.atmInt == 4 || w.uwp.atmInt == 7 || (w.uwp.atmInt >= 9 && w.uwp.atmInt <= 12)) && w.uwp.hydInt <= 2 {
			r += "He "
		}
		// Ocean world
		if w.uwp.sizeInt >= 10 && w.uwp.atmInt >= 3 && w.uwp.atmInt <= 12 && w.uwp.hydInt == 10 {
			r += "Oc "
		}
		// Dieback
		if w.uwp.popInt == 0 && w.uwp.govInt == 0 && w.uwp.lawInt == 0 && w.uwp.techInt > 0 {
			r += "Di "
		}
		// Pre-High
		if w.uwp.popInt == 8 {
			r += "Ph "
		}
		// Pre-Agricultural
		if w.uwp.atmInt >= 4 && w.uwp.atmInt <= 9 && w.uwp.hydInt >= 4 && w.uwp.hydInt <= 8 && (w.uwp.popInt == 4 || w.uwp.popInt == 8) {
			r += "Pa "
		}
		// Prison or Exile
		if (w.uwp.atmInt == 2 || w.uwp.atmInt == 3 || w.uwp.atmInt == 10 || w.uwp.atmInt == 11) && w.uwp.hydInt >= 1 && w.uwp.hydInt <= 5 && w.uwp.popInt >= 3 && w.uwp.popInt <= 6 && w.uwp.lawInt >= 6 && w.uwp.lawInt <= 9 {
			r += "Px "
		}
		// Pre-Industrial
		if (w.uwp.atmInt <= 2 || w.uwp.atmInt == 4 || w.uwp.atmInt == 7 || w.uwp.atmInt == 9) && (w.uwp.popInt == 7 || w.uwp.popInt == 8) {
			r += "Pi "
		}
		// Pre-Rich
		if (w.uwp.atmInt == 6 || w.uwp.atmInt == 8) && (w.uwp.popInt == 5 || w.uwp.popInt == 9) {
			r += "Pr "
		}
		// Rich
		if (w.uwp.atmInt == 6 || w.uwp.atmInt == 8) && w.uwp.popInt >= 6 && w.uwp.popInt <= 9 {
			r += "Ri "
		}
		// Frozen
		if w.habZoneVar >= 2 && w.uwp.sizeInt >= 2 && w.uwp.sizeInt <= 9 && w.uwp.hydInt != 0 {
			r += "Fr "
		}
		// Hot
		if w.habZoneVar == -1 {
			r += "Ho "
		}
		// Cold
		if w.habZoneVar == 1 {
			r += "Co "
		}
		// Locked
		if w.planetOrSat == mwTypeCloseSatellite {
			r += "Co "
		}
		// Tropic
		if w.uwp.sizeInt >= 6 && w.uwp.sizeInt <= 9 && w.uwp.atmInt >= 4 && w.uwp.atmInt <= 9 && w.uwp.hydInt >= 3 && w.uwp.hydInt <= 7 && w.habZoneVar == -1 {
			r += "Tr "
		}
		// Tundra
		if w.uwp.sizeInt >= 6 && w.uwp.sizeInt <= 9 && w.uwp.atmInt >= 4 && w.uwp.atmInt <= 9 && w.uwp.hydInt >= 3 && w.uwp.hydInt <= 7 && w.habZoneVar == 1 {
			r += "Tu "
		}

		// ---- Secondary
		// Twilight zone (Tz) - Orbit 0 or 1 (calculate later)
		// Farming (Fa) - HZ but not mainworld ATM 4-9, HYD 4-8, POP 2-6
		// Mining (Mi) - POP 2-6, Not MW, MW=In
		// Military Rule (Mr) - Ref
		// Penal Colony (Pe) - ATM 23AB, HYD 1-5, POP 3-6, GOV 6, LAW 6-9, Not MW
		// Reserve
		if w.uwp.popInt >= 1 && w.uwp.popInt <= 4 && w.uwp.govInt == 6 && w.uwp.lawInt >= 4 && w.uwp.lawInt <= 5 {
			r += "Re "
		}
		// // Colony
		// if uwp.popInt >= 5 && uwp.popInt <= 10 && uwp.govInt == 6 && uwp.lawInt <= 3 {
		// 	r += "Cy "
		// }
		// Far satellite
		if w.planetOrSat == mwTypeFarSatellite {
			r += "Sa "
		}
	}

	// ---- Population && Economic

	r = strings.TrimRight(r, " ")

	return
}

// extendWorld extends a basic world to T5SS standards and as a by-product, recalculates extensions.
// It returns the world worked on.
func (w *world) extendWorld() *world {

	w.genType = WgtT5ss

	var primary starDetail

	// Check that we have stars, if not generate them.
	if len(w.stars) == 0 {
		starSpectralFlux := Flux(0)
		starSizeFlux := Flux(0)

		//The "Primary" is the main or homeworld star in a star system.
		primary := determineStar(0, starSpectralFlux, starSizeFlux, true)
		w.stars = append(w.stars, &primary)

		// Get additional stars
		w.generateSystemStars(starSpectralFlux, starSizeFlux)
		// ---- Step F ---- WorldGen Additional Data (Stellar)
	} else {
		primary = *w.stars[0]
	}

	primaryDetail := getStellarDetail(primary.String())
	// Determine the world's habitable zone and orbit
	w.habZoneVar = determineHabitableZoneVariance(primary)
	w.orbit = primaryDetail.habitableZone + w.habZoneVar
	// Possibly adjust for a minimum orbit. Both habitable zone variance and orbit will need to change.
	if w.orbit < primaryDetail.minOrbit {
		w.habZoneVar = w.habZoneVar + primaryDetail.minOrbit
		w.orbit = primaryDetail.minOrbit
	}

	w.planetOrSat = determineMainworldType()
	w.mwSatGG = false
	if strings.Contains(w.planetOrSat, "Sa") {
		if Flux(0) <= 0 {
			w.mwSatGG = true
		}
		w.satOrbit = determineSatOrbit(w.mwSatGG, strings.Contains(w.planetOrSat, "Close"))
	}
	// Probably already have PBG
	//w.pbg = determinePBG(w.mwSatGG)
	w.worlds = 1 + w.pbg.gasGiants + w.pbg.planetoids + D6() + D6()
	w.determineExtensions()
	return w
}

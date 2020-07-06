package deprecated

/*
import (
	"fmt"
	"log"
	"strings"

	"github.com/atotto/clipboard"
)

// getTradeClassifications gets the Planetary, Population, Economic and Climate trade classifications for a world. They are returned
// as a string.
func (w world) getTradeClassificationsOld() (r string) {

	// ---- Planetary
	// Asteroid Belt & Vacuum
	if w.uwp.sizeInt == 0 && w.uwp.atmInt == 0 && w.uwp.hydInt == 0 {
		r += "As "
	} else if w.uwp.atmInt == 0 {
		r += "Va "
	}
	// Desert
	if w.uwp.atmInt >= 2 && w.uwp.atmInt <= 9 && w.uwp.hydInt == 0 {
		r += "De "
	}
	// Fluid
	if w.uwp.atmInt >= 10 && w.uwp.atmInt <= 12 && w.uwp.hydInt != 0 {
		r += "Fl "
	}
	// Garden World
	if w.uwp.sizeInt >= 6 && w.uwp.sizeInt <= 8 && (w.uwp.atmInt == 5 || w.uwp.atmInt == 6 || w.uwp.atmInt == 8) && w.uwp.hydInt >= 5 && w.uwp.hydInt <= 7 {
		r += "Ga "
	}
	// Hell World
	if w.uwp.sizeInt >= 3 && w.uwp.sizeInt <= 12 && (w.uwp.atmInt == 2 || w.uwp.atmInt == 4 || w.uwp.atmInt == 7 || (w.uwp.atmInt >= 9 && w.uwp.atmInt <= 12)) && w.uwp.hydInt <= 2 {
		r += "He "
	}
	// Ice-capped
	if w.uwp.atmInt <= 1 && w.uwp.hydInt != 0 {
		r += "Ic "
	}
	// Ocean world
	if w.uwp.sizeInt >= 10 && w.uwp.atmInt >= 3 && w.uwp.atmInt <= 12 && w.uwp.hydInt == 10 {
		r += "Oc "
	}
	// Water world
	if w.uwp.sizeInt >= 3 && w.uwp.sizeInt <= 9 && w.uwp.atmInt >= 3 && w.uwp.atmInt <= 12 && w.uwp.hydInt == 10 {
		r += "Wa "
	}
	// ---- Population
	// Dieback
	if w.uwp.popInt == 0 && w.uwp.techInt > 0 {
		r += "Di "
	}
	// Barren
	if w.uwp.popInt == 0 && w.uwp.govInt == 0 && w.uwp.lawInt == 0 && w.uwp.techInt == 0 {
		r += "Ba "
	}
	// Low Population
	if w.uwp.popInt >= 1 && w.uwp.popInt <= 3 {
		r += "Lo "
	}
	// Non-industrial
	if w.uwp.popInt >= 4 && w.uwp.popInt <= 6 {
		r += "Ni "
	}
	// Pre-High
	if w.uwp.popInt == 8 {
		r += "Ph "
	}
	// High-Population
	if w.uwp.popInt >= 9 {
		r += "Hi "
	}
	// ---- Economic
	// Pre-Agricultural
	if w.uwp.atmInt >= 4 && w.uwp.atmInt <= 9 && w.uwp.hydInt >= 4 && w.uwp.hydInt <= 8 && (w.uwp.popInt == 4 || w.uwp.popInt == 8) {
		r += "Pa "
	}
	// Agricultural
	if w.uwp.atmInt >= 4 && w.uwp.atmInt <= 9 && w.uwp.hydInt >= 4 && w.uwp.hydInt <= 8 && w.uwp.popInt >= 5 && w.uwp.popInt <= 7 {
		r += "Ag "
	}
	// Non-agricultural
	if w.uwp.atmInt <= 3 && w.uwp.hydInt <= 3 && w.uwp.popInt >= 6 {
		r += "Na "
	}
	// Prison or Exile
	if (w.uwp.atmInt == 2 || w.uwp.atmInt == 3 || w.uwp.atmInt == 10 || w.uwp.atmInt == 11) && w.uwp.hydInt >= 1 && w.uwp.hydInt <= 5 && w.uwp.popInt >= 3 && w.uwp.popInt <= 6 && w.uwp.lawInt >= 6 && w.uwp.lawInt <= 9 {
		r += "Px "
	}
	// Pre-Industrial
	if (w.uwp.atmInt <= 2 || w.uwp.atmInt == 4 || w.uwp.atmInt == 7 || w.uwp.atmInt == 9) && (w.uwp.popInt == 7 || w.uwp.popInt == 8) {
		r += "Pi "
	}
	// Industrial
	if (w.uwp.atmInt <= 2 || w.uwp.atmInt == 4 || w.uwp.atmInt == 7 || (w.uwp.atmInt >= 9 && w.uwp.atmInt <= 12)) && w.uwp.popInt >= 9 {
		r += "In "
	}
	// Poor
	if w.uwp.atmInt >= 2 && w.uwp.atmInt <= 5 && w.uwp.hydInt <= 3 {
		r += "Po "
	}
	// Pre-Rich
	if (w.uwp.atmInt == 6 || w.uwp.atmInt == 8) && (w.uwp.popInt == 5 || w.uwp.popInt == 9) {
		r += "Pr "
	}
	// Rich
	if (w.uwp.atmInt == 6 || w.uwp.atmInt == 8) && w.uwp.popInt >= 6 && w.uwp.popInt <= 9 {
		r += "Ri "
	}
	// ---- Climate
	// Frozen
	if w.habitableZone >= 2 && w.uwp.sizeInt >= 2 && w.uwp.sizeInt <= 9 && w.uwp.hydInt != 0 {
		r += "Fr "
	}
	// Hot
	if w.habitableZone == -1 {
		r += "Ho "
	}
	// Cold
	if w.habitableZone == 1 {
		r += "Co "
	}
	// Locked
	if w.planetOrSat == mwTypeCloseSatellite {
		r += "Co "
	}
	// Tropic
	if w.uwp.sizeInt >= 6 && w.uwp.sizeInt <= 9 && w.uwp.atmInt >= 4 && w.uwp.atmInt <= 9 && w.uwp.hydInt >= 3 && w.uwp.hydInt <= 7 && w.habitableZone == -1 {
		r += "Tr "
	}
	// Tundra
	if w.uwp.sizeInt >= 6 && w.uwp.sizeInt <= 9 && w.uwp.atmInt >= 4 && w.uwp.atmInt <= 9 && w.uwp.hydInt >= 3 && w.uwp.hydInt <= 7 && w.habitableZone == 1 {
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
		r += "Fa "
	}

	r = strings.TrimRight(r, " ")

	return
}

// determineBases determines the bases for a mainworld/system. Returns the Bases string.
// Set isAuto to true to not ask the user for input, isZhodani to true to determine
// Zhodani bases, and isBasic to true if using a CT03 setting.
//
// Note that for CT Book 3 puritans we have inverted the rolls (4 or less is equivalent
// to 10 or more - DMs taken into account). The percentage chances have not changed.
func (w world) determineBasesOld(isZhodani, isAuto, isBasic bool) (bases string) {
	// Bases
	navyBase := false
	scoutBase := false

	navyRoll := D6() + D6()
	scoutRoll := D6() + D6()

	switch w.uwp.starport {
	case "A":
		if navyRoll <= 6 {
			navyBase = true
			sLogLog("Navy base!")
		}
		if scoutRoll <= 4 {
			scoutBase = true
			sLogLog("Scout base!")
		}
	case "B":
		if navyRoll <= 5 && !isBasic {
			navyBase = true
			sLogLog("Navy base!")
		}
		if navyRoll == 6 && isBasic {
			navyBase = true
			sLogLog("Navy base!")
		}
		if scoutRoll <= 5 {
			scoutBase = true
			sLogLog("Scout base!")
		}
	case "C":
		if scoutRoll <= 6 {
			scoutBase = true
			sLogLog("Scout base!")
		}
	case "D":
		if scoutRoll <= 7 {
			scoutBase = true
			sLogLog("Scout base!")
		}
	}

	if isZhodani && scoutBase {
		scoutBase = false
	}

	// Ask the user for input regarding the bases
	if !isAuto {
		fmt.Println("\nThe world so far :")
		fmt.Println(w.String())

		baseString := "Bases : "
		if navyBase {
			baseString += "Navy"
		}
		if scoutBase {
			if navyBase {
				baseString += ", "
			}
			baseString += "Scout"
		}
		if !navyBase && !scoutBase {
			baseString += "No bases present"
		}
		fmt.Println(baseString)

		fmt.Println("\nIf you are happy with this, just hit enter for the next question, otherwise enter the bases that you want.")
		fmt.Println("Corsair-Aslan (C), Naval Depot (D), Embassy (E), Navy-Other (K), Miltary (M), Navy-Imp (N), ")
		fmt.Println("Clan-Aslan (R), Scout(S), Tlaukhu-Aslan(T), Exploration (V), Way Station (W).")
		fmt.Println("If you wish to remove all bases, choose ! (exclamation).")
		fmt.Println()
		bases = getChoice("Either enter the Base field (e.g. DS), or ! to clear, or Enter to accept as is : ")

	}
	if bases == "" {
		// Go with our calculated choices
		if navyBase {
			if isZhodani {
				bases = "K"
			} else {
				bases = "N"
			}
		}
		if scoutBase {
			bases += "S"
		}
	}
	if bases == "!" {
		bases = ""
	}
	return
}

// getAdditionalTCs asks the user for any additional Trade Classifications. Returns the updated TC field.
func (w world) getAdditionalTCs() string {
	fmt.Println()
	fmt.Println("Current UWP is " + w.uwp.String() + ", Trade Class is " + w.remarks + ", and zone is " + tableTravelZone(w.zone))
	fmt.Println()
	fmt.Println("You can add political or special TCs here. You will need to refer closely to")
	fmt.Println("the rules to ensure that the world fits within the correct classification.")
	fmt.Println("For instance, Subsector capitals (Cp) need to have a Class A starport, and")
	fmt.Println("colonies need to have population between 5 and A (incl.).")
	fmt.Println("The choices are:")
	fmt.Println("- Mr - Military rule")
	fmt.Println("- Cp - Subsector capital")
	fmt.Println("- Cs - Sector capital")
	fmt.Println("- Cx - Capital")
	fmt.Println("- Cy O:nnnn - Colony with owning world at Sector hex nnnn")
	fmt.Println("- Fo - Forbidden (Red Zone)")
	fmt.Println("- Pz - Puzzle (Amber zone & population 7+")
	fmt.Println("- Da - Dangerous (Amber zone & population 6-")
	fmt.Println("- Ab - Data repository")
	fmt.Println("- An - Ancient site")
	fmt.Println()
	tcs := getChoice("Enter the extra TCs here : ")

	if len(tcs) > 0 {
		return w.remarks + " " + tcs
	}
	return w.remarks
}

// generateT5World generates a single Traveller 5 Second Survey (T5SS) mainworld. If you wish this to be automatic, ie without questions
// being asked during generation, set isAuto to true. If the world is Zhodani, the generation process needs to be modified, and you should set
// isZhodani to true. The mainworld created is copied to the clipboard and to the worlds tab-delimited file.
// The function returns the world generated.
func generateT5WorldOld(isAuto, isZhodani bool) (w world) {

	if !isAuto {
		// Ask user for non-automated identification info for the world
		if getYesNoAnswer("Do you wish to use a common seed for repeat testing", false) {
			SeedForTesting(config.SeedValue)
			log.Println("Seeding... with", config.SeedValue)
		}

		w = w.getWorldIdentification()
		w.sectorAbbrev = getAbbreviationForSector(w.sector)

		if strings.Contains(w.allegiance, "Zh") {
			fmt.Println("You have entered a Zhodani allegiance. This places some constraints on world generation, e.g. maximum Tech Level.")
			isZhodani = !getYesNoAnswer("Do you wish to override applying Zhodani generation limits?", false)
		}

		if !getYesNoAnswer("The world will now be generated with these identifiers. Ready?", true) {
			return
		}
	} else {
		qms := "????"
		w.name = qms
		w.sectorAbbrev = qms
		w.sector = qms
		w.hexLoc = qms
		w.subsectorIndex = "?"
		if !isZhodani {
			w.allegiance = "Im"
		}
	}

	// ---- Step B ---- Basic System features
	starSpectralFlux := Flux(0)
	starSizeFlux := Flux(0)

	//The "Primary" is the main or homeworld star in a star system.
	primary := determineStar(0, starSpectralFlux, starSizeFlux, true)
	if primary.spectralType == "" {
		panic("Error: Primary Star generation has failed. Investigate.")
	}
	// Add the primary star to the world's information
	w.stars = append(w.stars, &primary)

	w.habitableZone = determineHabitableZoneVariance(primary)
	climate, _ := w.getClimate()
	w.planetOrSat = determineMainworldType()
	if !isAuto {
		fmt.Println("Mainworld type is " + w.planetOrSat)
	}
	w.mwSatGG = false
	if strings.Contains(w.planetOrSat, "Sa") {
		if Flux(0) <= 0 {
			w.mwSatGG = true
		}
		w.satOrbit = getSatOrbit(w.mwSatGG, strings.Contains(w.planetOrSat, "Close"))
	}
	w.pbg = determinePBG(w.mwSatGG)

	// ---- Step F ---- WorldGen Additional Data (Stellar)
	w.generateSystemStars(starSpectralFlux, starSizeFlux)

	// ---- Step C ---- Generate the UWP
	w.uwp = createWorld(wtMainworld, w.uwp, w.habitableZone)

	// Adjustments
	if isZhodani && w.uwp.popInt < 4 {
		w.uwp.popInt = 0
	}
	if w.uwp.popInt == 0 {
		w.pbg.populationDigit = 0
	}
	if isZhodani && w.uwp.techInt > 14 {
		w.uwp.techInt = 14
	}

	// ---- Step C ---- WorldGen Trade Classes and Zones
	w.remarks = w.getTradeClassificationsOld()
	w.determineZone()
	w.worlds = 1 + w.pbg.gasGiants + w.pbg.planetoids + D6() + D6()
	if !isAuto {
		w.remarks = w.getAdditionalTCs()
	}
	w.bases = w.determineBasesOld(isZhodani, isAuto, false)

	// ---- Step E ---- Extensions
	w.determineExtensions()

	worldOut := w.String()
	if !isAuto {
		fmt.Println()
		//fmt.Println("Homestar = " + w.stars[0].String())
		fmt.Printf("Homestar = %s\n", w.stars[0])
		//pf("Primary star habitable zone is orbit %d\n", getHabitableZoneDb(w.stars[0].String()))
		fmt.Printf("Primary star habitable zone is orbit %d\n", getHabitableZoneDb(w.stars[0].String()))
		fmt.Printf("Habitable Zone variance = %+d\n", w.habitableZone)
		fmt.Println("Mainworld Type = " + w.planetOrSat)
		// pf("    Planet core type = %s of density %f.\n", w.densityType, w.density)
		// pf("    Mass - %f.\n", w.mass)
		// pf("    Surface gravity - %fg\n", w.gravity)
		if w.planetOrSat != mwTypePlanet {
			if w.mwSatGG {
				fmt.Println("    Orbits a Gas Giant in orbit " + w.satOrbit)
			} else {
				fmt.Println("    Orbits a Big Planet in orbit " + w.satOrbit)
			}
		}
		fmt.Println("Climate = " + climate)
		fmt.Println("Star details:")
		for i := 0; i < len(w.stars); i++ {
			s := w.stars[i]
			//pf("  %d. %s.", i, s.toString())
			fmt.Printf("  %d. %s.", i, s)
			if i != 0 {
				fmt.Printf(" Orbit - %d.", s.orbit)
			}
			//pf(" HZ - %d.", getHabitableZoneDb(s.toString()))
			fmt.Printf(" HZ - %d.", getHabitableZoneDb(s.String()))
			//		pf(" Stellar mass - %f", s.mass)
			if s.companion != nil {
				//pf(" Companion - " + s.companion.toString())
				fmt.Printf(" Companion - %s", s.companion)
			}
			fmt.Printf("\n")
		}
		fmt.Println("\nWORLD DETAILS:")
		fmt.Println(worldOut)
		clipboard.WriteAll(worldOut)
		fmt.Println()
		fmt.Println("Your world has been copied to the clipboard.")
		if getYesNoAnswer("Do you wish to write to the world.tab file", true) {
			w.toFile(config.WorldOutputFile)
		}
	} else {
		fmt.Println(worldOut)
		w.toFile(config.WorldOutputFile)
	}
	return
}

// getWorldIdentification asks the user for identifying information for a star system, including world name, location and allegiance. It returns the new world.
func (w world) getWorldIdentification() world {

	fmt.Println("\nMAINWORLD IDENTIFICATION DETAILS")
	w.sectorAbbrev = getChoice("Enter the abbreviated Sector name (4 characters): ")
	sectors := getSectorsByName(w.sectorAbbrev)
	if len(sectors) == 1 {
		w.sector = sectors[0].name
		fmt.Println("Found sector " + w.sector)
	} else {
		if len(sectors) < 1 {
			fmt.Println("No matching sector found for abbreviation " + w.sectorAbbrev)
		} else {
			fmt.Printf("Multiple sectors found: ")
			for i := 0; i < len(sectors); i++ {
				fmt.Printf(sectors[i].name + " ")
			}
			fmt.Println()
		}
		w.sector = getChoice("Enter the full name of the Sector: ")
	}

	// Check the sector name for 4 characters.
	var hexLoc string
	for {
		flagContinue := false
		hexLoc = getChoice("Enter the Sector hex location (XXYY) : ")
		if !ValidateHexLoc(hexLoc) {
			fmt.Println("Please enter a valid Sector hex location (0101 to 3240)")
		} else {
			flagContinue = true
			w.hexLoc = hexLoc
		}
		if flagContinue {
			break
		}
	}
	w.subsectorIndex = ConvertHexLocToSubsector(w.hexLoc)
	fmt.Println("Subsector is " + w.subsectorIndex)
	w.name = getChoice("Enter the world name : ")

	// Get the world's allegiance, checking against the list from the database. Allow program to display all allegiances.
	for {
		flagContinue := false

		message := "Please enter a 4-character allegiance code (e.g. NaHu, ZhCo, ImDd), or ? for a list of valid codes."
		errMessage := "Please ensure you enter a valid allegiance code."

		fmt.Println(message)
		allegiance := getChoice("Enter the world's Allegiance : ")
		if allegiance == "?" {
			// Get list of all allegiances and print it out.
			fmt.Println("All allegiances:")
			fmt.Println()
			allegiances := getAllAllegianceCodes()
			fmt.Println(allegiances)
			fmt.Println()
		} else if len(allegiance) != 4 {
			fmt.Println(errMessage)
		} else {
			// Check the allegiance is valid, if so, we can continue on.
			if checkValidAllegianceCode(allegiance) {
				flagContinue = true
				w.allegiance = allegiance
			} else {
				fmt.Println(errMessage)
			}
		}
		if flagContinue {
			break
		}
	}

	return w
}

/*
// getHabitableZone returns the Orbit number for the Habitable zone of a star, based on its Type (Spectral Class) and Size (Luminosity).
func (s starDetail) getHabitableZone() int {

	// Brown Dwarfs do not have a habitable zone
	if s.spectralType == "BD" {
		return 0
	}
	// The other easy case is the Dwarf.
	if s.size == "D" {
		if s.spectralType == "O" {
			return 1
		}
		return 0
	}

	// All other spectral types have a habitable zone somewhere.
	switch s.spectralType {
	case "O":
		switch s.size {
		case "II":
			return 14
		case "III":
			return 13
		case "IV":
			return 12
		case "V":
			return 11
		default:
			return 15
		}
	case "B":
		switch s.size {
		case "II":
			return 12
		case "III":
			return 11
		case "IV":
			return 10
		case "V":
			return 9
		default:
			return 13
		}
	case "A":
		switch s.size {
		case "Ia":
			return 12
		case "Ib":
			return 11
		case "II":
			return 9
		default:
			return 7
		}
	case "F":
		switch s.size {
		case "Ia":
			return 11
		case "Ib":
			return 10
		case "II":
			return 9
		case "V":
			return 5
		case "VI":
			return 3
		default:
			return 6
		}
	case "G":
		switch s.size {
		case "Ia":
			return 12
		case "Ib":
			return 10
		case "II":
			return 9
		case "III":
			return 7
		case "IV":
			return 5
		case "V":
			return 3
		default:
			return 2
		}
	case "K":
		switch s.size {
		case "Ia":
			return 12
		case "Ib":
			return 10
		case "II":
			return 9
		case "III":
			return 8
		case "IV":
			return 5
		case "V":
			return 2
		default:
			return 1
		}
	default:
		// Only "M" left
		switch s.size {
		case "Ia":
			return 12
		case "Ib":
			return 11
		case "II":
			return 10
		case "III":
			return 9
		default:
			return 0
		}
	}
}
*/

/*
// getMass gets the Mass for the Star, expressed in Stellar Mass (ie Mass of Sol=1)
func (s starDetail) getMass() float64 {

	// This chart lists the mass of the star for indicative types and we interpolate the missing values.
	// The Key: "B0", "B5", "A0", "A5", "F0", "F5", "G0", "G5", "K0", "K5", "M0", "M5", "M9"
	// For White dwarfs: "B","A","F","G","K","M"
	// For Spectral O same as Spectral B.
	// For Brown Dwarf we just allocate a mass of 0.0352. This is the average of the range 2.5E+28 kg and 1.5E+29 kg, and ensures that it doesn't
	// change for a particular body once allocated.
	// Solar Mass is 1.989E+30 kg. Just so you know.
	massChart := map[string][]float64{
		"Ia":  []float64{60.0, 30.0, 18.0, 15.0, 13.0, 12.0, 12.0, 13.0, 14.0, 18.0, 20.0, 25.0, 30.0},
		"Ib":  []float64{50.0, 25.0, 16.0, 13.0, 12.0, 10.0, 10.0, 12.0, 13.0, 16.0, 16.0, 20.0, 25.0},
		"II":  []float64{30.0, 20.0, 14.0, 11.0, 10.0, 8.1, 8.1, 10.0, 11.0, 14.0, 14.0, 16.0, 18.0},
		"III": []float64{25.0, 15.0, 12.0, 9.0, 8.0, 5.0, 2.5, 3.2, 4.0, 5.0, 6.3, 7.4, 9.2},
		"IV":  []float64{20.00, 10.00, 6.00, 4.00, 2.50, 2.00, 1.75, 2.00, 2.30, 2.60, -1.0, -1.0, -1.0},
		"V":   []float64{18.0, 6.5, 3.2, 2.1, 1.7, 1.3, 1.04, 0.94, 0.825, 0.57, 0.489, 0.331, 0.215},
		"VI":  []float64{-1.0, -1.0, -1.0, -1.0, -1.0, 0.8, 0.6, 0.528, 0.43, 0.33, 0.154, 0.104, 0.058},
		"D":   []float64{0.260, 0.360, 0.420, 0.630, 0.830, 1.110},
	}

	// Cater for Brown Dwarfs.
	if s.spectralType == "BD" {
		return 0.0352
	}

	// If it's a Dwarf, just work out the correct one and return it: O,B = 0, A=1, F=2, G=3, K=4, M=5
	var baseIndex int
	switch s.spectralType {
	case "A":
		baseIndex = 1
	case "F":
		baseIndex = 2
	case "G":
		baseIndex = 3
	case "K":
		baseIndex = 4
	case "M":
		baseIndex = 5
	default:
		baseIndex = 0
	}
	if s.size == "D" {
		return massChart[s.size][baseIndex]
	}

	// This is how we're going to do this:
	// For any other size, we know which array we want (massChart[s.size][]), we just need to determine the index. We need a "conversion" between
	// spectral class+decimal and index, and whether there is any interpolation needed. We can get to the base index from:
	// O,B = 0, A = 2, F = 4, G = 6, K = 8, M = 10. We already have this fom the previous Dwarf calculation, except the dwarf indexes are half the
	// size of the other indexes.
	// Special case if M9, index = 12 get it and return
	// If decimal = 0, index stays same so get it and return
	// If decimal = 5, add one to index so get it and return
	// If decimal == 1 - 4, get the base index and the base index+1, interpolate and return.
	// IF decimal == 6 - 9, get the base index+1 and the base index + 2, interpolate and return.

	// Adjust the base index for the larger arrays of the non-dwarf stars
	baseIndex = baseIndex * 2

	// Special case if M9, index = 12 get it and return
	if s.spectralDecimal == 9 && s.spectralType == "M" {
		return massChart[s.size][baseIndex+2]
	}
	// If decimal = 0, index stays same so get it and return
	if s.spectralDecimal == 0 {
		return massChart[s.size][baseIndex]
	}
	// If decimal = 5, add one to index so get it and return
	if s.spectralDecimal == 5 {
		return massChart[s.size][baseIndex+1]
	}
	// Handle the other cases
	var lower, upper, result, diff, slots float64
	if s.spectralDecimal < 5 {
		lower = massChart[s.size][baseIndex]
		upper = massChart[s.size][baseIndex+1]
		slots = 5
		diff = float64(s.spectralDecimal)
	} else {
		lower = massChart[s.size][baseIndex+1]
		upper = massChart[s.size][baseIndex+2]
		diff = float64(s.spectralDecimal) - 5.0
		if s.spectralType == "M" {
			slots = 4.0
		} else {
			slots = 5.0
		}
	}
	// Interpolate
	result = lower + (diff*(upper-lower))/slots

	return result
}

*/

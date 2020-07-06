package main

// worlds.go contains code for finding, defining, and detailing worlds.

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/atotto/clipboard"
)

// world contains the full details for a world. It includes T5SS as well as World Builder's Handbook info.
type world struct {
	genType        WorldGenType  // The type of generation used for the world.
	id             int           // The ID from the database.
	name           string        // The name of the world.
	sectorAbbrev   string        // The name of the sector abbreviated.
	sector         string        // The full sector name.
	subsector      string        // The name of the subsector.
	subsectorIndex string        // The index of the subsector (A thru P).
	hexLoc         HexLoc        // The hex location in the sector.
	uwp            worldUwp      // The Universal World Profile for the world.
	bases          string        // The bases that may be present in the system.
	remarks        string        // Remarks are Trade Classifications.
	zone           TravelZone    // The world's Travel Zone, Green, Amber or Red.
	pbg            worldPBG      // The PBG indicator for the world, Population Digit, Planetoid Belts and Gas Giants.
	allegiance     string        // The Allegiance of the World.
	stars          []*starDetail // Details of stars in the system.
	importance     importanceExt // The World Importance Extension.
	economics      economicExt   // The Economic Extension.
	culture        cultureExt    // The Cultural Extension.
	nobility       string        // If Imperial, any nobility on the world.
	worlds         int           // The number of worlds in the Star System.
	ru             int           // The Resource Units for the world.
	orbit          int           // The orbit (of the primary star) the planet occupies if not a satellite, or the orbit of the central planet/gas giant if a satellite.
	worldType      string        // The world type (Mainworld, Hospitable, Wordlet, Inferno, Planetoid, RadWorld, Iceworld, Inner World, Stormworld, Bigworld).
	habZoneVar     int           // The variance in the primary star's habitable (0), inner (<0) or outer (>0) zone. At which the mainworld occupies.
	planetOrSat    string        // Whether the world orbits around a star or Gas Giant.
	mwSatGG        bool          // true if the mainworld orbits a gas Giant, false if it orbits a Big Planet. Ignored if the mainworld orbits a star.
	satOrbit       string        // The orbit if the mainworld is a satellite and orbits a central world.
}

// worldUwp Stores the Universal World Profile. Most numbers are stored as integer rather than strings.
type worldUwp struct {
	starport string // The star-/space-port type
	sizeInt  int    // Size of world integer from 0 to (usually F)
	atmInt   int    // Atmosphere integer from 0 to F
	hydInt   int    // Hydrographics percentage integer (in 10%s) from 0 to A (10)
	popInt   int    // Population integer from 0 to F. This is the exponent on 10**popInt
	govInt   int    // Government integer from 0 to F.
	lawInt   int    // Law level integer from 0 to J and possibly beyond
	techInt  int    // Tech Level integer from 0 to F and beyond.
}

// worldPBG stores the details of a PGB value for a world.
type worldPBG struct {
	populationDigit int // The Population digit
	planetoids      int // The number of Planetoid Belts in the System
	gasGiants       int // The number of Gas Giants in the System
}

// economicExt stores the economic extension for a world.
type economicExt struct {
	Resource       int // The Resource value
	Labour         int // The Labour value
	Infrastructure int // The Infrastructure value
	Efficiency     int // The Efficiency value
}

// cultureExt stores the Cultural extension for a world.
type cultureExt struct {
	Homogenity  int // The Homogenity value
	Acceptance  int // The Acceptance value
	Strangeness int // The Strangeness value
	Symbols     int // The Symbols value
}

// importanceExt stores the Importance extension for a world.
type importanceExt struct {
	Importance int // The importance value
}

// Constants for the type of world (main or otherwise)
const (
	wtMainworld  = "Mainworld"
	wtHospitable = "Hospitable"
	wtWorldlet   = "Worldlet"
	wtInferno    = "Inferno"
	wtPlanetoid  = "Planetoid"
	wtRadworld   = "RadWorld"
	wtIceworld   = "Iceworld"
	wtInnerWorld = "Inner World"
	wtStormWorld = "Stormworld"
	wtBigworld   = "Bigworld"
)

// Constants for the Mainworld basic type, either satellite or planet.
const (
	mwTypeFarSatellite   = "Satellite Far"
	mwTypeCloseSatellite = "Satellite Close"
	mwTypePlanet         = "Planet"
)

// Constants for Planet density
const (
	pdTypeHeavyCore  = "Heavy Core"
	pdTypeMoltenCore = "Molten Core"
	pdTypeRockyBody  = "Rocky Body"
	pdTypeIcyBody    = "Ice Body"
)

var basicAllegianceMap map[string]string // Contains the text strings for basic Allegiances
var mtSubsectorTrafficArr [4]string      // Contains the text strings for MegaTraveller subsector traffic
var mtSectorStarDensity [5]string        // Contains the text strings for MegaTraveller sector star density
var t5AllegianceMap map[string]string    // Contains the text strings for (basic) Traveller5 allegiances
var zoneMap map[string]string            // Maps short identifiers (R,A, or G) to their longer strings

// Constants for star sparcity (MegaTraveller)
const (
	ssBackwater = 0
	ssStandard  = 1
	ssMature    = 2
	ssCluster   = 3
)

// Constants for sector star density (MegaTraveller)
const (
	sdRift      = 0
	sdSparse    = 1
	sdScattered = 2
	sdStandard  = 3
	sdDense     = 4
)

func init() {
	basicAllegianceMap = make(map[string]string)

	basicAllegianceMap["Aslan"] = "As"
	basicAllegianceMap["Imperial"] = "Im"
	basicAllegianceMap["Vargr"] = "Va"
	basicAllegianceMap["Zhodani"] = "Zh"

	t5AllegianceMap = make(map[string]string)
	t5AllegianceMap["Imperial"] = "ImXX"
	t5AllegianceMap["Client State (Imp)"] = "CsIm"
	t5AllegianceMap["Non-Aligned"] = "NaHu"
	t5AllegianceMap["Vargr"] = "NaVa"
	t5AllegianceMap["Aslan"] = "AsXX"
	t5AllegianceMap["Zhodani"] = "ZhCo"
	t5AllegianceMap["Solomani"] = "SoCf"
	t5AllegianceMap["K'kree"] = "KkTw"
	t5AllegianceMap["Hiver"] = "HvFd"

	mtSubsectorTrafficArr[ssBackwater] = "Backwater"
	mtSubsectorTrafficArr[ssStandard] = "Standard"
	mtSubsectorTrafficArr[ssMature] = "Mature"
	mtSubsectorTrafficArr[ssCluster] = "Cluster"

	mtSectorStarDensity[sdRift] = "Rift (3%)"
	mtSectorStarDensity[sdSparse] = "Sparse (16%)"
	mtSectorStarDensity[sdScattered] = "Scattered (33%)"
	mtSectorStarDensity[sdStandard] = "Standard (50%)"
	mtSectorStarDensity[sdDense] = "Dense (66%)"
}

// ObjectBasicString returns the World as a string, showing only basic (CT03 & MT Basic) data.
func (w world) ObjectBasicString() (s string) {

	s = "[yellow]Name:[-] " + w.name + "\n"
	s += "[yellow]Sector:[-] " + w.hexLoc.String() + " " + w.sector + "\n"
	ssHex, _ := w.hexLoc.ConvertToSubsector()
	s += "[yellow]Subsec:[-] " + ssHex.String() + " " + w.subsector + " (" + w.subsectorIndex + ")\n\n"
	s += "[yellow]UWP:[-] " + w.uwp.String() + "\n"
	s += "[yellow]Bases:[-] " + w.bases + "\n"
	s += "[yellow]Zone:[-] " + w.zone.ColouredString() + "\n"
	s += "[yellow]Allegiance:[-]" + w.allegiance + "\n\n"

	// Determine Gas Giant

	// Based on world gen type, display the remaining details
	if w.genType == WgtCt03 {
		s += "[yellow]Gas Giant(s):[-] "
		if w.pbg.gasGiants == 0 {
			s += "not "
		}
		s += "present\n"
	} else if w.genType == WgtMtBasic || w.genType == WgtT5ss {
		s += fmt.Sprintf("[yellow]Population Mult:[-] %v\n", w.pbg.populationDigit)
		s += fmt.Sprintf("[yellow]Planetoid Belts:[-] %v\n", w.pbg.planetoids)
		s += fmt.Sprintf("[yellow]Gas Giants:[-] %v\n\n", w.pbg.gasGiants)
	}

	s += "[yellow]Trade Classifications:[-]\n"
	s += w.remarks + "\n\n"

	// Further for T5SS mainworlds
	if w.genType == WgtT5ss {
		if len(w.stars) != 0 {
			s += fmt.Sprintf("[yellow]Homestar:[-] %s\n", w.stars[0])
			s += fmt.Sprintf("[yellow]Mainworld Orbit:[-] %d\n", w.orbit)
			s += fmt.Sprintf("[yellow]Mainworld Type:[-] " + w.planetOrSat + "\n")
			if w.planetOrSat != mwTypePlanet {
				if w.mwSatGG {
					s += fmt.Sprintf("  Orbits Gas Giant in: " + w.satOrbit + "\n")
				} else {
					s += fmt.Sprintf("  Orbits Big Planet in: " + w.satOrbit + "\n")
				}
			}
			s += fmt.Sprintf("[yellow]Star Details:[-]\n")
			for i := 0; i < len(w.stars); i++ {
				st := w.stars[i]
				s += fmt.Sprintf("%d. %s.", i+1, st)
				if i == 0 {
					s += fmt.Sprintf(" (Prim.)")
				} else {
					s += fmt.Sprintf(" Orb: %d.", st.orbit)
				}
				s += fmt.Sprintf(" HZ: %d.", getHabitableZoneDb(st.String()))
				if st.companion != nil {
					s += fmt.Sprintf("\n   Comp.: %s", st.companion)
				}
				s += fmt.Sprintf("\n")
			}
			s += fmt.Sprintf("\n")
		} else {
			log.Printf("No stars detailed!")
		}
		// Display Extension information Ix, Ex, Cx, Nobility, Worlds, RU
		s += fmt.Sprintf("[yellow]Importance:[-] %s\n", w.importance.String())
		s += fmt.Sprintf("[yellow]Economics:[-] %s\n", w.economics.String())
		s += fmt.Sprintf("  [yellow]Resources:[-] %v\n", w.economics.Resource)
		s += fmt.Sprintf("  [yellow]Labor:[-] %v\n", w.economics.Labour)
		s += fmt.Sprintf("  [yellow]Infrastructure:[-] %v\n", w.economics.Infrastructure)
		s += fmt.Sprintf("  [yellow]Efficiency:[-] %+d\n", w.economics.Efficiency)
		s += fmt.Sprintf("[yellow]Cultural:[-] %s\n", w.culture.String())
		s += fmt.Sprintf("  [yellow]Homogenity:[-] %v\n", w.culture.Homogenity)
		s += fmt.Sprintf("  [yellow]Acceptance:[-] %v\n", w.culture.Acceptance)
		s += fmt.Sprintf("  [yellow]Strangeness:[-] %v\n", w.culture.Strangeness)
		s += fmt.Sprintf("  [yellow]Symbols:[-] %v\n", w.culture.Symbols)
		if w.nobility != "" {
			s += fmt.Sprintf("[yellow]Nobility:[-] %s\n", w.nobility)
		}
		s += fmt.Sprintf("[yellow]Worlds:[-] %v\n", w.worlds)
		s += fmt.Sprintf("[yellow]Resources:[-] %v\n", w.ru)
	} else {
		//log.Printf("Not a T5SS world!")
	}
	clipboard.WriteAll(s)

	return
}

// systemStarString prints out all the stars in a system, as would be expected in a mainworld listing. It returns the string.
func (w world) systemStarString() string {

	var s string

	for i := 0; i < len(w.stars); i++ {
		star := w.stars[i]
		s += star.String() + " "
		if star.companion != nil {
			s += star.companion.String() + " "
		}
	}
	return strings.TrimRight(s, " ")
}

// String outputs the world as a tab-delimited string, suitable for use in travellermap.com.
func (w world) String() (worldOut string) {

	worldOut = fmt.Sprintf("%s\t%s\t%s\t%s\t%s\t%s\t%s\t%s", w.sectorAbbrev, w.subsectorIndex, w.hexLoc.String(), w.name, w.uwp.String(), w.bases, w.remarks, w.zone)

	if w.genType == WgtCt03 {
		return
	}
	worldOut += fmt.Sprintf("\t%s\t%s", w.pbg.String(), w.allegiance)

	if w.genType == WgtMtBasic {
		return
	}

	worldOut += fmt.Sprintf("\t%s\t%s\t%s\t%s\t%s\t%d\t%d", w.systemStarString(), w.importance.String(),
		w.economics.String(), w.culture.String(), w.nobility, w.worlds, w.ru)
	return

}

// String returns the UWP string for a world based on its UWP structure.
func (u worldUwp) String() string {
	return u.starport + Ehex(u.sizeInt).String() + Ehex(u.atmInt).String() + Ehex(u.hydInt).String() +
		Ehex(u.popInt).String() + Ehex(u.govInt).String() + Ehex(u.lawInt).String() + "-" + Ehex(u.techInt).String()
}

// Gets Mainworld climate and Trade Classification (if any) from HabitableZone variance.
// Returns climate (text), trade classification.
func (w world) getClimate() (string, string) {

	variance := w.habZoneVar

	switch {
	case variance < 0:
		return "Hot. Tropic.", "Tr"
	case variance == 0:
		return "Temperate.", ""
	case variance == 1:
		return "Cold. Tundra", "Tu"
	default:
		return "Frozen.", "Fr"
	}
}

// String converts the integer importance extension to a string.
func (i importanceExt) String() string {
	return fmt.Sprintf("{ %+d }", i.Importance)
}

// String converts the Economics extension to a string.
func (e economicExt) String() (economicStr string) {
	economicStr = "(" + Ehex(e.Resource).String() + Ehex(e.Labour).String() + Ehex(e.Infrastructure).String()
	if e.Efficiency < 0 {
		economicStr += "-" + Ehex(-1*e.Efficiency).String()
	} else {
		economicStr += "+" + Ehex(e.Efficiency).String()
	}
	economicStr += ")"
	return
}

// calcRU calculates the Resource Units for an Economic Extension of a world. The value is returned as a positive or negative integer.
func (e economicExt) calcRU() (resourceUnits int) {
	resourceUnits = 1
	if e.Resource > 1 {
		resourceUnits = resourceUnits * e.Resource
	}
	if e.Labour > 1 {
		resourceUnits = resourceUnits * e.Labour
	}
	if e.Infrastructure > 1 {
		resourceUnits = resourceUnits * e.Infrastructure
	}
	if e.Efficiency != 0 {
		resourceUnits = resourceUnits * e.Efficiency
	}
	return
}

// String returns the Cultural Extension expressed as a string.
func (c cultureExt) String() string {
	return "[" + Ehex(c.Homogenity).String() + Ehex(c.Acceptance).String() + Ehex(c.Strangeness).String() + Ehex(c.Symbols).String() + "]"
}

// StringEsc returns the Cultural Extension expressed as a string, but escaping the tricky square brackets.
func (c cultureExt) StringEsc() string {
	return "[" + Ehex(c.Homogenity).String() + Ehex(c.Acceptance).String() + Ehex(c.Strangeness).String() + Ehex(c.Symbols).String() + "[]"
}

// getNobility gets the (Imperial) nobility for a world based on the trade classifications. Returns the nobility present as a string.
// You should ONLY call this function if the world has an allegiance of Imperium. Parameters are the remarks (Trade Classifications)
// and the world importance as an integer.
func (w *world) getNobility() string {

	if !strings.Contains(w.allegiance, "Im") {
		w.nobility = ""
		return w.nobility
	}

	w.nobility = "B"
	if strings.Contains(w.remarks, "Pa") || strings.Contains(w.remarks, "Pr") {
		w.nobility += "c"
	}
	if strings.Contains(w.remarks, "Ag") || strings.Contains(w.remarks, "Ri") {
		w.nobility += "C"
	}
	if strings.Contains(w.remarks, "Pi") {
		w.nobility += "D"
	}
	if strings.Contains(w.remarks, "Ph") {
		w.nobility += "e"
	}
	if strings.Contains(w.remarks, "In") || strings.Contains(w.remarks, "Hi") {
		w.nobility += "E"
	}
	if w.importance.Importance >= 4 {
		w.nobility += "f"
	}
	return w.nobility
}

// String returns the PBG value as a string.
func (p worldPBG) String() string {

	return Ehex(p.populationDigit).String() + Ehex(p.planetoids).String() + Ehex(p.gasGiants).String()

}

// toFile outputs the world to a given file. The file is appended to.
func (w world) toFile(file string) error {
	f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Printf("Unable to write to worlds file "+file+". Error: %v", err)
		return err
	}
	defer f.Close()
	if _, err := f.WriteString(w.String() + "\n"); err != nil {
		log.Printf("Write error : %v", err)
		return err
	}
	logPane.Log("World written to file : " + file)
	return nil
}

// parsePbg parsea a string into the components of a PBG structure. It returns the new worldPbg structure.
func parsePbg(pbg string) (wp worldPBG) {

	wp.populationDigit = int(EhexVal(string(pbg[0])))
	wp.planetoids = int(EhexVal(string(pbg[1])))
	wp.gasGiants = int(EhexVal(string(pbg[2])))

	return
}

// parseUwp parses a string into a worldUwp structure. It does not validate it, other than ensuring that valid
// eHex values are inserted. If an invalid eHex conversion is attempted, the corresponding field is set to -1.
func parseUwp(uwp string) (u worldUwp) {

	// UWP in the form "SsAHPGL-T". If S (starport) is "?" then it is likely that the rest of the UWP will
	// be "?"'s as well, and the UWP is invalid or a placeholder. The values in the string are eHex, with the
	// exception of Starport.

	if tableStarport(string(uwp[0])) != stpUnknown {
		u.starport = string(uwp[0])
	}
	u.sizeInt = int(EhexVal(string(uwp[1])))
	u.atmInt = int(EhexVal(string(uwp[2])))
	u.hydInt = int(EhexVal(string(uwp[3])))
	u.popInt = int(EhexVal(string(uwp[4])))
	u.govInt = int(EhexVal(string(uwp[5])))
	u.lawInt = int(EhexVal(string(uwp[6])))
	u.techInt = int(EhexVal(string(uwp[8])))
	return
}

// validate checks a worldUwp structure, returning true if valid or false otherwise.
func (u worldUwp) validate() bool {
	if tableStarport(u.starport) == stpUnknown {
		return false
	}
	if u.sizeInt < 0 || u.sizeInt > EhexMax {
		return false
	}
	if u.atmInt < 0 || u.atmInt > 15 {
		return false
	}
	if u.hydInt < 0 || u.hydInt > 10 {
		return false
	}
	if u.popInt < 0 || u.popInt > 15 {
		return false
	}
	if u.govInt < 0 || u.govInt > 15 {
		return false
	}
	if u.lawInt < 0 || u.lawInt > 18 {
		return false
	}
	if u.techInt < 0 || u.techInt > EhexMax {
		return false
	}
	return true
}

// parseImportanceExt takes a string representing an importance digit, and parses it into a
// importanceExt structure. The new importanceExt is returned or a blank one if it cannot be parsed.
func parseImportanceExt(s string) (ix importanceExt) {

	if len(s) < 3 {
		return
	}

	if myInt, err := strconv.Atoi(strings.Trim(s, "{ }")); err == nil {
		ix.Importance = myInt
	}
	return
}

// parseImportanceExt takes a string representing a world's Economic Extension, and parses it into an
// economicExt structure. The new economicExt is returned or a blank one if it cannot be parsed.
// There is no guarantee that all values will be valid.
func parseEconomicEx(s string) (ex economicExt) {

	if len(s) < 7 {
		return
	}

	// Remove parentheses, leaving something that looks like "ABC+D" (or "ABC-D")
	myString := strings.Trim(s, "()")
	chars := []rune(myString)
	ex.Resource = int(EhexVal(string(chars[0:1])))
	ex.Labour = int(EhexVal(string(chars[1:2])))
	ex.Infrastructure = int(EhexVal(string(chars[2:3])))

	// Note here we are possibly ignoring a bad value.
	ex.Efficiency, _ = strconv.Atoi(string(chars[3:]))

	return
}

// parseCultureEx takes a string representing the world's Cultural Extension, and parses it into a
// cultureEx structure. The new cultureEx is returned or a blank one if it cannot be parsed.
// There is no guarantee that all values will be valid.
func parseCultureEx(s string) (cx cultureExt) {

	if len(s) < 6 {
		return
	}

	// Remove brackets, leaving something that looks like ABCD
	myString := strings.Trim(s, "[]")
	chars := []rune(myString)
	cx.Homogenity = int(EhexVal(string(chars[0:1])))
	cx.Acceptance = int(EhexVal(string(chars[1:2])))
	cx.Strangeness = int(EhexVal(string(chars[2:3])))
	cx.Symbols = int(EhexVal(string(chars[3:4])))

	return
}

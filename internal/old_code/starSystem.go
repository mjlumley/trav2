package main

// starSystem.go contains code for star systems, generating and encoding/decoding in JSON.

// starSystem defines a struct for a Star System. Fields are exported to enable encoding/decoding into/from JSON.
// Examples are for Regina (where known or made up).
type starSystem struct {
	Name                 string      // The name of the Star System, which is the name of the Homeworld, eg "Regina"
	SectorAbbrev         string      // The canonical 4-letter abbreviation for the Sector, eg "Spin". Case insensitive (I think)
	Sector               string      // The full name of the Sector, eg "Spinward Marches"
	SectorHex            string      // The hex that the Star System occupies in the Sector, eg 1910
	UWP                  string      // The mainworld Universal World Profile, eg "A788899-C"
	TravelZone           string      // This will be "Red", "Amber" or (for Regina) "Green"
	Bases                string      // This will list all the bases, eg "NS" (for both Navy and Scout)
	TradeClassifications []string    // A slice of strings of the Trade Classifications for the Mainworld only, eg  {"Ri","Pa","Ph","An","Cp"}
	PopulationDigit      int         // The Population digit (multiplier), eg 7
	PlanetoidBelts       int         // The number of planetoid belts in the System, eg 0
	GasGiants            int         // The number of Gass Giants in the System, eg 3
	TotalWorlds          int         // Total number of worlds in the System, = Mainworld + gas giants + belts + other worlds (not satellites), eg 8
	Allegiance           string      // A four- or two-character string denoting mainworld allegiance, eg "Im" or "",
	Importance           int         // The Importance (integer value) for the mainworld, eg 4
	Economic             economicExt // The Economic extension for a world, eg {Resources:13,Labour:7,Infrastructure:14,Efficiency:4}
	Cultural             cultureExt  // The Culturaul extension for the mainworld, eg {Homogenity:9,Acceptance:12,Strangeness:6,Symbols:14}
	Nobility             string      // The Nobility string for the mainworld, eg "BcCeF"
	ResourceUnits        int         // The calculated resource Units for the mainworld, eg
	HabitalZoneVariance  int         // The variance from the Habitable Zone orbit, eg 0
	Climate              string      // The world climate, eg Temperate
	MainworldType        string      // Whether the world is a Planet or Close or Far Satellite of a Gas Giant or Big Planet, eg "Far Satellite"
	SatelliteOrbit       string      // Iff the world is a Satellite, the Orbit designator, eg "Arr"
	NativeStatus         string      // Indicates the type of Native Intelligent Life present, if any, eg "Natives"
}

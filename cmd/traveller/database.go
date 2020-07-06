package main

// database.go contains code for accessing the Sqlite database.

import (
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

// dbType constant defines what type of database we are using.
const dbType string = "sqlite3"

// worldDto is used for transferring a world to/from the database.
type worldDto struct {
	id             int
	sectorID       int
	sector         string
	sectorNameAbbr string
	subsectorIndex string
	subsector      string
	hexLoc         string
	name           string
	uwp            string
	bases          string
	remarks        string
	zone           string
	pbg            string
	allegiance     string
	stars          string
	importance     string
	economics      string
	culture        string
	nobility       string
	worlds         int
	ru             int
}

// stellarDto is used for grabbing star details from the database.
type stellarDto struct {
	id              int
	name            string
	luminosity      string
	spectral        string
	spectralDecimal int
	habitableZone   int
	minOrbit        int
	mass            float64
}

// initDb initialises the Database connection. Filename is in the dbFile string. It panics if not available.
// func initDb(file string) {

// 	// Check that the database file exists otherwise get out of here. We DON'T want to be creating an empty file.
// 	_, err := os.Stat(config.DatabaseFile)
// 	if os.IsNotExist(err) {
// 		panic(err)
// 	}
//
// 	// Get the database connection
// 	db, err := sql.Open(dbType, config.DatabaseFile)
// 	checkErr(err)
// 	defer db.Close()
//
// 	// If there is no database connection fail straight away.
// 	checkErr(db.Ping())
//
// 	// Do a test query.
// 	rows, err := db.Query("SELECT DISTINCT sector_id FROM world")
// 	checkErr(err)
// 	defer rows.Close()
// }

// // getAllAllegianceCodes retrieves all allegiance codes from the database and returns a slice of strings.
// func getAllAllegianceCodes() (ans []string) {

// 	// Get the database connection
// 	db, err := sql.Open(dbType, config.DatabaseFile)
// 	checkErr(err)
// 	defer db.Close()

// 	// Database query
// 	rows, err := db.Query("SELECT code FROM allegiance")
// 	checkErr(err)
// 	defer rows.Close()

// 	var code string

// 	// Append the codes to the slice.
// 	for rows.Next() {
// 		rows.Scan(&code)
// 		ans = append(ans, code)
// 	}
// 	return
// }

// getAllAllegiances retrieves all allegiance codes from the database and returns a slice of allegiance objects.
// func getAllAllegiances() (ans []allegiance) {

// 	//ans := make([]allegiance)

// 	// Get the database connection
// 	db, err := sql.Open(dbType, config.DatabaseFile)
// 	checkErr(err)
// 	defer db.Close()

// 	// Database query
// 	rows, err := db.Query("SELECT code, legacy_code, allegiance_name FROM allegiance ORDER BY code")
// 	checkErr(err)
// 	defer rows.Close()

// 	var code, legacyCode, aName string

// 	// Append the codes to the slice.
// 	for rows.Next() {
// 		rows.Scan(&code, &legacyCode, &aName)
// 		alleg := allegiance{code: code, legacyCode: legacyCode, name: aName}
// 		ans = append(ans, alleg)
// 	}
// 	return
// }

// checkValidAllegianceCode checks that the given allegiance code is valid and returns true if valid, false if not.
// func checkValidAllegianceCode(code string) bool {

// 	ans := getAllAllegianceCodes()

// 	for _, item := range ans {
// 		if item == code {
// 			return true
// 		}
// 	}
// 	return false
// }

// getAllSectors gets all the sectors from the database and returns a slice of sectors.
// func getAllSectors() (ss []sectorDTO) {

// 	// Get the database connection
// 	db, err := sql.Open(dbType, config.DatabaseFile)
// 	checkErr(err)
// 	defer db.Close()

// 	rows, err := db.Query("SELECT id, name, abbreviation FROM sector ORDER BY name")
// 	checkErr(err)
// 	defer rows.Close()

// 	var id int
// 	var name, abbrev string

// 	for rows.Next() {
// 		rows.Scan(&id, &name, &abbrev)
// 		ss = append(ss, sectorDTO{id: id, name: name, abbrev: abbrev})
// 	}
// 	return
// }

// getSectorsByName gets all sectors that match either the abbreviated name or fullname. It returns a slice of sectors.
// func getSectorsByName(search string) (ss []sectorDTO) {
// 	var queryString, lcSearch string

// 	// Get the database connection
// 	db, err := sql.Open(dbType, config.DatabaseFile)
// 	checkErr(err)
// 	defer db.Close()

// 	lcSearch = strings.ToLower(search)

// 	// SELECT id, name, abbreviation from sector where lower(abbreviation) = "spin" OR lower(name) LIKE "spin%"
// 	if len(search) < 5 {
// 		queryString = fmt.Sprintf("SELECT id, name, abbreviation, x_loc, y_loc FROM sector WHERE lower(abbreviation) = '%s' OR lower(name) LIKE '%s%%'", lcSearch, lcSearch)
// 	} else {
// 		queryString = fmt.Sprintf("SELECT id, name, abbreviation, x_loc, y_loc FROM sector WHERE lower(name) LIKE '%s%%'", lcSearch)
// 	}

// 	rows, err := db.Query(queryString)
// 	checkErr(err)
// 	defer rows.Close()

// 	var id, xLoc, yLoc int
// 	var name, abbrev string

// 	for rows.Next() {
// 		rows.Scan(&id, &name, &abbrev, &xLoc, &yLoc)
// 		ss = append(ss, sectorDTO{id: id, name: name, abbrev: abbrev, xLoc: xLoc, yLoc: yLoc})
// 	}

// 	return
// }

// getForeven gets world details for the Foreven sector from the database. It returns a slice of worlds.
// func getForeven() (s *sector) {

// 	s = &sector{name: "Foreven", abbrev: "Fore", saved: true, otu: true}

// 	// Get the database connection
// 	db, err := sql.Open(dbType, config.DatabaseFile)
// 	checkErr(err)
// 	defer db.Close()

// 	// ------ Collect Foreven sector details first
// 	queryString := fmt.Sprintf("SELECT id FROM sector WHERE sector.name = '%s'", s.name)
// 	//var sectorID int

// 	rows, err := db.Query(queryString)
// 	if !checkErr(err) {
// 		// Cannot continue if we can't find the sector in the first place!
// 		return nil
// 	}
// 	for rows.Next() {
// 		rows.Scan(&s.id)
// 	}
// 	rows.Close()

// 	// ------ Collect subsector information next
// 	queryString = fmt.Sprintf("SELECT subsector.id, subsector.name, language.name, subsector.subsector_index, capital_id FROM subsector, language WHERE "+
// 		"language.id = subsector.lang_id AND subsector.sector_id =%d", s.id)

// 	rows, err = db.Query(queryString)
// 	// We can possible ignore an error here
// 	checkErr(err)

// 	var ssID, capID int
// 	var ssName, langName, subsectorIndexString string
// 	for rows.Next() {
// 		rows.Scan(&ssID, &ssName, &langName, &subsectorIndexString, &capID)
// 		for i := 0; i < 16; i++ {
// 			if ssIndex[i] == subsectorIndexString {
// 				subsec := subsector{id: ssID, name: ssName, language: langName, capitalID: capID}
// 				s.subsectors[i] = subsec
// 			}
// 		}
// 	}
// 	rows.Close()

// 	// ------ Collect world details

// 	// Foreven details: Sector name = Foreven, Sector_Abbrev = Fore
// 	queryString = fmt.Sprintf("SELECT id, hex, name, UWP, bases, remarks, zone, PBG, allegiance, stars, importance, economics, culture, nobility, worlds, RU FROM world WHERE sector_id = %d", s.id)

// 	rows, err = db.Query(queryString)
// 	checkErr(err)

// 	var w worldDto

// 	for rows.Next() {
// 		rows.Scan(&w.id, &w.hexLoc, &w.name, &w.uwp, &w.bases, &w.remarks, &w.zone, &w.pbg, &w.allegiance, &w.stars, &w.importance, &w.economics, &w.culture, &w.nobility, &w.worlds, &w.ru)
// 		newWorld := w.convertToWorld()
// 		newWorld.sector = s.name
// 		newWorld.sectorAbbrev = s.abbrev
// 		newWorld.subsectorIndex = newWorld.hexLoc.GetIndex()
// 		if i := newWorld.hexLoc.IntIndex(); i != -1 {
// 			newWorld.subsector = s.subsectors[i].name
// 		}
// 		s.worlds = append(s.worlds, newWorld)
// 	}
// 	rows.Close()
// 	return
// }

// // getWorldsByName gets all worlds that match the world name. It returns a slice of worlds.
// func getWorldsByName(search string) (ws []world) {

// 	// Get the database connection
// 	db, err := sql.Open(dbType, config.DatabaseFile)
// 	checkErr(err)
// 	defer db.Close()

// 	// Prepare the search string.
// 	queryString := "select world.id, world.name as 'world_name', sector.name as 'sector_name', sector.abbreviation as 'sector_abbrev'," +
// 		" subsector.name as 'subsector_name', world.subsector_index as 'subsector_index', world.hex as 'hex_loc', world.UWP," +
// 		" world.bases, world.remarks, world.zone, world.pbg, world.allegiance, world.stars, world.importance," +
// 		" world.economics, world.culture, world.nobility, world.worlds, world.ru from world, sector," +
// 		" subsector where world.subsector_index = subsector.subsector_index AND subsector.sector_id = world.sector_id" +
// 		" and world.sector_id = sector.id and lower(world.name) like '%" + strings.ToLower(dbSanitise(search)) + "%'"

// 	rows, err := db.Query(queryString)
// 	checkErr(err)
// 	defer rows.Close()

// 	for rows.Next() {
// 		var w worldDto
// 		pw := &w
// 		rows.Scan(&pw.id, &pw.name, &pw.sector, &pw.sectorNameAbbr, &pw.subsector, &pw.subsectorIndex, &pw.hexLoc, &pw.uwp, &pw.bases, &pw.remarks, &pw.zone, &pw.pbg, &pw.allegiance, &pw.stars,
// 			&pw.importance, &pw.economics, &pw.culture, &pw.nobility, &pw.worlds, &pw.ru)

// 		ws = append(ws, w.convertToWorld())
// 	}

// 	return
// }

// // countSubsectorsToRetrieve counts the number of subsectors to be retrieved from travellermap.com.
// func countSubsectorsToRetrieve() (num int) {

// 	// Get the database connection
// 	db, err := sql.Open(dbType, config.DatabaseFile)
// 	checkErr(err)
// 	defer db.Close()

// 	rows, err := db.Query("SELECT count(*) FROM sector where is_detailed = 1 AND id NOT IN (select distinct sector_id from subsector)")
// 	checkErr(err)
// 	defer rows.Close()
// 	for rows.Next() {
// 		rows.Scan(&num)
// 	}
// 	return
// }

// // getDistinctSectors gets distinct sectors that have world information populated in the database. It returns a slice of sectors.
// func getDistinctSectors() (ss []sectorDTO) {

// 	// Get the database connection
// 	db, err := sql.Open(dbType, config.DatabaseFile)
// 	checkErr(err)
// 	defer db.Close()

// 	rows, err := db.Query("SELECT id, name, abbreviation, x_loc, y_loc FROM sector WHERE sector.id in (SELECT DISTINCT sector_id FROM world) and is_detailed=1")
// 	checkErr(err)
// 	defer rows.Close()

// 	var id, xLoc, yLoc int
// 	var name, abbrev string

// 	for rows.Next() {
// 		rows.Scan(&id, &name, &abbrev, &xLoc, &yLoc)
// 		ss = append(ss, sectorDTO{id: id, name: name, abbrev: abbrev, xLoc: xLoc, yLoc: yLoc})
// 	}
// 	return
// }

// // getHabitableZoneDb looks in the Stellar_detail table for habitable zone information about a star. It returns the habitable zone, which may be -1 for invalid.
// func getHabitableZoneDb(star string) int {
// 	starDto := getStellarDetail(star)
// 	return starDto.habitableZone
// }

// // getMassDb looks in the stellar_detail table for the mass of a specific star. It returns the mass if found (in standard Solar masses).
// func getMassDb(star string) float64 {
// 	starDto := getStellarDetail(star)
// 	return starDto.mass
// }

// // getStellarDetail gets info or a particular star from the stellar_detail table of the database. It returns the detail in a stellarDto.
// func getStellarDetail(star string) (s stellarDto) {

// 	// Get the database connection
// 	db, err := sql.Open(dbType, config.DatabaseFile)
// 	checkErr(err)
// 	defer db.Close()

// 	// Sanitise the string coming in
// 	star = dbSanitise(star)

// 	// Check length of string presented
// 	if len(star) < 2 {
// 		panic("Invalid parameter to getStellarDetail")
// 	}

// 	queryString := fmt.Sprintf("SELECT stellar_detail.id, stellar_detail.name, stellar_luminosity.name AS luminosity, stellar_spectral.name as spectral, "+
// 		"spectral_decimal, habitable_zone, min_zone, mass "+
// 		"FROM stellar_detail, stellar_luminosity, stellar_spectral "+
// 		"WHERE luminosity_id=stellar_luminosity.id AND spectral_id=stellar_spectral.id AND stellar_detail.name = '%s'", strings.ToUpper(star))

// 	rows, err := db.Query(queryString)
// 	checkErr(err)
// 	defer rows.Close()

// 	var id, decimal, habZone, minZone int
// 	var mass float64
// 	var name, luminosity, spectral string

// 	for rows.Next() {
// 		rows.Scan(&id, &name, &luminosity, &spectral, &decimal, &habZone, &minZone, &mass)
// 		s.id = id
// 		s.name = name
// 		s.luminosity = luminosity
// 		s.spectral = spectral
// 		s.spectralDecimal = decimal
// 		s.habitableZone = habZone
// 		s.minOrbit = minZone
// 		s.mass = mass
// 	}

// 	return
// }

/*

// copyWorldStagingToProd copies worlds from the world_staging table to the world table.
func copyWorldStagingToProd() {

	logPane.Log("Copying data from world_staging table to world table")

	// Get the database connection
	db, err := sql.Open(dbType, config.DatabaseFile)
	checkErr(err)
	defer db.Close()

	stmt, err := db.Prepare("INSERT INTO world (sector_id,subsector_index,hex,name,UWP,bases,remarks,zone,PBG,allegiance," +
		"stars,importance,economics,culture,nobility,worlds,RU) SELECT sector_id,subsector_index,hex,name,UWP,bases,remarks," +
		"zone,PBG,allegiance,stars,importance,economics,culture,nobility,worlds,RU FROM world_staging")
	checkErr(err)
	res, err := stmt.Exec()
	checkErr(err)

	// Get the number of rows copies
	affect, err := res.RowsAffected()
	stmt.Close()

	// Now mark those sectors with worlds present as is_detailed = true
	stmt, _ = db.Prepare("UPDATE sector SET is_detailed = 1 WHERE id IN (SELECT DISTINCT sector_id from world_staging)")
	stmt.Exec()
	stmt.Close()

	logPane.Log(fmt.Sprintf("%d rows deleted", affect))
	refocus(idMenuPaneStr)
}

*/

/*
// clearWorldStaging deletes all world data from the world_staging table.
func clearWorldStaging() {

	logPane.Log("Clearing all rows from world_staging table")

	// Get the database connection
	db, err := sql.Open(dbType, config.DatabaseFile)
	checkErr(err)
	defer db.Close()

	stmt, err := db.Prepare("DELETE FROM world_staging")
	checkErr(err)
	res, err := stmt.Exec()
	checkErr(err)
	// Get the number of rows copied
	affect, err := res.RowsAffected()
	stmt.Close()

	logPane.Log(fmt.Sprintf("%d rows deleted", affect))
	refocus(idMenuPaneStr)
}

*/

// checkErr checks an error on the database, optional string arguments can be provided. Returns true if OK, false if an error.
func checkErr(err error, args ...string) bool {
	if err != nil {
		//errString := fmt.Sprintf("Error\n%q: %s\n", err, args)
		//logPane.Log(errString)
		return false
	}
	return true
}

// dbSanitise dDoes a basic database sanitise on a text string prior to presenting to the database in a query. It returns the "clean" string.
func dbSanitise(text string) (ret string) {

	ret = strings.ReplaceAll(text, "%", "")
	ret = strings.ReplaceAll(ret, ";", "")
	ret = strings.ReplaceAll(ret, "?", "")
	ret = strings.ReplaceAll(ret, "*", "")
	ret = strings.ReplaceAll(ret, "_", "")
	ret = strings.ReplaceAll(ret, "&", "")
	ret = strings.ReplaceAll(ret, "$", "")
	ret = strings.ReplaceAll(ret, "#", "")
	ret = strings.ReplaceAll(ret, "@", "")
	ret = strings.ReplaceAll(ret, "!", "")
	ret = strings.ReplaceAll(ret, "'", "''")

	return
}

// convertToWorld converts a worldDto data transfer object to a World struct. It expects a fully completed object.
// It returns the new world struct.
// func (d worldDto) convertToWorld() (w world) {

// 	// First convert the easy
// 	w.id = d.id
// 	w.name = d.name
// 	w.sectorAbbrev = d.sectorNameAbbr
// 	w.sector = d.sector
// 	w.subsector = d.subsector
// 	w.subsectorIndex = d.subsectorIndex
// 	w.hexLoc = *NewHexLoc(d.hexLoc, true)
// 	w.uwp = parseUwp(d.uwp)
// 	w.bases = d.bases
// 	w.remarks = d.remarks
// 	// Ignore TravelZone errors. They'll just be converted to TzUnknown anyway
// 	var err error
// 	if w.zone, err = ZoneFromString(d.zone); err != nil {
// 		log.Printf(err.Error())
// 	}
// 	w.pbg = parsePbg(d.pbg)
// 	w.allegiance = d.allegiance
// 	w.stars = parseStars(d.stars)
// 	w.importance = parseImportanceExt(d.importance)
// 	w.economics = parseEconomicEx(d.economics)
// 	w.culture = parseCultureEx(d.culture)
// 	w.nobility = d.nobility
// 	w.worlds = d.worlds
// 	w.ru = d.ru

// 	return
// }

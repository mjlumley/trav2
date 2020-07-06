package main

// sqliteDb.go contains specific code for accessing the sqlite database.

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3" // Blank import used for importing sqlite3
)

// dbType constant defines what type of database we are using.
//const sdbType string = "sqlite3"

var dbFile string

// SetDbFile sets the name of the sqlite3 database file.
func SetDbFile(dbf string) {
	dbFile = dbf
}

// GetAllValidSectors gets all the sectors from the database and returns a slice of strings with their names.
func GetAllValidSectors() (ss []string, e error) {

	// Get the database connection
	db, e := sql.Open(dbType, dbFile)
	if e != nil {
		return nil, e
	}
	defer db.Close()

	rows, e := db.Query("SELECT DISTINCT name FROM sector where is_detailed=1 ORDER BY name")
	if e != nil {
		return nil, e
	}
	defer rows.Close()

	var name string

	for rows.Next() {
		rows.Scan(&name)
		ss = append(ss, name)
	}
	return ss, nil
}

// GetAllDetailedSectorAbbrev gets all the sectors from the database and returns a map with fullnames
// mapping to abbreviations. For example sm["Foreven"] = "Fore"
func GetAllDetailedSectorAbbrev() (sm map[string]string, e error) {

	sm = make(map[string]string)

	// Get the database connection
	db, e := sql.Open(dbType, dbFile)
	if e != nil {
		return nil, e
	}
	defer db.Close()

	rows, e := db.Query("SELECT DISTINCT name, abbreviation FROM sector where is_detailed=1 ORDER BY name")
	if e != nil {
		return nil, e
	}
	defer rows.Close()

	var name, abbreviation string

	for rows.Next() {
		rows.Scan(&name, &abbreviation)
		sm[name] = abbreviation
	}

	return sm, nil

}

// getSubsectorBySectorNameAndIndex gets the subsector that matches the sector and subsectorIndex.
// It returns a subsector object or a blank one with error set.
func getSubsectorBySectorNameAndIndex(sector, idx string) (ss subsector, e error) {

	// Get the database connection
	db, e := sql.Open(dbType, config.DatabaseFile)
	if e != nil {
		return
	}
	defer db.Close()

	queryString := "SELECT subsector.id, subsector.name, subsector.remarks, language.name, subsector.capital_id" +
		" FROM subsector,sector,language" +
		" WHERE subsector.sector_id = sector.id AND" +
		" subsector.lang_id = language.id AND" +
		" sector.name = '" + sector + "' AND subsector.subsector_index = '" + idx + "'"
	rows, e := db.Query(queryString)
	if e != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		sub := &ss
		rows.Scan(&sub.id, &sub.name, &sub.remarks, &sub.language, &sub.capitalID)
	}
	return
}

// GetAllMajorRaces gets all the major races from the database and returns a slice of strings with their names.
func GetAllMajorRaces() (rs []string, e error) {

	// Get the database connection
	db, e := sql.Open(dbType, dbFile)
	if e != nil {
		return nil, e
	}
	defer db.Close()

	rows, e := db.Query("SELECT DISTINCT race_name FROM race where is_major=1 ORDER BY race_name")
	if e != nil {
		return nil, e
	}
	defer rows.Close()

	var name string
	for rows.Next() {
		rows.Scan(&name)
		rs = append(rs, name)
	}
	return rs, nil
}

// GetAllCTSkills gets all the non-cascade skills from the database and returns a slice of strings.
func GetAllCTSkills() (ss []string, e error) {

	// Get the database connection
	db, e := sql.Open(dbType, dbFile)
	if e != nil {
		return nil, e
	}
	defer db.Close()

	rows, e := db.Query("SELECT skill_name FROM skill WHERE is_virtual=0 AND ruleset = (SELECT DISTINCT ID FROM ruleset WHERE abbreviation = 'CT') ORDER BY skill_name")
	if e != nil {
		return nil, e
	}
	defer rows.Close()

	var name string
	for rows.Next() {
		rows.Scan(&name)
		ss = append(ss, name)
	}
	return ss, nil
}

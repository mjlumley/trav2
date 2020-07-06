package main

// travellermap.com contains code for accessing travelmap.com and retrieving world/sector data.

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// tmf constants are strings that identify Traveller Map Fields in world output.
const (
	tmfName       = "Name"
	tmfSector     = "Sector"
	tmfSubSector  = "SS"
	tmfHexLoc     = "Hex"
	tmfUwp        = "UWP"
	tmfBases      = "Bases"
	tmfRemarks    = "Remarks"
	tmfZone       = "Zone"
	tmfPbg        = "PBG"
	tmfAllegiance = "Allegiance"
	tmfStars      = "Stars"
	tmfIx         = "{Ix}"
	tmfEx         = "(Ex)"
	tmfCx         = "[Cx]"
	tmfNobility   = "Nobility"
	tmfWorlds     = "W"
	tmfRu         = "RU"
)

// retrieveWorldData retrieves mainworld data from travellermap.com for the sectors that don't already
// have that data in the database.
func retrieveWorldData() {
	// List of sectors to retrieve
	var ss []sectorDTO

	// Get the database connection
	db, err := sql.Open(dbType, config.DatabaseFile)
	checkErr(err)
	defer db.Close()

	rows, err := db.Query("SELECT id,name,abbreviation,x_loc,y_loc from sector WHERE id NOT IN (select DISTINCT sector_id from world) AND is_detailed=1")
	checkErr(err)
	defer rows.Close()

	var id, xLoc, yLoc int
	var name, abbrev string
	for rows.Next() {
		rows.Scan(&id, &name, &abbrev, &xLoc, &yLoc)
		ss = append(ss, sectorDTO{id: id, name: name, abbrev: abbrev, xLoc: xLoc, yLoc: yLoc})
	}

	logPane.Log(fmt.Sprintf("Retrieving world data for %d sectors", len(ss)))

	if !getYesNoAnswer("Are you sure", true) {
		return
	}
	// Cycle through each sector attempting to get data from the website.
	for _, sector := range ss {

		// API call is in the form https://travellermap.com/api/sec?sx=xLoc&sy=yLoc&type=TabDelimted where xLoc,yLoc indicate (unique) sector location.
		myURL := "https://travellermap.com/api/sec?sx=" + strconv.Itoa(sector.xLoc) + "&sy=" + strconv.Itoa(sector.yLoc) + "&type=TabDelimited"
		logPane.Log(fmt.Sprintf("Retrieving id=%d - %s - %s %d,%d from %s", sector.id, sector.abbrev, sector.name, sector.xLoc, sector.yLoc, myURL))
		resp, err := http.Get(myURL)
		if err != nil {
			logPane.Log(err.Error())
			continue
		}
		defer resp.Body.Close()

		// Read the body of the data into a byte slice
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			logPane.Log(err.Error())
			continue
		}

		// Check the status code of the response - not sure if this actually works!
		if resp.StatusCode != 200 {
			logPane.Log(fmt.Sprintf("Failed for sectorId =%d- fail code:%d", sector.id, resp.StatusCode))
			continue
		}

		// Here we are assuming that the response is OK and we are intending to parse it. Change from bytes to string.
		strResp := string(body)

		// Split it on CRLF = \r\n (configurable)
		allLines := strings.Split(strResp, config.WebLineString)

		// This map indicates the order of fields in the retrieved Tab Delimited data
		m := make(map[string]int)

		// First line is the header line, split it on tabs (confiurable)
		markerLine := strings.Split(allLines[0], config.WebDelimiter)
		if len(markerLine) < config.WorldFields {
			// A panic is no good here we've got to go to the next sector.
			logPane.Log("Too few lines. Cannot continue with sector " + sector.abbrev)
			continue
		}
		// Get out the field names
		for n := 0; n < len(markerLine); n++ {
			m[markerLine[n]] = n
		}

		var w worldDto
		// Go through the remainder of the lines and get them into the database
		for i := 1; i < len(allLines); i++ {
			line := strings.Split(allLines[i], "\t")
			if len(line) > 5 {
				w.sectorID = sector.id

				// Probably have a good line here, lets do it.
				w.name = line[m[tmfName]]
				w.sectorNameAbbr = line[m[tmfSector]]    // A 4-character abbreviation
				w.subsectorIndex = line[m[tmfSubSector]] // It is actually the value A thru P
				w.hexLoc = line[m[tmfHexLoc]]            // The hexlocation
				w.uwp = line[m[tmfUwp]]
				w.bases = line[m[tmfBases]]
				w.remarks = line[m[tmfRemarks]]
				w.zone = line[m[tmfZone]]
				w.pbg = line[m[tmfPbg]]
				w.allegiance = line[m[tmfAllegiance]]
				w.stars = line[m[tmfStars]]

				// Following is "Second Survey" data and may or may not be present. We have to find out, as the DB does
				// not like having nulls input in integer fields when you specify the field for INSERT.
				w.importance = line[m[tmfIx]]
				w.economics = line[m[tmfEx]]
				w.culture = line[m[tmfCx]]
				w.nobility = line[m[tmfNobility]]

				// Handle the two integer fields
				if len(line[m[tmfWorlds]]) != 0 {
					tempInt, err := strconv.Atoi(line[m[tmfWorlds]])
					if err == nil {
						w.worlds = tempInt
					}
				}
				if len(line[m[tmfRu]]) != 0 {
					tempInt, err := strconv.Atoi(line[m[tmfRu]])
					if err == nil {
						w.ru = tempInt
					}
				}

				// Now determine whether the info is classic or second survey data
				if len(w.importance) < 2 {
					// Classic data, short INSERT statement
					stmt, err := db.Prepare("INSERT INTO world_staging " +
						"(sector_id, sector_code,subsector_index,hex,name,UWP,bases,remarks,zone,PBG,allegiance,stars,worlds,RU) " +
						"VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?)")
					stmt.Exec(w.sectorID, w.sectorNameAbbr, w.subsectorIndex, w.hexLoc, w.name, w.uwp, w.bases, w.remarks, w.zone, w.pbg, w.allegiance, w.stars, -1, w.ru)
					stmt.Close()
					if err != nil {
						logPane.Log(err.Error())
						continue
					}
				} else {
					// Second Survey data, full INSERT statement
					stmt, _ := db.Prepare("INSERT INTO world_staging " +
						"(sector_id,sector_code,subsector_index,hex,name,UWP,bases,remarks,zone,PBG,allegiance,stars,importance,economics,culture,nobility,worlds,RU) " +
						"VALUES (?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)")
					stmt.Exec(w.sectorID, w.sectorNameAbbr, w.subsectorIndex, w.hexLoc, w.name, w.uwp, w.bases, w.remarks, w.zone, w.pbg, w.allegiance, w.stars, w.importance, w.economics, w.culture, w.nobility, w.worlds, w.ru)
					stmt.Close()
					if err != nil {
						logPane.Log(err.Error())
						continue
					}
				}
			}
		}
		// Have a short break so we are not hammering the website
		time.Sleep(time.Duration(config.WebRetrieveWait) * time.Second)
	}

}

// retrieveSubsectorData retrieve subsector data from travellermap.com for the sectors that don't already
// have that data in the database.
func retrieveSubsectorData() {
	// List of sectors to retrieve
	var ss []sectorDTO

	// Get the database connection
	db, err := sql.Open(dbType, config.DatabaseFile)
	checkErr(err)
	defer db.Close()

	var id, xLoc, yLoc int
	var name, abbrev string

	rows, err := db.Query("SELECT id, name, abbreviation, x_loc, y_loc FROM sector where is_detailed = 1 AND id NOT IN (select distinct sector_id from subsector)")
	checkErr(err)
	defer rows.Close()
	for rows.Next() {
		rows.Scan(&id, &name, &abbrev, &xLoc, &yLoc)
		ss = append(ss, sectorDTO{id: id, name: name, abbrev: abbrev, xLoc: xLoc, yLoc: yLoc})
	}

	// For each sector, retrieve the data from the website and store it
	for _, sector := range ss {

		// API for retrieving https://travellermap.com/api/metadata?sx=sx&sy=sy
		myURL := "https://travellermap.com/api/metadata?sx=" + strconv.Itoa(sector.xLoc) + "&sy=" + strconv.Itoa(sector.yLoc)
		logPane.Log(fmt.Sprintf("Retrieving %d - %s - %s %d,%d from %s", sector.id, sector.abbrev, sector.name, sector.xLoc, sector.yLoc, myURL))
		resp, err := http.Get(myURL)
		if !checkErr(err) {
			continue
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if !checkErr(err) {
			continue
		}

		// Check the status code of the
		if resp.StatusCode != 200 {
			logPane.Log(fmt.Sprintf("Failed for id=%d code: %d", sector.id, resp.StatusCode))
			continue
		}

		// Base map[string] for the JSON data
		var dat map[string]interface{}

		if err := json.Unmarshal(body, &dat); err != nil {
			logPane.Log(fmt.Sprintf("Failed to unmarshall JSON for id=%d error: %s", sector.id, err.Error()))
		}
		// Extract the subsectors, an array
		sss := dat["Subsectors"].([]interface{})
		for i := 0; i < len(sss); i++ {

			// Beneath the Subsectors we have a map[string], for the individual subsector
			ssMap := sss[i].(map[string]interface{})

			var subSector subsectorDTO
			subSector.remarks = ""
			subSector.langID = 1
			subSector.capitalID = -1
			subSector.sectorID = sector.id

			for key, value := range ssMap {
				// Each value is an interface{} type, that is type asserted as a string

				// Determine the other details of this subsector
				switch key {
				case "Index":
					subSector.subsectorIndex = value.(string)
				case "Name":
					subSector.name = value.(string)
				default:
					// Ignore other keys, IndexNumber, Author
				}

			}
			// Check that we have all details to enter into database
			if subSector.subsectorIndex == "" || subSector.name == "" {
				// Cannot contine here
				logPane.Log("Incomplete data for subsector")
				continue
			}
			// Have details here to enter data in the database
			stmt, err := db.Prepare("INSERT INTO subsector (name, lang_id, sector_id, subsector_index, capital_id, remarks) VALUES (?,?,?,?,?,?)")
			stmt.Exec(subSector.name, subSector.langID, subSector.sectorID, subSector.subsectorIndex, subSector.capitalID, subSector.remarks)
			stmt.Close()
			if !checkErr(err) {
				continue
			}
		}
		// Have a short break so we are not hammering the website
		time.Sleep(time.Duration(config.WebRetrieveWait) * time.Second)
	}
}

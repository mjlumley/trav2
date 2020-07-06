package main

import (
	"fmt"
	"log"
	"sort"
	"strconv"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

// worldDialogs contains dialog boxes and other screen items for the world generation

func getWorldBasicData() {
	logPane.Log("Collecting basic world data")
	getWorldBasicDialog(WgtCt03)
}

func getWorldBasicMTData() {
	logPane.Log("Collecting basic MegaTraveller world data")
	getWorldBasicDialog(WgtMtBasic)
}

func getWorldBasicWbhData() {
	logPane.Log("Collecting basic World Builder's Handbook world data")
	getWorldBasicDialog(WgtMtWBH)
}

func getWorldBasicT5Data() {
	logPane.Log("Collecting basic Traveller5 world data")
	getWorldBasicDialog(WgtT5ss)
}

// getWorldBasicDialog collects the basic data for a world in a dialog box and then moves onto the generation function.
// Set the genType value with the type of generation needed from the const values in worldGen.go.
func getWorldBasicDialog(genType WorldGenType) {

	thisPage := "worldbasic"
	mainPageFocus()

	// Check genType is valid
	if genType != WgtCt03 && genType != WgtMtBasic && genType != WgtMtWBH && genType != WgtT5ss {
		logPane.Log(fmt.Sprintf("Invalid generation type: %v", int(genType)))
		return
	}

	sectors, err := GetAllValidSectors()
	if err != nil {
		logPane.Log("DB Error retrieving sectors: " + err.Error())
	}

	wbdPage := tview.NewForm()
	x, y, width, height := mainPage.GetRect()
	dialogWidth := 41
	dialogHeight := 11
	if genType == WgtMtBasic {
		dialogHeight = 15
	} else if genType == WgtT5ss {
		dialogHeight = 13
	}
	x = x + width/2 - dialogWidth/2
	y = y + height/2 - dialogHeight/2
	wbdPage.SetRect(x, y, dialogWidth, dialogHeight)
	wbdPage.SetBorder(true).SetTitle(" World Basic Data ").SetTitleAlign(tview.AlignCenter)
	nameInput := tview.NewInputField().SetLabel("World name").SetFieldWidth(24)
	sectorInput := tview.NewInputField().SetLabel("Sector").SetFieldWidth(24)
	sectorInput.SetAutocompleteFunc(func(currentText string) (entries []string) {
		// Ignore empty text
		if len(currentText) == 0 {
			return
		}
		for _, sector := range sectors {
			if strings.HasPrefix(strings.ToLower(sector), strings.ToLower(currentText)) {
				entries = append(entries, sector)
			}
		}
		if len(entries) <= 1 {
			entries = nil
		}
		return
	})
	hexInput := tview.NewInputField().SetLabel("Hex location").SetFieldWidth(4)
	hexInput.SetAcceptanceFunc(tview.InputFieldInteger)
	hexInput.SetDoneFunc(func(key tcell.Key) {
		hl := NewHexLoc(hexInput.GetText(), true)
		if hl == nil || !hl.IsValid() {
			if key != tcell.KeyEscape {
				logPane.Log("Invalid hex location!")
				hexInput.SetText("XXXX")
			}
		}
	})

	// This needs to be different for MT and T5
	var allegiances []string
	if genType == WgtMtBasic || genType == WgtMtWBH {
		allegiances = make([]string, 0, len(basicAllegianceMap))
		for k := range basicAllegianceMap {
			allegiances = append(allegiances, k)
		}
		sort.Strings(allegiances)
	} else if genType == WgtT5ss {
		allegiances = make([]string, 0, len(t5AllegianceMap))
		for k := range t5AllegianceMap {
			allegiances = append(allegiances, k)
		}
		sort.Strings(allegiances)
	}
	ssTraffic := make([]string, 0, len(mtSubsectorTrafficArr))
	for ss := range mtSubsectorTrafficArr {
		ssTraffic = append(ssTraffic, mtSubsectorTrafficArr[ss])
	}

	// These options only needed for MegaTraveller/WBH and T5
	allegianceInput := tview.NewDropDown().SetLabel("Allegiance").SetFieldWidth(24)
	allegianceInput.SetOptions(allegiances, nil)
	trafficInput := tview.NewDropDown().SetLabel("Sector age").SetFieldWidth(24)
	trafficInput.SetOptions(ssTraffic, nil)

	// Construct the form
	wbdPage.AddFormItem(nameInput)
	wbdPage.AddFormItem(sectorInput)
	wbdPage.AddFormItem(hexInput)
	if genType == WgtMtBasic || genType == WgtMtWBH {
		wbdPage.AddFormItem(allegianceInput)
		wbdPage.AddFormItem(trafficInput)
	} else if genType == WgtT5ss {
		wbdPage.AddFormItem(allegianceInput)
	}
	wbdPage.AddButton("Cancel", func() {
		mainPane.RemovePage(thisPage)
		mainPageFocus()
		refocus(idMenuPaneStr)
		return
	})
	wbdPage.AddButton("OK", func() {

		switch genType {
		case WgtMtBasic:
			_, allegianceText := allegianceInput.GetCurrentOption()
			_, trafficText := trafficInput.GetCurrentOption()
			logPane.Log("Generating basic MegaTraveller mainworld.")
			getNewMTWorld(nameInput.GetText(), hexInput.GetText(), sectorInput.GetText(), allegianceText, trafficText)
		case WgtMtWBH:
			_, allegianceText := allegianceInput.GetCurrentOption()
			_, trafficText := trafficInput.GetCurrentOption()
			logPane.Log("Generating detailed World Builder's Handbook MegaTraveller mainworld.")
			getNewWBHWorld(nameInput.GetText(), hexInput.GetText(), sectorInput.GetText(), allegianceText, trafficText)
		case WgtCt03:
			logPane.Log("Generating Classic Traveller mainworld.")
			getNewCT03World(nameInput.GetText(), hexInput.GetText(), sectorInput.GetText())
		case WgtT5ss:
			_, allegianceText := allegianceInput.GetCurrentOption()
			logPane.Log("Generating Traveller5 mainworld.")
			getNewT5World(nameInput.GetText(), hexInput.GetText(), sectorInput.GetText(), t5AllegianceMap[allegianceText])
		}
		mainPane.RemovePage(thisPage)
		mainPageFocus()
		return
	})
	mainPane.AddAndSwitchToPage(thisPage, wbdPage, false)
	refocus(idMainPaneStr)

}

// getNewMTWorld wraps the call to generate a new MegaTraveller basic world
// so it can be displayed.
func getNewMTWorld(name, hexLoc, sector, allegiance, traffic string) {
	w := generateMTWorld(name, hexLoc, sector, allegiance, traffic)

	if w.genType != WgtInvalid {
		// Output the result
		sidePane.DisplayObject(w.ObjectBasicString())
		log.Printf("World created: " + w.String())

		// Write to screen and file
		mainPage.Write([]byte(w.String() + "\n"))
		fn := config.WorldOutputFile
		if name != "" {
			fn = name + "_" + fn
		}
		fn = config.DataDir + fn
		w.toFile(fn)
	} else {
		logPane.Log("World generation failed")
	}

	// Return focus to menu pane
	refocus(idMenuPaneStr)

}

// getNewWBHWorld wraps the call to generate a new MegaTraveller World Builder's Handbook
// detailed world so it can be displayed.
func getNewWBHWorld(name, hexLoc, sector, allegiance, traffic string) {
	w := generateWBHWorld(name, hexLoc, sector, allegiance, traffic)

	if w.genType != WgtInvalid {

		// Output the result
		sidePane.DisplayObject(w.ObjectBasicString())
		log.Printf("World created: " + w.String())

		// Write to screen and file
		mainPage.Write([]byte(w.String() + "\n"))
		fn := config.WorldOutputFile
		if name != "" {
			fn = name + "_" + fn
		}
		fn = config.DataDir + fn
		w.toFile(fn)

	} else {
		logPane.Log("World generation failed")
	}
	// Return focus to menu pane
	refocus(idMenuPaneStr)

}

// getNewCT03World wraps the call to generate a new CT Book 3 world so it
// can be displayed on the main screen and in object window.
func getNewCT03World(name, hexLoc, sector string) {
	w := generateCT03World(name, hexLoc, sector)

	if w.genType != WgtInvalid {
		// Output the result
		sidePane.DisplayObject(w.ObjectBasicString())
		log.Printf("World created: " + w.String())

		// Write to screen and file
		mainPage.Write([]byte(w.String() + "\n"))
		fn := config.WorldOutputFile
		if name != "" {
			fn = name + "_" + fn
		}
		fn = config.DataDir + fn
		w.toFile(fn)

	} else {
		logPane.Log("World generation failed")
	}
	// Return focus to menu pane
	refocus(idMenuPaneStr)

}

// getNewT5World wraps the call the generate a new Traveller5 world and system
// so that it can be displayed on the main screen and object window.
func getNewT5World(name, hexLoc, sector, allegiance string) {

	w := generateT5World(name, hexLoc, sector, allegiance)

	if w.genType != WgtInvalid {
		// Output the result
		sidePane.DisplayObject(w.ObjectBasicString())
		log.Printf("World created: " + w.String())

		// Write to screen and file
		mainPage.Write([]byte(w.String() + "\n"))
		fn := config.WorldOutputFile
		if name != "" {
			fn = name + "_" + fn
		}
		fn = config.DataDir + fn
		w.toFile(fn)

	} else {
		logPane.Log("World generation failed")
	}

	// Return focus to menu pane
	refocus(idMenuPaneStr)

}

// getRandomSector collects the data for a random sector to be generated, populated
// with MegaTraveller basic mainworlds.
func getRandomSector() {

	collectSectorData := "collectSectorData"
	mainPageFocus()

	wbdPage := tview.NewForm()
	x, y, width, height := mainPage.GetRect()
	dialogWidth := 41
	dialogHeight := 14
	x = x + width/2 - dialogWidth/2
	y = y + height/2 - dialogHeight/2
	wbdPage.SetRect(x, y, dialogWidth, dialogHeight)
	wbdPage.SetBorder(true).SetTitle(" Generate Sector ").SetTitleAlign(tview.AlignCenter)
	nameInput := tview.NewInputField().SetLabel("Sector name").SetFieldWidth(24)
	allegiances := make([]string, 0, len(basicAllegianceMap))
	for k := range basicAllegianceMap {
		allegiances = append(allegiances, k)
	}
	sort.Strings(allegiances)
	ssDensity := make([]string, 0, len(mtSectorStarDensity))
	for ss := range mtSectorStarDensity {
		ssDensity = append(ssDensity, mtSectorStarDensity[ss])
	}
	ssTraffic := make([]string, 0, len(mtSubsectorTrafficArr))
	for ss := range mtSubsectorTrafficArr {
		ssTraffic = append(ssTraffic, mtSubsectorTrafficArr[ss])
	}

	// These options only needed for MegaTraveller
	allegianceInput := tview.NewDropDown().SetLabel("Allegiance").SetFieldWidth(24)
	allegianceInput.SetOptions(allegiances, nil)
	densityInput := tview.NewDropDown().SetLabel("Star density").SetFieldWidth(24)
	densityInput.SetOptions(ssDensity, nil)
	trafficInput := tview.NewDropDown().SetLabel("Sector age").SetFieldWidth(24)
	trafficInput.SetOptions(ssTraffic, nil)

	// Construct the form
	wbdPage.AddFormItem(nameInput)
	wbdPage.AddFormItem(allegianceInput)
	wbdPage.AddFormItem(densityInput)
	wbdPage.AddFormItem(trafficInput)
	wbdPage.AddButton("Cancel", func() {
		mainPane.RemovePage(collectSectorData)
		mainPageFocus()
		refocus(idMenuPaneStr)
		return
	})
	wbdPage.AddButton("OK", func() {
		//sLogLog("Clicked OK on Random Sector dialog")
		_, allegianceText := allegianceInput.GetCurrentOption()
		if allegianceText == "" {
			allegianceText = basicAllegianceMap["Imperial"]
		}
		_, densityText := densityInput.GetCurrentOption()
		if densityText == "" {
			densityText = mtSectorStarDensity[sdStandard]
		}
		_, trafficText := trafficInput.GetCurrentOption()
		if trafficText == "" {
			trafficText = mtSubsectorTrafficArr[ssStandard]
		}
		sectorText := nameInput.GetText()
		if sectorText == "" {
			sectorText = "Unknown"
		}
		mainPane.RemovePage(collectSectorData)
		mainPage.Clear()
		generateSector(sectorText, allegianceText, densityText, trafficText)
		return
	})
	mainPane.AddAndSwitchToPage(collectSectorData, wbdPage, false)
	refocus(idMainPaneStr)
}

// generateForeven generates all the missing worlds in the Foreven sector, with starting
// data taken from the database. The world generated (including those already provided),
// are written to the screen and to file
func generateForeven() {
	logPane.Log("Generating Foreven sector")

	sec := getForeven()
	log.Printf("Retrieved Foreven sector:")
	log.Printf("%v", sec)

	// Range over the locations and generate a world for each
	for i, w := range sec.worlds {
		//(sec.worlds[i]).genType = WgtT5ss
		msg := fmt.Sprintf("Subsector %v, Location %v", w.subsectorIndex, w.hexLoc.String())
		if w.name == "" || strings.Contains(w.uwp.String(), "?????") {
			if w.name == "" {
				w.name = "????"
			}

			// Handle allegiances. The generator is expected a Name not a code.
			if w.allegiance == "XXXX" {
				w.allegiance = "NaHu"
			}
			sec.worlds[i] = generateT5World(w.name, w.hexLoc.String(), w.sector, w.allegiance)
			log.Printf(msg + sec.worlds[i].String())
		} else {
			sec.worlds[i] = *w.extendWorld()
		}
	}

	tablePage := tview.NewTable()
	tablePage.SetBorder(false)

	header := strings.Split(headerOut[WgtT5ss], "\t")
	cols := len(header)
	rows := len(sec.worlds) + 1
	for r := 0; r < rows; r++ {
		if r == 0 {
			for c := 0; c < cols; c++ {
				tablePage.SetCell(r, c, tview.NewTableCell(header[c]).SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignCenter))
			}
		} else {
			updateT5TablePageRow(tablePage, r, &sec.worlds[r-1])
		}
	}
	tablePage.SetSelectable(true, false).SetFixed(1, 0)
	tablePage.SetEvaluateAllRows(true)

	// Handle operations on the table.
	tablePage.SetSelectionChangedFunc(func(row int, column int) {
		if row == 0 {
			sidePane.Clear()
			return
		}
		w := sec.worlds[row-1]
		sidePane.DisplayObject(w.ObjectBasicString())
	})
	tablePage.SetSelectedFunc(func(row int, column int) {
		if row == 0 {
			return
		}
		editSectorWorld(tablePage, sec, row-1)
	})
	tablePage.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEsc {
			// Here we want to check to save the sector.
			if sec.saved {
				mainPane.RemovePage(idSectorPageStr)
				mainPageFocus()
				refocus(idMenuPaneStr)
			} else {
				// Save Yes/No before removing the data!
				saveSectorConfirm(*sec)
			}
		}
	})
	tablePage.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		key := event.Key()

		if key == tcell.KeyCtrlS {
			// Output to the "sector" file.
			fn := config.DataDir + sec.name + ".tab"
			if err := sec.toFile(fn); err != nil {
				logPane.Log("Unable to write sector file")
			} else {
				sec.saved = true
				setSaveStatus(true)
				logPane.Log("Sector saved to file")
			}
			return nil
		} else if key == tcell.KeyBacktab || key == tcell.KeyTab {
			return handleTabs(idSectorPageStr, event)
		} else if key == tcell.KeyCtrlA {
			// Copy to the clipboard
			clipboard.WriteAll(sec.toTab())
			logPane.Log("Sector written to clipboard")
		} else if key == tcell.KeyF2 {
			// Edit the world entry
			row, _ := tablePage.GetSelection()
			if row == 0 {
				return nil
			}
			editSectorWorld(tablePage, sec, row-1)
		}
		return event

	})

	mainPane.AddAndSwitchToPage(idSectorPageStr, tablePage, true)
	setSaveStatus(false)
	refocus(idMainPaneStr)

}

// updateTablePageRow updates the specific row in the tablePage with the contents
// on the world. Can be used for initial construction or updates. Cut sick.
func updateT5TablePageRow(tp *tview.Table, row int, w *world) {

	tp.SetCell(row, 0, tview.NewTableCell(w.sectorAbbrev).SetTextColor(tcell.ColorWhite).SetAlign(tview.AlignLeft))
	tp.SetCell(row, 1, tview.NewTableCell(w.subsectorIndex).SetTextColor(tcell.ColorWhite).SetAlign(tview.AlignLeft))
	tp.SetCell(row, 2, tview.NewTableCell(w.hexLoc.String()).SetTextColor(tcell.ColorWhite).SetAlign(tview.AlignLeft))
	tp.SetCell(row, 3, tview.NewTableCell(w.name).SetTextColor(tcell.ColorWhite).SetAlign(tview.AlignLeft).SetReference(&w.name))
	tp.SetCell(row, 4, tview.NewTableCell(w.uwp.String()).SetTextColor(tcell.ColorWhite).SetAlign(tview.AlignLeft))
	tp.SetCell(row, 5, tview.NewTableCell(w.bases).SetTextColor(tcell.ColorWhite).SetAlign(tview.AlignLeft))
	tp.SetCell(row, 6, tview.NewTableCell(w.remarks).SetTextColor(tcell.ColorWhite).SetAlign(tview.AlignLeft))
	tp.SetCell(row, 7, tview.NewTableCell(w.zone.String()).SetTextColor(tcell.ColorWhite).SetAlign(tview.AlignLeft))
	tp.SetCell(row, 8, tview.NewTableCell(w.pbg.String()).SetTextColor(tcell.ColorWhite).SetAlign(tview.AlignLeft))
	tp.SetCell(row, 9, tview.NewTableCell(w.allegiance).SetTextColor(tcell.ColorWhite).SetAlign(tview.AlignLeft))
	tp.SetCell(row, 10, tview.NewTableCell(StarString(w.stars)).SetTextColor(tcell.ColorWhite).SetAlign(tview.AlignLeft))
	tp.SetCell(row, 11, tview.NewTableCell(w.importance.String()).SetTextColor(tcell.ColorWhite).SetAlign(tview.AlignLeft))
	tp.SetCell(row, 12, tview.NewTableCell(w.economics.String()).SetTextColor(tcell.ColorWhite).SetAlign(tview.AlignLeft))
	tp.SetCell(row, 13, tview.NewTableCell(w.culture.StringEsc()).SetTextColor(tcell.ColorWhite).SetAlign(tview.AlignLeft))
	tp.SetCell(row, 14, tview.NewTableCell(w.nobility).SetTextColor(tcell.ColorWhite).SetAlign(tview.AlignLeft))
	tp.SetCell(row, 15, tview.NewTableCell(strconv.Itoa(w.worlds)).SetTextColor(tcell.ColorWhite).SetAlign(tview.AlignLeft))
	tp.SetCell(row, 16, tview.NewTableCell(strconv.Itoa(w.ru)).SetTextColor(tcell.ColorWhite).SetAlign(tview.AlignLeft))

}

// generateSector creates a new sector full of mainworlds, randomly determining
// if a system is present in each hex and generating the world for that.
func generateSector(sectorName, allegiance, density, traffic string) {

	logPane.Log("Generating random sector")

	var locs []HexLoc

	for x := 1; x <= 32; x++ {
		for y := 1; y <= 40; y++ {

			// Get the actual hex location
			hexLoc := HexLoc{x: x, y: y, sector: true}

			switch density {
			case mtSectorStarDensity[sdRift]:
				if D6()+D6() <= 2 {
					locs = append(locs, hexLoc)
				}
			case mtSectorStarDensity[sdSparse]:
				if D6() >= 6 {
					locs = append(locs, hexLoc)
				}
			case mtSectorStarDensity[sdScattered]:
				if D6() >= 5 {
					locs = append(locs, hexLoc)
				}
			case mtSectorStarDensity[sdStandard]:
				if D6() >= 4 {
					locs = append(locs, hexLoc)
				}
			case mtSectorStarDensity[sdDense]:
				if D6() >= 3 {
					locs = append(locs, hexLoc)
				}
			}
		}
	}

	// We have a slice of HexLocs, sort them before populating them.
	sort.Sort(ByLoc(locs))

	// A blank sector to put all those worlds into.
	sec := sector{id: -1, name: sectorName, saved: false, otu: false}

	// Range over the locations and generate a world for each.
	for _, h := range locs {
		msg := fmt.Sprintf("Subsector %v, Location %v", h.GetIndex(), h.String())
		w := generateMTWorld("????", h.String(), sectorName, allegiance, traffic)
		sec.worlds = append(sec.worlds, w)
		log.Printf(msg + w.String())
		//s += w.String() + "\n"
	}

	// Create and display a table page for output of the list of worlds generated.
	tablePage := tview.NewTable()
	tablePage.SetBorder(false)
	header := strings.Split(headerOut[WgtMtBasic], "\t")
	cols := len(header)
	rows := len(sec.worlds) + 1
	for r := 0; r < rows; r++ {
		if r == 0 {
			for c := 0; c < cols; c++ {
				tablePage.SetCell(r, c, tview.NewTableCell(header[c]).SetTextColor(tcell.ColorYellow).SetAlign(tview.AlignCenter))
			}
		} else {
			updateTablePageRow(tablePage, r, &sec.worlds[r-1])
		}
	}
	tablePage.SetSelectable(true, false).SetFixed(1, 0)
	tablePage.SetEvaluateAllRows(true)

	// Handle operations on the table.
	tablePage.SetSelectionChangedFunc(func(row int, column int) {
		if row == 0 {
			sidePane.Clear()
			return
		}
		w := sec.worlds[row-1]
		sidePane.DisplayObject(w.ObjectBasicString())
	})
	tablePage.SetSelectedFunc(func(row int, column int) {
		if row == 0 {
			return
		}
		editSectorWorld(tablePage, &sec, row-1)
	})
	tablePage.SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEsc {
			// Here we want to check to save the sector.
			if sec.saved {
				mainPane.RemovePage(idSectorPageStr)
				mainPageFocus()
				refocus(idMenuPaneStr)
			} else {
				// Save Yes/No before removing the data!
				saveSectorConfirm(sec)
			}
		}
	})
	tablePage.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		key := event.Key()

		if key == tcell.KeyCtrlS {
			// Output to the "sector" file.
			fn := config.DataDir + sec.name + ".tab"
			if err := sec.toFile(fn); err != nil {
				logPane.Log("Unable to write sector file")
			} else {
				sec.saved = true
				setSaveStatus(true)
				logPane.Log("Sector saved to file")
			}
			return nil
		} else if key == tcell.KeyBacktab || key == tcell.KeyTab {
			return handleTabs(idSectorPageStr, event)
		} else if key == tcell.KeyCtrlA {
			// Copy to the clipboard
			clipboard.WriteAll(sec.toTab())
			logPane.Log("Sector written to clipboard")
		} else if key == tcell.KeyF2 {
			// Edit the world entry
			row, _ := tablePage.GetSelection()
			if row == 0 {
				return nil
			}
			editSectorWorld(tablePage, &sec, row-1)
		}
		return event

	})

	mainPane.AddAndSwitchToPage(idSectorPageStr, tablePage, true)
	setSaveStatus(false)
	refocus(idMainPaneStr)
}

// updateTablePageRow updates the specific row in the tablePage with the contents
// on the world. Can be used for initial construction or updates. Cut sick.
func updateTablePageRow(tp *tview.Table, row int, w *world) {
	tp.SetCell(row, 0, tview.NewTableCell(w.sectorAbbrev).SetTextColor(tcell.ColorWhite).SetAlign(tview.AlignLeft))
	tp.SetCell(row, 1, tview.NewTableCell(w.subsectorIndex).SetTextColor(tcell.ColorWhite).SetAlign(tview.AlignLeft))
	tp.SetCell(row, 2, tview.NewTableCell(w.hexLoc.String()).SetTextColor(tcell.ColorWhite).SetAlign(tview.AlignLeft))
	tp.SetCell(row, 3, tview.NewTableCell(w.name).SetTextColor(tcell.ColorWhite).SetAlign(tview.AlignLeft).SetReference(&w.name))
	tp.SetCell(row, 4, tview.NewTableCell(w.uwp.String()).SetTextColor(tcell.ColorWhite).SetAlign(tview.AlignLeft))
	tp.SetCell(row, 5, tview.NewTableCell(w.bases).SetTextColor(tcell.ColorWhite).SetAlign(tview.AlignLeft))
	tp.SetCell(row, 6, tview.NewTableCell(w.remarks).SetTextColor(tcell.ColorWhite).SetAlign(tview.AlignLeft))
	tp.SetCell(row, 7, tview.NewTableCell(w.zone.String()).SetTextColor(tcell.ColorWhite).SetAlign(tview.AlignLeft))
	tp.SetCell(row, 8, tview.NewTableCell(w.pbg.String()).SetTextColor(tcell.ColorWhite).SetAlign(tview.AlignLeft))
	tp.SetCell(row, 9, tview.NewTableCell(w.allegiance).SetTextColor(tcell.ColorWhite).SetAlign(tview.AlignLeft))
}

// editSectorWorld edits an individual world within a sector. idx is the index into the
// slice of worlds within the sector.
func editSectorWorld(tablePage *tview.Table, s *sector, idx int) {

	thisPage := "editWorld"

	w := s.worlds[idx]

	edwPage := tview.NewForm()
	x, y, width, height := mainPage.GetRect()
	dialogWidth := 41
	dialogHeight := 14
	x = x + width/2 - dialogWidth/2
	y = y + height/2 - dialogHeight/2
	edwPage.SetRect(x, y, dialogWidth, dialogHeight)
	edwPage.SetBorder(true).SetTitle(" Edit World ").SetTitleAlign(tview.AlignCenter)

	/*
	 * Fields we want to edit here:
	 * x Name
	 * - Bases (from a combination of codes)
	 * x Travel Zone (radio button style R,A,G) (or Drop-down selection)
	 * x Allegiance
	 * - PBG (3 separate fields)
	 */

	// Name
	nameInput := tview.NewInputField().SetLabel("World name").SetFieldWidth(24)
	nameInput.SetText(w.name)

	// Travel Zone
	zoneInput := tview.NewDropDown().SetLabel("Travel zone").SetFieldWidth(24)
	zoneInput.SetOptions([]string{TzGreen.Desc(), TzAmber.Desc(), TzRed.Desc()}, nil)
	zoneInput.SetCurrentOption(int(w.zone))

	// Allegiance
	allegianceInput := tview.NewDropDown().SetLabel("Allegiance").SetFieldWidth(24)
	allegiances := make([]string, 0, len(basicAllegianceMap))
	for k := range basicAllegianceMap {
		allegiances = append(allegiances, k)
	}
	sort.Strings(allegiances)
	allegianceIndex := -1
	for k, v := range allegiances {
		if basicAllegianceMap[v] == w.allegiance {
			allegianceIndex = k
		}
	}
	allegianceInput.SetOptions(allegiances, nil)
	if allegianceIndex != -1 {
		allegianceInput.SetCurrentOption(allegianceIndex)
	}

	// Construct the form
	edwPage.AddFormItem(nameInput)
	edwPage.AddFormItem(zoneInput)
	edwPage.AddFormItem(allegianceInput)
	edwPage.AddButton("Cancel", func() {
		mainPane.RemovePage(thisPage)
		refocus(idMainPaneStr)
		return
	})
	edwPage.AddButton("Save", func() {
		logPane.Log("Saving world to sector")
		// Get new values
		w.name = nameInput.GetText()
		x, _ := zoneInput.GetCurrentOption()
		w.zone = TravelZone(x)
		_, y := allegianceInput.GetCurrentOption()
		w.allegiance = basicAllegianceMap[y]
		s.worlds[idx] = w
		mainPane.RemovePage(thisPage)
		// Update the display with the changed world
		updateTablePageRow(tablePage, idx+1, &s.worlds[idx])
		sidePane.DisplayObject(w.ObjectBasicString())
		mainPane.SwitchToPage(idSectorPageStr)
		s.saved = false
		setSaveStatus(false)
		// The refocus will do a redraw.
		//app.ForceDraw()
		refocus(idMainPaneStr)
		return
	})
	mainPane.AddAndSwitchToPage(thisPage, edwPage, false)
	refocus(idMainPaneStr)
}

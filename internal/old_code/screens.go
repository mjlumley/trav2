package main

// screens.go contains the code for screen handling

import (
	"fmt"
	"log"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

var (
	mainPane   *tview.Pages       // The main window, central to the application
	sidePane   *ObjectPane        // Used for displaying an object during generation, such as a user or world
	logPane    *LogPane           // Used for output of log and info messages from generation
	menuPane   *tview.TreeView    // The menu column on the left allows choosing a command
	titleBar   *tview.TextView    // Displays the application title and level
	contextBar *tview.TextView    // Displays contextual usage information
	saveBar    *tview.TextView    // Displays save status for objects generated in the application
	app        *tview.Application // The application object
	grid       *tview.Grid        // The grid that manages the windows
)

// A set of strings to identify the usable primitives.

// idMenuPaneStr is the identifier for the idMenuPaneStr pane.
const idMenuPaneStr string = "menu"

// idMainPaneStr is the identifier for the main pane.
const idMainPaneStr string = "main"

// idObjectPaneStr is the identifier for the RHS object pane.
const idObjectPaneStr string = "object"

// idLogPaneStr is the identifier for the logging pane.
const idLogPaneStr string = "log"

// idSectorPageStr is the identifier for the sector page on the main pane.
const idSectorPageStr string = "sectorPage"

// idGeneralPageStr is the identifier for the main page on the main pane (!!)
const idGeneralPageStr string = "mainpage"

// sbContext is the status bar base text
const sbContext string = "[yellow]Ctrl-C[-]: Exit | [yellow]F5[-]: [\"" + idMenuPaneStr + "\"]Menu[\"\"] | [yellow]F6[-]: [\"" + idMainPaneStr + "\"]Main[\"\"] | [yellow]F7[-]: [\"" +
	idObjectPaneStr + "\"]Object[\"\"] | [yellow]F8[-]: [\"" + idLogPaneStr + "\"]Log[\"\"]"

var mainPage *ScrollableTextView

// notMain is not the main function, but sets up the application and prime screen elements.
func notMain() {

	// Create the application and setup the screen.
	log.Printf("Creating application")
	app = tview.NewApplication()
	app.EnableMouse(true)

	// The main window
	log.Printf("Creating Main pages pane")
	mainPane = tview.NewPages()
	mainPane.SetBorder(true).SetTitle(" Main ").SetTitleAlign(tview.AlignCenter)
	log.Printf("Creating new ScollableTextView")
	mainPage = NewScrollableTextView()
	mainPage.SetDynamicColors(true).SetTextAlign(tview.AlignLeft) //.SetBorder(true).SetTitle(" Main ").SetTitleAlign(tview.AlignCenter)
	mainPage.SetScrollable(true).ScrollToEnd().SetWordWrap(true)
	mainPage.SetChangedFunc(func() {
		log.Printf("mainPage SetChangedFunc()")
		app.ForceDraw()
	})
	mainPane.AddAndSwitchToPage(idGeneralPageStr, mainPage, true)
	mainPage.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		key := event.Key()
		if key == tcell.KeyTab || key == tcell.KeyBacktab {
			return handleTabs(idGeneralPageStr, event)
		}
		return event
	})

	// The side Object for progress display of objects (characters, worlds etc)
	log.Printf("Creating ObjectPane")
	sidePane = NewObjectPane(idObjectPaneStr)
	sidePane.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		key := event.Key()
		if key == tcell.KeyTab || key == tcell.KeyBacktab {
			return handleTabs(sidePane.GetIdentifier(), event)
		}
		return event
	})
	sidePane.SetChangedFunc(func() {
		log.Printf("Side Pane SetChangedFunc")
		app.ForceDraw()
	})

	// The log section displays ongoing progress and information
	log.Printf("Creating LogPane")
	logPane = NewLogPane(idLogPaneStr)
	logPane.SetChangedFunc(func() {
		log.Printf("Log Pane SetChangedFunc")
		app.ForceDraw()
		log.Printf("Finished drawing in LogPane SetChangedFunc")
	})
	logPane.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		key := event.Key()
		if key == tcell.KeyBackspace || key == tcell.KeyDelete {
			log.Printf("DEL/BS received in Log Window. Clearing.")
			logPane.Clear()
			return nil
		} else if key == tcell.KeyTab || key == tcell.KeyBacktab {
			return handleTabs(idLogPaneStr, event)
		}
		return event
	})

	// The title bar displays the app name and level
	log.Printf("Creating TextView for TitleBar")
	titleBar = tview.NewTextView()
	titleBar.SetDynamicColors(true).SetTextAlign(tview.AlignCenter).SetText("The Traveller's Tool")
	titleBar.SetBackgroundColor(tcell.ColorTeal)
	titleBar.SetScrollable(false).SetWordWrap(false)

	// The status bar displays help information and other alert
	log.Printf("Creating TextView for ContextBar")
	contextBar = tview.NewTextView()
	contextBar.SetDynamicColors(true).SetTextAlign(tview.AlignLeft).SetText(sbContext).SetRegions(true)
	contextBar.SetBackgroundColor(tcell.ColorTeal)
	contextBar.SetScrollable(false).SetWordWrap(false)
	contextBar.SetChangedFunc(func() {
		log.Printf("ContextBar entered SetChangedFunc")
		app.ForceDraw()
	})

	// The save bar indicates whether an object has been saved (blank) or not (asterisk shown).
	log.Printf("Creating TextView for SaveBar")
	saveBar = tview.NewTextView()
	saveBar.SetDynamicColors(true).SetTextAlign(tview.AlignLeft).SetText(" ")
	saveBar.SetBackgroundColor(tcell.ColorTeal)
	saveBar.SetScrollable(false).SetWordWrap(false)
	saveBar.SetChangedFunc(func() {
		log.Printf("saveBar SetChangedFunc()")
		app.ForceDraw()
	})

	// The menu is where commands are given
	log.Printf("Creating TreeView for MenuPane")
	menuPane = tview.NewTreeView().SetTopLevel(1).SetAlign(false).SetGraphics(true)
	menuPane.SetBorder(true)
	log.Printf("About to initMenuPane")
	initMenuPane()

	// The grid is the display manager
	log.Printf("Creating Grid")
	grid = tview.NewGrid().
		SetRows(1, 0, 10, 1).
		SetColumns(24, 0, 36, 1).
		SetBorders(false)
	grid.AddItem(titleBar, 0, 0, 1, 4, 0, 0, false)
	grid.AddItem(menuPane, 1, 0, 2, 1, 0, 0, true)
	grid.AddItem(mainPane, 1, 1, 1, 1, 0, 0, false)
	grid.AddItem(sidePane, 1, 2, 1, 2, 0, 0, false)
	grid.AddItem(logPane, 2, 1, 1, 3, 0, 0, false)
	grid.AddItem(contextBar, 3, 0, 1, 3, 0, 0, false)
	grid.AddItem(saveBar, 3, 3, 1, 1, 1, 1, false)

	log.Printf("About to start the main grid/window.")
	app.SetInputCapture(globalKeyProcess)
	if err := app.SetRoot(grid, true).SetFocus(menuPane).Run(); err != nil {
		panic(err)
	}
	log.Printf("Running... - Unlikely we'll EVER get here.")
}

// refocus changes the pane that receives the focus. This changes the text on the context bar as well.
func refocus(newPane string) tview.Primitive {
	log.Printf("refocus(newPane) : %v", newPane)

	if newPane == "" {
		panic("Empty value to refocus()")
	}

	log.Printf("Refocus to : " + newPane)

	switch newPane {
	case idMenuPaneStr:
		menuStatusText(menuPane.GetCurrentNode())
		contextBar.Highlight(idMenuPaneStr)
		app.SetFocus(menuPane)
		return menuPane
	case idMainPaneStr:
		barText := sbContext
		name, _ := mainPane.GetFrontPage()
		if name == idSectorPageStr {
			barText = sbContext + " | [yellow]Ctrl-S[-] to save sector, [yellow]Ctrl-A[-] to copy"
		}
		contextBar.SetText(barText)
		contextBar.Highlight(idMainPaneStr)
		app.SetFocus(mainPane)
		return mainPane
	case idObjectPaneStr:
		contextBar.SetText(sbContext)
		contextBar.Highlight(idObjectPaneStr)
		app.SetFocus(sidePane)
		return sidePane
	case idLogPaneStr:
		contextBar.SetText(sbContext)
		contextBar.Highlight(idLogPaneStr)
		app.SetFocus(logPane)
		return logPane
	default:
		panic("Unknown pane : " + newPane)
	}
}

// setSavedStatus displays an asterisk at the bottom right hand corner of the context bar to
// indicate that an object is saved or unsaved.
func setSaveStatus(saved bool) {
	saveBar.Clear()
	if !saved {
		saveBar.SetBackgroundColor(tcell.ColorRed)
		saveBar.Write([]byte("*"))
	} else {
		saveBar.SetBackgroundColor(tcell.ColorTeal)
	}
	return
}

// mainPageFocus ensures that the main page on the central, main pane, is upper-most and showing.
func mainPageFocus() {
	mainPane.SendToFront(idGeneralPageStr)
	app.ForceDraw()
}

// getCTBook01CharacterBasicData collects the basic data required for a Classic Traveller Book 01
// character and then passes the collected data onto the generator.
func getCTBook01CharacterBasicData() {
	logPane.Log("Collecting basic character data for Book 01 Classic Traveller")

	// Collect list of races from the database
	races, err := GetAllMajorRaces()
	if err != nil {
		logPane.Log("DB Error retrieving races: " + err.Error())
	}

	// Get a new character to start the process
	c := NewCT01Char()

	// Create and position the Dialog box
	cbdPage := tview.NewForm()
	dialogWidth := 52
	dialogHeight := 18
	x, y, width, height := mainPage.GetRect()
	x = x + width/2 - dialogWidth/2
	y = y + height/2 - dialogHeight/2
	cbdPage.SetRect(x, y, dialogWidth, dialogHeight)
	cbdPage.SetBorder(true).SetTitle("CT Book 1 Character Basic Data").SetTitleAlign(tview.AlignCenter)
	nameInput := tview.NewInputField().SetLabel("Name").SetFieldWidth(24)
	uppField := NewReadonlyField().SetLabel("UPP")
	uppField.SetText(c.UPP())
	serviceInput := tview.NewDropDown().SetLabel("Preferred Service").SetFieldWidth(24)
	var svcOpts []string
	for i := 1; i < 7; i++ {
		svcOpts = append(svcOpts, CareerCT01(i).String())
	}
	serviceInput.SetOptions(svcOpts, nil)
	raceInput := tview.NewInputField().SetLabel("Race").SetFieldWidth(24)
	raceInput.SetAutocompleteFunc(func(currentText string) (entries []string) {
		// Ignore empty text
		if len(currentText) == 0 {
			return
		}
		for _, race := range races {
			if strings.HasPrefix(strings.ToLower(race), strings.ToLower(currentText)) {
				entries = append(entries, race)
			}
		}
		if len(entries) <= 1 {
			entries = nil
		}
		return
	})
	sexInput := tview.NewInputField().SetLabel("Sex").SetFieldWidth(24).SetText("Male")
	surviveInput := tview.NewCheckbox().SetLabel("Dies on survival fail?")
	surviveInput.SetChecked(false)
	cbdPage.AddFormItem(nameInput).
		AddFormItem(uppField).
		AddFormItem(serviceInput).
		AddFormItem(raceInput).
		AddFormItem(sexInput).
		AddFormItem(surviveInput)
	cbdPage.SetCancelFunc(cancelFromCbd).
		AddButton("Cancel", cancelFromCbd).
		AddButton("OK", func() {
			mainPane.RemovePage("charbasic")
			_, serviceChosen := serviceInput.GetCurrentOption()
			logPane.Log(fmt.Sprintf("Results := Name: %v, Service: %v, Race: %v, Sex: %v, Dies?: %v", nameInput.GetText(), serviceChosen, raceInput.GetText(), sexInput.GetText(), surviveInput.IsChecked()))
			c.SetName(nameInput.GetText()).SetRace(raceInput.GetText()).SetSex(sexInput.GetText())
			c.GenerateCT01Character(serviceChosen, surviveInput.IsChecked())
			refocus(idMenuPaneStr)
			return
		})
	mainPane.AddPage("charbasic", cbdPage, false, false)
	mainPane.ShowPage("charbasic")
	// Should be no need to redraw since the refocus will cause redraw.
	//app.ForceDraw()
	refocus(idMainPaneStr)
}

// This cancels from the Character Basic Data dialog box.
func cancelFromCbd() {
	mainPane.RemovePage("charbasic")
	refocus(idMenuPaneStr)
	return
}

// areYouSure displays a modal box asking the user if they REALLY want to continue. You supply a function to continue on with, whether Yes or No is default, and the question to ask.
func askYesNoQuestion(defaultYes bool, questionStr string, nextFunc func()) {
	logPane.Log("Confirming Yes or No")

	var previousFocus tview.Primitive

	previousFocus = app.GetFocus()

	questionPage := tview.NewModal().SetText(questionStr)
	questionPage.AddButtons([]string{"Yes", "No"})
	questionPage.SetTitleAlign(tview.AlignCenter)
	if defaultYes {
		questionPage.SetFocus(0)
	} else {
		questionPage.SetFocus(1)
	}
	questionPage.SetDoneFunc(func(buttonIndex int, buttonLabel string) {
		if buttonLabel == "Yes" {
			mainPane.RemovePage("yesno")
			nextFunc()
			return
		}
		mainPane.RemovePage("yesno")
		if previousFocus != nil {
			app.SetFocus(previousFocus)
		}
		return
	})
	mainPane.AddPage("yesno", questionPage, false, false)
	mainPane.ShowPage("yesno")
	// Should be no need to redraw since the refocus will cause redraw.
	//app.ForceDraw()
	refocus(idMainPaneStr)
}

// globalKeyProcess process key events for the application.
func globalKeyProcess(event *tcell.EventKey) *tcell.EventKey {

	key := event.Key()
	switch key {
	case tcell.KeyCtrlC: // Catch the Ctrl-C for exit
		areYouSureExit()
		return nil
	case tcell.KeyF5: // F5 for the menu pane
		refocus(idMenuPaneStr)
		return nil
	case tcell.KeyF6: // F6 for the main pane
		refocus(idMainPaneStr)
		return nil
	case tcell.KeyF7: // F7 for the object/side pane
		refocus(idObjectPaneStr)
		return nil
	case tcell.KeyF8: // F8 for the object/side pane
		refocus(idLogPaneStr)
		return nil
	case tcell.KeyF1: // F1for help
		doHelp()
		return nil
	case tcell.KeyCtrlA: // Copy all text to the clipboard
		badWindow := "Unable to copy text off that window."
		pane := app.GetFocus()
		if pane == menuPane {
			logPane.Log(badWindow)
			return nil
		}
		if sp, ok := pane.(*tview.TextView); ok {
			clipboard.WriteAll(sp.GetText(true))
			logPane.Log("Window text written to clipboard.")
		} else if sp, ok := pane.(*ScrollableTextView); ok {
			clipboard.WriteAll(sp.GetText(true))
			logPane.Log("Window text written to clipboard.")
		} else if _, ok := pane.(*tview.Table); ok {
			logPane.Log("Passing onto lower-level handler")
			return event
		}
		return nil
	default:
		return event
	}
}

var comms chan *tcell.EventKey
var choice chan string

// captureMain handles key events in the main box
func captureMain(event *tcell.EventKey) *tcell.EventKey {
	logPane.Log("Key event received for " + event.Name())

	if event.Key() == tcell.KeyEscape {
		// We want to finish here
		logPane.Log("Received an ESC - EXITING!")
		refocus(idMenuPaneStr)
	}
	// TODO: You've got to be kidding!
	//mainPage.Write([]byte(string([]rune{event.Rune()})))
	comms <- event
	log.Printf("Adding a key to the channel %s", event.Name())
	return event
}

func doTest() {
	logPane.Log("Testing functionality - renewed")
	mainPage.Clear()
	comms = make(chan *tcell.EventKey)
	choice = make(chan string)

	mainPage.Write([]byte("Here we testing go!\n"))

	mainPage.SetInputCapture(captureMain)
	refocus(idMainPaneStr)

	mainPage.Write([]byte("I'm looking for a 1, 2 or 3."))
	logPane.Log("Forcing draw on app")
	app.ForceDraw()
	log.Printf("About to enter go func")
	go func() {
		var retString string
	WaitLoop:
		for {

			key := <-comms

			// We are going to consume it regardless
			switch key.Name() {
			case "Rune[1]":
				retString = "Received a 1"
				break WaitLoop
			case "Rune[2]":
				retString = "Received a 2"
				break WaitLoop
			case "Rune[3]":
				retString = "Received a 3"
				break WaitLoop
			default:
				logPane.Log("Ignoring " + key.Name())
			}
		}
		mainPage.SetInputCapture(nil)
		logPane.Log(retString)
		refocus(idMenuPaneStr)
		return
	}()
	refocus(idMenuPaneStr)
}

// copyStageToProd moves all world data in the world_staging table to the world table.
func copyStageToProd() {
	logPane.Log("copyStageToProd()")
	askYesNoQuestion(false, "Are you sure you want to copy data from 'world_staging' to 'worlds' tables?", copyWorldStagingToProd)
}

// deleteStaged deletes all worlds from the world_staging table.
func deleteStaged() {
	logPane.Log("deleteStaged()")
	// Ask the user to confirm this is what they want to do
	askYesNoQuestion(false, "Are you sure you want to delete all staged world data", clearWorldStaging)
}

// missingWorlds run the Missing Worlds to retrieve world data from travellermap.com.
func missingWorlds() {
	logPane.Log("missingWorlds()")
	askYesNoQuestion(false, "Are you sure you want to retrieve data from travellermap.com? (This will take some time!)", retrieveWorldData)
}

// missingSubs runs the missing subsectors function to retrieve sector data from travellemap.com
func missingSubs() {
	logPane.Log("missingSubs()")
	str := fmt.Sprintf("Do you wish to retrieve missing data for %d sectors from travellermap.com? (This will take some time!)", countSubsectorsToRetrieve())
	askYesNoQuestion(false, str, retrieveSubsectorData)
}

// handleTabs handles "tabbing" around the application screen. The screens are navigated in the
// order: menu -> main -> object -> log -> menu. Note that the main pane can be any one of a
// variety of pages, and they really should handle the tabbing themselves. Refocussing on the
// main pane should send focu to the uppermost page.
func handleTabs(source string, event *tcell.EventKey) *tcell.EventKey {
	key := event.Key()
	log.Printf("Tab event received: %v in source: %v", key, source)

	switch source {
	case idMenuPaneStr:
		if key == tcell.KeyBacktab {
			refocus(idLogPaneStr)
			return nil
		} else if key == tcell.KeyTab {
			refocus(idMainPaneStr)
			return nil
		}
	case idMainPaneStr, idGeneralPageStr, idSectorPageStr:
		if key == tcell.KeyBacktab {
			refocus(idMenuPaneStr)
			return nil
		} else if key == tcell.KeyTab {
			refocus(idObjectPaneStr)
			return nil
		}
	case idObjectPaneStr:
		if key == tcell.KeyBacktab {
			refocus(idMainPaneStr)
			return nil
		} else if key == tcell.KeyTab {
			refocus(idLogPaneStr)
			return nil
		}
	case idLogPaneStr:
		if key == tcell.KeyBacktab {
			refocus(idObjectPaneStr)
			return nil
		} else if key == tcell.KeyTab {
			refocus(idMenuPaneStr)
			return nil
		}
	default:
		return event
	}
	return event
}

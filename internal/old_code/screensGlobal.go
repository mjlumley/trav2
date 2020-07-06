package main

import (
	"fmt"
	"log"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

// screensGlobal contains windows, screens, dialogs, etc that are needed globally.

// doAbout shows an About dialog box, with application version displayed.
func doAbout() {
	log.Printf("doAbout()")
	logPane.Log("Running About box")

	versionString := fmt.Sprintf("The Traveller's Tool\n\n\u2261 Version %s", appVersion)
	aboutPage := tview.NewModal().SetText(versionString)
	aboutPage.SetTitle(" About ")
	aboutPage.AddButtons([]string{"OK"})
	aboutPage.SetDoneFunc(func(buttonIndex int, buttonLabel string) {
		app.SetRoot(grid, true)
		return
	})
	app.SetRoot(aboutPage, false)
}

// doHelp displays a modal generic Help page.
func doHelp() {
	helpString :=
		`Welcome to [yellow]The Traveller's Tool[-]

Keyboard Shortcuts
  Arrow keys = Navigate the menu
  Enter = Expand/collapse menu or execute the menu option
  F1 = Help
  F5 = Switch to the menu tree
  F6 = Switch to the main (central) window
  F7 = Switch to the object pane on right-hand side
  F8 = Switch to the logging pane on bottom right
  Tab/Shift-Tab = Move back and forth between panes
  Ctrl-A = Copy all text in window to clipboard
  Ctrl-S = Save data (where applicable)
  Ctrl-C = Exit the application

Instructions
To use the application, select the command from the cascading menu. Hints will appear on the Status Bar at the bottom while navigating the menu. Output (for instance, word generation) will appear in the central window. During some generation (for example characters and worlds), the object being built will appear in the "object pane" on the right-hand side.

In most windows, you can use [red]Ctrl-A[-] to select and copy all text to the clipboard. To do this, select the window (F6/F7/F8) and then hit Ctrl-A. The windows text should be copied to the clipboard.

Hit [red]<enter>[-] to return to the application.`

	x, y, width, height := grid.GetRect()
	dialogWidth := 88
	dialogHeight := 32
	x = x + width/2 - dialogWidth/2
	y = y + height/2 - dialogHeight/2

	helpPage := tview.NewTextView().SetText(helpString)
	helpPage.SetRect(x, y, dialogWidth, dialogHeight)
	helpPage.SetTitle(" Help ").SetTitleAlign(tview.AlignCenter)
	helpPage.SetDoneFunc(func(key tcell.Key) {
		app.SetRoot(grid, true)
		return
	})
	helpPage.SetBorderColor(tcell.ColorLightGray)
	helpPage.SetBorderPadding(1, 1, 3, 3)
	helpPage.SetBackgroundColor(tcell.ColorBlue)
	helpPage.SetWordWrap(true)
	helpPage.SetDynamicColors(true)
	helpPage.SetBorder(true)
	app.SetRoot(helpPage, false)
}

// areYouSureExit displays a modal box asking the user if they REALLY wish to exit.
func areYouSureExit() {
	logPane.Log("Confirming exit")

	exitPage := tview.NewModal().SetText("Do you want to exit the application?")
	exitPage.SetTitle(" Confirm exit ").SetTitleAlign(tview.AlignCenter)
	exitPage.AddButtons([]string{"Cancel", "Exit"})
	exitPage.SetDoneFunc(func(buttonIndex int, buttonLabel string) {
		if buttonLabel == "Exit" {
			app.Stop()
			return
		}
		app.SetRoot(grid, true)
		return
	})
	app.SetRoot(exitPage, false)
}

// saveSectorConfirm asks the user to confirm that they wish to CLEAR and exit
// the sector generator window
func saveSectorConfirm(s sector) {
	logPane.Log("Confirm save sector before exit")

	thisPage := "savesector"

	confirmPage := tview.NewModal().SetText("The sector has not been saved to file. Do you wish to Cancel, Save or Exit (unsaved)?")
	confirmPage.AddButtons([]string{"Cancel", "Save", "Don't save"})
	confirmPage.SetTitle(" Save Sector? ").SetTitleAlign(tview.AlignCenter)
	confirmPage.SetDoneFunc(func(buttonIndex int, buttonLabel string) {
		mainPane.RemovePage(thisPage)
		mainPane.ShowPage(idSectorPageStr)
		if buttonLabel == "Don't save" {
			log.Printf("Hit 'Don't save' in confirm dialog for sector")
			//mainPane.SetBorder(false)
			mainPane.RemovePage(idSectorPageStr)
			mainPageFocus()
			setSaveStatus(true)
			refocus(idMenuPaneStr)
		} else if buttonLabel == "Save" {
			log.Printf("Hit 'Save' in confirm dialog for sector")
			fn := config.DataDir + s.name + ".tab"
			if err := s.toFile(fn); err != nil {
				logPane.Log("Unable to write sector file")
			} else {
				s.saved = true
				logPane.Log("Sector saved to file")
			}
			setSaveStatus(true)
			//mainPane.SetBorder(false)
			mainPane.RemovePage(idSectorPageStr)
			mainPageFocus()
			refocus(idMenuPaneStr)
		}
	})
	mainPane.AddAndSwitchToPage(thisPage, confirmPage, false)
}

// notImplemented displays a dialog box indicating that the current function is yet to be coded.
// It always returns to the menu pane.
func notImplemented() {
	logPane.Log("Running Not Implemented Box")

	thisPage := "notimplemented"

	niPage := tview.NewModal().SetText("This function has not been implemented yet.")
	niPage.AddButtons([]string{"OK"})
	niPage.SetDoneFunc(func(buttonIndex int, buttonLabel string) {
		mainPane.RemovePage(thisPage)
		refocus(idMenuPaneStr)
		return
	})
	niPage.SetTitle(" Not Implemented ").SetTitleAlign(tview.AlignCenter)
	mainPane.AddAndSwitchToPage(thisPage, niPage, false)
	refocus(idMainPaneStr)
}

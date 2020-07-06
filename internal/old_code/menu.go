package main

import (
	"log"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

// menu.go contains code for the display and handling of the menu.
//
// The "Menu" is actually a tview TreeView with the root object hidden.
// This appears like an indented list, with the ability to open and close
// the sub-menus.

var (
	rootNode *tview.TreeNode   // The (hidden) root menu node.
	menuMap  map[string]string // This map contains the text strings and context strings for the menu options.
)

const mainMenuText = "Menu" // This is actually not used for anything other than identification.

// init initialises the menu contents, creating all the sub-menu items and context strings.
func init() {

	log.Printf("Initialising the menu")

	menuMap = make(map[string]string)

	// Root node - simply says "Menu", and regardless, it is invisible!
	rootNode = initNewTreeNode(mainMenuText, "Choose a menu option").SetSelectable(true)

	// This will help constructing the character generation context string.
	cgs := " character generation"

	// Construct the Main menu hierarchy. This is nested, so watch those trailing dots.
	rootNode.
		AddChild(initNewTreeNode("Characters", "Characters menu options").SetSelectable(true).
			AddChild(initNewTreeNode("Classic Trav", "Classic Traveller"+cgs).SetSelectedFunc(notImplemented).SetSelectable(true).
				AddChild(initNewTreeNode("Book 1", "Generate [yellow]Book 1 - Characters & Combat[-] characters").SetSelectedFunc(getCTBook01CharacterBasicData).SetSelectable(true)).
				AddChild(initNewTreeNode("Book 4-8", "Detailed characters from Books 4 to Book 8").SetSelectedFunc(notImplemented).SetSelectable(true)).
				AddChild(initNewTreeNode("Citizens", "Supplement 4 [yellow]Citizens of the Imperium[-] characters").SetSelectedFunc(notImplemented).SetSelectable(true)).
				AddChild(initNewTreeNode("Any", "Generate character using any CT career").SetSelectedFunc(notImplemented).SetSelectable(true))).
			AddChild(initNewTreeNode("MegaTraveller", "MegaTraveller"+cgs).SetSelectedFunc(notImplemented).SetSelectable(true)).
			AddChild(initNewTreeNode("TNE", "Traveller New Era"+cgs).SetSelectedFunc(notImplemented).SetSelectable(true)).
			AddChild(initNewTreeNode("Traveller 4", "Marc Miller's Traveller 4 (T4)"+cgs).SetSelectedFunc(notImplemented).SetSelectable(true)).
			AddChild(initNewTreeNode("Mongoose", "Mongoose Traveller (1st ed)"+cgs).SetSelectedFunc(notImplemented).SetSelectable(true)).
			AddChild(initNewTreeNode("Traveller5", "Traveller5"+cgs).SetSelectedFunc(notImplemented).SetSelectable(true)).
			AddChild(initNewTreeNode("Mongoose2", "Mongoose Traveller 2nd Edition"+cgs).SetSelectedFunc(notImplemented).SetSelectable(true)).
			CollapseAll()).
		AddChild(initNewTreeNode("Worlds", "Worlds menu options").SetSelectable(true).
			AddChild(initNewTreeNode("Basic", "Generate a standard [yellow]Book 3[-] world").SetSelectedFunc(getWorldBasicData).SetSelectable(true)).
			AddChild(initNewTreeNode("Expanded", "Generate a Classic Traveller [yellow]Book 6 Scouts[-] system").SetSelectedFunc(notImplemented).SetSelectable(true)).
			AddChild(initNewTreeNode("MT Basic", "Generate a [yellow]basic MegaTraveller[-] world").SetSelectedFunc(getWorldBasicMTData).SetSelectable(true)).
			AddChild(initNewTreeNode("MT Extended", "Generate an [yellow]extended MegaTraveller[-] system").SetSelectedFunc(notImplemented).SetSelectable(true)).
			AddChild(initNewTreeNode("DGP World Builder", "Generate MegaTraveller [yellow]World Builder's Handbook[-] world details").SetSelectedFunc(getWorldBasicWbhData).SetSelectable(true)).
			AddChild(initNewTreeNode("T5 2nd Survey", "Generate a Traveller5 Second Survey world").SetSelectedFunc(getWorldBasicT5Data).SetSelectable(true)).
			AddChild(initNewTreeNode("T5 System", "Generate a Traveller5 Second Survey system").SetSelectedFunc(notImplemented).SetSelectable(true)).
			AddChild(initNewTreeNode("Foreven", "Populate the Foreven sector").SetSelectedFunc(generateForeven).SetSelectable(true)).
			AddChild(initNewTreeNode("Random Sector", "Generate a new random sector full of [yellow]basic MegaTraveller[-] worlds").SetSelectedFunc(getRandomSector).SetSelectable(true)).
			CollapseAll()).
		AddChild(initNewTreeNode("Language", "Language menu options").SetSelectable(true).
			AddChild(initNewTreeNode("Aslan (Trokh)", "Generate Aslan words").SetSelectedFunc(generateAslanWords).SetSelectable(true)).
			AddChild(initNewTreeNode("Darrian", "Generate Darrian words").SetSelectedFunc(generateDarrianWords).SetSelectable(true)).
			AddChild(initNewTreeNode("Droyne", "Generate Droyne words").SetSelectedFunc(generateDroyneWords).SetSelectable(true)).
			AddChild(initNewTreeNode("K'kree", "Generate K'kree words").SetSelectedFunc(generateKkreeWords).SetSelectable(true)).
			AddChild(initNewTreeNode("Vargr (Gvegh)", "Generate Vargr words").SetSelectedFunc(generateVargrWords).SetSelectable(true)).
			AddChild(initNewTreeNode("Vilani", "Generate Vilani words").SetSelectedFunc(generateVilaniWords).SetSelectable(true)).
			AddChild(initNewTreeNode("Zhodani", "Generate Zhodani words").SetSelectedFunc(generateZhodaniWords).SetSelectable(true)).
			CollapseAll()).
		AddChild(initNewTreeNode("Search", "Search menu options").SetSelectable(true).
			AddChild(initNewTreeNode("Find character", "Search for a character in the database").SetSelectedFunc(notImplemented).SetSelectable(true)).
			AddChild(initNewTreeNode("Find world", "Search for a world in the database").SetSelectedFunc(notImplemented).SetSelectable(true)).
			CollapseAll()).
		AddChild(initNewTreeNode("Database", "Database menu options").SetSelectable(true).
			AddChild(initNewTreeNode("Missing worlds", "Retrieve 'missing' world data from travellermap.com").SetSelectedFunc(missingWorlds).SetSelectable(true)).
			AddChild(initNewTreeNode("Missing subsector", "Retrieve 'missing' subsector data from travellermap.com").SetSelectedFunc(missingSubs).SetSelectable(true)).
			AddChild(initNewTreeNode("Deploy staged", "Copy staged world data to production").SetSelectedFunc(copyStageToProd).SetSelectable(true)).
			AddChild(initNewTreeNode("Delete staged", "Delete staged world data").SetSelectedFunc(deleteStaged).SetSelectable(true)).
			CollapseAll()).
		AddChild(initNewTreeNode("About", "Display info about this application").SetSelectedFunc(doAbout).SetSelectable(true)).
		AddChild(initNewTreeNode("Test", "Test screen/app functionality").SetSelectedFunc(doTest).SetSelectable(true)).
		AddChild(initNewTreeNode("Exit", "Exit the application").SetSelectedFunc(areYouSureExit).SetSelectable(true))

}

// initMenuPane finalises initialisation of the menu Pane. It creates the
// function to handle selection of menu items (setting the context bar text),
// sets the current node of the menu to the top of the tree, and sets the
// function that is called when a menu node is actioned.
func initMenuPane() {

	log.Printf("Initialising the menu pane")
	// Finally the menu pane
	menuPane.SetSelectedFunc(func(n *tview.TreeNode) {
		if n.GetChildren() == nil || n == rootNode {
			return
		}
		if n.IsExpanded() {
			n.Collapse()
			menuStatusText(n)
			return
		}
		n.Expand()
		menuStatusText(n)
	})
	menuPane.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		key := event.Key()
		if key == tcell.KeyTab || key == tcell.KeyBacktab {
			return handleTabs(idMenuPaneStr, event)
		}
		return event
	})
	baseNode := (rootNode.GetChildren())[0]
	log.Printf("Setting root node for menu")
	menuPane.SetRoot(rootNode).SetCurrentNode(baseNode)
	menuStatusText(baseNode)
	log.Printf("menuPane.SetChangedFunc")
	menuPane.SetChangedFunc(menuStatusText)

}

// initNewTreeNode creates a new TreeNode with the specified text and
// sets up the contextBar text for that menu item. It returns the new
// TreeNode created.
func initNewTreeNode(text string, sbText string) (tn *tview.TreeNode) {

	tn = tview.NewTreeNode(text)
	menuMap[text] = sbText
	return
}

// menuStatusText changes the Status Bar (sbContext) text depending on
// which TreeNode is provided to it. It can be called directly by providing
//a TreeNode, or called as a handler function for the TreeView.
func menuStatusText(node *tview.TreeNode) {
	log.Printf("Enter menuStatusText")
	if node == nil {
		log.Printf("Node is nil - cannot set status text.")
		return
	}
	if contextBar == nil {
		panic("Context bar has not been initialised")
	}

	getOpen := "<enter> to open menu"
	getClose := "<enter> to close menu"

	sbText := sbContext + " | "

	if node.GetText() == mainMenuText {
		sbText += menuMap[mainMenuText]
	} else if len(node.GetChildren()) > 0 {
		if node.IsExpanded() {
			sbText += getClose
		} else {
			sbText += getOpen
		}
		sbText += " " + menuMap[node.GetText()]
	} else {
		sbText += menuMap[node.GetText()]
	}
	log.Printf("Setting context bar text to %s", sbText)
	contextBar.SetText(sbText)
	//log.Printf("Why isn't highlighting working??")
	contextBar.Highlight("menu")
	log.Printf("About to return from menuStatusText")
	return
}

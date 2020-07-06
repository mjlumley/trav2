package deprecated

// userMenus.go contains all the usermenus and some interaction.

/*

// menuTopLevel runs the top level menu, asking what option the user would like to do. Returns when the
// user wants to exit.
func menuTopLevel() {

	for {
		callClear()
		fmt.Println("Traveller top level menu")
		fmt.Println()
		fmt.Println("1. Generate a Character")
		fmt.Println("2. Find a world")
		fmt.Println("3. Find a character")
		fmt.Println("8. Generate Foreven sector (T5SS)")
		fmt.Println("x. Exit")

		switch getChoice("\nYour choice : ") {
		case "x", "X":
			fmt.Println("Exiting...")
			return
		case "1":
			menuGenerateCharacter()
		case "2":
			//findAWorld()
		case "3":
			//findACharacter()
		case "8":
			//generateForeven()
		default:
			fmt.Println("Incorrect choice.")
		}
	}
}

// -- Character Generation

// menuGenerateCharacter runs the Character Generation menu, which is one-level down from the top level
func menuGenerateCharacter() {
	for {
		callClear()
		fmt.Println("Character Generation menu")
		fmt.Println()
		fmt.Println("1. Classic Traveller")
		fmt.Println("x. Return to previous level")

		switch getChoice("\nYour Character Generation choice : ") {
		case "x", "X":
			fmt.Println("Returning...")
			return
		case "1":
			generateCTBook01Characters()
			enterToContinue()
		default:
			fmt.Println("Incorrect choice.")
			enterToContinue()
		}

	}
}
*/

/*

// findAWorld searches through the database for a world that matches the users search terms.
func findAWorld() {

	for {
		breakFlag := false
		callClear()
		fmt.Println("Find-a-world menu")
		fmt.Println()
		fmt.Println("This searches for worlds based on search characteristics, and may be useful for")
		fmt.Println("finding a homeworld for your character.")
		fmt.Println()
		fmt.Println("1. Search by name")
		fmt.Println("x. Return to previous level")

		switch getChoice("\nYour choice : ") {
		case "x", "X":
			fmt.Println("Returning...")
			breakFlag = true
		case "1":
			findWorldByName()
			enterToContinue()
		default:
			fmt.Println("Incorrect choice.")
			enterToContinue()
		}
		if breakFlag {
			break
		}
	}
}
*/

/*
// findWorldByName finds a world by its name.
func findWorldByName() {
	callClear()
	fmt.Println("Find a world by name")
	fmt.Println()
	fmt.Println("The search will be done on partial results: 'tur' will return 'Tureded', 'Arcturus' and others.")
	fmt.Println("Do not use wildcards or other punctuation, the only exception being apostrophes.")
	fmt.Println("They will be stripped anyway.")

	worldSearch := getChoice("Enter the world name (leave blank to exit): ")
	if len(worldSearch) < 1 {
		return
	}

	//A slice of world should be returned
	worlds := getWorldsByName(worldSearch)
	if len(worlds) <= 0 {
		fmt.Println("\nNo worlds found!")
	} else {
		fmt.Println("\nFound", len(worlds), "worlds.")
		pageCount := 0
		for i := 0; i < len(worlds); i++ {
			fmt.Println(i, ": ", worlds[i].String())
			pageCount++
			if pageCount == config.languageNumItems {
				switch getChoice("\nPress <enter> to continue, x to exit") {
				case "x", "X":
					fmt.Println("Returning...")
					return
				default:
					pageCount = 0
				}
			}
		}
	}
	return
}
*/

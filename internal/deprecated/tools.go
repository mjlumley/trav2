package deprecated

// tools.go contains tools for use throughout the entire application.

/* // clear provides a map for storing clear funcs (although one is only needed!).
var clear map[string]func() //create a map for storing clear funcs

// init performs initialisation specific to these tools.
func init() {
	//Initialize the clear map - for screen clears
	clear = make(map[string]func())
	clear["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

/////////////////
// Tools
/////////////////

// inputFieldHexLoc validates whether the hex location entered is valid or not.
func inputFieldHexloc(text string, ch rune) bool {
	if !ValidateHexLoc(text) {
		return false
	}
	return true
}

// callClear clears the console window. Unbelievably, this actually works!
func callClear() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

// getChoice gets a users input choice. It returns the string with the users choice, trimmed
// for spaces and carriage returns/line feeds/newlines.
func getChoice(questionString string) (choice string) {

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf(questionString)
	choice, _ = reader.ReadString('\n')

	switch config.os {
	case "windows":
		choice = strings.Replace(choice, "\r\n", "", -1)
	default:
		choice = strings.Replace(choice, "\n", "", -1)
	}

	return
}

// enterToContinue asks for the user to press enter to continue. Actually, anything will do.
func enterToContinue() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Press <Enter> to continue")
	_, _ = reader.ReadString('\n')

}

// getYesNoAnswer gets a yes/no answer. Provide a question in the string, and whether the default is yes (true) or no (false).
// Returns true for yes or false for no.
func getYesNoAnswer(question string, defaultYes bool) bool {
	for {
		if defaultYes {
			question += " [Y/n]?"
		} else {
			question += " [y/N]?"
		}
		r := getChoice(question)
		if r == "" {
			return defaultYes
		}
		r = strings.ToLower(r)
		switch r {
		case "y", "yes":
			return true
		case "n", "no":
			return false
		default:
			fmt.Println("Incorrect choice - please enter 'y' or 'n'")
		}
	}

}
*/

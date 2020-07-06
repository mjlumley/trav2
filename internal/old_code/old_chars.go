package main

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/dustin/go-humanize"
)

// The characters.go source file contains code that allows searching for and refining characters in the database.

var (
	// gcCascadeTable contains the cascade skills for Gun Combat
	gcCascadeTable [10]string

	// bcCascadeTbale contains the cascade skills for Blade Combat
	bcCascadeTable [11]string

	// veCascadeTable contains the cascade skllls for Vehicle
	veCascadeTable [10]string
)

// charct stores information about a Traveller.
type charct struct {

	// Basics
	name string // The character's name
	age  int    // The character's age
	race string // The character's race

	// Characteristics
	strength     int // strength digit
	dexterity    int // dexterity digit
	endurance    int // endurance digit
	intelligence int // intelligence digit
	education    int // education digit
	social       int // social standing digit

	// Basic data
	credits int     // credits indicates amount of money held
	skills  []Skill // skills lists all the skills a character has
	terms   int     // terms indicates how many terms (4 year periods) the character has served
	service string  // the (final) profession that the character served in

	// Other, optional parts
	sex       string   // sex is the characters sex.
	homeworld string   // homeworld indicates the characters homeworld.
	benefits  []string // benefits is a list of benefits the character has received on retirement.
	rank      string   // rank is the final rank of the character's profession.

	// Character record
	history []string // history shows a log of the character generation process.
}

// Constants for service character served in
const (
	// NONE signifies no service
	NONE int = 0

	// NAVY indicates the Navy service
	NAVY int = 1

	// MARINES indicates the Marines service
	MARINES int = 2

	// ARMY indicates the Army service
	ARMY int = 3

	// SCOUTS indicates the Scouts service
	SCOUTS int = 4

	// MERCHANTS indicates Merchant service
	MERCHANTS int = 5

	// OTHER indicates service in another field
	OTHER int = 6
)

// TAS is the string forTravellers' Aid Society benefit
const TAS string = "Travellers' Aid Society"

var service [7]string // Character's Service (0 = None, 1 = Navy, etc)
var rank [7][7]string // Character's rank

// init intialises the structures needed for this module.
func init() {

	service[NONE] = "None"
	service[NAVY] = "Navy"
	service[MARINES] = "Marines"
	service[ARMY] = "Army"
	service[SCOUTS] = "Scouts"
	service[MERCHANTS] = "Merchants"
	service[OTHER] = "Other"

	rank[NAVY][0] = "Spacehand"
	rank[NAVY][1] = "Ensign"
	rank[NAVY][2] = "Lieutenant"
	rank[NAVY][3] = "Lt Cmdr"
	rank[NAVY][4] = "Commander"
	rank[NAVY][5] = "Captain"
	rank[NAVY][6] = "Admiral"

	rank[MARINES][0] = "Marine"
	rank[MARINES][1] = "Lieutenant"
	rank[MARINES][2] = "Captain"
	rank[MARINES][3] = "Force Cmdr"
	rank[MARINES][4] = "Lt Colonel"
	rank[MARINES][5] = "Colonel"
	rank[MARINES][6] = "Brigadier"

	rank[ARMY][0] = "Soldier"
	rank[ARMY][1] = "Lieutenant"
	rank[ARMY][2] = "Captain"
	rank[ARMY][3] = "Major"
	rank[ARMY][4] = "Lt Colonel"
	rank[ARMY][5] = "Colonel"
	rank[ARMY][6] = "General"

	rank[MERCHANTS][0] = "Merchant"
	rank[MERCHANTS][1] = "4th Officer"
	rank[MERCHANTS][2] = "3rd Officer"
	rank[MERCHANTS][3] = "2nd Officer"
	rank[MERCHANTS][4] = "1st Officer"
	rank[MERCHANTS][5] = "Captain"
	rank[MERCHANTS][6] = "Captain"

	gcCascadeTable[0] = "Body Pistol"
	gcCascadeTable[1] = "Auto Pistol"
	gcCascadeTable[2] = "Revolver"
	gcCascadeTable[3] = "Carbine"
	gcCascadeTable[4] = "Rifle"
	gcCascadeTable[5] = "Auto Rifle"
	gcCascadeTable[6] = "Shotgun"
	gcCascadeTable[7] = "SMG"
	gcCascadeTable[8] = "Laser Carbine"
	gcCascadeTable[9] = "Laser Rifle"

	bcCascadeTable[0] = "Dagger"
	bcCascadeTable[1] = "Blade"
	bcCascadeTable[2] = "Foil"
	bcCascadeTable[3] = "Sword"
	bcCascadeTable[4] = "Cutlass"
	bcCascadeTable[5] = "Broadsword"
	bcCascadeTable[6] = "Bayonet"
	bcCascadeTable[7] = "Spear"
	bcCascadeTable[8] = "Halberd"
	bcCascadeTable[9] = "Pike"
	bcCascadeTable[10] = "Cudgel"

	veCascadeTable[0] = "Helicopter"
	veCascadeTable[1] = "Propeller-driven Fixed Wing"
	veCascadeTable[2] = "Jet-driven Fixed Wing"
	veCascadeTable[3] = "Grav Vehicle"
	veCascadeTable[4] = "Tracked Vehicle"
	veCascadeTable[5] = "Wheeled Vehicle"
	veCascadeTable[6] = "Small Watercraft"
	veCascadeTable[7] = "Large Watercraft"
	veCascadeTable[8] = "Hovercraft"
	veCascadeTable[9] = "Submersible"

}

// findACharacter finds a character in the database.
func findACharacter() {
	notImplemented()
	enterToContinue()
}

// introChargenInfo prints some basic character generation information applicable to all generation systems.
func introChargenInfo() {
	fmt.Println("Character generation creates a basic character, following the sequences in the")
	fmt.Println("respective rulebooks. It does not fully flesh out the character, rather it allows")
	fmt.Println("you to go through the process and create the basics. A 'generation history' will")
	fmt.Println("allow you to flesh out the character at a later stage. This basic info can be")
	fmt.Println("stored in the database or printed to screen.")
	fmt.Println()
	fmt.Println("You have the option of generating all rolls yourself or having the program do the")
	fmt.Println("rolls for you.")
	fmt.Println()
	chargenHelp()
}

// chargenHelp provides help for the character generation.
func chargenHelp() {
	fmt.Println("While generating characters, you can use the following commands:")
	fmt.Println(" 0-F - input the choice from the menu (upper- or lower-case for letters)")
	fmt.Println("  <  - move to next page (when choices are many)")
	fmt.Println("  >  - move to previous page (when choices are many)")
	fmt.Println("  !  - display the character sheet as it currently stands")
	fmt.Println("  ?  - display this help information")
	fmt.Println("  x  - exit the generation sequence. You will be given a chance to confirm.")
	fmt.Println()
	fmt.Println("Note that these commands may not always apply.")
	fmt.Println()
}

// generateCTBook01Characters generates characters using the Classic Traveller Book 01 rules.
func generateCTBook01Characters() {
	callClear()
	fmt.Println("Classic Traveller Book 01 Character Generation")
	fmt.Println()
	introChargenInfo()

	// Generate character from here
	if getYesNoAnswer("Do you wish to perform your own dice rolls", false) {
		notImplemented()
		return
	}

	dieIfSurvivalFail := getYesNoAnswer("Do you wish to enforce death on survival fail (yes), or convert to an injury (no) : ", true)

	for {
		generateCTBook01Character(dieIfSurvivalFail)
		if !getYesNoAnswer("Generate another character ", true) {
			break
		}
	}

}

// generateCTBook01Character generates a character using the Classic Traveller Book 01 rules. Set dieIfSurvivalFail to true if you wish to use the
// hard survival fail=death rules or false if instead survival fail=injury and muster out.
func generateCTBook01Character(dieIfSurvivalFail bool) {

	// Personal Development Table. Index is [serviceNumber-1][roll]
	pdTable := [6][6]string{{"+1 Str", "+1 Dex", "+1 End", "+1 Int", "+1 Edu", "+1 Soc"},
		{"+1 Str", "+1 Dex", "+1 End", "Gambling", "Brawling", "Blade Combat"},
		{"+1 Str", "+1 Dex", "+1 End", "Gambling", "+1 Edu", "Brawling"},
		{"+1 Str", "+1 Dex", "+1 End", "+1 Int", "+1 Edu", "Gun Combat"},
		{"+1 Str", "+1 Dex", "+1 End", "+1 Str", "Blade Combat", "Bribery"},
		{"+1 Str", "+1 Dex", "+1 End", "Blade Combat", "Brawling", "-1 Social"}}
	// Service Skills Table. Index is [serviceNumber-1][roll]
	ssTable := [6][6]string{{"Ship's Boat", "Vacc Suit", "Forward Observer", "Gunnery", "Blade Combat", "Gun Combat"},
		{"Vehicle", "Vacc Suit", "Blade Combat", "Gun Combat", "Blade Combat", "Gun Combat"},
		{"Vehicle", "Air/Raft", "Gun Combat", "Forward Observer", "Blade Combat", "Gun Combat"},
		{"Vehicle", "Vacc Suit", "Mechanical", "Navigation", "Electronics", "Jack of All Trades"},
		{"Vehicle", "Vacc Suit", "Jack of All Trades", "Steward", "Electronics", "Gun Combat"},
		{"Vehicle", "Gambling", "Brawling", "Bribery", "Blade Combat", "Gun Combat"}}
	// Advanced Education Table. Index is [serviceNumber-1][roll]
	aeTable := [6][6]string{{"Vacc Suit", "Mechanical", "Electronic", "Engineering", "Gunnery", "Jack of All Trades"},
		{"Vehicle", "Mechanical", "Electronic", "Tactics", "Blade Combat", "Gun Combat"},
		{"Vehicle", "Mechanical", "Electronic", "Tactics", "Blade Combat", "Gun Combat"},
		{"Vehicle", "Mechanical", "Electronic", "Jack of All Trades", "Gunnery", "Medical"},
		{"Streetwise", "Mechanical", "Electronic", "Navigation", "Gunnery", "Medical"},
		{"Streetwise", "Mechanical", "Electronic", "Gambling", "Brawling", "Forgery"}}
	// Higher Education Table. Index is [serviceNumber-1][roll]
	heTable := [6][6]string{{"Medical", "Navigation", "Engineering", "Computer", "Pilot", "Admin"},
		{"Medical", "Tactics", "Tactics", "Computer", "Leader", "Admin"},
		{"Medical", "Tactics", "Tactics", "Computer", "Leader", "Admin"},
		{"Medical", "Navigation", "Engineering", "Computer", "Pilot", "Jack of All Trades"},
		{"Medical", "Navigation", "Engineering", "Computer", "Pilot", "Admin"},
		{"Medical", "Forgery", "Electronics", "Computer", "Streetwise", "Jack of All Trades"}}

	// enlsitment provides the required rolls for enlistment
	enlistment := [7]int{0, 8, 9, 5, 7, 7, 3}
	// survival provides the required rolls for survival
	survival := [7]int{0, 5, 6, 5, 7, 5, 5}
	// commission provides the required rolls for commission
	commission := [7]int{0, 10, 9, 5, 0, 4, 0}
	// promotion provides the required rolls for promotion
	promotion := [7]int{0, 8, 9, 6, 0, 10, 0}
	// reenlist provides the required rolls for re-enlistment
	reenlist := [7]int{0, 6, 6, 7, 3, 4, 5}

	// Basic stuff
	c := charct{name: "John Doe", race: "Human", sex: "Male", age: 18}
	c.strength = D6() + D6()
	c.dexterity = D6() + D6()
	c.endurance = D6() + D6()
	c.intelligence = D6() + D6()
	c.education = D6() + D6()
	c.social = D6() + D6()
	rankInt := 0

	serviceNumber := NONE // This will be 0=Draftee, 1=Navy, etc, 6=Other
	isDraftee := true     // Needed for determining 1st term commission

	// Get Service choice
	for {
		breakFlag := false
		callClear()
		fmt.Println()
		fmt.Printf("%s\n", c)
		fmt.Println()
		fmt.Println("Enlistment")
		fmt.Println("----------")
		fmt.Println("0. Volunteer for draft (no commission during first term)")
		fmt.Println("1. Navy      8+ (DM+1:Int 8+, DM+2:Edu 9+)")
		fmt.Println("2. Marines   9+ (DM+1:Int 8+, DM+2:Str 8+)")
		fmt.Println("3. Army      5+ (DM+1:Dex 6+, DM+2:End 5+)")
		fmt.Println("4. Scouts    7+ (DM+1:Int 6+, DM+2:Str 8+)")
		fmt.Println("5. Merchants 7+ (DM+1:Str 7+, DM+2:Int 6+)")
		fmt.Println("6. Other     3+")
		fmt.Println()
		fmt.Println("N. Edit the character's name")
		fmt.Println("S. Edit the character's sex")
		fmt.Println("R. Edit the character's race")
		fmt.Println("!. Display the character sheet")
		fmt.Println("?. Display help")
		fmt.Println("x. Exit character generation")
		fmt.Println()

		roll := D6() + D6()
		switch getChoice("Select a Service or option : ") {
		case "0":
			breakFlag = true
		case "1":
			// Navy
			if c.intelligence >= 8 {
				roll++
			}
			if c.education >= 9 {
				roll += 2
			}
			if roll >= enlistment[NAVY] {
				serviceNumber = NAVY
			}
			breakFlag = true
		case "2":
			// Marines
			if c.intelligence >= 8 {
				roll++
			}
			if c.strength >= 8 {
				roll += 2
			}
			if roll >= enlistment[MARINES] {
				serviceNumber = MARINES
			}
			breakFlag = true
		case "3":
			// Army
			if c.dexterity >= 6 {
				roll++
			}
			if c.endurance >= 5 {
				roll += 2
			}
			if roll >= enlistment[ARMY] {
				serviceNumber = ARMY
			}
			breakFlag = true
		case "4":
			// Scouts
			if c.intelligence >= 6 {
				roll++
			}
			if c.strength >= 8 {
				roll += 2
			}
			if roll >= enlistment[SCOUTS] {
				serviceNumber = SCOUTS
			}
			breakFlag = true
		case "5":
			// Merchants
			if c.strength >= 7 {
				roll++
			}
			if c.intelligence >= 6 {
				roll += 2
			}
			if roll >= enlistment[MERCHANTS] {
				serviceNumber = MERCHANTS
			}
			breakFlag = true
		case "6":
			// Other
			if roll >= enlistment[OTHER] {
				serviceNumber = OTHER
			}
			breakFlag = true
		case "x", "X":
			fmt.Println("Returning...")
			return // NONE signifies no service

		case "!":
			fmt.Println(c.String())
		case "?":
			chargenHelp()
		case "N":
			c.name = getChoice("Character's name : ")
		case "S":
			c.sex = getChoice("Character's sex : ")
		case "R":
			c.race = getChoice("Character's race : ")
		default:
			breakFlag = false
		}
		if breakFlag {
			break
		}
	}

	// Deal with the draft
	if serviceNumber > NONE {
		isDraftee = false
		fmt.Println("Enlistment has succeeded in the " + service[serviceNumber])
	} else {
		serviceNumber = D6()
		isDraftee = true
		fmt.Println("You have been drafted into the " + service[serviceNumber])
	}
	c.service = service[serviceNumber]
	c.rank = rank[serviceNumber][0]

	// Determines if the character dies/injures during a term.
	doesSurvive := false

	// The character is in a service and must go through each term from here
	currentTerm := 0
	for {
		currentTerm++
		fmt.Printf("%s term in the %s\n", humanize.Ordinal(currentTerm), service[serviceNumber])
		fmt.Println("------------------------------")

		// Check survival
		roll := D6() + D6()
		doesSurvive = false
		dm := 0
		switch serviceNumber {
		case NAVY:
			if c.intelligence >= 7 {
				dm = 2
			}
		case MARINES:
			if c.endurance >= 8 {
				dm = 2
			}
		case ARMY:
			if c.education >= 6 {
				dm = 2
			}
		case SCOUTS:
			if c.endurance >= 9 {
				dm = 2
			}
		case MERCHANTS:
			if c.intelligence >= 7 {
				dm = 2
			}
		case OTHER:
			if c.intelligence >= 9 {
				dm = 2
			}
		default: // Increase age
			fmt.Println("Unknown service number", serviceNumber)
			panic("Exiting")
		}
		if (roll + dm) >= survival[serviceNumber] {
			doesSurvive = true
			fmt.Println("You have survived this term.")
		} else {
			c.ageCharacter(2)
			doesSurvive = false
			if dieIfSurvivalFail {
				fmt.Println("You have been killed in service.")
				fmt.Println(c.String())
				return
			}
			fmt.Println("You have been injured in service.")
			break
		}

		numSkills := 1 // Number of skills to obtain this term
		if currentTerm == 1 {
			numSkills++
		}

		// Some rank and service skills
		if currentTerm == 1 {
			switch serviceNumber {
			case MARINES:
				fmt.Println("Marines automatically receive Cutlass-1")
				c.addSkill("Cutlass", true)
			case ARMY:
				fmt.Println("Soldiers automatically receive Rifle-1")
				c.addSkill("Rifle", true)
			case SCOUTS:
				fmt.Println("Scouts automatically receive Pilot-1")
				c.addSkill("Pilot", true)
			}
		}

		// Check for commission
		if rankInt == 0 && !(isDraftee && currentTerm == 1) && serviceNumber != NONE && serviceNumber != SCOUTS && serviceNumber != OTHER {
			dm = 0
			roll = D6() + D6()
			switch serviceNumber {
			case NAVY:
				if c.social >= 9 {
					dm = 2
				}
			case MARINES:
				if c.education >= 7 {
					dm = 2
				}
			case ARMY:
				if c.endurance >= 7 {
					dm = 2
				}
			case MERCHANTS:
				if c.intelligence >= 6 {
					dm = 2
				}
			}
			if (roll + dm) > commission[serviceNumber] {
				rankInt = 1
				c.rank = rank[serviceNumber][rankInt]
				fmt.Println("Received a commission. New rank is " + c.rank)
				numSkills++

				// Check for rank and service skills
				switch serviceNumber {
				case MARINES:
					// Marine lieutenants receive Revolver-1
					fmt.Println("Marine lieutenants automatically receive Revolver-1")
					c.addSkill("Revolver", true)
				case ARMY:
					// Army lieutenants receive SMG-1
					fmt.Println("Army lieutenants automatically receive SMG-1")
					c.addSkill("SMG", true)
				}
			}
		}

		// Check for promotion
		if rankInt > 0 {
			gotPromotion := false
			roll = D6() + D6()
			dm = 0
			switch serviceNumber {
			case NAVY:
				if c.education >= 8 {
					dm = 1
				}
				if (roll+dm) >= promotion[NAVY] && rankInt < 6 {
					gotPromotion = true
				}
			case MARINES:
				if c.social >= 8 {
					dm = 1
				}
				if (roll+dm) >= promotion[MARINES] && rankInt < 6 {
					gotPromotion = true
				}
			case ARMY:
				if c.education >= 7 {
					dm = 1
				}
				if (roll+dm) >= promotion[ARMY] && rankInt < 6 {
					gotPromotion = true
				}
			case MERCHANTS:
				if c.intelligence >= 9 {
					dm = 1
				}
				if (roll+dm) >= promotion[MERCHANTS] && rankInt < 5 {
					gotPromotion = true
				}
			}

			// Check old rank and new rank here. Might have received a promotion but no further progress can be made.
			if gotPromotion {
				rankInt++
				c.rank = rank[serviceNumber][rankInt]
				fmt.Println("Received a promotion. New rank is " + c.rank)
				numSkills++

				// Determine if there are any rank and service skills
				switch serviceNumber {
				case NAVY:
					// Navy Captains & Admirals get +1 social
					switch rankInt {
					case 5, 6:
						c.social++
						fmt.Println("Navy " + rank[serviceNumber][rankInt] + "s receive +1 Social standing")
						c.attributeChange("Soc", 1)
					}
				case MERCHANTS:
					// Merchant 1st Officer gets Pilot-1
					if rankInt == 4 {
						fmt.Println("Merchant 1st Officers automatically receive Pilot-1")
						c.addSkill("Pilot", true)
					}
				}
			}

		}

		// Get skills for this term
		fmt.Printf("You are eligible for %d skills this term.\n", numSkills)
		enterToContinue()

		for i := 0; i < numSkills; i++ {
			breakFlag := false
			for {
				callClear()
				fmt.Printf("Skill selection in term %d for %s", currentTerm, c.name)
				fmt.Println()
				fmt.Println()
				fmt.Println(c.String())
				fmt.Println()
				fmt.Println("0. Randomly select a table")
				fmt.Println("1. Personal Development table (Hit 'A' to show table)")
				fmt.Println("2. Service Skills table       (Hit 'B' to show table)")
				fmt.Println("3. Advanced Education table   (Hit 'C' to show table)")
				if c.education >= 8 {
					fmt.Println("4. Higher Education table     (Hit 'D' to show table)")
				}
				fmt.Println("N. Edit the character's name")
				fmt.Println("S. Edit the character's sex")
				fmt.Println("R. Edit the character's race")
				fmt.Println("!. Display the character sheet")
				fmt.Println("?. Display help")
				fmt.Println()

				choice := getChoice("Select a table or option : ")

				// If random choice then select a table
				if choice == "0" {
					if c.education >= 8 {
						roll = Dice(4)
					} else {
						roll = Dice(3)
					}
					choice = fmt.Sprintf("%d", roll)
				}
				roll = D6()
				newSkill := ""
				switch choice {
				case "1":
					newSkill = pdTable[serviceNumber-1][roll-1]
					breakFlag = true
				case "2":
					newSkill = ssTable[serviceNumber-1][roll-1]
					breakFlag = true
				case "3":
					newSkill = aeTable[serviceNumber-1][roll-1]
					breakFlag = true
				case "4":
					if c.education < 8 {
						fmt.Println("Unknown option")
					} else {
						newSkill = heTable[serviceNumber-1][roll-1]
						breakFlag = true
					}
				case "A":
					fmt.Println()
					fmt.Println("Personal Development table")
					fmt.Println()
					for i, s := range pdTable[serviceNumber-1] {
						fmt.Printf("%d.  %s\n", i+1, s)
					}
					fmt.Println()
					enterToContinue()
				case "B":
					fmt.Println()
					fmt.Println("Service Skills table")
					fmt.Println()
					for i, s := range ssTable[serviceNumber-1] {
						fmt.Printf("%d.  %s\n", i+1, s)
					}
					fmt.Println()
					enterToContinue()
				case "C":
					fmt.Println()
					fmt.Println("Advanced Education table")
					fmt.Println()
					for i, s := range aeTable[serviceNumber-1] {
						fmt.Printf("%d.  %s\n", i+1, s)
					}
					fmt.Println()
					enterToContinue()
				case "D":
					if c.education < 8 {
						fmt.Println("Unknown option")
					} else {
						fmt.Println()
						fmt.Println("Higher Education table")
						fmt.Println()
						for i, s := range heTable[serviceNumber-1] {
							fmt.Printf("%d.  %s\n", i+1, s)
						}
						fmt.Println()
						enterToContinue()
					}
				case "!":
					fmt.Println(c.String())
				case "?":
					chargenHelp()
					enterToContinue()
				case "N":
					c.name = getChoice("Character's name : ")
				case "S":
					c.sex = getChoice("Character's sex : ")
				case "R":
					c.race = getChoice("Character's race : ")
				default:
					fmt.Println("Unknown option")
					breakFlag = false
				}

				// We have rolled for a skill
				if breakFlag {
					c.addSkill(newSkill, false)
					enterToContinue()
					break
				}
			}
		}

		// Increase age
		if c.ageCharacter(4) == -1 {
			// Character has dies as a result of aging! Immediately exit
			fmt.Println("Character has died as a result of aging")
			fmt.Println(c.String())
			enterToContinue()
			return
		}

		fmt.Println("Term of service completed!")
		c.terms = currentTerm
		fmt.Println(c.String())

		// Check whether the player wants the character to re-enlist
		wantReenlist := getYesNoAnswer("Do you wish to re-enlist", true)
		roll = D6() + D6()

		if roll == 12 {
			// Character is forced to reenlist
			fmt.Println("The " + service[serviceNumber] + " requires to remain in service.")

		} else if !wantReenlist {
			// They do not want to re-enlist. They should muster-out/retire
			fmt.Println("You have elected not to re-enlist.")
			break

		} else if roll >= reenlist[serviceNumber] {
			// They have made their reenlistment roll and want to stay
			if c.terms >= 7 {
				// Normally not allowed to re-enlist past 7 terms
				fmt.Printf("You have completed %d terms. You are required to retire.\n", c.terms)
			}
			fmt.Println("You have suceeded in re-enlisting.")
		} else if roll < reenlist[serviceNumber] {
			// They failed their reenlistment roll but want to stay
			fmt.Println("You have failed to re-enlistment.")
			break
		}

	}

	// Mustering out benefits
	c.musterOut(serviceNumber, rankInt)

	// Retirement pay
	if doesSurvive && c.terms >= 5 && serviceNumber != 4 && serviceNumber != 6 {
		retirementPay := (c.terms - 3) * 2000
		rpString := humanize.Comma(int64(retirementPay))
		rp := "Retirement pay: Cr" + rpString + " per annum"
		c.credits += retirementPay
		c.benefits = append(c.benefits, rp)
		fmt.Println("You have received of retirement payment of Cr" + rpString + ", and will receive annual payments of this amount.")
	}

	fmt.Println("-------")
	fmt.Println(c.String())
	fmt.Println("-------")

}

// musterOut gets the characters mustering out benefits, improvements and skills. It requires the serviceNumber (1-6) and rank (0-6).
func (c *charct) musterOut(serviceNumber int, rankInt int) {

	if c.terms == 0 {
		fmt.Println("No complete terms. No mustering out benefits.")
		return
	}

	var benefitsDM int // Dice Modifier for the Benefits table
	var cashDM int     // Dice Modifier for the Cash table
	var cashRolls int  // Number of cash rolls to do (max 3)

	// Benefits table is [service-1][dice roll+benefitsDM]
	benefitsTable := [6][7]string{{"Low Passage", "+1 Int", "+2 Edu", "Blade", "TAS", "High Passage", "+2 Soc"}, {"Low Passage", "+2 Int", "+1 Edu", "Blade", "TAS", "High Passage", "+2 Soc"},
		{"Low Passage", "+1 Int", "+2 Edu", "Gun", "High Passage", "Mid Passage", "+1 Soc"}, {"Low Passage", "+2 Int", "+2 Edu", "Blade", "Gun", "Scout Ship", ""},
		{"Low Passage", "+1 Int", "+1 Edu", "Gun", "Blade", "Low Passage", "Free Trader"}, {"Low Passage", "+1 Int", "+1 Edu", "Gun", "High Passage", "", ""}}

	// Cash Table is [service-1][dice roll+cashDM]
	cashTable := [6][7]int{{1000, 5000, 5000, 10000, 20000, 50000, 50000}, {2000, 5000, 5000, 10000, 20000, 30000, 40000},
		{2000, 5000, 10000, 10000, 10000, 20000, 30000}, {20000, 20000, 30000, 30000, 50000, 50000, 50000},
		{1000, 5000, 10000, 20000, 20000, 40000, 40000}, {1000, 5000, 10000, 10000, 10000, 50000, 100000}}

	// Determine number of benefits that the character receives
	num := c.terms

	// More because of rank
	switch rankInt {
	case 1, 2:
		num++
	case 3, 4:
		num += 2
	case 5, 6:
		num += 3
		benefitsDM = 1
	}
	if FindSkill(c.skills, "Gambling") >= 0 {
		cashDM++
	}
	cashRolls = 3
	if cashRolls > num {
		cashRolls = num
	}

	// Roll each benefit here
	for ben := 0; ben < num; ben++ {
		fmt.Println()
		fmt.Println(c.String())
		fmt.Println()
		fmt.Printf("Your character has %d benefits. A maximum of %d of these may be taken from the Cash table.\n\n", num-ben, cashRolls)

		// Display the tables
		fmt.Println("Benefits table")
		fmt.Println("--------------")
		for i := 0; i < 7; i++ {
			benefit := benefitsTable[serviceNumber-1][i]
			fmt.Printf("%d. %s\n", i+1, benefit)
		}
		fmt.Println("----------")
		useCashTable := false
		if cashRolls > 0 {
			fmt.Println("Cash table")
			fmt.Println("----------")
			for i := 0; i < 7; i++ {
				cash := cashTable[serviceNumber-1][i]
				fmt.Printf("%d. Cr%s\n", i+1, humanize.Comma(int64(cash)))
			}
			fmt.Println("----------")
			useCashTable = getYesNoAnswer("Do you wish to roll on the Cash table", true)
		}

		roll := D6() - 1
		if useCashTable {
			roll += cashDM
			credits := cashTable[serviceNumber-1][roll]
			cashRolls--
			c.credits = c.credits + credits
			fmt.Printf("You received Cr%s.\n", humanize.Comma(int64(credits)))
		} else {
			roll += benefitsDM
			benefit := benefitsTable[serviceNumber-1][roll]
			fmt.Printf("You received : %s\n", benefit)

			// Now allocate the benefits
			if benefit == "" {
				// No benefit
				continue
			} else if strings.Contains(benefit, "+") {
				c.addSkill(benefit, false)
				continue

			} else if benefit == "Blade" {
				// Cascade blade weapon or benefit
				for {
					fmt.Println("Choose a specific blade weapon")
					for x := 0; x < len(bcCascadeTable); x++ {
						fmt.Printf("%s. %s\n", Ehex(x+1).String(), bcCascadeTable[x])
					}
					val := int(EhexVal(getChoice("Your choice (number/letter) : ")))
					if val > 0 && val <= len(bcCascadeTable) {
						benefit = bcCascadeTable[val-1]
						break
					} else {
						fmt.Println("Invalid choice")
						fmt.Println()
					}
				}
				// We have the specific weapon
				if findBenefit(c.benefits, benefit) >= 0 && getYesNoAnswer("You already have this weapon. Would you like to increase your skill instead", true) {
					c.addSkill(benefit, false)
					continue
				}
				c.benefits = append(c.benefits, benefit)
				continue
			} else if benefit == "Gun" {
				// Cascade gun weapon or benefit
				for {
					fmt.Println("Choose a specific gun weapon")
					for x := 0; x < len(gcCascadeTable); x++ {
						fmt.Printf("%s. %s\n", Ehex(x+1).String(), gcCascadeTable[x])
					}
					val := int(EhexVal(getChoice("Your choice (number/letter) : ")))
					if val > 0 && val <= len(gcCascadeTable) {
						benefit = gcCascadeTable[val-1]
						break
					} else {
						fmt.Println("Invalid choice")
						fmt.Println()
					}
				}
				// We have the specific weapon
				if findBenefit(c.benefits, benefit) >= 0 && getYesNoAnswer("You already have this weapon. Would you like to increase your skill instead", true) {
					c.addSkill(benefit, false)
					continue
				}
				c.benefits = append(c.benefits, benefit)
				continue
			} else if benefit == "TAS" {
				// Check if already have, as can only have one
				if findBenefit(c.benefits, TAS) == -1 {
					c.benefits = append(c.benefits, TAS)
					continue
				}
				fmt.Println("You have already received membership of the " + TAS)
				ben--
				continue
			} else if benefit == "Scout Ship" {
				// Check if already have, as can only have one
				if findBenefit(c.benefits, benefit) == -1 {
					c.benefits = append(c.benefits, benefit)
					continue
				}
				fmt.Println("You have already received a Scout Ship")
				ben--
				continue
			} else if benefit == "Free Trader" {
				// Check if already have this
				if findBenefit(c.benefits, benefit) == -1 {
					c.benefits = append(c.benefits, benefit)
					continue
				}
				c.benefits = append(c.benefits, "Free Trader:+10 years payments")
				continue
			} else {
				c.benefits = append(c.benefits, benefit)
			}
		}
	}

}

// findBenefit looks through a slice of benefits provided in bb for a benefit matching the search string
// and returns the index if found, or -1 if not.
func findBenefit(bb []string, search string) int {

	foundLoc := -1
	for i, benefit := range bb {
		if benefit == search {
			foundLoc = i
			break
		}
	}

	return foundLoc
}

// showUPP returns the Universal Personality Profile for a character.
func (c charct) showUPP() string {

	return Ehex(c.strength).String() + Ehex(c.dexterity).String() + Ehex(c.endurance).String() +
		Ehex(c.intelligence).String() + Ehex(c.education).String() + Ehex(c.social).String()

}

// String returns the character as a string.
func (c charct) String() string {
	var s string
	var num int

	raceString := ""
	if c.race != "" {
		if c.sex != "" {
			raceString = c.race + " " + c.sex
		} else {
			raceString = c.race
		}
	}

	s = fmt.Sprintf(c.name+" ("+raceString+") - "+c.service+" "+c.rank+"  "+c.showUPP()+"  Terms:%d  Age:%d  Cr%s", c.terms, c.age, humanize.Comma(int64(c.credits)))
	num = len(c.skills)
	if num > 0 {
		s += "\n"
		for i, skill := range c.skills {
			s += skill.String()
			if i < (num - 1) {
				s += ", "
			}
		}
	}
	num = len(c.benefits)
	if num > 0 {
		s += "\n"
		for i, ben := range c.benefits {
			s += ben
			if i < (num - 1) {
				s += ", "
			}
		}
	}
	return s
}

// ObjectString outputs the character as a string suitable for display on a TextView.
func (c charct) ObjectString() (s string) {
	var num int

	s = "[yellow]Name:[-] " + c.name + "\n"
	s = "[yellow]UPP:[-] " + c.showUPP() + "\n\n"
	s += "[yellow]Race:[-] " + c.race + "\n"
	s += "[yellow]Sex:[-] " + c.sex + "\n"
	s += "[yellow]Service:[-] " + c.service + "\n"
	s += "[yellow]Rank:[-] " + c.rank + "\n"
	s += fmt.Sprintf("[yellow]Terms:[-] %v\n", c.terms)
	s += fmt.Sprintf("[yellow]Age:[-] %v\n", c.age)
	s += fmt.Sprintf("[yellow]Cr:[-] %v\n", humanize.Comma(int64(c.credits)))

	num = len(c.skills)
	s += "[yellow]Skills :-[-]\n"
	if num > 0 {
		for _, skill := range c.skills {
			s += "  " + skill.String() + "\n"
		}
	} else {
		s += "  none\n"

	}
	num = len(c.benefits)
	s += "[yellow]Benefits :-[-]\n"
	if num > 0 {
		s += "\n"
		for _, ben := range c.benefits {
			s += "  " + ben + "\n"
		}
	} else {
		s += "  none\n"

	}
	return
}

// attributeChange handles an attribute change to a character, whether up or down. The attribute to change is indicated as a 3-character
// string (eg Str, Int, Soc, etc), and the amount of change in the integer. The new value is returned or -1 if the character has died in an aging crisis.
// Keeping within attribute bounds (0 - 15) and handling aging crisises is all done.
func (c *charct) attributeChange(a string, v int) int {
	switch a {
	case "Str":
		c.strength += v
		if c.strength <= 0 {
			if c.agingCrisis() == -1 {
				return -1
			}
			c.strength = 1
			fmt.Println("Strength set to 1 due to aging crisis.")
			return 1
		} else if c.strength > 15 {
			fmt.Println("Strength is limited to 15")
			c.strength = 15
			return 15
		}
		fmt.Println("New strength is", c.strength)
		return c.strength
	case "Dex":
		c.dexterity += v
		if c.dexterity <= 0 {
			if c.agingCrisis() == -1 {
				return -1
			}
			c.dexterity = 1
			fmt.Println("Dexterity set to 1 due to aging crisis.")
			return 1
		} else if c.dexterity > 15 {
			fmt.Println("Dexterity is limited to 15")
			c.dexterity = 15
			return 15
		}
		fmt.Println("New dexterity is", c.dexterity)
		return c.dexterity
	case "End":
		c.endurance += v
		if c.endurance <= 0 {
			if c.agingCrisis() == -1 {
				return -1
			}
			c.endurance = 1
			fmt.Println("Endurance set to 1 due to aging crisis.")
			return 1
		} else if c.endurance > 15 {
			fmt.Println("Endurance is limited to 15")
			c.endurance = 15
			return 15
		}
		fmt.Println("New endurance is", c.endurance)
		return c.endurance
	case "Int":
		c.intelligence += v
		if c.intelligence <= 0 {
			if c.agingCrisis() == -1 {
				return -1
			}
			c.intelligence = 1
			fmt.Println("Intelligence set to 1 due to aging crisis.")
			return 1
		} else if c.intelligence > 15 {
			fmt.Println("Intelligence is limited to 15")
			c.intelligence = 15
			return 15
		}
		fmt.Println("New intelligence is", c.intelligence)
		return c.intelligence
	case "Edu":
		c.education += v
		if c.education <= 0 {
			if c.agingCrisis() == -1 {
				return -1
			}
			c.education = 1
			fmt.Println("Education set to 1 due to aging crisis.")
			return 1
		} else if c.education > 15 {
			fmt.Println("Education is limited to 15")
			c.education = 15
			return 15
		}
		fmt.Println("New education is", c.education)
		return c.education
	case "Soc":
		c.social += v
		if c.social <= 0 {
			if c.agingCrisis() == -1 {
				return -1
			}
			c.social = 1
			fmt.Println("Social standing set to 1 due to aging crisis.")
			return 1
		} else if c.social > 15 {
			fmt.Println("Social standing is limited to 15")
			c.social = 15
			return 15
		}
		fmt.Println("New social standing is", c.social)
		return c.social
	default:
		panic("Unknown characteristic adjustment for " + a + ".")
	}
}

// ageCharacter ages the character by y years. All rolls are taken. If the new age takes a character over a certain age threshhold then
// saving throws are made for aging effects. The function returns -1 if the character dies as a result of an aging crisis, otherwise the
// new age is returned.
func (c *charct) ageCharacter(y int) int {

	isDead := false

	if y <= 0 {
		return y
	}

	c.age += y
	fmt.Printf("Age increased by %d years.\n", y)

	if c.age >= 34 && c.age%4 == 2 {

		// Do some aging
		if c.age < 50 {
			// Strength
			if D6()+D6() < 8 {
				fmt.Println("Strength reduced by aging.")
				if c.attributeChange("Str", -1) == -1 {
					isDead = true
				}
			}
			// Dexterity
			if D6()+D6() < 7 {
				fmt.Println("Dexterity reduced by aging.")
				c.attributeChange("Dex", -1)
			}
			// Endurance
			if D6()+D6() < 8 {
				fmt.Println("Endurance reduced by aging.")
				c.attributeChange("End", -1)
			}
		} else if c.age < 66 {
			// Strength
			if D6()+D6() < 9 {
				fmt.Println("Strength reduced by aging.")
				if c.attributeChange("Str", -1) == -1 {
					isDead = true
				}
			}
			// Dexterity
			if D6()+D6() < 8 {
				fmt.Println("Dexterity reduced by aging.")
				c.attributeChange("Dex", -1)
			}
			// Endurance
			if D6()+D6() < 9 {
				fmt.Println("Endurance reduced by aging.")
				c.attributeChange("End", -1)
			}

		} else {
			// Strength
			if D6()+D6() < 9 {
				fmt.Println("Strength reduced by aging.")
				if c.attributeChange("Str", -1) == -2 {
					isDead = true
				}
			}
			// Dexterity
			if D6()+D6() < 9 {
				fmt.Println("Dexterity reduced by aging.")
				c.attributeChange("Dex", -2)
			}
			// Endurance
			if D6()+D6() < 9 {
				fmt.Println("Endurance reduced by aging.")
				c.attributeChange("End", -2)
			}
			// Intelligence
			if D6()+D6() < 9 {
				fmt.Println("Intelligence reduced by aging.")
				c.attributeChange("End", -1)
			}

		}
	}
	if isDead {
		return -1
	}
	return c.age
}

// agingCrisis handles an aging crisis in a character. It returns -1 on character death, or the number of months the character has aged with slow drug.
func (c charct) agingCrisis() int {
	fmt.Println("Character is undergoing an aging crisis")
	medicLevel := 0
	for {
		val, err := strconv.Atoi(getChoice("Input the Skill Level of any medic in attendance : "))
		if err == nil {
			medicLevel = val
			break
		} else {
			fmt.Println(err)
		}
	}
	roll := D6() + D6() + medicLevel
	if roll >= 8 {
		// Survives
		months := D6()
		fmt.Println("Character has survived aging crisis with", months, "slow drug aging.")
		return months

	}
	// Dies
	fmt.Println("Character has died in aging crisis!")
	return -13
}

// addSkill adds a skill or attribute change to a character. s is the string of the skill to be added to the character.
// If the skill is already present, it will be incremented. If skill is in the form "+1 Str" (eg) then an attribute
// will be updated. If atLevelOne is true, the skill will be added at level 1 if they do not already have the skill.
func (c *charct) addSkill(s string, atLevelOne bool) {

	if strings.ContainsAny(s, "+-") {
		// Attribute change
		parts := strings.Fields(s)
		amt, err := strconv.Atoi(parts[0])
		if err != nil {
			panic("Unable to convert characteristic adjustment from skill table for " + s)
		}
		c.attributeChange(parts[1], amt)
		return
	}
	// Handle cascade skills
	var cascade []string
	switch s {
	case "Gun Combat":
		for _, item := range gcCascadeTable {
			cascade = append(cascade, item)
		}
	case "Blade Combat":
		for _, item := range bcCascadeTable {
			cascade = append(cascade, item)
		}
	case "Vehicle":
		for _, item := range veCascadeTable {
			cascade = append(cascade, item)
		}
	}
	if cascade != nil {
		for {
			fmt.Println()
			fmt.Println("Choose a specific skill for " + s)
			fmt.Println()
			for x, item := range cascade {
				fmt.Printf("%s. %s\n", Ehex(x+1).String(), item)
			}
			val := int(EhexVal(getChoice("Your specific skill (choose number/letter) : ")))
			if val > 0 && val <= len(cascade) {
				s = cascade[val-1]
				break
			} else {
				fmt.Println("Invalid choice")
				fmt.Println()
			}
		}
	}

	// Assign new skill, then continue on
	i := FindSkill(c.skills, s)
	if i > -1 {
		// Existing skill we need to increment
		if !atLevelOne {
			c.skills[i].Level++
			fmt.Printf("New skill level is %s-%d\n", c.skills[i].Name, c.skills[i].Level)
		}
		return
	}
	// New skill we need to add
	c.skills = append(c.skills, Skill{Name: s, Level: 1})
	fmt.Printf("Added new skill %s-1\n", s)
	return
}

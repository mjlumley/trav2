package main

import (
	"fmt"
	"log"

	"github.com/dustin/go-humanize"
)

// classicBook1.go contains a code for character generation specific to
// ClassicTraveller Book 1 (and possibly related types)

// CareerCT01 represents the "service" the character enters.
type CareerCT01 int

// CT01Char stores information about a Classic Traveller Book 01 character.
type CT01Char struct {

	// Basics
	name string // The character's name
	age  int    // The character's age
	race string // The character's race

	// Characteristics
	strength     Ehex // strength digit
	dexterity    Ehex // dexterity digit
	endurance    Ehex // endurance digit
	intelligence Ehex // intelligence digit
	education    Ehex // education digit
	social       Ehex // social standing digit

	// Basic data
	credits int        // credits indicates amount of money held
	skills  []Skill    // skills lists all the skills a character has
	terms   int        // terms indicates how many terms (4 year periods) the character has served
	service CareerCT01 // the (final) profession that the character served in

	// Other, optional parts
	sex       string   // sex is the characters sex.
	homeworld string   // homeworld indicates the characters homeworld.
	benefits  []string // benefits is a list of benefits the character has received on retirement.
	rank      string   // rank is the final rank of the character's profession.

	// Character record
	history []string // history shows a log of the character generation process.
}

// Constants for the Career service the characters serves in.
const (
	// NONE signifies no service
	None CareerCT01 = iota

	// NAVY indicates the Navy service - CT Book 1
	Navy

	// MARINES indicates the Marines service - CT Book 1
	Marines

	// ARMY indicates the Army service - CT Book 1
	Army

	// SCOUTS indicates the Scouts service - CT Book 1
	Scouts

	// MERCHANTS indicates Merchant service - CT Book 1
	Merchants

	// OTHER indicates service in another field - CT Book 1
	Other
)

var serviceStr [7]string // Character's Service (0 = None, 1 = Navy, etc)
var iRank [7][7]string   // Character's iRank

// init intialises the structures needed for this module.
func init() {

	serviceStr[None] = "None"
	serviceStr[Navy] = "Navy"
	serviceStr[Marines] = "Marines"
	serviceStr[Army] = "Army"
	serviceStr[Scouts] = "Scouts"
	serviceStr[Merchants] = "Merchants"
	serviceStr[Other] = "Other"

	iRank[Navy][0] = "Spacehand"
	iRank[Navy][1] = "Ensign"
	iRank[Navy][2] = "Lieutenant"
	iRank[Navy][3] = "Lt Cmdr"
	iRank[Navy][4] = "Commander"
	iRank[Navy][5] = "Captain"
	iRank[Navy][6] = "Admiral"

	iRank[Marines][0] = "Marine"
	iRank[Marines][1] = "Lieutenant"
	iRank[Marines][2] = "Captain"
	iRank[Marines][3] = "Force Cmdr"
	iRank[Marines][4] = "Lt Colonel"
	iRank[Marines][5] = "Colonel"
	iRank[Marines][6] = "Brigadier"

	iRank[Army][0] = "Soldier"
	iRank[Army][1] = "Lieutenant"
	iRank[Army][2] = "Captain"
	iRank[Army][3] = "Major"
	iRank[Army][4] = "Lt Colonel"
	iRank[Army][5] = "Colonel"
	iRank[Army][6] = "General"

	iRank[Scouts][0] = "Scout"

	iRank[Merchants][0] = "Merchant"
	iRank[Merchants][1] = "4th Officer"
	iRank[Merchants][2] = "3rd Officer"
	iRank[Merchants][3] = "2nd Officer"
	iRank[Merchants][4] = "1st Officer"
	iRank[Merchants][5] = "Captain"
	iRank[Merchants][6] = "Captain"

	iRank[Other][0] = "Citizen"
}

// String converts the service into a String.
func (c CareerCT01) String() string {
	//return [...]string{"None", "Navy", "Marines", "Army", "Scouts", "Merchants", "Other"}[c]
	return serviceStr[c]
}

// Val converts the service into an integer.
func (c CareerCT01) Val() int {
	return int(c)
}

//var mainPage *tview.TextView    // The main page which to output progress information.
//var sidePane *tview.TextView    // The "object" page which to display the character.
//var sLogPane *ScrollableTextView // The logging page.

// UPP returns the Universal Personality Profile for a character as a string.
func (c CT01Char) UPP() string {

	return c.strength.String() + c.dexterity.String() + c.endurance.String() + c.intelligence.String() + c.education.String() + c.social.String()
}

// ObjectString outputs the character as a string suitable for display on a TextView.
func (c CT01Char) ObjectString() (s string) {

	var num int // General use as an int number.

	s = "[yellow]Name:[-] " + c.name + "\n"
	s += "[yellow]UPP:[-] " + c.UPP() + "\n\n"
	s += "[yellow]Race:[-] " + c.race + "\n"
	s += "[yellow]Sex:[-] " + c.sex + "\n"
	s += "[yellow]Service:[-] " + c.service.String() + "\n"
	s += "[yellow]Rank:[-] " + c.rank + "\n"
	s += fmt.Sprintf("[yellow]Terms:[-] %d\n", c.terms)
	s += fmt.Sprintf("[yellow]Age:[-] %d\n", c.age)
	s += fmt.Sprintf("[yellow]Cr:[-] %s\n", humanize.Comma(int64(c.credits)))

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

// SetPages sets up the generator with the pages to output to. These need to be set
// before using the generator or it will return.
//func SetPages(main, side *tview.TextView, logp *ScrollableTextView) {
//	mainPage = main
//	sidePane = side
//	sLogPane = logp
//}

// NewCT01Char returns a new CT Book 1 character
func NewCT01Char() (c *CT01Char) {

	character := CT01Char{
		age:          18,
		strength:     Ehex(D6() + D6()),
		dexterity:    Ehex(D6() + D6()),
		endurance:    Ehex(D6() + D6()),
		intelligence: Ehex(D6() + D6()),
		education:    Ehex(D6() + D6()),
		social:       Ehex(D6() + D6()),
		service:      None,
	}

	return &character

}

// GenerateCT01Character creates a character based on the given information, name, Service (as String), race, sex, and
// whether to kill the character if he/she fails a survival roll during generation.
func (c *CT01Char) GenerateCT01Character(service string, dieOnFail bool) {

	if mainPage == nil || sidePane == nil || logPane == nil {
		log.Println("GenerateCT01Character called before initialising with SetPages().")
		return
	}
	if c == nil {
		log.Printf("Attempt to generate a character without creating a new one! Please call NewCT01Char() first")
		return
	}

	var isDraftee bool
	mainPage.Clear()
	c.service, isDraftee = c.enlist(service)
	sidePane.Clear()
	sidePane.Write([]byte(c.ObjectString()))
	rankInt := 0
	c.rank = iRank[c.service][rankInt]
	if isDraftee {
		//
	}
	sidePane.Clear()
	sidePane.Write([]byte(c.ObjectString()))

	// The character is in a service and must go through each term from here.
	currentTerm := 0

	for {
		currentTerm++
		mainPageOutLn(humanize.Ordinal(currentTerm) + " term in in the " + c.service.String() + "\n")
		survives := checkSurvival(c)
		if survives {
			mainPageOutLn("You have survived this term.")
		} else {
			if dieOnFail {
				mainPageOutLn("You have been killed in service.")
				return
			}
			mainPageOutLn("You have been injured in service and must leave")
		}
		numSkills := 1 // Number of skills to obtain this term
		if currentTerm == 1 {
			numSkills++

			// Also receive rank and ser3vice skills.
			switch c.service {
			case Marines:
				mainPageOutLn("New Marines automatically receive Cutlass-1")
				c.addSkill("Cutlass", true)
			case Army:
				mainPageOutLn("New Soldiers automatically receive Rifle-1")
				c.addSkill("Rifle", true)
			case Scouts:
				mainPageOutLn("New Scouts automatically receive Pilot-1")
				c.addSkill("Pilot", true)

			}
		}

	}

}

// SetName sets the character's name.
func (c *CT01Char) SetName(name string) *CT01Char {
	c.name = name
	return c
}

// SetRace sets the character's race.
func (c *CT01Char) SetRace(race string) *CT01Char {
	c.race = race
	return c
}

// SetSex sets the character's sex.
func (c *CT01Char) SetSex(sex string) *CT01Char {
	c.sex = sex
	return c
}

// addSkill adds a skill or attribute change to a character. s is the string of the skill to be added to the character.
// If the skill is already present, it will be incremented. If skill is in the form "+1 Str" (eg) then an attribute
// will be updated. If atLevelOne is true, the skill will be added at level 1 if they do not already have the skill.
// Note that only valid and concrete skills can be added.
func (c *CT01Char) addSkill(s string, atLevelOne bool) {

}

// attributeChange handles an attribute change to a character, whether up or down. The attribute to change is indicated as a 3-character
// string (eg Str, Int, Soc, etc), and the amount of change in the integer. The new value is returned or -1 if the character has died in an aging crisis.
// Keeping within attribute bounds (0 - 15) and handling aging crisises is all done.
func (c *CT01Char) attributeChange(a string, v int) Ehex {
	switch a {
	case "Str":
		c.strength += Ehex(v)
		if c.strength <= 0 {
			if c.agingCrisis() == -1 {
				return -1
			}
			c.strength = 1
			mainPageOutLn("Strength set to 1 due to aging crisis.")
			return 1
		} else if c.strength > 15 {
			mainPageOutLn("Strength is limited to 15")
			c.strength = 15
			return 15
		}
		mainPageOutLn(fmt.Sprintf("New strength is %d", c.strength))
		return c.strength
	case "Dex":
		c.dexterity += Ehex(v)
		if c.dexterity <= 0 {
			if c.agingCrisis() == -1 {
				return -1
			}
			c.dexterity = 1
			mainPageOutLn("Dexterity set to 1 due to aging crisis.")
			return 1
		} else if c.dexterity > 15 {
			mainPageOutLn("Dexterity is limited to 15")
			c.dexterity = 15
			return 15
		}
		mainPageOutLn(fmt.Sprintf("New dexterity is %d", c.dexterity))
		return c.dexterity
	case "End":
		c.endurance += Ehex(v)
		if c.endurance <= 0 {
			if c.agingCrisis() == -1 {
				return -1
			}
			c.endurance = 1
			mainPageOutLn("Endurance set to 1 due to aging crisis.")
			return 1
		} else if c.endurance > 15 {
			mainPageOutLn("Endurance is limited to 15")
			c.endurance = 15
			return 15
		}
		mainPageOutLn(fmt.Sprintf("New endurance is %d", c.endurance))
		return c.endurance
	case "Int":
		c.intelligence += Ehex(v)
		if c.intelligence <= 0 {
			if c.agingCrisis() == -1 {
				return -1
			}
			c.intelligence = 1
			mainPageOutLn("Intelligence set to 1 due to aging crisis.")
			return 1
		} else if c.intelligence > 15 {
			mainPageOutLn("Intelligence is limited to 15")
			c.intelligence = 15
			return 15
		}
		mainPageOutLn(fmt.Sprintf("New intelligence is %d", c.intelligence))
		return c.intelligence
	case "Edu":
		c.education += Ehex(v)
		if c.education <= 0 {
			if c.agingCrisis() == -1 {
				return -1
			}
			c.education = 1
			mainPageOutLn("Education set to 1 due to aging crisis.")
			return 1
		} else if c.education > 15 {
			mainPageOutLn("Education is limited to 15")
			c.education = 15
			return 15
		}
		mainPageOutLn(fmt.Sprintf("New education is %d", c.education))
		return c.education
	case "Soc":
		c.social += Ehex(v)
		if c.social <= 0 {
			if c.agingCrisis() == -1 {
				return -1
			}
			c.social = 1
			mainPageOutLn("Social standing set to 1 due to aging crisis.")
			return 1
		} else if c.social > 15 {
			mainPageOutLn("Social standing is limited to 15")
			c.social = 15
			return 15
		}
		mainPageOutLn(fmt.Sprintf("New social standing is %d", c.social))
		return c.social
	default:
		//TODO: Fix this to do something useful instead of just panic.
		panic("Unknown characteristic adjustment for " + a + ".")
	}
}

// agingCrisis handles an aging crisis in a character. It returns -1 on character death, or the number of months the character has aged with slow drug.
func (c CT01Char) agingCrisis() int {
	mainPageOutLn("Character is undergoing an aging crisis. Assuming medic has Medic-2.")

	medicLevel := 2
	// for {
	// 	val, err := strconv.Atoi(getChoice("Input the Skill Level of any medic in attendance : "))
	// 	if err == nil {
	// 		medicLevel = val
	// 		break
	// 	} else {
	// 		fmt.Println(err)
	// 	}
	// }

	// TODO: Fix this to take into account medic's skill level. Here we have just assumed Medic-2.
	roll := D6() + D6() + medicLevel
	if roll >= 8 {
		// Survives
		months := D6()
		mainPageOutLn(fmt.Sprintf("Character has survived aging crisis with %d slow drug aging.", months))
		return months

	}
	// Dies
	mainPageOutLn("Character has died in aging crisis!")
	return -1
}

// enlist enlists the character in a service. It attempts to enlist in the service of
// choice. It returns the new service (CareerCT01) that the character has been enlisted
// in and should be copied into the character. If the choice is invalid, CareerCT01 None
// is returned. Also returned is whether the character is a draftee or not
func (c *CT01Char) enlist(choice string) (s CareerCT01, draftee bool) {

	s = None
	draftee = false

	// enlistment provides the required rolls for enlistment
	enlistment := [6]int{8, 9, 5, 7, 7, 3}

	d1 := D6()
	d2 := D6()
	dm := 0

	mainPageOutLn(fmt.Sprintf("For enlistment you rolled a %d + %d for total %d", d1, d2, d1+d2))

	// Assign dice modifiers.
	myService := 0
	switch choice {
	case "Navy":
		if c.intelligence >= 8 {
			dm++
			mainPageOutLn(choice + " enlistment: DM +1 for intelligence.")
		}
		if c.education >= 9 {
			dm += 2
			mainPageOutLn(choice + " enlistment: DM +2 for education.")
		}
		myService = int(Navy)
	case "Marines":
		if c.intelligence >= 8 {
			dm++
			mainPageOutLn(choice + " enlistment: DM +1 for intelligence.")
		}
		if c.strength >= 8 {
			dm += 2
			mainPageOutLn(choice + " enlistment: DM +2 for strength.")
		}
		myService = int(Marines)
	case "Army":
		if c.dexterity >= 6 {
			dm++
			mainPageOutLn(choice + " enlistment: DM +1 for dexterity.")
		}
		if c.endurance >= 5 {
			dm += 2
			mainPageOutLn(choice + " enlistment: DM +2 for endurance.")
		}
		myService = int(Army)
	case "Scouts":
		if c.intelligence >= 6 {
			dm++
			mainPageOutLn(choice + " enlistment: DM +1 for intelligence.")
		}
		if c.strength >= 8 {
			dm += 2
			mainPageOutLn(choice + " enlistment: DM +2 for strength.")
		}
		myService = int(Scouts)
	case "Merchants":
		if c.strength >= 7 {
			dm++
			mainPageOutLn(choice + " enlistment: DM +1 for strength.")
		}
		if c.intelligence >= 6 {
			dm += 2
			mainPageOutLn(choice + " enlistment: DM +2 for intelligence.")
		}
		myService = int(Merchants)
	case "Other":
		// No DMs to add
		myService = int(Other)
	default:
		str := "Invalid enlistment choice " + choice + "!"
		logPane.Write([]byte(str))
		log.Println(str)
		return
	}

	// Attempt the enlistment
	if d1+d2+dm >= enlistment[myService-1] {
		s = CareerCT01(myService)
		draftee = false
		mainPageOutLn("Enlistment has succeeded into the " + s.String() + ".")
	} else {
		draftee = true
		s = CareerCT01(D6())
		mainPageOutLn("Enlistment into the " + choice + " has failed. You have been drafted into the " + s.String() + ".")
	}
	return
}

// Checks if the character survives their term of service.
func checkSurvival(c *CT01Char) bool {

	// survival provides the required rolls for survival
	survival := [6]int{5, 6, 5, 7, 5, 5}

	d1 := D6()
	d2 := D6()
	dm := 0

	mainPageOutLn(fmt.Sprintf("For survival you rolled a %d + %d for total %d", d1, d2, d1+d2))

	switch c.service {
	case Navy:
		if c.intelligence >= 7 {
			dm = 2
			mainPageOutLn(c.service.String() + " survival: DM +2 for intelligence.")
		}
	case Marines:
		if c.endurance >= 8 {
			dm = 2
			mainPageOutLn(c.service.String() + " survival: DM +2 for endurance.")
		}
	case Army:
		if c.education >= 6 {
			dm = 2
			mainPageOutLn(c.service.String() + " survival: DM +2 for education.")
		}
	case Scouts:
		if c.endurance >= 9 {
			dm = 2
			mainPageOutLn(c.service.String() + " survival: DM +2 for endurance.")
		}
	case Merchants:
		if c.intelligence >= 7 {
			dm = 2
			mainPageOutLn(c.service.String() + " survival: DM +2 for intelligence.")
		}
	case Other:
		if c.intelligence >= 9 {
			dm = 2
			mainPageOutLn(c.service.String() + " survival: DM +2 for intelligence.")
		}
	default: // Increase age
		str := "Invalid service " + c.service.String() + "!"
		logPane.Write([]byte(str))
		log.Println(str)
		return false
	}

	if d1+d2+dm >= survival[c.service-1] {
		return true
	}
	return false
}

// mainPageOut writes a string to the Main Page.
func mainPageOut(s string) {
	mainPage.Write([]byte(s))
}

// mainPageOutLn writes a string a newline to the Main Page.
func mainPageOutLn(s string) {
	mainPage.Write([]byte(s + "\n"))
}

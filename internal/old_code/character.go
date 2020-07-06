package main

// character.go contains code objects and code specific to characters.

// Career represents the "service" the character enters.
type Career int

// character stores information about a Traveller character.
type character struct {

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
	credits int     // credits indicates amount of money held
	skills  []Skill // skills lists all the skills a character has
	terms   int     // terms indicates how many terms (4 year periods) the character has served
	service Career  // the (final) profession that the character served in

	// Other, optional parts
	sex       string   // sex is the characters sex.
	homeworld world    // homeworld indicates the characters homeworld.
	benefits  []string // benefits is a list of benefits the character has received on retirement.
	rank      string   // rank is the final rank of the character's profession.
	title     string   // Any title that the character possesses

	// Character record
	history []string // history shows a log of the character generation process.
}

package main

// ruleset.go contains information on the different Traveller rulesets

// Ruleset describes a distinct set of rules for Traveller.
type Ruleset int

// Constants defined for the specific rulesets.
const (
	RulesAll Ruleset = iota
	RulesClassic
	RulesMegaTraveller
	RulesTNE
	RulesTraveller4
	RulesMongoose
	RulesTraveller5
)

// String returns a string describing the Ruleset.
func (r Ruleset) String() string {
	return [...]string{"All", "Classic Traveller", "MegaTraveller", "The New Era", "Traveller 4", "Mongoose Traveller", "Traveller5", "Mongoose Traveller 2"}[r]
}

// Abbr returns an abbreviation for the Ruleset.
func (r Ruleset) Abbr() string {
	return [...]string{"All", "CT", "MT", "TNE", "T4", "MgT", "T5", "MgT2"}[r]
}

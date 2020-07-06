package main

import "fmt"

// skill.go contains code for dealing with skills

// Skill defines a particular skill, including level, held by a Traveller.
type Skill struct {
	Name  string // The skill name
	Level int    // The skill level
}

// String returns the skill as a string.
func (s Skill) String() string {
	return fmt.Sprintf("%s-%d", s.Name, s.Level)

}

// FindSkill looks through a slice of skills provided in ss for a skill matching the search string
// and returns the index if found, or -1 if not.
func FindSkill(ss []Skill, search string) int {

	for i, skill := range ss {
		if skill.Name == search {
			return i
		}
	}

	return -1
}

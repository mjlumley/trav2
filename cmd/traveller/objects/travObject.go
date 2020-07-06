// Package obj provides custom objects designed for use with Traveller, Characters, Worlds, and others.
package obj

// TravObject defines a generic traveller object that can be loaded, displayed, edited and saved.
type TravObject interface {

	// Draws the object on the screen.
	Draw()

	// Loads the object from the database.
	Load(idx int) *TravObject

	// Saves the object to the database.
	Save()

	// Displays a window for editing the object on the screen.
	Edit()

	// Returns whether the object has been edited or not.
	IsDirty() bool

	// The Traveller "Standard" output for the object.
	String() string

	// Gets the Database ID of the object.
	GetID() int
}

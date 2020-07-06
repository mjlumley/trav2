package main

import (
	"github.com/rivo/tview"
)

// ObjectPane.go details an "Object Pane".

// ObjectPane is a tview.TextView that is used for displaying
// details of generated "objects", such as characters and worlds.
type ObjectPane struct {
	tview.TextView

	identifier string // Identifies the ObjectPane throughout the application.
}

// NewObjectPane creates and intialises a new ObjectPane.
func NewObjectPane(id string) *ObjectPane {

	// Create an isntance
	o := ObjectPane{TextView: *tview.NewTextView(), identifier: id}
	o.SetDynamicColors(true).SetTextAlign(tview.AlignLeft).SetBorder(true).SetTitle(" Object ")

	return &o
}

// GetIdentifier gets the idnetifier for this object pane.
func (o *ObjectPane) GetIdentifier() string {
	return o.identifier
}

// DisplayObject displays the object (the string) in the object window.
// The object pane is first cleared. It does not force a draw of the screen.
func (o *ObjectPane) DisplayObject(obj string) {
	o.TextView.Clear()
	o.TextView.Write([]byte(obj))
}

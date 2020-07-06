package main

// LogPane.go contains code for managing LogPane objects.

import (
	"fmt"
	"log"

	"github.com/rivo/tview"
)

/* // LogPane describes a window pane for writing log output to.
type LogPane struct {
	ScrollableTextView

	identifier string
}

// NewLogPane returns a new LogPane object.
func NewLogPane(id string) *LogPane {
	lg := LogPane{ScrollableTextView: *NewScrollableTextView(), identifier: id}

	lg.ScrollableTextView.SetTextAlign(tview.AlignLeft).SetBorder(true).SetTitle(" Log Viewer ").SetTitleAlign(tview.AlignLeft)
	lg.SetScrollable(true).ScrollToEnd().SetWordWrap(true)

	return &lg
}

// Log logs a message to the log pane and to whatever logging is done.
func (lg *LogPane) Log(s string) {
	lg.Write([]byte(fmt.Sprintf("%s\n", s)))
	lg.ScrollToEnd()
	log.Printf("%s\n", s)
}
*/

// LogPane describes a window pane for writing log output to.
type LogPane struct {
	tview.TextView

	identifier string
}

// NewLogPane returns a new LogPane object.
func NewLogPane(id string) *LogPane {
	lg := LogPane{TextView: *tview.NewTextView(), identifier: id}

	lg.TextView.SetTextAlign(tview.AlignLeft).SetBorder(true).SetTitle(" Log Viewer ").SetTitleAlign(tview.AlignLeft)
	lg.SetScrollable(true).ScrollToEnd().SetWordWrap(true)

	return &lg
}

// Log logs a message to the log pane and to whatever logging is done.
func (lg *LogPane) Log(s string) {
	lg.Write([]byte(fmt.Sprintf("%s\n", s)))
	lg.ScrollToEnd()
	log.Printf("%s\n", s)
}

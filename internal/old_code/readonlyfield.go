package main

// readonlyfield.go provides a read-only field to be used on place of the Input Field on a form.

import (
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

// ReadonlyField implements a simple textfield for display readonly text.
type ReadonlyField struct {
	*tview.Box

	// The text that is contained in the field
	text string

	// The text to be displayed before the input area.
	label string

	// The screen width of the label area. A value of 0 means use the width of
	// the label text.
	labelWidth int

	// The label color.
	labelColor tcell.Color

	// The background color of the input area.
	fieldBackgroundColor tcell.Color

	// The text color of the input area.
	fieldTextColor tcell.Color

	// An optional function which is called when the user indicated that they
	// are done entering text. The key which was pressed is provided (tab,
	// shift-tab, or escape).
	done func(tcell.Key)

	// A callback function set by the Form class and called when the user leaves
	// this form item.
	finished func(tcell.Key)
}

// NewReadonlyField returns a new read-only field.
func NewReadonlyField() *ReadonlyField {
	return &ReadonlyField{
		Box:                  tview.NewBox(),
		labelColor:           tview.Styles.SecondaryTextColor,
		fieldBackgroundColor: tview.Styles.ContrastBackgroundColor,
		fieldTextColor:       tview.Styles.PrimaryTextColor,
	}
}

// SetText sets the current text of the readonly field.
func (r *ReadonlyField) SetText(text string) *ReadonlyField {
	r.text = text
	return r
}

// SetLabel sets the text to be displayed before the input area.
func (r *ReadonlyField) SetLabel(label string) *ReadonlyField {
	r.label = label
	return r
}

// GetLabel returns the text to be displayed before the input area.
//
// Note this is required to extend FormItem.
func (r *ReadonlyField) GetLabel() string {
	return r.label
}

// SetLabelWidth sets the screen width of the label. A value of 0 will cause the
// primitive to use the width of the label string.
func (r *ReadonlyField) SetLabelWidth(width int) *ReadonlyField {
	r.labelWidth = width
	return r
}

// SetLabelColor sets the color of the label.
func (r *ReadonlyField) SetLabelColor(color tcell.Color) *ReadonlyField {
	r.labelColor = color
	return r
}

// SetFieldBackgroundColor sets the background color of the input area.
func (r *ReadonlyField) SetFieldBackgroundColor(color tcell.Color) *ReadonlyField {
	r.fieldBackgroundColor = color
	return r
}

// SetFieldTextColor sets the text color of the input area.
func (r *ReadonlyField) SetFieldTextColor(color tcell.Color) *ReadonlyField {
	r.fieldTextColor = color
	return r
}

// SetFormAttributes sets attributes shared by all form items.
//
// Note this is required to extend FormItem.
func (r *ReadonlyField) SetFormAttributes(labelWidth int, labelColor, bgColor, fieldTextColor, fieldBgColor tcell.Color) tview.FormItem {
	r.labelWidth = labelWidth
	r.labelColor = labelColor
	r.Box.SetBackgroundColor(bgColor)
	r.fieldTextColor = fieldTextColor
	r.fieldBackgroundColor = fieldBgColor
	return r
}

// GetFieldWidth returns this primitive's field width.
//
// Note this is required to extend FormItem.
func (r *ReadonlyField) GetFieldWidth() int {
	return len(r.text)
}

// SetDoneFunc sets a handler which is called when the user is done using the
// checkbox. The callback function is provided with the key that was pressed,
// which is one of the following:
//
//   - KeyEscape: Abort text input.
//   - KeyTab: Move to the next field.
//   - KeyBacktab: Move to the previous field.
func (r *ReadonlyField) SetDoneFunc(handler func(key tcell.Key)) *ReadonlyField {
	r.done = handler
	return r
}

// SetFinishedFunc sets a callback invoked when the user leaves this form item.
// Note this is required in order to extend FormItem.
func (r *ReadonlyField) SetFinishedFunc(handler func(key tcell.Key)) tview.FormItem {
	r.finished = handler
	return r
}

// Draw draws this primitive onto the screen.
func (r *ReadonlyField) Draw(screen tcell.Screen) {
	r.Box.Draw(screen)

	// Prepare
	x, y, width, height := r.GetInnerRect()
	rightLimit := x + width
	if height < 1 || rightLimit <= x {
		return
	}

	// Draw label.
	if r.labelWidth > 0 {
		labelWidth := r.labelWidth
		if labelWidth > rightLimit-x {
			labelWidth = rightLimit - x
		}
		tview.Print(screen, r.label, x, y, labelWidth, tview.AlignLeft, r.labelColor)
		x += labelWidth
	} else {
		_, drawnWidth := tview.Print(screen, r.label, x, y, rightLimit-x, tview.AlignLeft, r.labelColor)
		x += drawnWidth
	}

	// Draw field.
	tview.Print(screen, r.text, x, y, len(r.text), tview.AlignLeft, r.fieldTextColor)
}

// InputHandler returns the handler for this primitive.
func (r *ReadonlyField) InputHandler() func(event *tcell.EventKey, setFocus func(p tview.Primitive)) {
	return r.WrapInputHandler(func(event *tcell.EventKey, setFocus func(p tview.Primitive)) {
		// Process key event.
		switch key := event.Key(); key {
		case tcell.KeyEnter, tcell.KeyTab, tcell.KeyBacktab, tcell.KeyEscape: // We're done.
			if r.done != nil {
				r.done(key)
			}
			if r.finished != nil {
				r.finished(key)
			}
		}
	})
}

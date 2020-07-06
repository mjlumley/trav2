package main

import (
	"strings"

	"github.com/gdamore/tcell"
	"github.com/rivo/tview"
)

// ScrollableTextView is a TextView which can display scrollbars. It embeds the TextView struct
// so can be used as a drop-in replacement for a TextView, and so is identical with the exception
// of the scrollbars. By default, vertical scrollbars are displayed, and horizontal bars are
// not. This matches the settings for the embedded TextView. Refer to tview.TextView for more details.
//
// Scrollbars
//
// There are two types of scrollbar, verticalBar and horizontalBar. These can
// be turned on or off by using SetHorizontalBar() and SetVerticalBar(). The scrollbars will only
// display if there is enough data in the TextView's text for them. Note that it only turns on or
// off the scrollbar display, it does not disable or enable the ability of the TextView to scroll!
//
// So, note that the setting of the actual scrolling or wrapping of the TextView is separate to the
// appearance of the scrollbars.
//
type ScrollableTextView struct {
	*tview.TextView

	// If set to true, the TextView will display a horizontal scrollbar.
	horizontalBar bool

	// If set to true, the TextView will display a vertical scrollbar.
	verticalBar bool
}

// NewScrollableTextView returns a new scrollable text view.
func NewScrollableTextView() *ScrollableTextView {
	return &ScrollableTextView{
		TextView:      tview.NewTextView(),
		horizontalBar: false,
		verticalBar:   true}
}

// SetHorizontalBar sets the flag that, if true, leads to a horizontal
// scrollbar being displayed. If false, the horizontal bar will not be
// displayed even if the text is wider than the display.
func (s *ScrollableTextView) SetHorizontalBar(barOn bool) *ScrollableTextView {
	s.horizontalBar = barOn
	return s
}

// SetVerticalBar sets the flag that, if true, leads to a vertical
// scrollbar being displayed. If false, the vertical bar will not be
// displayed even if the text is longer than the display.
func (s *ScrollableTextView) SetVerticalBar(barOn bool) *ScrollableTextView {
	s.verticalBar = barOn
	return s
}

// GetHorizontalBar returns the value of the horizontal bar flag, true if
// displayed, false if not.
func (s *ScrollableTextView) GetHorizontalBar() bool {
	return s.horizontalBar
}

// GetVerticalBar returns the value of the vertical bar flag, true if
// displayed, false if not.
func (s *ScrollableTextView) GetVerticalBar() bool {
	return s.verticalBar
}

// SetScrollable sets the flag that decides whether or not the TextView is
// scrollable. If true, text is kept in a buffer and can be navigated. If false,
// the last line will always be visible. You need to use SetVerticalBar(true) to
// additionally display the bar, this only passes the value onto the underlying
// TextView, but will turn the display of the bar off additionally if the value
// of scrollable is false.
func (s *ScrollableTextView) SetScrollable(scrollable bool) *ScrollableTextView {
	if !scrollable {
		s.SetVerticalBar(false)
	}
	s.TextView.SetScrollable(scrollable)
	return s
}

// SetWrap sets the flag that, if true, leads to lines that are longer than the
// available width being wrapped onto the next line. If false, any characters
// beyond the available width are not displayed.
//
// If wrap is turned on, then the display of the bar is turned off.  If wrap
// is turned off, then horizontal bars MAY be displayed
// (use SetHorizontalBar(true)). By default, wrap is on and bars are hidden.
func (s *ScrollableTextView) SetWrap(wrap bool) *ScrollableTextView {
	if wrap {
		s.SetHorizontalBar(false)
	}
	s.TextView.SetWrap(wrap)
	return s
}

// SetWordWrap sets the flag that, if true and if the "wrap" flag is also true
// (see SetWrap()), wraps the line at spaces or after punctuation marks. Note
// that trailing spaces will not be printed.
//
// This flag is ignored if the "wrap" flag is false.
//
// If word wrap is turned on, then the display of the bar is turned off.  If wrap
// is turned off, then horizontal bars MAY be displayed
// (use SetHorizontalBar(true)). By default, wrap is on (word wrap is not), and
// bars are hidden.
func (s *ScrollableTextView) SetWordWrap(wrapOnWords bool) *ScrollableTextView {
	if wrapOnWords {
		s.SetHorizontalBar(false)
	}
	s.TextView.SetWordWrap(wrapOnWords)
	return s
}

// Draw draws the scrollable text view on the screen. It calls the embedded TextView.
func (s *ScrollableTextView) Draw(screen tcell.Screen) {

	x, y, width, height := s.TextView.GetRect()
	s.TextView.Draw(screen)

	// Quick return
	if !s.verticalBar && !s.horizontalBar {
		return
	}

	// Calculate the parameters required for the scrollbar calculation
	_, _, iWidth, iHeight := s.TextView.GetInnerRect()
	text := s.TextView.GetText(true)
	buffer := strings.Split(strings.Replace(text, "\r\n", "\n", -1), "\n")
	var longestLine int
	for _, str := range buffer {
		if len(str) > longestLine {
			longestLine = len(str)
		}
	}
	numLines := len(buffer)
	rowOffset, colOffset := s.TextView.GetScrollOffset()

	// Draw a vertical scroll bar
	if s.verticalBar && numLines > iHeight {
		screen.SetContent(x+width-1, y+1, tcell.RuneUArrow, nil, tcell.StyleDefault.Foreground(tview.Styles.BorderColor).Background(tview.Styles.PrimitiveBackgroundColor))
		for ny := y + 2; ny < y+height-2; ny++ {
			screen.SetContent(x+width-1, ny, tcell.RuneCkBoard, nil, tcell.StyleDefault.Foreground(tview.Styles.BorderColor).Background(tview.Styles.PrimitiveBackgroundColor))
		}
		screen.SetContent(x+width-1, y+height-2, tcell.RuneDArrow, nil, tcell.StyleDefault.Foreground(tview.Styles.BorderColor).Background(tview.Styles.PrimitiveBackgroundColor))

		// Determine the position of the "handle"
		barHeight := height - 4
		vBarLoc := (rowOffset * barHeight) / (numLines - iHeight)
		if rowOffset+iHeight >= numLines {
			vBarLoc = barHeight - 1
		}
		screen.SetContent(x+width-1, y+2+vBarLoc, tcell.RuneBlock, nil, tcell.StyleDefault.Foreground(tview.Styles.BorderColor).Background(tview.Styles.PrimitiveBackgroundColor))
	}

	// Draw a horizontal scroll bar
	if s.horizontalBar && longestLine > iWidth {
		screen.SetContent(x+1, y+height-1, tcell.RuneLArrow, nil, tcell.StyleDefault.Foreground(tview.Styles.BorderColor).Background(tview.Styles.PrimitiveBackgroundColor))
		for nx := x + 2; nx < x+width-2; nx++ {
			screen.SetContent(nx, y+height-1, tcell.RuneCkBoard, nil, tcell.StyleDefault.Foreground(tview.Styles.BorderColor).Background(tview.Styles.PrimitiveBackgroundColor))
		}
		screen.SetContent(x+width-2, y+height-1, tcell.RuneRArrow, nil, tcell.StyleDefault.Foreground(tview.Styles.BorderColor).Background(tview.Styles.PrimitiveBackgroundColor))

		// Detemine the position of the "handle"
		barWidth := width - 4
		hBarLoc := (colOffset * barWidth) / (longestLine - iWidth)
		if colOffset+iWidth >= longestLine {
			hBarLoc = barWidth - 1
		}
		screen.SetContent(x+2+hBarLoc, y+height-1, tcell.RuneBlock, nil, tcell.StyleDefault.Foreground(tview.Styles.BorderColor).Background(tview.Styles.PrimitiveBackgroundColor))
	}
}

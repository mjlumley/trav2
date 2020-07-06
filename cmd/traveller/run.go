package main

import (
	"bytes"
	"fmt"
	"time"

	"github.com/inkyblackness/imgui-go"
)

// Platform covers mouse/keyboard/gamepad inputs, cursor shape, timing, windowing.
type Platform interface {
	// ShouldStop is regularly called as the abort condition for the program loop.
	ShouldStop() bool
	// ProcessEvents is called once per render loop to dispatch any pending events.
	ProcessEvents()
	// DisplaySize returns the dimension of the display.
	DisplaySize() [2]float32
	// FramebufferSize returns the dimension of the framebuffer.
	FramebufferSize() [2]float32
	// NewFrame marks the begin of a render pass. It must update the imgui IO state according to user input (mouse, keyboard, ...)
	NewFrame()
	// PostRender marks the completion of one render pass. Typically this causes the display buffer to be swapped.
	PostRender()
	// ClipboardText returns the current text of the clipboard, if available.
	ClipboardText() (string, error)
	// SetClipboardText sets the text as the current text of the clipboard.
	SetClipboardText(text string)
}

type gClipboard struct {
	platform Platform
}

func (board gClipboard) Text() (string, error) {
	return board.platform.ClipboardText()
}

func (board gClipboard) SetText(text string) {
	board.platform.SetClipboardText(text)
}

// Renderer covers rendering imgui draw data.
type Renderer interface {
	// PreRender causes the display buffer to be prepared for new output.
	PreRender(clearColor [3]float32)
	// Render draws the provided imgui draw data.
	Render(displaySize [2]float32, framebufferSize [2]float32, drawData imgui.DrawData)
}

const (
	millisPerSecond = 1000
	sleepDuration   = time.Millisecond * 25
)

// Run implements the main program loop of the demo. It returns when the platform signals to stop.
// This demo application shows some basic features of ImGui, as well as exposing the standard demo window.
func Run(p Platform, r Renderer) {
	imgui.CurrentIO().SetClipboard(gClipboard{platform: p})

	showDemoWindow := false
	showGoDemoWindow := false
	showDebugWindow := false
	showLogWindow := false
	showWordgenWindow := false
	clearColor := [3]float32{0.0, 0.0, 0.0}
	f := float32(0)
	counter := 0

	languageSelection := 0
	languagePatterns := false

	doExitPopup := false
	doNotImplementedPopup := false
	doAbout := false

	doExit := false
	autoscrollLog := true

	var langText bytes.Buffer

	for !p.ShouldStop() {
		p.ProcessEvents()

		// Signal start of a new frame
		p.NewFrame()
		imgui.NewFrame()

		// 0. Main screen setup.
		// MainMenuBar
		if imgui.BeginMainMenuBar() {
			if imgui.BeginMenu("File") {
				if imgui.MenuItem("Open") {
					doNotImplementedPopup = true
				}
				if imgui.MenuItem("Backup") {
					doNotImplementedPopup = true
				}
				imgui.Separator()
				if imgui.MenuItem("Settings") {
					doNotImplementedPopup = true
				}
				imgui.Separator()
				// If quit is selected, run up an "quit if you are sure" window and exit.
				if imgui.MenuItemV("Quit", "Alt-F4", false, true) {
					doExitPopup = true
				}
				imgui.EndMenu()
			}
			if imgui.BeginMenu("Edit") {
				if imgui.MenuItem("Copy") {
					doNotImplementedPopup = true
				}
				if imgui.MenuItem("Paste") {
					doNotImplementedPopup = true
				}
				if imgui.MenuItem("Select All") {
					doNotImplementedPopup = true
				}
				imgui.Separator()
				if imgui.MenuItemV("Log Window", "Alt-L", showLogWindow, true) {
					showLogWindow = !showLogWindow
				}
				if imgui.MenuItem("Object Window") {
					doNotImplementedPopup = true
				}
				imgui.EndMenu()
			}
			if imgui.BeginMenu("Generate") {
				if imgui.MenuItem("Characters") {
					doNotImplementedPopup = true
				}
				if imgui.MenuItemV("Language", "", showWordgenWindow, true) {
					showWordgenWindow = !showWordgenWindow
				}
				if imgui.MenuItem("Worlds") {
					doNotImplementedPopup = true
				}
				imgui.EndMenu()
			}
			if imgui.BeginMenu("Tools") {
				if imgui.MenuItem("Manage Campaign") {
					doNotImplementedPopup = true
				}
				if imgui.MenuItem("Search") {
					doNotImplementedPopup = true
				}
				if imgui.MenuItem("ImGui-Go Debug") {
					showDebugWindow = true
				}
				imgui.EndMenu()
			}
			if imgui.BeginMenu("Help") {
				if imgui.MenuItem("Help") {
					doNotImplementedPopup = true
				}
				if imgui.MenuItem("About") {
					doAbout = true
				}
				imgui.EndMenu()
			}
			imgui.EndMainMenuBar()
		}

		// 1. Show the debug demo window.
		// Tip: if we don't call imgui.Begin()/imgui.End() the widgets automatically appears in a window called "Debug".
		if showDebugWindow {
			imgui.BeginV("ImGui-Go Debug", &showDebugWindow, 0)
			//imgui.Text("ภาษาไทย测试조선말")                   // To display these, you'll need to register a compatible font
			imgui.Text("Hello, world!")                  // Display some text
			imgui.SliderFloat("float", &f, 0.0, 1.0)     // Edit 1 float using a slider from 0.0f to 1.0f
			imgui.ColorEdit3("clear color", &clearColor) // Edit 3 floats representing a color

			imgui.Checkbox("Demo Window", &showDemoWindow) // Edit bools storing our window open/close state
			imgui.Checkbox("Go Demo Window", &showGoDemoWindow)

			if imgui.Button("Button") { // Buttons return true when clicked (most widgets return true when edited/activated)
				counter++
			}
			imgui.SameLine()
			imgui.Text(fmt.Sprintf("counter = %d", counter))

			imgui.Text(fmt.Sprintf("Application average %.3f ms/frame (%.1f FPS)",
				millisPerSecond/imgui.CurrentIO().Framerate(), imgui.CurrentIO().Framerate()))
			imgui.End()
		}

		// 3. Show the ImGui demo window. Most of the sample code is in imgui.ShowDemoWindow().
		// Read its code to learn more about Dear ImGui!
		if showDemoWindow {
			// Normally user code doesn't need/want to call this because positions are saved in .ini file anyway.
			// Here we just want to make the demo initial state a bit more friendly!
			const demoX = 650
			const demoY = 20
			imgui.SetNextWindowPosV(imgui.Vec2{X: demoX, Y: demoY}, imgui.ConditionFirstUseEver, imgui.Vec2{})

			imgui.ShowDemoWindow(&showDemoWindow)
		}
		if showGoDemoWindow {
			Show(&showGoDemoWindow)
		}

		// 4. Show the Log window
		if showLogWindow {

			// Set window location and size first time only.
			imgui.SetNextWindowPosV(imgui.Vec2{X: 0, Y: windowHeight - 240}, imgui.ConditionFirstUseEver, imgui.Vec2{})
			imgui.SetNextWindowSizeV(imgui.Vec2{X: windowWidth, Y: 240}, imgui.ConditionFirstUseEver)

			// Start
			imgui.BeginV("Traveller Log", &showLogWindow, 0)

			// Popup for log options
			if imgui.BeginPopup("Log Options") {
				imgui.Checkbox("Auto-scroll", &autoscrollLog)
				imgui.EndPopup()
			}

			if imgui.Button("Options") {
				imgui.OpenPopup("Log Options")
			}
			imgui.SameLine()
			clear := imgui.Button("Clear")
			imgui.Separator()
			if clear {
				memLog.Reset()
			}
			imgui.BeginChildV("logscroll", imgui.Vec2{}, false, imgui.WindowFlagsHorizontalScrollbar)
			imgui.PushStyleVarVec2(imgui.StyleVarItemSpacing, imgui.Vec2{})
			imgui.Text(memLog.String())
			imgui.PopStyleVar()
			if autoscrollLog && imgui.GetScrollY() >= imgui.GetScrollMaxX() {
				imgui.SetScrollHereY(1.0)
			}
			imgui.EndChild()
			imgui.End()
		}

		// 5. Show the Word Generation window
		if showWordgenWindow {

			imgui.SetNextWindowPosV(imgui.Vec2{X: 0, Y: 40}, imgui.ConditionFirstUseEver, imgui.Vec2{})
			imgui.SetNextWindowSizeV(imgui.Vec2{X: 480, Y: 480}, imgui.ConditionFirstUseEver)

			// Start the Word Generation Window
			imgui.BeginV("Word Generation", &showWordgenWindow, 0)
			// Dropdown for language selection
			if imgui.BeginComboV("Language", Language(languageSelection).String(), 0) {
				for n := LanguageAslan; n < LanguageNone; n++ {
					isSelected := Language(languageSelection) == n
					if imgui.SelectableV(n.String(), isSelected, 0, imgui.Vec2{}) {
						languageSelection = int(n)
					}
					if isSelected {
						imgui.SetItemDefaultFocus()
					}
				}
				imgui.EndCombo()
			}
			imgui.Checkbox("Show patterns", &languagePatterns)
			imgui.SameLine()
			HelpMarker("Will show the consonant/verb/syllable pattern of the word.")
			if imgui.Button("Generate") {
				langText.WriteString(getLanguageWord(languageSelection, languagePatterns) + "\n")
			}
			imgui.SameLine()
			if imgui.Button("Clear") {
				langText.Reset()
			}
			imgui.Separator()
			imgui.BeginChildV("langscroll", imgui.Vec2{}, false, imgui.WindowFlagsHorizontalScrollbar)
			imgui.PushStyleVarVec2(imgui.StyleVarItemSpacing, imgui.Vec2{})
			imgui.Text(langText.String())
			imgui.PopStyleVar()
			if imgui.GetScrollY() >= imgui.GetScrollMaxX() {
				imgui.SetScrollHereY(1.0)
			}
			imgui.EndChild()
			imgui.End()
		}

		// For not implemented features
		if doNotImplementedPopup {
			imgui.OpenPopup("Not Implemented")
			doNotImplementedPopup = false
		}
		if imgui.BeginPopupModalV("Not Implemented", nil, imgui.WindowFlagsAlwaysAutoResize) {
			imgui.Text("This feature has not been implemented yet!")
			if imgui.Button("OK") {
				imgui.CloseCurrentPopup()
			}
			imgui.EndPopup()
		}

		// The About window. Another modal.
		if doAbout {
			imgui.OpenPopup("About")
			doAbout = false
		}
		if imgui.BeginPopupModalV("About", nil, imgui.WindowFlagsAlwaysAutoResize) {
			imgui.Text((fmt.Sprintf("The Travellers Tool\n\nVersion %s\n", appVersion)))
			if imgui.Button("OK") {
				imgui.CloseCurrentPopup()
			}
			imgui.EndPopup()
		}

		// Handle the exit
		if doExitPopup {
			imgui.OpenPopup("Are you sure?")
			doExitPopup = false
		}
		if imgui.BeginPopupModalV("Are you sure?", nil, imgui.WindowFlagsAlwaysAutoResize) {
			imgui.Text("Are you sure you wish to exit?")
			if imgui.Button("OK") {
				doExit = true
				imgui.CloseCurrentPopup()
			}
			imgui.SetItemDefaultFocus()
			imgui.SameLine()
			if imgui.Button("Cancel") {
				imgui.CloseCurrentPopup()
			}
			imgui.EndPopup()
		}

		// Rendering
		imgui.Render() // This call only creates the draw data list. Actual rendering to framebuffer is done below.

		r.PreRender(clearColor)
		// At this point, the application could perform its own rendering...
		// app.RenderScene()

		r.Render(p.DisplaySize(), p.FramebufferSize(), imgui.RenderedDrawData())
		p.PostRender()

		// sleep to avoid 100% CPU usage for this demo
		<-time.After(sleepDuration)

		// Are we exiting?
		if doExit {
			return
		}
	}
}

func getLanguageWord(lidx int, showPattern bool) (retValue string) {

	lang := Language(lidx)

	retValue, pattern := lang.GenWord()

	if showPattern {
		retValue += " - [" + lang.String() + "] - (" + pattern + ")"
	}

	return
}

// HelpMarker displaya a little (?) mark which shows a tooltip when hovered.
// In your own code you may want to display an actual icon if you are using a merged icon fonts (see docs/FONTS.md)
func HelpMarker(desc string) {

	//imgui.PushStyleColor(imgui.StyleColorText,)  //     colors[ImGuiCol_TextDisabled]           = ImVec4(0.50f, 0.50f, 0.50f, 1.00f);
	imgui.PushStyleColor(imgui.StyleColorText, imgui.Vec4{X: 0.5, Y: 0.5, Z: 0.5, W: 1.0})
	imgui.Text("(?)") // ImGui::TextDisabled
	imgui.PopStyleColor()
	if imgui.IsItemHovered() {
		imgui.BeginTooltip()
		imgui.PushTextWrapPosV(imgui.FontSize() * 35.0)
		imgui.Text(desc)
		imgui.PopTextWrapPos()
		imgui.EndTooltip()
	}
}

/*
   PushStyleColor(ImGuiCol_Text, GImGui->Style.Colors[ImGuiCol_TextDisabled]);
   TextV(fmt, args);
   PopStyleColor();
*/

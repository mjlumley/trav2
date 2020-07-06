package main

import (
	"fmt"
	"log"

	"github.com/inkyblackness/imgui-go"
)

type windowFlags struct {
	noTitlebar     bool
	noScrollbar    bool
	noMenu         bool
	noMove         bool
	noResize       bool
	noCollapse     bool
	noNav          bool
	noBackground   bool
	noBringToFront bool
}

func (f windowFlags) combined() int {
	flags := 0
	if f.noTitlebar {
		flags |= imgui.WindowFlagsNoTitleBar
	}
	if f.noScrollbar {
		flags |= imgui.WindowFlagsNoScrollbar
	}
	if !f.noMenu {
		flags |= imgui.WindowFlagsMenuBar
	}
	if f.noMove {
		flags |= imgui.WindowFlagsNoMove
	}
	if f.noResize {
		flags |= imgui.WindowFlagsNoResize
	}
	if f.noCollapse {
		flags |= imgui.WindowFlagsNoCollapse
	}
	if f.noNav {
		flags |= imgui.WindowFlagsNoNav
	}
	if f.noBackground {
		flags |= imgui.WindowFlagsNoBackground
	}
	if f.noBringToFront {
		flags |= imgui.WindowFlagsNoBringToFrontOnFocus
	}
	return flags
}

var window = struct {
	flags   windowFlags
	noClose bool

	widgets widgets
	layout  layout
	popups  popups
	columns columns
	misc    misc
}{}

func bulletText(text string) {
	imgui.Bullet()
	imgui.Text(text)
}

// Show demonstrates most ImGui features that were ported to Go.
// This function tries to recreate the original demo window as closely as possible.
//
// In theory, if both windows would provide the identical functionality, then the wrapper would be complete.
func Show(keepOpen *bool) {
	imgui.SetNextWindowPosV(imgui.Vec2{X: 650, Y: 20}, imgui.ConditionFirstUseEver, imgui.Vec2{})
	imgui.SetNextWindowSizeV(imgui.Vec2{X: 550, Y: 680}, imgui.ConditionFirstUseEver)

	if window.noClose {
		keepOpen = nil
	}
	if !imgui.BeginV("ImGui-Go Demo", keepOpen, window.flags.combined()) {
		// Early out if the window is collapsed, as an optimization.
		imgui.End()
		return
	}

	// Use fixed width for labels (by passing a negative value), the rest goes to widgets.
	// We choose a width proportional to our font size.
	imgui.PushItemWidth(imgui.FontSize() * -12)

	// MenuBar
	if imgui.BeginMenuBar() {
		if imgui.BeginMenu("File") {
			showMainMenuFile()
			imgui.EndMenu()
		}
		if imgui.BeginMenu("Edit") {
			if imgui.MenuItemV("Undo", "CTRL+Z", false, true) {
				log.Print("Undo")
			}
			imgui.EndMenu()
		}
		if imgui.BeginMenu("Generate") {
			imgui.EndMenu()
		}
		imgui.EndMenuBar()
	}

	/*
			        if (ImGui::BeginMenu("File"))
		        {
		            ShowExampleMenuFile();
		            ImGui::EndMenu();
		        }
		        if (ImGui::BeginMenu("Edit"))
		        {
		            if (ImGui::MenuItem("Undo", "CTRL+Z")) {}
		            if (ImGui::MenuItem("Redo", "CTRL+Y", false, false)) {}  // Disabled item
		            ImGui::Separator();
		            if (ImGui::MenuItem("Cut", "CTRL+X")) {}
		            if (ImGui::MenuItem("Copy", "CTRL+C")) {}
		            if (ImGui::MenuItem("Paste", "CTRL+V")) {}
		            ImGui::EndMenu();
		        }
		        ImGui::EndMainMenuBar();

	*/

	imgui.Text(fmt.Sprintf("dear imgui says hello. (%s)", imgui.Version()))
	imgui.Spacing()

	if imgui.CollapsingHeader("Help") {
		imgui.Text("ABOUT THIS DEMO:")
		bulletText("Sections below are demonstrating many aspects of the wrapper.")
		bulletText("This demo may not be complete. Refer to the \"native\" demo window for a full overview.")
		bulletText("The \"Examples\" menu above leads to more demo contents.")
		bulletText("The \"Tools\" menu above gives access to: About Box, Style Editor,\n" +
			"and Metrics (general purpose Dear ImGui debugging tool).")
		imgui.Separator()

		imgui.Text("PROGRAMMER GUIDE:")
		bulletText("See the demo.Show() code in internal/demo/Window.go. <- you are here!")
		bulletText("See comments in imgui.cpp.")
		bulletText("See example applications in the examples/ folder.")
		bulletText("Read the FAQ at http://www.dearimgui.org/faq/")
		bulletText("Set 'io.ConfigFlags |= NavEnableKeyboard' for keyboard controls.")
		bulletText("Set 'io.ConfigFlags |= NavEnableGamepad' for gamepad controls.")
		imgui.Separator()

		imgui.Text("USER GUIDE:")
		showUserGuide()
	}

	// MISSING: Configuration

	if imgui.CollapsingHeader("Window options") {
		imgui.Checkbox("No titlebar", &window.flags.noTitlebar)
		imgui.SameLineV(150, -1)
		imgui.Checkbox("No scrollbar", &window.flags.noScrollbar)
		imgui.SameLineV(300, -1)
		imgui.Checkbox("No menu", &window.flags.noMenu)
		imgui.Checkbox("No move", &window.flags.noMove)
		imgui.SameLineV(150, -1)
		imgui.Checkbox("No resize", &window.flags.noResize)
		imgui.SameLineV(300, -1)
		imgui.Checkbox("No collapse", &window.flags.noCollapse)
		imgui.Checkbox("No close", &window.noClose)
		imgui.SameLineV(150, -1)
		imgui.Checkbox("No nav", &window.flags.noNav)
		imgui.SameLineV(300, -1)
		imgui.Checkbox("No background", &window.flags.noBackground)
		imgui.Checkbox("No bring to front", &window.flags.noBringToFront)
	}

	// All demo contents
	window.widgets.show()
	window.layout.show()
	window.popups.show()
	window.columns.show()
	window.misc.show()

	// End of ShowDemoWindow()
	imgui.End()
}

func showMainMenuFile() {

	/*
	   ImGui::MenuItem("(dummy menu)", NULL, false, false);
	   if (ImGui::MenuItem("New")) {}
	   if (ImGui::MenuItem("Open", "Ctrl+O")) {}
	   if (ImGui::BeginMenu("Open Recent"))
	   {
	       ImGui::MenuItem("fish_hat.c");
	       ImGui::MenuItem("fish_hat.inl");
	       ImGui::MenuItem("fish_hat.h");
	       if (ImGui::BeginMenu("More.."))
	       {
	           ImGui::MenuItem("Hello");
	           ImGui::MenuItem("Sailor");
	           if (ImGui::BeginMenu("Recurse.."))
	           {
	               ShowExampleMenuFile();
	               ImGui::EndMenu();
	           }
	           ImGui::EndMenu();
	       }
	       ImGui::EndMenu();
	   }
	   if (ImGui::MenuItem("Save", "Ctrl+S")) {}
	   if (ImGui::MenuItem("Save As..")) {}

	   ImGui::Separator();
	   if (ImGui::BeginMenu("Options"))
	   {
	       static bool enabled = true;
	       ImGui::MenuItem("Enabled", "", &enabled);
	       ImGui::BeginChild("child", ImVec2(0, 60), true);
	       for (int i = 0; i < 10; i++)
	           ImGui::Text("Scrolling Text %d", i);
	       ImGui::EndChild();
	       static float f = 0.5f;
	       static int n = 0;
	       ImGui::SliderFloat("Value", &f, 0.0f, 1.0f);
	       ImGui::InputFloat("Input", &f, 0.1f);
	       ImGui::Combo("Combo", &n, "Yes\0No\0Maybe\0\0");
	       ImGui::EndMenu();
	   }

	   if (ImGui::BeginMenu("Colors"))
	   {
	       float sz = ImGui::GetTextLineHeight();
	       for (int i = 0; i < ImGuiCol_COUNT; i++)
	       {
	           const char* name = ImGui::GetStyleColorName((ImGuiCol)i);
	           ImVec2 p = ImGui::GetCursorScreenPos();
	           ImGui::GetWindowDrawList()->AddRectFilled(p, ImVec2(p.x+sz, p.y+sz), ImGui::GetColorU32((ImGuiCol)i));
	           ImGui::Dummy(ImVec2(sz, sz));
	           ImGui::SameLine();
	           ImGui::MenuItem(name);
	       }
	       ImGui::EndMenu();
	   }

	   // Here we demonstrate appending again to the "Options" menu (which we already created above)
	   // Of course in this demo it is a little bit silly that this function calls BeginMenu("Options") twice.
	   // In a real code-base using it would make senses to use this feature from very different code locations.
	   if (ImGui::BeginMenu("Options")) // <-- Append!
	   {
	       static bool b = true;
	       ImGui::Checkbox("SomeOption", &b);
	       ImGui::EndMenu();
	   }

	   if (ImGui::BeginMenu("Disabled", false)) // Disabled
	   {
	       IM_ASSERT(0);
	   }
	   if (ImGui::MenuItem("Checked", NULL, true)) {}
	   if (ImGui::MenuItem("Quit", "Alt+F4")) {}

	*/

}

func showUserGuide() {
	bulletText("Double-click on title bar to collapse window.")
	bulletText("Click and drag on lower corner to resize window\n(double-click to auto fit window to its contents).")
	bulletText("CTRL+Click on a slider or drag box to input value as text.")
	bulletText("TAB/SHIFT+TAB to cycle through keyboard editable fields.")

	// MISSING: Allow FontUserScaling

	bulletText("While inputing text:\n")
	imgui.Indent()
	bulletText("CTRL+Left/Right to word jump.")
	bulletText("CTRL+A or double-click to select all.")
	bulletText("CTRL+X/C/V to use clipboard cut/copy/paste.")
	bulletText("CTRL+Z,CTRL+Y to undo/redo.")
	bulletText("ESCAPE to revert.")
	bulletText("You can apply arithmetic operators +,*,/ on numerical values.\nUse +- to subtract.")
	imgui.Unindent()
	bulletText("With keyboard navigation enabled:")
	imgui.Indent()
	bulletText("Arrow keys to navigate.")
	bulletText("Space to activate a widget.")
	bulletText("Return to input text into a widget.")
	bulletText("Escape to deactivate a widget, close popup, exit child window.")
	bulletText("Alt to jump to the menu layer of a window.")
	bulletText("CTRL+Tab to select a window.")
	imgui.Unindent()
}

type widgets struct {
	buttonClicked int
	check         bool
	radio         int
}

// nolint: nestif
func (widgets *widgets) show() {
	if !imgui.CollapsingHeader("Widgets") {
		return
	}

	if imgui.TreeNode("Basic") {
		if imgui.Button("Button") {
			widgets.buttonClicked++
		}
		if widgets.buttonClicked&1 != 0 {
			imgui.SameLine()
			imgui.Text("Thanks for clicking me!")
		}

		imgui.Checkbox("checkbox", &widgets.check)

		if imgui.RadioButton("radio a", widgets.radio == 0) {
			widgets.radio = 0
		}
		imgui.SameLine()
		if imgui.RadioButton("radio b", widgets.radio == 1) {
			widgets.radio = 1
		}
		imgui.SameLine()
		if imgui.RadioButton("radio c", widgets.radio == 2) {
			widgets.radio = 2
		}

		imgui.TreePop()
	}
}

type layout struct {
}

func (layout *layout) show() {

}

type popups struct {
}

func (popups *popups) show() {

}

type columns struct {
}

func (columns *columns) show() {

}

type misc struct {
}

func (misc *misc) show() {

}

package ui

import (
	"fmt"

	"github.com/jesseduffield/gocui"
	"github.com/youpele52/lazysetup/pkg/colors"
	"github.com/youpele52/lazysetup/pkg/constants"
	"github.com/youpele52/lazysetup/pkg/models"
)

// layoutMultiPanel renders the four-panel layout used in PageMultiPanel
// Left side (stacked vertically):
//   - Panel 1: Package Manager selection (top)
//   - Panel 2: Action selection (middle) - Install/Update/Delete
//   - Panel 3: Tools selection (bottom)
//
// Right side (full height):
//   - Panel 0: Status/results display (read-only)
//
// Bottom status bar shows keybinding hints
func layoutMultiPanel(g *gocui.Gui, state *models.State, maxX, maxY int) error {
	// Delete old views
	for _, viewName := range []string{constants.ViewMenu, constants.ViewResult, constants.ViewTools, "installing", "results"} {
		g.DeleteView(viewName)
	}

	leftPanelWidth := maxX / 3
	panelHeight := maxY - 3
	activePanel := state.GetActivePanel()

	// Calculate heights for left panels (3 panels stacked)
	packageManagerHeight := panelHeight / 3
	actionHeight := panelHeight / 4
	toolsStartY := packageManagerHeight + actionHeight + 2

	// Panel 1: Package Manager (top-left)
	if v, err := g.SetView(constants.PanelPackageManager, 0, 0, leftPanelWidth, packageManagerHeight); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Wrap = true
		v.SelBgColor = colors.HighlightBg
		v.SelFgColor = colors.HighlightFg
	}

	if v, err := g.View(constants.PanelPackageManager); err == nil {
		v.Title = "[1]-" + constants.TitlePackageManager
		if activePanel == models.PanelPackageManager {
			v.FgColor = colors.ActiveBorderColor
		} else {
			v.FgColor = colors.TextPrimary
		}
		v.Clear()
		for i, method := range state.InstallMethods {
			marker := constants.RadioUnselected
			if i == state.SelectedIndex {
				marker = constants.RadioSelected
			}

			if activePanel == models.PanelPackageManager {
				if i == state.SelectedIndex {
					fmt.Fprintf(v, "%s%s %s%s\n", colors.ANSIMagenta, marker, method, colors.ANSIReset)
				} else {
					fmt.Fprintf(v, "%s%s %s%s\n", colors.ANSIGreen, marker, method, colors.ANSIReset)
				}
			} else {
				if i == state.SelectedIndex {
					fmt.Fprintf(v, "%s%s %s%s\n", colors.ANSIMagenta, marker, method, colors.ANSIReset)
				} else {
					fmt.Fprintf(v, "%s %s\n", marker, method)
				}
			}
		}
		if activePanel == models.PanelPackageManager {
			v.SetCursor(0, state.SelectedIndex)
		}
	}

	// Panel 2: Action (middle-left)
	actions := []string{"Install", "Update", "Delete"}
	if v, err := g.SetView(constants.PanelAction, 0, packageManagerHeight+1, leftPanelWidth, packageManagerHeight+actionHeight+1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Wrap = true
		v.SelBgColor = colors.HighlightBg
		v.SelFgColor = colors.HighlightFg
	}

	if v, err := g.View(constants.PanelAction); err == nil {
		v.Title = "[2]-" + constants.TitleAction
		if activePanel == models.PanelAction {
			v.FgColor = colors.ActiveBorderColor
		} else {
			v.FgColor = colors.TextPrimary
		}
		v.Clear()
		for i, action := range actions {
			marker := constants.RadioUnselected
			if i == state.ActionIndex {
				marker = constants.RadioSelected
			}

			if activePanel == models.PanelAction {
				if i == state.ActionIndex {
					fmt.Fprintf(v, "%s%s %s%s\n", colors.ANSIMagenta, marker, action, colors.ANSIReset)
				} else {
					fmt.Fprintf(v, "%s%s %s%s\n", colors.ANSIGreen, marker, action, colors.ANSIReset)
				}
			} else {
				if i == state.ActionIndex {
					fmt.Fprintf(v, "%s%s %s%s\n", colors.ANSIMagenta, marker, action, colors.ANSIReset)
				} else {
					fmt.Fprintf(v, "%s %s\n", marker, action)
				}
			}
		}
		if activePanel == models.PanelAction {
			v.SetCursor(0, state.ActionIndex)
		}
	}

	// Panel 3: Tools Selection (bottom-left)
	if v, err := g.SetView(constants.PanelTools, 0, toolsStartY, leftPanelWidth, panelHeight); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Wrap = true
		v.SelBgColor = colors.HighlightBg
		v.SelFgColor = colors.HighlightFg
	}

	if v, err := g.View(constants.PanelTools); err == nil {
		v.Title = "[3]-" + constants.TitleTools
		if activePanel == models.PanelTools {
			v.FgColor = colors.ActiveBorderColor
		} else {
			v.FgColor = colors.TextPrimary
		}
		v.Clear()
		for i, tool := range state.Tools {
			selected := state.SelectedTools[tool]
			var marker string
			if selected {
				marker = constants.CheckboxSelected
			} else {
				marker = constants.CheckboxUnselected
			}

			if activePanel == models.PanelTools {
				if i == state.ToolsIndex {
					fmt.Fprintf(v, "%s%s %s%s\n", colors.ANSIMagenta, marker, tool, colors.ANSIReset)
				} else if selected {
					fmt.Fprintf(v, "%s%s %s%s\n", colors.ANSIMagenta, marker, tool, colors.ANSIReset)
				} else {
					fmt.Fprintf(v, "%s%s %s%s\n", colors.ANSIGreen, marker, tool, colors.ANSIReset)
				}
			} else {
				if selected {
					fmt.Fprintf(v, "%s%s %s%s\n", colors.ANSIMagenta, marker, tool, colors.ANSIReset)
				} else {
					fmt.Fprintf(v, "%s %s\n", marker, tool)
				}
			}
		}
		if activePanel == models.PanelTools {
			v.SetCursor(0, state.ToolsIndex)
		}
	}

	// Panel 0: Status/Results (right - full height, read-only)
	if v, err := g.SetView(constants.PanelProgress, leftPanelWidth+1, 0, maxX-1, panelHeight); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Wrap = true
		v.Highlight = false
	}

	if v, err := g.View(constants.PanelProgress); err == nil {
		v.Title = "[0]-" + constants.TitleStatus
		if activePanel == models.PanelStatus {
			v.FgColor = colors.ActiveBorderColor
		} else {
			v.FgColor = colors.TextPrimary
		}
		v.Clear()
		installationDone := state.GetInstallationDone()
		installStartTime := state.GetInstallStartTime()
		if installStartTime > 0 && !installationDone {
			// Show installation progress
			spinnerFrame := state.GetSpinnerFrame()
			installOutput := state.GetInstallOutput()
			selectedMethod := state.GetSelectedMethod()
			currentTool := state.GetCurrentTool()
			installingIndex := state.GetInstallingIndex()
			selectedTools := state.GetSelectedTools()
			selectedAction := state.GetSelectedAction()
			message := BuildInstallationProgressMessage(selectedMethod, currentTool, installingIndex, len(selectedTools), installationDone, spinnerFrame, installOutput, selectedAction)
			fmt.Fprint(v, message)
		} else if installationDone {
			// Show results
			results := state.GetInstallResults()
			message := BuildInstallationResultsMessage(results)
			fmt.Fprint(v, message)
		} else {
			// Show logo by default
			fmt.Fprint(v, constants.Logo)
		}
	}

	// Status bar at bottom
	if v, err := g.SetView("status_bar", 0, panelHeight+1, maxX-1, maxY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.FgColor = colors.TextPrimary
		v.Wrap = false
	}

	if v, err := g.View("status_bar"); err == nil {
		v.Clear()
		fmt.Fprintf(v, "Tab/Numbers: Switch panels | ↑↓: Navigate | Space: Toggle | Enter: Confirm | Esc: Back | Ctrl+C: Quit")
	}

	return nil
}

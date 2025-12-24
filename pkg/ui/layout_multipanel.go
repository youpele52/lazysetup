package ui

import (
	"fmt"

	"github.com/jesseduffield/gocui"
	"github.com/youpele52/lazysetup/pkg/colors"
	"github.com/youpele52/lazysetup/pkg/constants"
	"github.com/youpele52/lazysetup/pkg/models"
)

// layoutMultiPanel renders the three-panel layout used in PageMultiPanel
// Panel 0 (right): Installation method selection
// Panel 1 (left, bottom): Tool selection with checkboxes
// Panel 2 (right, full height): Status/results display
// Bottom status bar shows keybinding hints
func layoutMultiPanel(g *gocui.Gui, state *models.State, maxX, maxY int) error {
	// Delete old views
	for _, viewName := range []string{constants.ViewMenu, constants.ViewResult, constants.ViewTools, "installing", "results"} {
		g.DeleteView(viewName)
	}

	leftPanelWidth := maxX / 3
	panelHeight := maxY - 3

	// Panel 1: Installation Methods (top-left)
	installationHeight := panelHeight / 2
	activePanel := state.GetActivePanel()
	if v, err := g.SetView(constants.PanelInstallation, 0, 0, leftPanelWidth, installationHeight); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Wrap = true
		v.SelBgColor = colors.HighlightBg
		v.SelFgColor = colors.HighlightFg
	}

	// Update title every frame based on active panel
	if v, err := g.View(constants.PanelInstallation); err == nil {
		v.Title = "[1]-" + constants.TitleInstallation
	}

	if v, err := g.View(constants.PanelInstallation); err == nil {
		v.Clear()
		for i, method := range state.InstallMethods {
			marker := constants.RadioUnselected
			if i == state.SelectedIndex {
				marker = constants.RadioSelected
			}

			if activePanel == models.PanelInstallation {
				// Active panel: green text by default
				if i == state.SelectedIndex {
					// Cursor position: magenta text (will have magenta background from highlight)
					fmt.Fprintf(v, "%s%s %s%s\n", colors.ANSIMagenta, marker, method, colors.ANSIReset)
				} else {
					// Unselected item: green text
					fmt.Fprintf(v, "%s%s %s%s\n", colors.ANSIGreen, marker, method, colors.ANSIReset)
				}
			} else {
				// Inactive panel
				if i == state.SelectedIndex {
					fmt.Fprintf(v, "%s%s %s%s\n", colors.ANSIMagenta, marker, method, colors.ANSIReset)
				} else {
					fmt.Fprintf(v, "%s %s\n", marker, method)
				}
			}
		}
		if activePanel == models.PanelInstallation {
			v.SetCursor(0, state.SelectedIndex)
		}
	}

	// Panel 2: Tools Selection (bottom-left)
	if v, err := g.SetView(constants.PanelTools, 0, installationHeight+1, leftPanelWidth, panelHeight); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "[2]-" + constants.TitleToolSelection
		v.Wrap = true
		v.SelBgColor = colors.HighlightBg
		v.SelFgColor = colors.HighlightFg
	}

	if v, err := g.View(constants.PanelTools); err == nil {
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
				// Active panel: green text by default
				if i == state.ToolsIndex {
					// Cursor position: magenta text (will have magenta background from highlight)
					fmt.Fprintf(v, "%s%s %s%s\n", colors.ANSIMagenta, marker, tool, colors.ANSIReset)
				} else if selected {
					// Selected item: magenta text
					fmt.Fprintf(v, "%s%s %s%s\n", colors.ANSIMagenta, marker, tool, colors.ANSIReset)
				} else {
					// Unselected item: green text
					fmt.Fprintf(v, "%s%s %s%s\n", colors.ANSIGreen, marker, tool, colors.ANSIReset)
				}
			} else {
				// Inactive panel
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
		v.Title = "[0]-" + constants.PanelStatus
		v.Wrap = true
		v.Highlight = false
	}

	// Update border color every frame based on active panel
	if v, err := g.View(constants.PanelProgress); err == nil {
		if activePanel == models.PanelProgress {
			v.FgColor = colors.ActiveBorderColor
		} else {
			v.FgColor = colors.TextPrimary
		}
	}

	if v, err := g.View(constants.PanelProgress); err == nil {
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
			message := BuildInstallationProgressMessage(selectedMethod, currentTool, installingIndex, len(selectedTools), installationDone, spinnerFrame, installOutput)
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

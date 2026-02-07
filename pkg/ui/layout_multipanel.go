package ui

import (
	"fmt"
	"time"

	"github.com/jesseduffield/gocui"
	"github.com/youpele52/lazysetup/pkg/colors"
	"github.com/youpele52/lazysetup/pkg/constants"
	"github.com/youpele52/lazysetup/pkg/models"
	"github.com/youpele52/lazysetup/pkg/ui/messages"
)

// PanelHeights holds calculated heights for each panel
type PanelHeights struct {
	PackageManagerHeight int
	ActionHeight         int
	ToolsHeight          int
}

// calculatePanelHeights computes optimal panel heights based on content needs
// Uses responsive design with breakpoints for different terminal sizes:
// - Large (45+ rows): Spacious mode - all panels get ideal space
// - Medium (25-44 rows): Balanced mode - compressed but comfortable
// - Small (<25 rows): Compact mode - minimize top panels, maximize Tools
// Package Manager: 9 items, Action: 4 items, Tools: 27 items
func calculatePanelHeights(totalHeight int, minHeight int) PanelHeights {
	const (
		packageManagerItems = 9
		actionItems         = 4
		toolsItems          = 27
		spacing             = 2

		// Responsive breakpoints
		largeTerminal  = 45 // Spacious mode
		mediumTerminal = 25 // Balanced mode
		// Below mediumTerminal = Compact mode
	)

	// Ideal heights (items + borders/padding)
	packageManagerIdeal := packageManagerItems + 2 // 11 rows
	actionIdeal := actionItems + 2                  // 6 rows

	var packageManagerHeight, actionHeight, toolsHeight int

	if totalHeight >= largeTerminal {
		// SPACIOUS MODE (45+ rows)
		// Give all panels ideal/comfortable space
		packageManagerHeight = packageManagerIdeal
		actionHeight = actionIdeal
		toolsHeight = totalHeight - packageManagerHeight - actionHeight - spacing

	} else if totalHeight >= mediumTerminal {
		// BALANCED MODE (25-44 rows)
		// Compress top panels moderately, give Tools good space
		packageManagerHeight = min(packageManagerIdeal, totalHeight/4)
		actionHeight = min(actionIdeal, totalHeight/6)

		// Give remaining space to Tools panel
		remaining := totalHeight - packageManagerHeight - actionHeight - spacing
		if remaining < minHeight {
			// Adjust if remaining is too small
			packageManagerHeight = totalHeight / 3
			actionHeight = totalHeight / 4
			remaining = totalHeight - packageManagerHeight - actionHeight - spacing
		}
		toolsHeight = remaining

	} else {
		// COMPACT MODE (<25 rows)
		// Minimize top panels, maximize Tools (has 27 items - most content)
		// Package Manager: just enough for 9 items
		packageManagerHeight = max(minHeight, min(7, totalHeight/4))

		// Action: minimal (only 4 items)
		actionHeight = max(minHeight, min(6, totalHeight/6))

		// Tools: all remaining space (27 items need it)
		toolsHeight = totalHeight - packageManagerHeight - actionHeight - spacing

		// Safety check: ensure minimum heights
		if toolsHeight < minHeight {
			// Extreme case - fall back to equal distribution
			packageManagerHeight = totalHeight / 3
			actionHeight = totalHeight / 3
			toolsHeight = totalHeight - packageManagerHeight - actionHeight - spacing
		}
	}

	return PanelHeights{
		PackageManagerHeight: packageManagerHeight,
		ActionHeight:         actionHeight,
		ToolsHeight:          toolsHeight,
	}
}

// min returns the minimum of two integers
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// max returns the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

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

	// Calculate minimum space requirements
	// 3 panels minimum + 2 spacing lines between panels + 1 status bar
	minPanelHeight := 5 // Minimum 5 rows per panel (title + borders + at least 3 content lines)
	panelSpacing := 2   // Space between stacked panels
	statusBarHeight := 2
	requiredHeight := (minPanelHeight * 3) + panelSpacing + statusBarHeight

	// Check if terminal is too small to display all panels
	if maxY < requiredHeight {
		// Terminal too small - show error in a simplified view
		return renderTerminalTooSmallView(g, maxX, maxY, minPanelHeight, requiredHeight)
	}

	// Calculate heights for left panels using proportional allocation
	// This gives more space to panels with more items (Tools has 27 items)
	heights := calculatePanelHeights(panelHeight, minPanelHeight)
	packageManagerHeight := heights.PackageManagerHeight
	actionHeight := heights.ActionHeight
	toolsHeight := heights.ToolsHeight
	toolsStartY := packageManagerHeight + actionHeight + panelSpacing

	// Render left-side panels
	if err := renderPackageManagerPanel(PackageManagerParams{
		PanelParams: PanelParams{
			Gui:            g,
			State:          state,
			ActivePanel:    activePanel,
			LeftPanelWidth: leftPanelWidth,
		},
		Height: packageManagerHeight,
	}); err != nil {
		return err
	}
	if err := renderActionPanel(ActionPanelParams{
		PanelParams: PanelParams{
			Gui:            g,
			State:          state,
			ActivePanel:    activePanel,
			LeftPanelWidth: leftPanelWidth,
		},
		PackageManagerY: packageManagerHeight,
		ActionHeight:    actionHeight,
	}); err != nil {
		return err
	}
	if err := renderToolsPanel(ToolsPanelParams{
		PanelParams: PanelParams{
			Gui:            g,
			State:          state,
			ActivePanel:    activePanel,
			LeftPanelWidth: leftPanelWidth,
		},
		ToolsStartY: toolsStartY,
		ToolsHeight: toolsHeight, // FIXED: was PanelHeight (wrong boundary)
	}); err != nil {
		return err
	}

	// Panel 0: Status/Results (right - full height, read-only, scrollable)
	if v, err := g.SetView(constants.PanelProgress, leftPanelWidth+1, 0, maxX-1, panelHeight); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Wrap = true
		v.Highlight = false
		v.Autoscroll = false
	}

	if v, err := g.View(constants.PanelProgress); err == nil {
		v.Title = "[0]-" + constants.TitleStatus
		if activePanel == models.PanelStatus {
			v.FgColor = colors.ActiveBorderColor
		} else {
			v.FgColor = colors.TextPrimary
		}

		installationDone := state.GetInstallationDone()
		installStartTime := state.GetInstallStartTime()
		results := state.GetInstallResults()
		lastRenderedCount := state.LastRenderedResultCount
		completionTime := state.ActionCompletionTime
		errorMsg := state.Error

		// Show update notification if available (hide after 10 seconds)
		if state.UpdateMessage != "" {
			elapsed := time.Now().Unix() - state.UpdateMessageTime
			if elapsed < 10 {
				v.Clear()
				fmt.Fprintf(v, "%s%s%s\n\n", colors.ANSIYellow, state.UpdateMessage, colors.ANSIReset)
				fmt.Fprint(v, constants.Logo)
				return nil
			} else {
				// Clear message after timeout
				state.UpdateMessage = ""
			}
		}

		// Show validation errors
		if errorMsg != "" {
			v.Clear()
			fmt.Fprintf(v, "%s%s%s\n", colors.ANSIRed, errorMsg, colors.ANSIReset)
			return nil
		}

		// Auto-clear results after 40 seconds
		if installationDone && completionTime > 0 {
			elapsed := time.Now().Unix() - completionTime
			if elapsed >= 40 {
				// Clear and reset
				v.Clear()
				state.InstallationDone = false
				state.ActionCompletionTime = 0
				state.LastRenderedResultCount = 0
				fmt.Fprint(v, constants.Logo)
				return nil
			}
		}

		if installStartTime > 0 && !installationDone {
			// Show installation progress - clear and update every frame for spinner
			v.Clear()
			params := messages.ProgressMessageParams{
				SelectedMethod:   state.GetSelectedMethod(),
				CurrentTool:      state.GetCurrentTool(),
				InstallingIndex:  state.GetInstallingIndex(),
				TotalTools:       len(state.GetSelectedTools()),
				InstallationDone: installationDone,
				SpinnerFrame:     state.GetSpinnerFrame(),
				InstallOutput:    state.GetInstallOutput(),
				Action:           state.GetSelectedAction(),
			}
			message := messages.BuildInstallationProgressMessage(params)
			fmt.Fprint(v, message)
		} else if installationDone && len(results) > lastRenderedCount {
			// Append only new results (rolling credits style)
			// Latest results at top, oldest at bottom
			selectedAction := state.GetSelectedAction()
			newResults := results[lastRenderedCount:]
			message := messages.BuildNewResultsMessage(newResults, selectedAction)
			fmt.Fprint(v, message)
			state.LastRenderedResultCount = len(results)
		} else if len(results) == 0 {
			// Show logo only if no results exist
			v.Clear()

			// Show Nix disclaimer if Nix is selected as package manager
			if state.GetSelectedMethod() == "Nix" {
				fmt.Fprintf(v, "%s%s%s\n\n", colors.ANSIYellow, constants.NixDisclaimer, colors.ANSIReset)
			}

			fmt.Fprint(v, constants.Logo)
		}
		// If results exist, preserve them for scrolling
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
		if state.GetIsSearchMode() {
			// Search mode status
			fmt.Fprintf(v, "Type to filter | ↑↓: Nav | Space: Toggle | /: Exit | Esc: Exit | ⏎: Confirm")
		} else if state.GetActivePanel() == models.PanelTools {
			// Normal tools panel status
			if state.UpdateAvailable {
				fmt.Fprintf(v, "Tab/0-3: Panels | ↑↓: Nav | g/w: First | G/s: Last | /: Search | Space: Toggle | ⏎: Confirm | C: Clear | U: Update | Ctrl+C: Quit")
			} else {
				fmt.Fprintf(v, "Tab/0-3: Panels | ↑↓: Nav | g/w: First | G/s: Last | /: Search | Space: Toggle | ⏎: Confirm | C: Clear | Esc: Back | Ctrl+C: Quit")
			}
		} else {
			// Other panels status
			if state.UpdateAvailable {
				fmt.Fprintf(v, "Tab/0-3: Panels | ↑↓: Nav | g/w: First | G/s: Last | Space: Toggle | ⏎: Confirm | C: Clear | U: Update | Ctrl+C: Quit")
			} else {
				fmt.Fprintf(v, "Tab/0-3: Panels | ↑↓: Nav | g/w: First | G/s: Last | Space: Toggle | ⏎: Confirm | C: Clear | Esc: Back | Ctrl+C: Quit")
			}
		}
	}

	// Render sudo confirmation popup if needed
	if state.GetShowSudoConfirm() {
		if err := renderSudoConfirmPopup(g, maxX, maxY, state); err != nil {
			return err
		}
	} else {
		// Delete popup if not needed
		g.DeleteView(constants.PopupConfirm)
	}

	return nil
}

// renderTerminalTooSmallView displays an error message when the terminal
// is too small to accommodate the minimum panel requirements.
// Clears all panels and shows only an error message with resize instructions.
func renderTerminalTooSmallView(g *gocui.Gui, maxX, maxY, minPanelHeight, requiredHeight int) error {
	// Clear all existing views
	for _, viewName := range []string{
		constants.ViewMenu, constants.ViewResult, constants.ViewTools,
		"installing", "results", constants.PanelPackageManager,
		constants.PanelAction, constants.PanelTools, "status_bar",
		constants.PanelProgress,
	} {
		g.DeleteView(viewName)
	}

	// Create a full-screen error view
	viewName := "terminal_too_small"
	if v, err := g.SetView(viewName, 0, 0, maxX-1, maxY-1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Wrap = true
		v.Title = "Error"
		v.FgColor = colors.FailureColor
	}

	if v, err := g.View(viewName); err == nil {
		v.Clear()
		fmt.Fprintf(v, "\n\n%s", colors.ANSIRed)
		fmt.Fprintf(v, "  ╔════════════════════════════════════════════════════════╗\n")
		fmt.Fprintf(v, "  ║                                                        ║\n")
		fmt.Fprintf(v, "  ║     TERMINAL WINDOW TOO SMALL                          ║\n")
		fmt.Fprintf(v, "  ║                                                        ║\n")
		fmt.Fprintf(v, "  ║  Current size: %dx%d rows                              ║\n", maxX, maxY)
		fmt.Fprintf(v, "  ║  Minimum required: %d rows                             ║\n", requiredHeight)
		fmt.Fprintf(v, "  ║                                                        ║\n")
		fmt.Fprintf(v, "  ║  Please resize your terminal window to at least        ║\n")
		fmt.Fprintf(v, "  ║  80 columns x %d rows to use lazysetup.                ║\n", requiredHeight)
		fmt.Fprintf(v, "  ║                                                        ║\n")
		fmt.Fprintf(v, "  ║  Press Ctrl+C to quit.                                 ║\n")
		fmt.Fprintf(v, "  ║                                                        ║\n")
		fmt.Fprintf(v, "  ╚════════════════════════════════════════════════════════╝\n")
		fmt.Fprintf(v, "%s", colors.ANSIReset)
	}

	return nil
}

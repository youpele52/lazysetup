package ui

import (
	"fmt"
	"time"

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
		PanelHeight: panelHeight,
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
			params := ProgressMessageParams{
				SelectedMethod:   state.GetSelectedMethod(),
				CurrentTool:      state.GetCurrentTool(),
				InstallingIndex:  state.GetInstallingIndex(),
				TotalTools:       len(state.GetSelectedTools()),
				InstallationDone: installationDone,
				SpinnerFrame:     state.GetSpinnerFrame(),
				InstallOutput:    state.GetInstallOutput(),
				Action:           state.GetSelectedAction(),
			}
			message := BuildInstallationProgressMessage(params)
			fmt.Fprint(v, message)
		} else if installationDone && len(results) > lastRenderedCount {
			// Append only new results (rolling credits style)
			// Latest results at top, oldest at bottom
			selectedAction := state.GetSelectedAction()
			newResults := results[lastRenderedCount:]
			message := BuildNewResultsMessage(newResults, selectedAction)
			fmt.Fprint(v, message)
			state.LastRenderedResultCount = len(results)
		} else if len(results) == 0 {
			// Show logo only if no results exist
			v.Clear()
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
		if state.UpdateAvailable {
			fmt.Fprintf(v, "Tab/0-3: Panels | ↑↓: Nav | Space: Toggle | ⏎: Confirm | C: Clear | U: Update | Ctrl+C: Quit")
		} else {
			fmt.Fprintf(v, "Tab/0-3: Panels | ↑↓: Nav | Space: Toggle | ⏎: Confirm | C: Clear | Esc: Back | Ctrl+C: Quit")
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

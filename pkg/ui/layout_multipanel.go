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
		} else if installationDone {
			// Show results
			results := state.GetInstallResults()
			selectedAction := state.GetSelectedAction()
			message := BuildInstallationResultsMessage(results, selectedAction)
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

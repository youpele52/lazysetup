package ui

import (
	"fmt"

	"github.com/jesseduffield/gocui"
	"github.com/youpele52/lazysetup/pkg/colors"
	"github.com/youpele52/lazysetup/pkg/constants"
	"github.com/youpele52/lazysetup/pkg/models"
)

// Layout returns a gocui layout function that renders the appropriate UI based on current page
// It dispatches to specific layout functions (menu, selection, tools, installing, results, multipanel)
// based on state.CurrentPage value
func Layout(state *models.State) func(*gocui.Gui) error {
	return func(g *gocui.Gui) error {
		maxX, maxY := g.Size()

		switch state.CurrentPage {
		case models.PageMenu:
			return layoutMenuPage(g, state, maxX, maxY)
		case models.PageSelection:
			return layoutSelectionPage(g, state, maxX, maxY)
		case models.PageTools:
			return layoutToolsPage(g, state, maxX, maxY)
		case models.PageInstalling:
			return layoutInstallingPage(g, state, maxX, maxY)
		case models.PageResults:
			return layoutResultsPage(g, state, maxX, maxY)
		case models.PageMultiPanel:
			return layoutMultiPanel(g, state, maxX, maxY)
		default:
			return layoutMenuPage(g, state, maxX, maxY)
		}
	}
}

func layoutMenuPage(g *gocui.Gui, state *models.State, maxX, maxY int) error {
	if err := g.DeleteView(constants.ViewResult); err != nil && err != gocui.ErrUnknownView {
		return err
	}

	contentHeight := len(state.InstallMethods) + 2
	startY := maxY/2 - contentHeight/2
	endY := startY + contentHeight

	if v, err := g.SetView(constants.ViewMenu, maxX/4, startY, 3*maxX/4, endY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = constants.TitleInstallation
		v.Highlight = true
		v.SelBgColor = colors.HighlightBg
		v.SelFgColor = colors.HighlightFg
		v.FgColor = colors.TextPrimary
		v.Wrap = true

		if err := g.SetCurrentView(constants.ViewMenu); err != nil {
			return err
		}
	}

	if v, err := g.View(constants.ViewMenu); err == nil {
		v.Clear()
		for i, method := range state.InstallMethods {
			if i == state.SelectedIndex {
				fmt.Fprintf(v, "%s %s\n", constants.RadioSelected, method)
			} else {
				fmt.Fprintf(v, "%s %s\n", constants.RadioUnselected, method)
			}
		}
		v.SetCursor(0, state.SelectedIndex)
	}

	return nil
}

func layoutSelectionPage(g *gocui.Gui, state *models.State, maxX, maxY int) error {
	if err := g.DeleteView(constants.ViewMenu); err != nil && err != gocui.ErrUnknownView {
		return err
	}

	contentHeight := 4
	if state.CheckStatus != "" {
		contentHeight++
	}
	if state.Error != "" {
		contentHeight += 2
	}
	startY := maxY/2 - contentHeight/2
	endY := startY + contentHeight

	if v, err := g.SetView(constants.ViewResult, maxX/4, startY, 3*maxX/4, endY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = constants.TitleSelection
		v.FgColor = colors.TextPrimary
		v.Wrap = true

		if err := g.SetCurrentView(constants.ViewResult); err != nil {
			return err
		}
	}

	if v, err := g.View(constants.ViewResult); err == nil {
		v.Clear()
		fmt.Fprintf(v, constants.MessageSelected, state.SelectedMethod)
		if state.CheckStatus != "" {
			fmt.Fprintf(v, "Status: %s\n", state.CheckStatus)
		}
		if state.Error != "" {
			fmt.Fprintf(v, "Error: %s\n", state.Error)
		}
	}

	return nil
}

func layoutToolsPage(g *gocui.Gui, state *models.State, maxX, maxY int) error {
	if err := g.DeleteView(constants.ViewMenu); err != nil && err != gocui.ErrUnknownView {
		return err
	}

	contentHeight := len(state.Tools) + 2
	startY := maxY/2 - contentHeight/2
	endY := startY + contentHeight

	if v, err := g.SetView(constants.ViewTools, maxX/4, startY, 3*maxX/4, endY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = constants.TitleToolSelection
		v.Highlight = true
		v.SelBgColor = colors.HighlightBg
		v.SelFgColor = colors.HighlightFg
		v.FgColor = colors.TextPrimary
		v.Wrap = true

		if err := g.SetCurrentView(constants.ViewTools); err != nil {
			return err
		}
	}

	if v, err := g.View(constants.ViewTools); err == nil {
		v.Clear()
		for _, tool := range state.Tools {
			selected := state.SelectedTools[tool]
			var marker string
			if selected {
				marker = constants.CheckboxSelected
			} else {
				marker = constants.CheckboxUnselected
			}
			fmt.Fprintf(v, "%s %s\n", marker, tool)
		}
		v.SetCursor(0, state.ToolsIndex)
	}

	return nil
}

func layoutInstallingPage(g *gocui.Gui, state *models.State, maxX, maxY int) error {
	if err := g.DeleteView(constants.ViewTools); err != nil && err != gocui.ErrUnknownView {
		return err
	}

	contentHeight := 10
	if len(state.InstallOutput) > 0 {
		contentHeight += 3
	}
	startY := maxY/2 - contentHeight/2
	endY := startY + contentHeight

	if v, err := g.SetView(constants.ViewInstalling, maxX/4, startY, 3*maxX/4, endY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = constants.TitleInstalling
		v.FgColor = colors.TextPrimary
		v.Wrap = true

		if err := g.SetCurrentView(constants.ViewInstalling); err != nil {
			return err
		}
	}

	if v, err := g.View(constants.ViewInstalling); err == nil {
		v.Clear()
		message := BuildInstallationProgressMessage(state.SelectedMethod, state.CurrentTool, state.InstallingIndex, len(state.Tools), state.InstallationDone, state.SpinnerFrame, state.InstallOutput)
		fmt.Fprint(v, message)
	}

	return nil
}

func layoutResultsPage(g *gocui.Gui, state *models.State, maxX, maxY int) error {
	if err := g.DeleteView(constants.ViewInstalling); err != nil && err != gocui.ErrUnknownView {
		return err
	}

	contentHeight := 4 + len(state.InstallResults) + 2
	for _, result := range state.InstallResults {
		if !result.Success && result.Error != "" {
			contentHeight++
		}
	}
	startY := maxY/2 - contentHeight/2
	endY := startY + contentHeight

	if v, err := g.SetView(constants.ViewResults, maxX/4, startY, 3*maxX/4, endY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = constants.ResultsSummaryTitle
		v.FgColor = colors.TextPrimary
		v.Wrap = true

		if err := g.SetCurrentView(constants.ViewResults); err != nil {
			return err
		}
	}

	if v, err := g.View(constants.ViewResults); err == nil {
		v.Clear()
		message := BuildInstallationResultsMessage(state.InstallResults)
		fmt.Fprint(v, message)
	}

	return nil
}

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
	if v, err := g.SetView(constants.PanelInstallation, 0, 0, leftPanelWidth, installationHeight); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "[1]" + " " + constants.TitleInstallation
		v.FgColor = colors.TextPrimary
		v.Wrap = true
		if state.ActivePanel == models.PanelInstallation {
			v.SelBgColor = colors.HighlightBg
			v.SelFgColor = colors.HighlightFg
		}
	}

	if v, err := g.View(constants.PanelInstallation); err == nil {
		v.Clear()
		for i, method := range state.InstallMethods {
			isSelected := i == state.SelectedIndex && state.ActivePanel == models.PanelInstallation
			isConfirmed := state.SelectedMethod == method

			if isSelected {
				fmt.Fprintf(v, "%s%s %s%s\n", colors.ANSIMagenta, constants.RadioSelected, method, colors.ANSIReset)
			} else if isConfirmed {
				fmt.Fprintf(v, "%s%s %s%s\n", colors.ANSIMagenta, constants.RadioSelected, method, colors.ANSIReset)
			} else {
				fmt.Fprintf(v, "%s %s\n", constants.RadioUnselected, method)
			}
		}
		if state.ActivePanel == models.PanelInstallation {
			v.SetCursor(0, state.SelectedIndex)
		}
	}

	// Panel 2: Tools Selection (bottom-left)
	if v, err := g.SetView(constants.PanelTools, 0, installationHeight+1, leftPanelWidth, panelHeight); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "[2]" + " " + constants.TitleToolSelection
		v.FgColor = colors.TextPrimary
		v.Wrap = true
		if state.ActivePanel == models.PanelTools {
			v.SelBgColor = colors.HighlightBg
			v.SelFgColor = colors.HighlightFg
		}
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
			if i == state.ToolsIndex && state.ActivePanel == models.PanelTools {
				fmt.Fprintf(v, "%s%s %s%s\n", colors.ANSIMagenta, marker, tool, colors.ANSIReset)
			} else {
				fmt.Fprintf(v, "%s %s\n", marker, tool)
			}
		}
		if state.ActivePanel == models.PanelTools {
			v.SetCursor(0, state.ToolsIndex)
		}
	}

	// Panel 0: Status/Results (right - full height, read-only)
	if v, err := g.SetView(constants.PanelProgress, leftPanelWidth+1, 0, maxX-1, panelHeight); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "[0]" + " " + constants.PanelStatus
		v.FgColor = colors.TextPrimary
		v.Wrap = true
		v.Highlight = false
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

// getSpinner returns the current spinner character from the animation frame
// Frames cycle through 10 characters: ⠋ ⠙ ⠹ ⠸ ⠼ ⠴ ⠦ ⠧ ⠇ ⠏
// The spinner is wrapped in magenta ANSI color codes
func getSpinner(frame int) string {
	spins := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	spinner := spins[frame%len(spins)]
	return fmt.Sprintf("%s%s%s", colors.ANSIMagenta, spinner, colors.ANSIReset)
}

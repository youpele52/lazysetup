package ui

import (
	"fmt"

	"github.com/jesseduffield/gocui"
	"github.com/youpele52/lazysetup/pkg/colors"
	"github.com/youpele52/lazysetup/pkg/constants"
	"github.com/youpele52/lazysetup/pkg/models"
)

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

	if v, err := g.SetView("tools", maxX/4, startY, 3*maxX/4, endY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Select Tools"
		v.Highlight = true
		v.SelBgColor = colors.HighlightBg
		v.SelFgColor = colors.HighlightFg
		v.FgColor = colors.TextPrimary
		v.Wrap = true

		if err := g.SetCurrentView("tools"); err != nil {
			return err
		}
	}

	if v, err := g.View("tools"); err == nil {
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
	if err := g.DeleteView("tools"); err != nil && err != gocui.ErrUnknownView {
		return err
	}

	contentHeight := 10
	if len(state.InstallOutput) > 0 {
		contentHeight += 3
	}
	startY := maxY/2 - contentHeight/2
	endY := startY + contentHeight

	if v, err := g.SetView("installing", maxX/4, startY, 3*maxX/4, endY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "Installing"
		v.FgColor = colors.TextPrimary
		v.Wrap = true

		if err := g.SetCurrentView("installing"); err != nil {
			return err
		}
	}

	if v, err := g.View("installing"); err == nil {
		v.Clear()
		message := BuildInstallationProgressMessage(state.SelectedMethod, state.CurrentTool, state.InstallingIndex, len(state.Tools), state.InstallationDone, state.SpinnerFrame, state.InstallOutput)
		fmt.Fprint(v, message)
	}

	return nil
}

func layoutResultsPage(g *gocui.Gui, state *models.State, maxX, maxY int) error {
	if err := g.DeleteView("installing"); err != nil && err != gocui.ErrUnknownView {
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

	if v, err := g.SetView("results", maxX/4, startY, 3*maxX/4, endY); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = constants.ResultsSummaryTitle
		v.FgColor = colors.TextPrimary
		v.Wrap = true

		if err := g.SetCurrentView("results"); err != nil {
			return err
		}
	}

	if v, err := g.View("results"); err == nil {
		v.Clear()
		message := BuildInstallationResultsMessage(state.InstallResults)
		fmt.Fprint(v, message)
	}

	return nil
}

func layoutMultiPanel(g *gocui.Gui, state *models.State, maxX, maxY int) error {
	// Delete old views
	for _, viewName := range []string{constants.ViewMenu, constants.ViewResult, "tools", "installing", "results"} {
		g.DeleteView(viewName)
	}

	leftPanelWidth := maxX / 3
	panelHeight := maxY - 3

	// Panel 1: Installation Methods (top-left)
	installationHeight := panelHeight / 2
	if v, err := g.SetView("panel_installation", 0, 0, leftPanelWidth, installationHeight); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "[1] Installation"
		v.FgColor = colors.TextPrimary
		v.Wrap = true
		if state.ActivePanel == models.PanelInstallation {
			v.SelBgColor = colors.HighlightBg
			v.SelFgColor = colors.HighlightFg
		}
	}

	if v, err := g.View("panel_installation"); err == nil {
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
	if v, err := g.SetView("panel_tools", 0, installationHeight+1, leftPanelWidth, panelHeight); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "[2] Select Tools"
		v.FgColor = colors.TextPrimary
		v.Wrap = true
		if state.ActivePanel == models.PanelTools {
			v.SelBgColor = colors.HighlightBg
			v.SelFgColor = colors.HighlightFg
		}
	}

	if v, err := g.View("panel_tools"); err == nil {
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

	// Panel 0: Progress/Results (right - full height, read-only)
	if v, err := g.SetView("panel_progress", leftPanelWidth+1, 0, maxX-1, panelHeight); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Title = "[0] Progress"
		v.FgColor = colors.TextPrimary
		v.Wrap = true
		v.Highlight = false
	}

	if v, err := g.View("panel_progress"); err == nil {
		v.Clear()
		if state.InstallStartTime > 0 && !state.InstallationDone {
			// Show installation progress
			message := BuildInstallationProgressMessage(state.SelectedMethod, state.CurrentTool, state.InstallingIndex, len(state.Tools), state.InstallationDone, state.SpinnerFrame, state.InstallOutput)
			fmt.Fprint(v, message)
		} else if state.InstallationDone {
			// Show results
			message := BuildInstallationResultsMessage(state.InstallResults)
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

func getSpinner(frame int) string {
	spins := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	spinner := spins[frame%len(spins)]
	return fmt.Sprintf("%s%s%s", colors.ANSIMagenta, spinner, colors.ANSIReset)
}

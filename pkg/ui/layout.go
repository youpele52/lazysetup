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
		v.FgColor = colors.AccentText
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
		v.FgColor = colors.AccentText
		v.Wrap = true

		if err := g.SetCurrentView("installing"); err != nil {
			return err
		}
	}

	if v, err := g.View("installing"); err == nil {
		v.Clear()
		fmt.Fprintf(v, "Current Tool: %s\n", state.CurrentTool)
		fmt.Fprintf(v, "Progress: %d/%d\n\n", state.InstallingIndex, len(state.Tools))
		fmt.Fprintf(v, "====================\n")

		if !state.InstallationDone {
			spinner := getSpinner(state.SpinnerFrame)
			fmt.Fprintf(v, "%s Installing...\n", spinner)
		}

		fmt.Fprintf(v, "%s", state.InstallOutput)
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
		v.Title = "Installation Results"
		v.FgColor = colors.AccentText
		v.Wrap = true

		if err := g.SetCurrentView("results"); err != nil {
			return err
		}
	}

	if v, err := g.View("results"); err == nil {
		v.Clear()
		fmt.Fprintf(v, "Installation Summary\n")
		fmt.Fprintf(v, "====================\n\n")

		successCount := 0
		failureCount := 0

		for _, result := range state.InstallResults {
			if result.Success {
				fmt.Fprintf(v, "✓ %s - Success\n", result.Tool)
				successCount++
			} else {
				fmt.Fprintf(v, "✗ %s - Failed\n", result.Tool)
				if result.Error != "" {
					fmt.Fprintf(v, "  Error: %s\n", result.Error)
				}
				failureCount++
			}
		}

		fmt.Fprintf(v, "\n====================\n")
		fmt.Fprintf(v, "Total: %d Success, %d Failed\n", successCount, failureCount)
	}

	return nil
}

func getSpinner(frame int) string {
	spins := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	return spins[frame%len(spins)]
}

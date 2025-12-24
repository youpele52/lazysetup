package ui

import (
	"fmt"

	"github.com/jesseduffield/gocui"
	"github.com/youpele52/lazysetup/pkg/colors"
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

// getSpinner returns the current spinner character from the animation frame
// Frames cycle through 10 characters: ⠋ ⠙ ⠹ ⠸ ⠼ ⠴ ⠦ ⠧ ⠇ ⠏
// The spinner is wrapped in magenta ANSI color codes
func getSpinner(frame int) string {
	spins := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	spinner := spins[frame%len(spins)]
	return fmt.Sprintf("%s%s%s", colors.ANSIMagenta, spinner, colors.ANSIReset)
}

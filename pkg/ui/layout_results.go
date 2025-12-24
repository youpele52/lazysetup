package ui

import (
	"fmt"

	"github.com/jesseduffield/gocui"
	"github.com/youpele52/lazysetup/pkg/colors"
	"github.com/youpele52/lazysetup/pkg/constants"
	"github.com/youpele52/lazysetup/pkg/models"
)

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
		v.Title = "[4]—" + constants.ResultsSummaryTitle
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

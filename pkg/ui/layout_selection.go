package ui

import (
	"fmt"

	"github.com/jesseduffield/gocui"
	"github.com/youpele52/lazysetup/pkg/colors"
	"github.com/youpele52/lazysetup/pkg/constants"
	"github.com/youpele52/lazysetup/pkg/models"
)

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

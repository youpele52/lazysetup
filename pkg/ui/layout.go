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
		default:
			return layoutMenuPage(g, state, maxX, maxY)
		}
	}
}

func layoutMenuPage(g *gocui.Gui, state *models.State, maxX, maxY int) error {
	if err := g.DeleteView(constants.ViewResult); err != nil && err != gocui.ErrUnknownView {
		return err
	}

	if v, err := g.SetView(constants.ViewMenu, maxX/4, maxY/4, 3*maxX/4, maxY/2); err != nil {
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
				fmt.Fprintf(v, "● %s\n", method)
			} else {
				fmt.Fprintf(v, "○ %s\n", method)
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

	if v, err := g.SetView(constants.ViewResult, maxX/4, maxY/4, 3*maxX/4, maxY-2); err != nil {
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

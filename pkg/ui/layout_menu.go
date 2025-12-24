package ui

import (
	"fmt"

	"github.com/jesseduffield/gocui"
	"github.com/youpele52/lazysetup/pkg/colors"
	"github.com/youpele52/lazysetup/pkg/constants"
	"github.com/youpele52/lazysetup/pkg/models"
)

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
		v.Title = "[1]-" + constants.TitleInstallation
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

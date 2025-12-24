package ui

import (
	"fmt"

	"github.com/jesseduffield/gocui"
	"github.com/youpele52/lazysetup/pkg/colors"
	"github.com/youpele52/lazysetup/pkg/constants"
	"github.com/youpele52/lazysetup/pkg/models"
)

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
		v.Title = "[3]-" + constants.TitleTools
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
		for i, tool := range state.Tools {
			selected := state.SelectedTools[tool]
			var marker string
			if selected {
				marker = constants.CheckboxSelected
			} else {
				marker = constants.CheckboxUnselected
			}
			// Active panel: green text by default
			if i == state.ToolsIndex {
				// Cursor position: magenta text (will have magenta background from highlight)
				fmt.Fprintf(v, "%s%s %s%s\n", colors.ANSIMagenta, marker, tool, colors.ANSIReset)
			} else if selected {
				// Selected item: magenta text
				fmt.Fprintf(v, "%s%s %s%s\n", colors.ANSIMagenta, marker, tool, colors.ANSIReset)
			} else {
				// Unselected item: green text
				fmt.Fprintf(v, "%s%s %s%s\n", colors.ANSIGreen, marker, tool, colors.ANSIReset)
			}
		}
		v.SetCursor(0, state.ToolsIndex)
	}

	return nil
}

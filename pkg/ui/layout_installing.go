package ui

import (
	"fmt"

	"github.com/jesseduffield/gocui"
	"github.com/youpele52/lazysetup/pkg/colors"
	"github.com/youpele52/lazysetup/pkg/constants"
	"github.com/youpele52/lazysetup/pkg/models"
)

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
		v.Title = "[3]—" + constants.TitleInstalling
		v.FgColor = colors.TextPrimary
		v.Wrap = true

		if err := g.SetCurrentView(constants.ViewInstalling); err != nil {
			return err
		}
	}

	if v, err := g.View(constants.ViewInstalling); err == nil {
		v.Clear()
		params := ProgressMessageParams{
			SelectedMethod:   state.SelectedMethod,
			CurrentTool:      state.CurrentTool,
			InstallingIndex:  state.InstallingIndex,
			TotalTools:       len(state.Tools),
			InstallationDone: state.InstallationDone,
			SpinnerFrame:     state.SpinnerFrame,
			InstallOutput:    state.InstallOutput,
			Action:           state.SelectedAction,
		}
		message := BuildInstallationProgressMessage(params)
		fmt.Fprint(v, message)
	}

	return nil
}

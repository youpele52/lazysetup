package installer

import (
	"github.com/jesseduffield/gocui"
	"github.com/youpele52/lazysetup/pkg/models"
)

func CursorUp(state *models.State) func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		if state.SelectedIndex > 0 {
			state.SelectedIndex--
		}
		return nil
	}
}

func CursorDown(state *models.State) func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		if state.SelectedIndex < len(state.InstallMethods)-1 {
			state.SelectedIndex++
		}
		return nil
	}
}

func SelectMethod(state *models.State) func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		state.SelectedMethod = state.InstallMethods[state.SelectedIndex]
		return nil
	}
}

func GoBack(state *models.State) func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		state.Reset()
		return nil
	}
}

func Quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

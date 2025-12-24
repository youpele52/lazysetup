package ui

import (
	"log"

	"github.com/jesseduffield/gocui"
	"github.com/youpele52/lazysetup/pkg/handlers"
	"github.com/youpele52/lazysetup/pkg/models"
)

// SetupKeybindings configures all keyboard shortcuts for the application
// Bindings: Ctrl+C (quit), Tab (next panel), 0/1/2 (jump to panel),
// Arrow keys (navigate), Space (toggle), Enter (confirm/execute), Esc (back/abort)
func SetupKeybindings(g *gocui.Gui, state *models.State) {
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, handlers.Quit); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyTab, gocui.ModNone, handlers.NextPanel(state)); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", '0', gocui.ModNone, handlers.SwitchToPanel(state, models.PanelProgress)); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", '1', gocui.ModNone, handlers.SwitchToPanel(state, models.PanelInstallation)); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", '2', gocui.ModNone, handlers.SwitchToPanel(state, models.PanelTools)); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyArrowUp, gocui.ModNone, handlers.MultiPanelCursorUp(state)); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyArrowDown, gocui.ModNone, handlers.MultiPanelCursorDown(state)); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeySpace, gocui.ModNone, handlers.MultiPanelToggleTool(state)); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyEnter, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		if state.GetCurrentPage() == models.PageMultiPanel {
			if state.GetActivePanel() == models.PanelInstallation {
				return handlers.MultiPanelSelectMethod(state)(g, v)
			} else if state.GetActivePanel() == models.PanelTools {
				return handlers.MultiPanelStartInstallation(state)(g, v)
			}
		}
		return nil
	}); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyEsc, gocui.ModNone, handlers.GoBack(state)); err != nil {
		log.Panicln(err)
	}
}

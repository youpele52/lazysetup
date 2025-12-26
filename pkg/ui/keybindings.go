package ui

import (
	"log"

	"github.com/jesseduffield/gocui"
	"github.com/youpele52/lazysetup/pkg/handlers"
	"github.com/youpele52/lazysetup/pkg/models"
)

// SetupKeybindings configures all keyboard shortcuts for the application
// Bindings: Ctrl+C (quit), Tab (next panel), 0/1/2/3 (jump to panel),
// Arrow keys (navigate), Space (toggle), Enter (confirm/execute), Esc (back/abort)
func SetupKeybindings(g *gocui.Gui, state *models.State) {
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, handlers.Quit); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyTab, gocui.ModNone, handlers.NextPanel(state)); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", '0', gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		if state.GetShowSudoConfirm() {
			state.AppendPasswordInput('0')
			return nil
		}
		return handlers.SwitchToPanel(state, models.PanelStatus)(g, v)
	}); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", '1', gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		if state.GetShowSudoConfirm() {
			state.AppendPasswordInput('1')
			return nil
		}
		return handlers.SwitchToPanel(state, models.PanelPackageManager)(g, v)
	}); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", '2', gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		if state.GetShowSudoConfirm() {
			state.AppendPasswordInput('2')
			return nil
		}
		return handlers.SwitchToPanel(state, models.PanelAction)(g, v)
	}); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", '3', gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		if state.GetShowSudoConfirm() {
			state.AppendPasswordInput('3')
			return nil
		}
		return handlers.SwitchToPanel(state, models.PanelTools)(g, v)
	}); err != nil {
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
		// Handle sudo confirmation popup first
		if state.GetShowSudoConfirm() {
			return handlers.ConfirmSudoPopup(state)(g, v)
		}

		if state.GetCurrentPage() == models.PageMultiPanel {
			switch state.GetActivePanel() {
			case models.PanelPackageManager:
				return handlers.MultiPanelSelectMethod(state)(g, v)
			case models.PanelAction:
				return handlers.MultiPanelSelectAction(state)(g, v)
			case models.PanelTools:
				return handlers.MultiPanelExecuteAction(state)(g, v)
			}
		}
		return nil
	}); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyEsc, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		// Handle sudo confirmation popup first
		if state.GetShowSudoConfirm() {
			return handlers.CancelSudoPopup(state)(g, v)
		}
		return handlers.GoBack(state)(g, v)
	}); err != nil {
		log.Panicln(err)
	}

	// Backspace for password input
	if err := g.SetKeybinding("", gocui.KeyBackspace, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		if state.GetShowSudoConfirm() {
			state.BackspacePasswordInput()
		}
		return nil
	}); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyBackspace2, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		if state.GetShowSudoConfirm() {
			state.BackspacePasswordInput()
		}
		return nil
	}); err != nil {
		log.Panicln(err)
	}

	// Character input for password - bind printable characters (excluding 0-3 which are handled above)
	for _, char := range "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ456789!@#$%^&*()-_=+[]{}|;:',.<>?/`~\"\\  " {
		c := char // capture for closure
		if err := g.SetKeybinding("", c, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
			if state.GetShowSudoConfirm() {
				state.AppendPasswordInput(c)
			}
			return nil
		}); err != nil {
			log.Panicln(err)
		}
	}
}

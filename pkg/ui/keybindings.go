package ui

import (
	"fmt"
	"log"

	"github.com/jesseduffield/gocui"
	"github.com/youpele52/lazysetup/pkg/constants"
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

	if err := g.SetKeybinding("", gocui.KeyArrowUp, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		if state.GetShowSudoConfirm() {
			return nil
		}
		// Scroll up in status panel ONLY if active
		if state.GetActivePanel() == models.PanelStatus {
			if v, err := g.View(constants.PanelProgress); err == nil {
				ox, oy := v.Origin()
				if oy > 0 {
					v.SetOrigin(ox, oy-1)
				}
			}
			return nil
		}
		// Navigate in other panels
		return handlers.MultiPanelCursorUp(state)(g, v)
	}); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeyArrowDown, gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		if state.GetShowSudoConfirm() {
			return nil
		}
		// Scroll down in status panel ONLY if active
		if state.GetActivePanel() == models.PanelStatus {
			if v, err := g.View(constants.PanelProgress); err == nil {
				ox, oy := v.Origin()
				v.SetOrigin(ox, oy+1)
			}
			return nil
		}
		// Navigate in other panels
		return handlers.MultiPanelCursorDown(state)(g, v)
	}); err != nil {
		log.Panicln(err)
	}

	// Jump to first item shortcuts: 'g' (vim-style) and 'w'
	if err := g.SetKeybinding("", 'g', gocui.ModNone, handlers.JumpToFirst(state)); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("", 'w', gocui.ModNone, handlers.JumpToFirst(state)); err != nil {
		log.Panicln(err)
	}

	// Jump to last item shortcuts: 'G' (vim-style) and 's'
	if err := g.SetKeybinding("", 'G', gocui.ModNone, handlers.JumpToLast(state)); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("", 's', gocui.ModNone, handlers.JumpToLast(state)); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("", gocui.KeySpace, gocui.ModNone, handlers.MultiPanelToggleTool(state)); err != nil {
		log.Panicln(err)
	}

	// Clear status screen with 'c' key
	if err := g.SetKeybinding("", 'c', gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		if state.GetShowSudoConfirm() {
			state.AppendPasswordInput('c')
			return nil
		}
		// Clear status screen and reset all state
		if v, err := g.View(constants.PanelProgress); err == nil {
			v.Clear()
			state.SetInstallationDone(false)
			state.ActionCompletionTime = 0
			state.LastRenderedResultCount = 0
			state.SetInstallStartTime(0)
			state.Error = ""
			fmt.Fprint(v, constants.Logo)
		}
		return nil
	}); err != nil {
		log.Panicln(err)
	}

	// Update application with 'u' key
	if err := g.SetKeybinding("", 'u', gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		if state.GetShowSudoConfirm() {
			state.AppendPasswordInput('u')
			return nil
		}
		// Trigger update if available
		if state.UpdateAvailable {
			go handlers.ExecuteUpdate(state)
		}
		return nil
	}); err != nil {
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

	// Character input for password - bind all printable ASCII characters
	// Range 33-126 covers all printable ASCII except space (32)
	for i := 33; i <= 126; i++ {
		char := rune(i)
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

	// Also bind space character for passwords
	if err := g.SetKeybinding("", ' ', gocui.ModNone, func(g *gocui.Gui, v *gocui.View) error {
		if state.GetShowSudoConfirm() {
			state.AppendPasswordInput(' ')
			return nil
		}
		return handlers.MultiPanelToggleTool(state)(g, v)
	}); err != nil {
		log.Panicln(err)
	}

	// Open website with 'w' key
	if err := g.SetKeybinding("", 'w', gocui.ModNone, handlers.OpenWebsite); err != nil {
		log.Panicln(err)
	}
}

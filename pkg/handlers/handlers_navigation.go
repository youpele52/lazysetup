package handlers

import (
	"github.com/jesseduffield/gocui"
	"github.com/youpele52/lazysetup/pkg/config"
	"github.com/youpele52/lazysetup/pkg/models"
)

// NextPanel switches to the next panel (Tab key)
// Cycles through: PackageManager(0) -> Action(2) -> Tools(3) -> PackageManager(0)
// Status panel (1) is skipped as it's read-only
func NextPanel(state *models.State) func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		currentPage := state.GetCurrentPage()
		if currentPage == models.PageMultiPanel {
			activePanel := state.GetActivePanel()
			var newPanel models.Panel
			switch activePanel {
			case models.PanelPackageManager:
				newPanel = models.PanelAction
			case models.PanelAction:
				newPanel = models.PanelTools
			case models.PanelTools:
				newPanel = models.PanelPackageManager
			default:
				newPanel = models.PanelPackageManager
			}
			state.SetActivePanel(newPanel)
		}
		return nil
	}
}

// PrevPanel switches to the previous panel (Shift+Tab)
// Cycles through: PackageManager(0) <- Action(2) <- Tools(3) <- PackageManager(0)
// Status panel (1) is skipped as it's read-only
func PrevPanel(state *models.State) func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		currentPage := state.GetCurrentPage()
		if currentPage == models.PageMultiPanel {
			activePanel := state.GetActivePanel()
			var newPanel models.Panel
			switch activePanel {
			case models.PanelPackageManager:
				newPanel = models.PanelTools
			case models.PanelAction:
				newPanel = models.PanelPackageManager
			case models.PanelTools:
				newPanel = models.PanelAction
			default:
				newPanel = models.PanelPackageManager
			}
			state.SetActivePanel(newPanel)
		}
		return nil
	}
}

// SwitchToPanel switches to a specific panel (0, 1, 2 keys)
func SwitchToPanel(state *models.State, panel models.Panel) func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		currentPage := state.GetCurrentPage()
		if currentPage == models.PageMultiPanel {
			state.SetActivePanel(panel)
		}
		return nil
	}
}

// MultiPanelCursorUp moves cursor up in the active panel
func MultiPanelCursorUp(state *models.State) func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		currentPage := state.GetCurrentPage()
		if currentPage == models.PageMultiPanel {
			activePanel := state.GetActivePanel()
			switch activePanel {
			case models.PanelPackageManager:
				state.PackageManagerScroll.ScrollUp()
			case models.PanelAction:
				state.ActionScroll.ScrollUp()
				state.SelectedAction = models.ActionType(state.ActionScroll.Cursor)
			case models.PanelTools:
				// Use search scroll if in search mode
				if state.GetIsSearchMode() {
					state.ToolsSearchScroll.ScrollUp()
				} else {
					state.ToolsScroll.ScrollUp()
				}
			}
		}
		return nil
	}
}

// MultiPanelCursorDown moves cursor down in the active panel
func MultiPanelCursorDown(state *models.State) func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		currentPage := state.GetCurrentPage()
		if currentPage == models.PageMultiPanel {
			activePanel := state.GetActivePanel()
			switch activePanel {
			case models.PanelPackageManager:
				state.PackageManagerScroll.ScrollDown()
			case models.PanelAction:
				state.ActionScroll.ScrollDown()
				state.SelectedAction = models.ActionType(state.ActionScroll.Cursor)
			case models.PanelTools:
				// Use search scroll if in search mode
				if state.GetIsSearchMode() {
					state.ToolsSearchScroll.ScrollDown()
				} else {
					state.ToolsScroll.ScrollDown()
				}
			}
		}
		return nil
	}
}

// MultiPanelToggleTool toggles tool selection in the Tools panel
func MultiPanelToggleTool(state *models.State) func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		if state.GetCurrentPage() != models.PageMultiPanel {
			return nil
		}
		if state.GetActivePanel() != models.PanelTools {
			return nil
		}

		// Determine which tool list to use
		var toolList []string
		var cursorPos int

		if state.GetIsSearchMode() {
			toolList = state.GetFilteredTools()
			cursorPos = state.ToolsSearchScroll.Cursor
		} else {
			toolList = state.Tools
			cursorPos = state.ToolsScroll.Cursor
		}

		// Validate cursor position
		if cursorPos >= len(toolList) || cursorPos < 0 {
			return nil
		}

		// Toggle the tool at cursor position
		tool := toolList[cursorPos]
		state.SelectedTools[tool] = !state.SelectedTools[tool]

		return nil
	}
}

// CursorUp moves cursor up in single-panel views
func CursorUp(state *models.State) func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		state.PackageManagerScroll.ScrollUp()
		return nil
	}
}

// CursorDown moves cursor down in single-panel views
func CursorDown(state *models.State) func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		state.PackageManagerScroll.ScrollDown()
		return nil
	}
}

// ToolsCursorUp moves cursor up in tools list
func ToolsCursorUp(state *models.State) func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		state.ToolsScroll.ScrollUp()
		return nil
	}
}

// ToolsCursorDown moves cursor down in tools list
func ToolsCursorDown(state *models.State) func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		state.ToolsScroll.ScrollDown()
		return nil
	}
}

// ToggleTool toggles tool selection in single-panel view
func ToggleTool(state *models.State) func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		tool := state.Tools[state.ToolsScroll.Cursor]
		state.SelectedTools[tool] = !state.SelectedTools[tool]
		return nil
	}
}

// JumpToFirst jumps to the first item in the active panel
// Supports both 'g' (vim-style) and 'w' shortcuts
// Disabled when sudo confirmation popup is shown
func JumpToFirst(state *models.State) func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		if state.GetShowSudoConfirm() {
			return nil // Disable in password mode
		}

		currentPage := state.GetCurrentPage()
		if currentPage == models.PageMultiPanel {
			activePanel := state.GetActivePanel()
			switch activePanel {
			case models.PanelPackageManager:
				state.PackageManagerScroll.JumpToFirst()
			case models.PanelAction:
				state.ActionScroll.JumpToFirst()
				state.SelectedAction = models.ActionType(0)
			case models.PanelTools:
				if state.GetIsSearchMode() {
					state.ToolsSearchScroll.JumpToFirst()
				} else {
					state.ToolsScroll.JumpToFirst()
				}
			}
		}
		return nil
	}
}

// JumpToLast jumps to the last item in the active panel
// Supports both 'G' (vim-style) and 's' shortcuts
// Disabled when sudo confirmation popup is shown
func JumpToLast(state *models.State) func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		if state.GetShowSudoConfirm() {
			return nil // Disable in password mode
		}

		currentPage := state.GetCurrentPage()
		if currentPage == models.PageMultiPanel {
			activePanel := state.GetActivePanel()
			switch activePanel {
			case models.PanelPackageManager:
				state.PackageManagerScroll.JumpToLast()
			case models.PanelAction:
				state.ActionScroll.JumpToLast()
				state.SelectedAction = models.ActionType(len(config.Actions) - 1)
			case models.PanelTools:
				if state.GetIsSearchMode() {
					state.ToolsSearchScroll.JumpToLast()
				} else {
					state.ToolsScroll.JumpToLast()
				}
			}
		}
		return nil
	}
}

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
				if state.SelectedIndex > 0 {
					state.SelectedIndex--
				}
			case models.PanelAction:
				if state.ActionIndex > 0 {
					state.ActionIndex--
					state.SelectedAction = models.ActionType(state.ActionIndex)
				}
			case models.PanelTools:
				if state.ToolsIndex > 0 {
					state.ToolsIndex--
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
				if state.SelectedIndex < len(state.InstallMethods)-1 {
					state.SelectedIndex++
				}
			case models.PanelAction:
				if state.ActionIndex < len(config.Actions)-1 {
					state.ActionIndex++
					state.SelectedAction = models.ActionType(state.ActionIndex)
				}
			case models.PanelTools:
				if state.ToolsIndex < len(state.Tools)-1 {
					state.ToolsIndex++
				}
			}
		}
		return nil
	}
}

// MultiPanelToggleTool toggles tool selection in the Tools panel
func MultiPanelToggleTool(state *models.State) func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		if state.GetCurrentPage() == models.PageMultiPanel && state.GetActivePanel() == models.PanelTools {
			tool := state.Tools[state.ToolsIndex]
			state.SelectedTools[tool] = !state.SelectedTools[tool]
		}
		return nil
	}
}

// CursorUp moves cursor up in single-panel views
func CursorUp(state *models.State) func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		if state.SelectedIndex > 0 {
			state.SelectedIndex--
		}
		return nil
	}
}

// CursorDown moves cursor down in single-panel views
func CursorDown(state *models.State) func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		if state.SelectedIndex < len(state.InstallMethods)-1 {
			state.SelectedIndex++
		}
		return nil
	}
}

// ToolsCursorUp moves cursor up in tools list
func ToolsCursorUp(state *models.State) func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		if state.ToolsIndex > 0 {
			state.ToolsIndex--
		}
		return nil
	}
}

// ToolsCursorDown moves cursor down in tools list
func ToolsCursorDown(state *models.State) func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		if state.ToolsIndex < len(state.Tools)-1 {
			state.ToolsIndex++
		}
		return nil
	}
}

// ToggleTool toggles tool selection in single-panel view
func ToggleTool(state *models.State) func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		tool := state.Tools[state.ToolsIndex]
		state.SelectedTools[tool] = !state.SelectedTools[tool]
		return nil
	}
}

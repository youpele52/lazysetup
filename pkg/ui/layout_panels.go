package ui

import (
	"fmt"

	"github.com/jesseduffield/gocui"
	"github.com/youpele52/lazysetup/pkg/colors"
	"github.com/youpele52/lazysetup/pkg/constants"
	"github.com/youpele52/lazysetup/pkg/models"
)

// renderPackageManagerPanel renders the Package Manager selection panel
func renderPackageManagerPanel(g *gocui.Gui, state *models.State, activePanel models.Panel, leftPanelWidth, packageManagerHeight int) error {
	if v, err := g.SetView(constants.PanelPackageManager, 0, 0, leftPanelWidth, packageManagerHeight); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Wrap = true
		v.SelBgColor = colors.HighlightBg
		v.SelFgColor = colors.HighlightFg
	}

	if v, err := g.View(constants.PanelPackageManager); err == nil {
		v.Title = "[1]-" + constants.TitlePackageManager
		if activePanel == models.PanelPackageManager {
			v.FgColor = colors.ActiveBorderColor
		} else {
			v.FgColor = colors.TextPrimary
		}
		v.Clear()
		for i, method := range state.InstallMethods {
			marker := constants.RadioUnselected
			if i == state.SelectedIndex {
				marker = constants.RadioSelected
			}

			if activePanel == models.PanelPackageManager {
				if i == state.SelectedIndex {
					fmt.Fprintf(v, "%s%s %s%s\n", colors.ANSIMagenta, marker, method, colors.ANSIReset)
				} else {
					fmt.Fprintf(v, "%s%s %s%s\n", colors.ANSIGreen, marker, method, colors.ANSIReset)
				}
			} else {
				if i == state.SelectedIndex {
					fmt.Fprintf(v, "%s%s %s%s\n", colors.ANSIMagenta, marker, method, colors.ANSIReset)
				} else {
					fmt.Fprintf(v, "%s %s\n", marker, method)
				}
			}
		}
		if activePanel == models.PanelPackageManager {
			v.SetCursor(0, state.SelectedIndex)
		}
	}
	return nil
}

// renderActionPanel renders the Action selection panel
func renderActionPanel(g *gocui.Gui, state *models.State, activePanel models.Panel, leftPanelWidth, packageManagerHeight, actionHeight int) error {
	actions := []string{"Install", "Update", "Uninstall"}
	if v, err := g.SetView(constants.PanelAction, 0, packageManagerHeight+1, leftPanelWidth, packageManagerHeight+actionHeight+1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Wrap = true
		v.SelBgColor = colors.HighlightBg
		v.SelFgColor = colors.HighlightFg
	}

	if v, err := g.View(constants.PanelAction); err == nil {
		v.Title = "[2]-" + constants.TitleAction
		if activePanel == models.PanelAction {
			v.FgColor = colors.ActiveBorderColor
		} else {
			v.FgColor = colors.TextPrimary
		}
		v.Clear()
		for i, action := range actions {
			marker := constants.RadioUnselected
			if i == state.ActionIndex {
				marker = constants.RadioSelected
			}

			if activePanel == models.PanelAction {
				if i == state.ActionIndex {
					fmt.Fprintf(v, "%s%s %s%s\n", colors.ANSIMagenta, marker, action, colors.ANSIReset)
				} else {
					fmt.Fprintf(v, "%s%s %s%s\n", colors.ANSIGreen, marker, action, colors.ANSIReset)
				}
			} else {
				if i == state.ActionIndex {
					fmt.Fprintf(v, "%s%s %s%s\n", colors.ANSIMagenta, marker, action, colors.ANSIReset)
				} else {
					fmt.Fprintf(v, "%s %s\n", marker, action)
				}
			}
		}
		if activePanel == models.PanelAction {
			v.SetCursor(0, state.ActionIndex)
		}
	}
	return nil
}

// renderToolsPanel renders the Tools selection panel
func renderToolsPanel(g *gocui.Gui, state *models.State, activePanel models.Panel, leftPanelWidth, toolsStartY, panelHeight int) error {
	if v, err := g.SetView(constants.PanelTools, 0, toolsStartY, leftPanelWidth, panelHeight); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Wrap = true
		v.SelBgColor = colors.HighlightBg
		v.SelFgColor = colors.HighlightFg
	}

	if v, err := g.View(constants.PanelTools); err == nil {
		v.Title = "[3]-" + constants.TitleTools
		if activePanel == models.PanelTools {
			v.FgColor = colors.ActiveBorderColor
		} else {
			v.FgColor = colors.TextPrimary
		}
		v.Clear()
		for i, tool := range state.Tools {
			marker := constants.CheckboxUnselected
			if state.SelectedTools[tool] {
				marker = constants.CheckboxSelected
			}

			if activePanel == models.PanelTools {
				if i == state.ToolsIndex {
					fmt.Fprintf(v, "%s%s %s%s\n", colors.ANSIMagenta, marker, tool, colors.ANSIReset)
				} else {
					fmt.Fprintf(v, "%s%s %s%s\n", colors.ANSIGreen, marker, tool, colors.ANSIReset)
				}
			} else {
				if i == state.ToolsIndex {
					fmt.Fprintf(v, "%s%s %s%s\n", colors.ANSIMagenta, marker, tool, colors.ANSIReset)
				} else {
					fmt.Fprintf(v, "%s %s\n", marker, tool)
				}
			}
		}
		if activePanel == models.PanelTools {
			v.SetCursor(0, state.ToolsIndex)
		}
	}
	return nil
}

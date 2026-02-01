package ui

import (
	"fmt"

	"github.com/jesseduffield/gocui"
	"github.com/youpele52/lazysetup/pkg/colors"
	"github.com/youpele52/lazysetup/pkg/config"
	"github.com/youpele52/lazysetup/pkg/constants"
	"github.com/youpele52/lazysetup/pkg/models"
)

// PanelParams groups common parameters for panel rendering
type PanelParams struct {
	Gui            *gocui.Gui
	State          *models.State
	ActivePanel    models.Panel
	LeftPanelWidth int
}

// PackageManagerParams groups parameters for renderPackageManagerPanel
type PackageManagerParams struct {
	PanelParams
	Height int
}

// ActionPanelParams groups parameters for renderActionPanel
type ActionPanelParams struct {
	PanelParams
	PackageManagerY int
	ActionHeight    int
}

// ToolsPanelParams groups parameters for renderToolsPanel
type ToolsPanelParams struct {
	PanelParams
	ToolsStartY int
	PanelHeight int
}

// renderPackageManagerPanel renders the Package Manager selection panel
func renderPackageManagerPanel(params PackageManagerParams) error {
	g := params.Gui
	state := params.State
	activePanel := params.ActivePanel
	leftPanelWidth := params.LeftPanelWidth
	packageManagerHeight := params.Height
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

		// Calculate visible count based on panel height (accounting for borders)
		visibleCount := packageManagerHeight - 2
		if visibleCount < 1 {
			visibleCount = 1
		}

		// Update scroll state - set visible count but don't adjust offset
		// Offset is managed by navigation functions (ScrollUp/Down, JumpToFirst/Last)
		state.PackageManagerScroll.VisibleCount = visibleCount
		state.PackageManagerScroll.ItemCount = len(state.InstallMethods)

		// Set scroll origin
		v.SetOrigin(0, state.PackageManagerScroll.Offset)

		// Only render visible items
		startIdx := state.PackageManagerScroll.Offset
		endIdx := startIdx + visibleCount
		if endIdx > len(state.InstallMethods) {
			endIdx = len(state.InstallMethods)
		}

		for i := startIdx; i < endIdx; i++ {
			method := state.InstallMethods[i]
			marker := constants.RadioUnselected
			if i == state.PackageManagerScroll.Cursor {
				marker = constants.RadioSelected
			}

			if activePanel == models.PanelPackageManager {
				if i == state.PackageManagerScroll.Cursor {
					fmt.Fprintf(v, "%s%s %s%s\n", colors.ANSIMagenta, marker, method, colors.ANSIReset)
				} else {
					fmt.Fprintf(v, "%s%s %s%s\n", colors.ANSIGreen, marker, method, colors.ANSIReset)
				}
			} else {
				if i == state.PackageManagerScroll.Cursor {
					fmt.Fprintf(v, "%s%s %s%s\n", colors.ANSIMagenta, marker, method, colors.ANSIReset)
				} else {
					fmt.Fprintf(v, "%s %s\n", marker, method)
				}
			}
		}
		if activePanel == models.PanelPackageManager {
			v.SetCursor(0, state.PackageManagerScroll.Cursor-state.PackageManagerScroll.Offset)
		}
	}
	return nil
}

// renderActionPanel renders the Action selection panel
func renderActionPanel(params ActionPanelParams) error {
	g := params.Gui
	state := params.State
	activePanel := params.ActivePanel
	leftPanelWidth := params.LeftPanelWidth
	packageManagerHeight := params.PackageManagerY
	actionHeight := params.ActionHeight
	actions := config.Actions
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

		// Calculate visible count based on panel height (accounting for borders)
		visibleCount := actionHeight - 2
		if visibleCount < 1 {
			visibleCount = 1
		}

		// Update scroll state
		state.ActionScroll.VisibleCount = visibleCount
		state.ActionScroll.ItemCount = len(actions)

		// Set scroll origin
		v.SetOrigin(0, state.ActionScroll.Offset)

		// Only render visible items
		startIdx := state.ActionScroll.Offset
		endIdx := startIdx + visibleCount
		if endIdx > len(actions) {
			endIdx = len(actions)
		}

		for i := startIdx; i < endIdx; i++ {
			action := actions[i]
			marker := constants.RadioUnselected
			if i == state.ActionScroll.Cursor {
				marker = constants.RadioSelected
			}

			if activePanel == models.PanelAction {
				if i == state.ActionScroll.Cursor {
					fmt.Fprintf(v, "%s%s %s%s\n", colors.ANSIMagenta, marker, action, colors.ANSIReset)
				} else {
					fmt.Fprintf(v, "%s%s %s%s\n", colors.ANSIGreen, marker, action, colors.ANSIReset)
				}
			} else {
				if i == state.ActionScroll.Cursor {
					fmt.Fprintf(v, "%s%s %s%s\n", colors.ANSIMagenta, marker, action, colors.ANSIReset)
				} else {
					fmt.Fprintf(v, "%s %s\n", marker, action)
				}
			}
		}
		if activePanel == models.PanelAction {
			v.SetCursor(0, state.ActionScroll.Cursor-state.ActionScroll.Offset)
		}
	}
	return nil
}

// renderToolsPanel renders the Tools selection panel
func renderToolsPanel(params ToolsPanelParams) error {
	g := params.Gui
	state := params.State
	activePanel := params.ActivePanel
	leftPanelWidth := params.LeftPanelWidth
	toolsStartY := params.ToolsStartY
	panelHeight := params.PanelHeight
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

		// Calculate visible tool count based on panel height (accounting for borders)
		visibleCount := panelHeight - toolsStartY - 2
		if visibleCount < 1 {
			visibleCount = 1
		}

		// Update scroll state - set visible count but don't adjust offset
		// Offset is managed by navigation functions (ScrollUp/Down, JumpToFirst/Last)
		state.ToolsScroll.VisibleCount = visibleCount
		state.ToolsScroll.ItemCount = len(state.Tools)

		// Set scroll origin
		v.SetOrigin(0, state.ToolsScroll.Offset)

		// Only render visible tools
		startIdx := state.ToolsScroll.Offset
		endIdx := startIdx + visibleCount
		if endIdx > len(state.Tools) {
			endIdx = len(state.Tools)
		}

		for i := startIdx; i < endIdx; i++ {
			tool := state.Tools[i]
			marker := constants.CheckboxUnselected
			if state.SelectedTools[tool] {
				marker = constants.CheckboxSelected
			}

			if activePanel == models.PanelTools {
				if i == state.ToolsScroll.Cursor {
					fmt.Fprintf(v, "%s%s %s%s\n", colors.ANSIMagenta, marker, tool, colors.ANSIReset)
				} else {
					fmt.Fprintf(v, "%s%s %s%s\n", colors.ANSIGreen, marker, tool, colors.ANSIReset)
				}
			} else {
				if i == state.ToolsScroll.Cursor {
					fmt.Fprintf(v, "%s%s %s%s\n", colors.ANSIMagenta, marker, tool, colors.ANSIReset)
				} else {
					fmt.Fprintf(v, "%s %s\n", marker, tool)
				}
			}
		}
		if activePanel == models.PanelTools {
			v.SetCursor(0, state.ToolsScroll.Cursor-state.ToolsScroll.Offset)
		}
	}
	return nil
}

// renderSudoConfirmPopup renders a centered password input popup for sudo
func renderSudoConfirmPopup(g *gocui.Gui, maxX, maxY int, state *models.State) error {
	popupWidth := 50
	popupHeight := 8
	x0 := (maxX - popupWidth) / 2
	y0 := (maxY - popupHeight) / 2
	x1 := x0 + popupWidth
	y1 := y0 + popupHeight

	if v, err := g.SetView(constants.PopupConfirm, x0, y0, x1, y1); err != nil {
		if err != gocui.ErrUnknownView {
			return err
		}
		v.Wrap = true
		v.FgColor = colors.TextPrimary
		v.Editable = false
	}

	if v, err := g.View(constants.PopupConfirm); err == nil {
		v.Title = constants.SudoConfirmTitle
		v.Clear()
		// Create masked password display
		passwordInput := state.GetPasswordInput()
		maskedPassword := ""
		for range passwordInput {
			maskedPassword += constants.PasswordMask
		}
		fmt.Fprintf(v, constants.SudoConfirmMessage, maskedPassword)
	}

	// Bring popup to front
	g.SetViewOnTop(constants.PopupConfirm)
	g.SetCurrentView(constants.PopupConfirm)

	return nil
}

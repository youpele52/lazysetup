package ui

import (
	"fmt"
	"strings"

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
	ToolsHeight int // RENAMED: was PanelHeight (misleading - this is height, not boundary)
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

		// Get actual rendered dimensions from the view (excludes borders automatically)
		_, visibleCount := v.Size()
		if visibleCount < 1 {
			visibleCount = 1
		}

		// Update scroll state - set visible count but don't adjust offset
		// Offset is managed by navigation functions (ScrollUp/Down, JumpToFirst/Last)
		state.PackageManagerScroll.VisibleCount = visibleCount
		state.PackageManagerScroll.ItemCount = len(state.InstallMethods)

		// Safety check: ensure cursor is visible
		if state.PackageManagerScroll.Cursor >= state.PackageManagerScroll.ItemCount {
			state.PackageManagerScroll.Cursor = state.PackageManagerScroll.ItemCount - 1
		}
		if state.PackageManagerScroll.Cursor < 0 {
			state.PackageManagerScroll.Cursor = 0
		}
		// Ensure offset shows the cursor
		if state.PackageManagerScroll.Cursor < state.PackageManagerScroll.Offset {
			state.PackageManagerScroll.Offset = state.PackageManagerScroll.Cursor
		} else if state.PackageManagerScroll.Cursor >= state.PackageManagerScroll.Offset+visibleCount {
			state.PackageManagerScroll.Offset = state.PackageManagerScroll.Cursor - visibleCount + 1
		}

		// Calculate scroll offset to keep cursor visible
		offset := state.PackageManagerScroll.Offset
		cursor := state.PackageManagerScroll.Cursor

		// Ensure cursor is within bounds
		if cursor >= len(state.InstallMethods) {
			cursor = len(state.InstallMethods) - 1
			state.PackageManagerScroll.Cursor = cursor
		}
		if cursor < 0 {
			cursor = 0
			state.PackageManagerScroll.Cursor = cursor
		}

		// Adjust offset to keep cursor visible
		if cursor < offset {
			offset = cursor
		} else if cursor >= offset+visibleCount {
			offset = cursor - visibleCount + 1
		}
		state.PackageManagerScroll.Offset = offset

		// Set scroll position
		v.SetOrigin(0, offset)

		// Render ALL methods (gocui's SetOrigin will handle which ones are visible)
		for i := 0; i < len(state.InstallMethods); i++ {
			method := strings.ToLower(state.InstallMethods[i])
			marker := constants.RadioUnselected
			if i == cursor {
				marker = constants.RadioSelected
			}

			if activePanel == models.PanelPackageManager {
				if i == cursor {
					fmt.Fprintf(v, "%s%s %s%s\n", colors.ANSIMagenta, marker, method, colors.ANSIReset)
				} else {
					fmt.Fprintf(v, "%s%s %s%s\n", colors.ANSIGreen, marker, method, colors.ANSIReset)
				}
			} else {
				if i == cursor {
					fmt.Fprintf(v, "%s%s %s%s\n", colors.ANSIMagenta, marker, method, colors.ANSIReset)
				} else {
					fmt.Fprintf(v, "%s %s\n", marker, method)
				}
			}
		}
		if activePanel == models.PanelPackageManager {
			// Set cursor position RELATIVE to visible area
			v.SetCursor(0, cursor-offset)
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

		// Get actual rendered dimensions from the view (excludes borders automatically)
		_, visibleCount := v.Size()
		if visibleCount < 1 {
			visibleCount = 1
		}

		// Update scroll state
		state.ActionScroll.VisibleCount = visibleCount
		state.ActionScroll.ItemCount = len(actions)

		// Safety check: ensure cursor is visible
		if state.ActionScroll.Cursor >= state.ActionScroll.ItemCount {
			state.ActionScroll.Cursor = state.ActionScroll.ItemCount - 1
		}
		if state.ActionScroll.Cursor < 0 {
			state.ActionScroll.Cursor = 0
		}
		// Ensure offset shows the cursor
		if state.ActionScroll.Cursor < state.ActionScroll.Offset {
			state.ActionScroll.Offset = state.ActionScroll.Cursor
		} else if state.ActionScroll.Cursor >= state.ActionScroll.Offset+visibleCount {
			state.ActionScroll.Offset = state.ActionScroll.Cursor - visibleCount + 1
		}

		// Calculate scroll offset to keep cursor visible
		offset := state.ActionScroll.Offset
		cursor := state.ActionScroll.Cursor

		// Ensure cursor is within bounds
		if cursor >= len(actions) {
			cursor = len(actions) - 1
			state.ActionScroll.Cursor = cursor
		}
		if cursor < 0 {
			cursor = 0
			state.ActionScroll.Cursor = cursor
		}

		// Adjust offset to keep cursor visible
		if cursor < offset {
			offset = cursor
		} else if cursor >= offset+visibleCount {
			offset = cursor - visibleCount + 1
		}
		state.ActionScroll.Offset = offset

		// Set scroll position
		v.SetOrigin(0, offset)

		// Render ALL actions (gocui's SetOrigin will handle which ones are visible)
		for i := 0; i < len(actions); i++ {
			action := strings.ToLower(actions[i])
			marker := constants.RadioUnselected
			if i == cursor {
				marker = constants.RadioSelected
			}

			if activePanel == models.PanelAction {
				if i == cursor {
					fmt.Fprintf(v, "%s%s %s%s\n", colors.ANSIMagenta, marker, action, colors.ANSIReset)
				} else {
					fmt.Fprintf(v, "%s%s %s%s\n", colors.ANSIGreen, marker, action, colors.ANSIReset)
				}
			} else {
				if i == cursor {
					fmt.Fprintf(v, "%s%s %s%s\n", colors.ANSIMagenta, marker, action, colors.ANSIReset)
				} else {
					fmt.Fprintf(v, "%s %s\n", marker, action)
				}
			}
		}
		if activePanel == models.PanelAction {
			// Set cursor position RELATIVE to visible area
			v.SetCursor(0, cursor-offset)
		}
	}
	return nil
}

// renderToolsPanel renders the Tools selection panel with scrollbar indicator
func renderToolsPanel(params ToolsPanelParams) error {
	g := params.Gui
	state := params.State
	activePanel := params.ActivePanel
	leftPanelWidth := params.LeftPanelWidth
	toolsStartY := params.ToolsStartY
	toolsHeight := params.ToolsHeight

	// SAFETY: Ensure height is positive
	if toolsHeight < 3 {
		toolsHeight = 3 // Minimum height to show anything
	}

	// Calculate correct end Y coordinate (bottom of tools panel)
	toolsEndY := toolsStartY + toolsHeight - 1

	if v, err := g.SetView(constants.PanelTools, 0, toolsStartY, leftPanelWidth, toolsEndY); err != nil {
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

		// Get actual rendered dimensions from the view (excludes borders automatically)
		_, visibleCount := v.Size()
		if visibleCount < 1 {
			visibleCount = 1
		}

		// Update scroll state - set visible count but don't adjust offset
		// Offset is managed by navigation functions (ScrollUp/Down, JumpToFirst/Last)
		state.ToolsScroll.VisibleCount = visibleCount
		state.ToolsScroll.ItemCount = len(state.Tools)

		// Safety check: ensure cursor is visible
		if state.ToolsScroll.Cursor >= state.ToolsScroll.ItemCount {
			state.ToolsScroll.Cursor = state.ToolsScroll.ItemCount - 1
		}
		if state.ToolsScroll.Cursor < 0 {
			state.ToolsScroll.Cursor = 0
		}
		// Ensure offset shows the cursor
		if state.ToolsScroll.Cursor < state.ToolsScroll.Offset {
			state.ToolsScroll.Offset = state.ToolsScroll.Cursor
		} else if state.ToolsScroll.Cursor >= state.ToolsScroll.Offset+visibleCount {
			state.ToolsScroll.Offset = state.ToolsScroll.Cursor - visibleCount + 1
		}

		// Calculate scroll offset to keep cursor visible
		offset := state.ToolsScroll.Offset
		cursor := state.ToolsScroll.Cursor

		// Ensure cursor is within bounds
		if cursor >= len(state.Tools) {
			cursor = len(state.Tools) - 1
			state.ToolsScroll.Cursor = cursor
		}
		if cursor < 0 {
			cursor = 0
			state.ToolsScroll.Cursor = cursor
		}

		// Adjust offset to keep cursor visible
		if cursor < offset {
			offset = cursor
		} else if cursor >= offset+visibleCount {
			offset = cursor - visibleCount + 1
		}
		state.ToolsScroll.Offset = offset

		// Set scroll position
		v.SetOrigin(0, offset)

		// Render ALL tools (gocui's SetOrigin will handle which ones are visible)
		for i := 0; i < len(state.Tools); i++ {
			tool := state.Tools[i]
			marker := constants.CheckboxUnselected
			if state.SelectedTools[tool] {
				marker = constants.CheckboxSelected
			}

			if activePanel == models.PanelTools {
				if i == cursor {
					fmt.Fprintf(v, "%s%s %s%s\n", colors.ANSIMagenta, marker, constants.GetToolDisplayName(tool), colors.ANSIReset)
				} else {
					fmt.Fprintf(v, "%s%s %s%s\n", colors.ANSIGreen, marker, constants.GetToolDisplayName(tool), colors.ANSIReset)
				}
			} else {
				if i == cursor {
					fmt.Fprintf(v, "%s%s %s%s\n", colors.ANSIMagenta, marker, constants.GetToolDisplayName(tool), colors.ANSIReset)
				} else {
					fmt.Fprintf(v, "%s %s\n", marker, constants.GetToolDisplayName(tool))
				}
			}
		}
		if activePanel == models.PanelTools {
			// Set cursor position RELATIVE to visible area
			v.SetCursor(0, cursor-offset)
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

package handlers

import (
	"testing"

	"github.com/youpele52/lazysetup/pkg/config"
	"github.com/youpele52/lazysetup/pkg/models"
)

// TestNextPanel_PanelSwitching tests panel navigation with Tab key.
// Priority: P2 - Navigation depends on correct panel switching.
// Tests that NextPanel cycles through panels correctly.
func TestNextPanel_PanelSwitching(t *testing.T) {
	t.Run("cycles through panels correctly", func(t *testing.T) {
		state := models.NewState()
		state.SetCurrentPage(models.PageMultiPanel)
		state.SetActivePanel(models.PanelPackageManager)

		handler := NextPanel(state)
		_ = handler(nil, nil)

		if state.GetActivePanel() != models.PanelAction {
			t.Errorf("Expected PanelAction, got %v", state.GetActivePanel())
		}

		_ = handler(nil, nil)
		if state.GetActivePanel() != models.PanelTools {
			t.Errorf("Expected PanelTools, got %v", state.GetActivePanel())
		}

		_ = handler(nil, nil)
		if state.GetActivePanel() != models.PanelPackageManager {
			t.Errorf("Expected PanelPackageManager, got %v", state.GetActivePanel())
		}
	})
}

// TestPrevPanel_PanelSwitching tests panel navigation with Shift+Tab.
// Priority: P2 - Navigation depends on correct panel switching.
// Tests that PrevPanel cycles through panels in reverse.
func TestPrevPanel_PanelSwitching(t *testing.T) {
	t.Run("cycles through panels in reverse", func(t *testing.T) {
		state := models.NewState()
		state.SetCurrentPage(models.PageMultiPanel)
		state.SetActivePanel(models.PanelPackageManager)

		handler := PrevPanel(state)
		_ = handler(nil, nil)

		if state.GetActivePanel() != models.PanelTools {
			t.Errorf("Expected PanelTools, got %v", state.GetActivePanel())
		}
	})
}

// TestCursorMovement_IndexBounds tests cursor movement stays within bounds.
// Priority: P2 - Prevents out-of-bounds panics.
// Tests that cursor up/down respects list boundaries.
func TestCursorMovement_IndexBounds(t *testing.T) {
	t.Run("cursor up at 0 stays at 0", func(t *testing.T) {
		state := models.NewState()
		state.PackageManagerScroll.Cursor = 0

		handler := CursorUp(state)
		_ = handler(nil, nil)

		if state.PackageManagerScroll.Cursor != 0 {
			t.Errorf("Expected index 0, got %d", state.PackageManagerScroll.Cursor)
		}
	})

	t.Run("cursor down at max stays at max", func(t *testing.T) {
		state := models.NewState()
		maxIndex := len(state.InstallMethods) - 1
		state.PackageManagerScroll.Cursor = maxIndex

		handler := CursorDown(state)
		_ = handler(nil, nil)

		if state.PackageManagerScroll.Cursor != maxIndex {
			t.Errorf("Expected index %d, got %d", maxIndex, state.PackageManagerScroll.Cursor)
		}
	})

	t.Run("cursor moves within bounds", func(t *testing.T) {
		state := models.NewState()
		state.PackageManagerScroll.Cursor = 1

		handlerUp := CursorUp(state)
		_ = handlerUp(nil, nil)
		if state.PackageManagerScroll.Cursor != 0 {
			t.Errorf("Expected index 0, got %d", state.PackageManagerScroll.Cursor)
		}

		handlerDown := CursorDown(state)
		_ = handlerDown(nil, nil)
		if state.PackageManagerScroll.Cursor != 1 {
			t.Errorf("Expected index 1, got %d", state.PackageManagerScroll.Cursor)
		}
	})
}

// TestToolsCursorMovement_IndexBounds tests tools cursor stays within bounds.
// Priority: P2 - Prevents out-of-bounds panics in tools list.
// Tests that tools cursor respects list boundaries.
func TestToolsCursorMovement_IndexBounds(t *testing.T) {
	t.Run("tools cursor up at 0 stays at 0", func(t *testing.T) {
		state := models.NewState()
		state.ToolsScroll.Cursor = 0

		handler := ToolsCursorUp(state)
		_ = handler(nil, nil)

		if state.ToolsScroll.Cursor != 0 {
			t.Errorf("Expected index 0, got %d", state.ToolsScroll.Cursor)
		}
	})

	t.Run("tools cursor down at max stays at max", func(t *testing.T) {
		state := models.NewState()
		maxIndex := len(state.Tools) - 1
		state.ToolsScroll.Cursor = maxIndex

		handler := ToolsCursorDown(state)
		_ = handler(nil, nil)

		if state.ToolsScroll.Cursor != maxIndex {
			t.Errorf("Expected index %d, got %d", maxIndex, state.ToolsScroll.Cursor)
		}
	})
}

// TestToggleTool_TogglesSelection tests tool selection toggle.
// Priority: P2 - Core user interaction for selecting tools.
// Tests that ToggleTool correctly toggles tool selection state.
func TestToggleTool_TogglesSelection(t *testing.T) {
	t.Run("toggles tool selection", func(t *testing.T) {
		state := models.NewState()
		state.ToolsScroll.Cursor = 0
		tool := state.Tools[0]
		state.SelectedTools[tool] = false

		handler := ToggleTool(state)
		_ = handler(nil, nil)

		if !state.SelectedTools[tool] {
			t.Error("Expected tool to be selected after toggle")
		}

		_ = handler(nil, nil)
		if state.SelectedTools[tool] {
			t.Error("Expected tool to be deselected after second toggle")
		}
	})
}

// TestSwitchToPanel_DirectSwitch tests direct panel switching.
// Priority: P2 - Navigation depends on direct panel access.
// Tests that SwitchToPanel sets the correct panel.
func TestSwitchToPanel_DirectSwitch(t *testing.T) {
	t.Run("switches to specific panel", func(t *testing.T) {
		state := models.NewState()
		state.SetCurrentPage(models.PageMultiPanel)
		state.SetActivePanel(models.PanelPackageManager)

		handler := SwitchToPanel(state, models.PanelTools)
		_ = handler(nil, nil)

		if state.GetActivePanel() != models.PanelTools {
			t.Errorf("Expected PanelTools, got %v", state.GetActivePanel())
		}
	})
}

// TestJumpToFirst_VimNavigation tests vim-style jump to first item.
// Priority: P1 - Vim-style navigation is a core user feature.
// Tests that JumpToFirst resets cursor to position 0 for all panels.
func TestJumpToFirst_VimNavigation(t *testing.T) {
	t.Run("jumps to first in package manager panel", func(t *testing.T) {
		state := models.NewState()
		state.SetCurrentPage(models.PageMultiPanel)
		state.SetActivePanel(models.PanelPackageManager)
		state.PackageManagerScroll.Cursor = 5

		handler := JumpToFirst(state)
		_ = handler(nil, nil)

		if state.PackageManagerScroll.Cursor != 0 {
			t.Errorf("Expected cursor 0, got %d", state.PackageManagerScroll.Cursor)
		}
	})

	t.Run("jumps to first in tools panel", func(t *testing.T) {
		state := models.NewState()
		state.SetCurrentPage(models.PageMultiPanel)
		state.SetActivePanel(models.PanelTools)
		state.ToolsScroll.Cursor = 10

		handler := JumpToFirst(state)
		_ = handler(nil, nil)

		if state.ToolsScroll.Cursor != 0 {
			t.Errorf("Expected cursor 0, got %d", state.ToolsScroll.Cursor)
		}
	})

	t.Run("disabled when sudo confirm is shown", func(t *testing.T) {
		state := models.NewState()
		state.SetCurrentPage(models.PageMultiPanel)
		state.SetActivePanel(models.PanelPackageManager)
		state.PackageManagerScroll.Cursor = 5
		state.SetShowSudoConfirm(true)

		handler := JumpToFirst(state)
		_ = handler(nil, nil)

		if state.PackageManagerScroll.Cursor != 5 {
			t.Errorf("Expected cursor to remain 5 when disabled, got %d", state.PackageManagerScroll.Cursor)
		}
	})
}

// TestJumpToLast_VimNavigation tests vim-style jump to last item.
// Priority: P1 - Vim-style navigation is a core user feature.
// Tests that JumpToLast moves cursor to the last item for all panels.
func TestJumpToLast_VimNavigation(t *testing.T) {
	t.Run("jumps to last in package manager panel", func(t *testing.T) {
		state := models.NewState()
		state.SetCurrentPage(models.PageMultiPanel)
		state.SetActivePanel(models.PanelPackageManager)
		state.PackageManagerScroll.Cursor = 0
		lastIndex := len(state.InstallMethods) - 1

		handler := JumpToLast(state)
		_ = handler(nil, nil)

		if state.PackageManagerScroll.Cursor != lastIndex {
			t.Errorf("Expected cursor %d, got %d", lastIndex, state.PackageManagerScroll.Cursor)
		}
	})

	t.Run("jumps to last in tools panel", func(t *testing.T) {
		state := models.NewState()
		state.SetCurrentPage(models.PageMultiPanel)
		state.SetActivePanel(models.PanelTools)
		state.ToolsScroll.Cursor = 0
		lastIndex := len(state.Tools) - 1

		handler := JumpToLast(state)
		_ = handler(nil, nil)

		if state.ToolsScroll.Cursor != lastIndex {
			t.Errorf("Expected cursor %d, got %d", lastIndex, state.ToolsScroll.Cursor)
		}
	})

	t.Run("updates selected action when in action panel", func(t *testing.T) {
		state := models.NewState()
		state.SetCurrentPage(models.PageMultiPanel)
		state.SetActivePanel(models.PanelAction)
		state.ActionScroll.Cursor = 0
		lastActionIndex := len(config.Actions) - 1

		handler := JumpToLast(state)
		_ = handler(nil, nil)

		if state.ActionScroll.Cursor != lastActionIndex {
			t.Errorf("Expected cursor %d, got %d", lastActionIndex, state.ActionScroll.Cursor)
		}
	})

	t.Run("disabled when sudo confirm is shown", func(t *testing.T) {
		state := models.NewState()
		state.SetCurrentPage(models.PageMultiPanel)
		state.SetActivePanel(models.PanelPackageManager)
		state.PackageManagerScroll.Cursor = 0
		state.SetShowSudoConfirm(true)

		handler := JumpToLast(state)
		_ = handler(nil, nil)

		if state.PackageManagerScroll.Cursor != 0 {
			t.Errorf("Expected cursor to remain 0 when disabled, got %d", state.PackageManagerScroll.Cursor)
		}
	})
}

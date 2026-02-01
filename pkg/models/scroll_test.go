package models

import (
	"testing"
)

// TestPanelScrollState_ScrollUp tests cursor movement upward.
// Priority: P1 - Core navigation functionality.
func TestPanelScrollState_ScrollUp(t *testing.T) {
	t.Run("scroll up decreases cursor", func(t *testing.T) {
		state := &PanelScrollState{
			Cursor:    5,
			ItemCount: 10,
		}
		state.ScrollUp()
		if state.Cursor != 4 {
			t.Errorf("Expected cursor 4, got %d", state.Cursor)
		}
	})

	t.Run("scroll up at 0 stays at 0", func(t *testing.T) {
		state := &PanelScrollState{
			Cursor:    0,
			ItemCount: 10,
		}
		state.ScrollUp()
		if state.Cursor != 0 {
			t.Errorf("Expected cursor 0, got %d", state.Cursor)
		}
	})

	t.Run("scroll up adjusts offset when cursor goes above", func(t *testing.T) {
		state := &PanelScrollState{
			Offset:       5,
			Cursor:       5,
			ItemCount:    20,
			VisibleCount: 10,
		}
		state.ScrollUp()
		if state.Offset != 4 {
			t.Errorf("Expected offset 4, got %d", state.Offset)
		}
	})
}

// TestPanelScrollState_ScrollDown tests cursor movement downward.
// Priority: P1 - Core navigation functionality.
func TestPanelScrollState_ScrollDown(t *testing.T) {
	t.Run("scroll down increases cursor", func(t *testing.T) {
		state := &PanelScrollState{
			Cursor:    5,
			ItemCount: 10,
		}
		state.ScrollDown()
		if state.Cursor != 6 {
			t.Errorf("Expected cursor 6, got %d", state.Cursor)
		}
	})

	t.Run("scroll down at max stays at max", func(t *testing.T) {
		state := &PanelScrollState{
			Cursor:    9,
			ItemCount: 10,
		}
		state.ScrollDown()
		if state.Cursor != 9 {
			t.Errorf("Expected cursor 9, got %d", state.Cursor)
		}
	})

	t.Run("scroll down adjusts offset when cursor goes below visible", func(t *testing.T) {
		state := &PanelScrollState{
			Offset:       0,
			Cursor:       9,
			ItemCount:    20,
			VisibleCount: 10,
		}
		state.ScrollDown()
		// Cursor moves to 10, which is at the boundary (offset + visible = 0 + 10 = 10)
		// EnsureCursorVisible adjusts offset to keep cursor in view
		if state.Cursor != 10 {
			t.Errorf("Expected cursor 10, got %d", state.Cursor)
		}
		// Offset should be adjusted so cursor 10 is visible (offset = 10 - 10 + 1 = 1)
		if state.Offset != 1 {
			t.Errorf("Expected offset 1, got %d", state.Offset)
		}
	})
}

// TestPanelScrollState_JumpToFirst tests vim-style jump to first item.
// Priority: P1 - Vim-style navigation feature.
func TestPanelScrollState_JumpToFirst(t *testing.T) {
	t.Run("jump to first resets cursor and offset", func(t *testing.T) {
		state := &PanelScrollState{
			Offset:       50,
			Cursor:       75,
			ItemCount:    100,
			VisibleCount: 10,
		}
		state.JumpToFirst()
		if state.Cursor != 0 {
			t.Errorf("Expected cursor 0, got %d", state.Cursor)
		}
		if state.Offset != 0 {
			t.Errorf("Expected offset 0, got %d", state.Offset)
		}
	})

	t.Run("jump to first with empty list", func(t *testing.T) {
		state := &PanelScrollState{
			Offset:       0,
			Cursor:       0,
			ItemCount:    0,
			VisibleCount: 10,
		}
		state.JumpToFirst()
		if state.Cursor != 0 {
			t.Errorf("Expected cursor 0, got %d", state.Cursor)
		}
	})
}

// TestPanelScrollState_JumpToLast tests vim-style jump to last item.
// Priority: P1 - Vim-style navigation feature.
func TestPanelScrollState_JumpToLast(t *testing.T) {
	t.Run("jump to last sets cursor to last item", func(t *testing.T) {
		state := &PanelScrollState{
			Offset:       0,
			Cursor:       0,
			ItemCount:    100,
			VisibleCount: 10,
		}
		state.JumpToLast()
		if state.Cursor != 99 {
			t.Errorf("Expected cursor 99, got %d", state.Cursor)
		}
	})

	t.Run("jump to last adjusts offset when items exceed visible", func(t *testing.T) {
		state := &PanelScrollState{
			Offset:       0,
			Cursor:       0,
			ItemCount:    100,
			VisibleCount: 10,
		}
		state.JumpToLast()
		expectedOffset := 90 // ItemCount - VisibleCount
		if state.Offset != expectedOffset {
			t.Errorf("Expected offset %d, got %d", expectedOffset, state.Offset)
		}
	})

	t.Run("jump to last keeps offset at 0 when all items fit", func(t *testing.T) {
		state := &PanelScrollState{
			Offset:       0,
			Cursor:       0,
			ItemCount:    5,
			VisibleCount: 10,
		}
		state.JumpToLast()
		if state.Offset != 0 {
			t.Errorf("Expected offset 0, got %d", state.Offset)
		}
	})

	t.Run("jump to last estimates visible count when not set", func(t *testing.T) {
		state := &PanelScrollState{
			Offset:       0,
			Cursor:       0,
			ItemCount:    100,
			VisibleCount: 0,
		}
		state.JumpToLast()
		if state.Cursor != 99 {
			t.Errorf("Expected cursor 99, got %d", state.Cursor)
		}
		// Should estimate 5 items and set offset to 95
		expectedOffset := 95
		if state.Offset != expectedOffset {
			t.Errorf("Expected offset %d (estimate), got %d", expectedOffset, state.Offset)
		}
	})

	t.Run("jump to last with empty list", func(t *testing.T) {
		state := &PanelScrollState{
			Offset:       0,
			Cursor:       0,
			ItemCount:    0,
			VisibleCount: 10,
		}
		state.JumpToLast()
		if state.Cursor != -1 {
			t.Errorf("Expected cursor -1 (ItemCount - 1), got %d", state.Cursor)
		}
	})
}

// TestPanelScrollState_EnsureCursorVisible tests offset adjustment for cursor visibility.
// Priority: P1 - Prevents cursor from being off-screen.
func TestPanelScrollState_EnsureCursorVisible(t *testing.T) {
	t.Run("adjusts offset when cursor above visible area", func(t *testing.T) {
		state := &PanelScrollState{
			Offset:       10,
			Cursor:       5,
			ItemCount:    50,
			VisibleCount: 10,
		}
		state.EnsureCursorVisible()
		if state.Offset != 5 {
			t.Errorf("Expected offset 5, got %d", state.Offset)
		}
	})

	t.Run("adjusts offset when cursor below visible area", func(t *testing.T) {
		state := &PanelScrollState{
			Offset:       0,
			Cursor:       15,
			ItemCount:    50,
			VisibleCount: 10,
		}
		state.EnsureCursorVisible()
		if state.Offset != 6 {
			t.Errorf("Expected offset 6, got %d", state.Offset)
		}
	})

	t.Run("does not adjust when cursor already visible", func(t *testing.T) {
		state := &PanelScrollState{
			Offset:       0,
			Cursor:       5,
			ItemCount:    50,
			VisibleCount: 10,
		}
		state.EnsureCursorVisible()
		if state.Offset != 0 {
			t.Errorf("Expected offset 0, got %d", state.Offset)
		}
	})

	t.Run("clamps offset to max", func(t *testing.T) {
		state := &PanelScrollState{
			Offset:       100,
			Cursor:       45,
			ItemCount:    50,
			VisibleCount: 10,
		}
		state.EnsureCursorVisible()
		maxOffset := 40 // ItemCount - VisibleCount
		if state.Offset > maxOffset {
			t.Errorf("Expected offset clamped to max %d, got %d", maxOffset, state.Offset)
		}
	})

	t.Run("clamps offset to 0 when negative", func(t *testing.T) {
		state := &PanelScrollState{
			Offset:       -10,
			Cursor:       5,
			ItemCount:    50,
			VisibleCount: 10,
		}
		state.EnsureCursorVisible()
		if state.Offset != 0 {
			t.Errorf("Expected offset 0, got %d", state.Offset)
		}
	})
}

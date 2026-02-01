package models

// PanelScrollState tracks scroll position for any panel.
// Used by all selectable panels (package managers, actions, tools) for unified scroll management.
type PanelScrollState struct {
	Offset       int // First visible item index
	Cursor       int // Currently selected item index
	ItemCount    int // Total number of items
	VisibleCount int // How many items fit in view (calculated at render time)
}

// ScrollUp moves cursor up and adjusts offset.
// Priority: P1 - Core navigation functionality.
func (s *PanelScrollState) ScrollUp() {
	if s.Cursor > 0 {
		s.Cursor--
		s.EnsureCursorVisible()
	}
}

// ScrollDown moves cursor down and adjusts offset.
// Priority: P1 - Core navigation functionality.
func (s *PanelScrollState) ScrollDown() {
	if s.Cursor < s.ItemCount-1 {
		s.Cursor++
		s.EnsureCursorVisible()
	}
}

// JumpToFirst moves to first item.
// Priority: P1 - Vim-style navigation feature.
func (s *PanelScrollState) JumpToFirst() {
	s.Cursor = 0
	s.Offset = 0
}

// JumpToLast moves to last item with consistent padding.
// Priority: P1 - Vim-style navigation feature.
func (s *PanelScrollState) JumpToLast() {
	s.Cursor = s.ItemCount - 1

	// Fallback: estimate visible count if not set yet
	visibleCount := s.VisibleCount
	if visibleCount <= 0 {
		// Estimate based on typical terminal - at least show 5 items
		visibleCount = 5
	}

	// Only adjust offset if we have more items than visible space
	if s.ItemCount > visibleCount {
		// Calculate offset so last item appears at the bottom, matching top padding
		s.Offset = s.ItemCount - visibleCount
		if s.Offset < 0 {
			s.Offset = 0
		}
	} else {
		// If all items fit, show from beginning
		s.Offset = 0
	}
}

// EnsureCursorVisible adjusts offset so cursor is in view.
// Priority: P1 - Prevents cursor from being off-screen.
func (s *PanelScrollState) EnsureCursorVisible() {
	if s.Cursor < s.Offset {
		s.Offset = s.Cursor
	} else if s.Cursor >= s.Offset+s.VisibleCount {
		s.Offset = s.Cursor - s.VisibleCount + 1
	}
	// Clamp offset
	maxOffset := s.ItemCount - s.VisibleCount
	if maxOffset < 0 {
		maxOffset = 0
	}
	if s.Offset > maxOffset {
		s.Offset = maxOffset
	}
	if s.Offset < 0 {
		s.Offset = 0
	}
}

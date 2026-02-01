package models

// PanelScrollState tracks scroll position for any panel
type PanelScrollState struct {
	Offset       int // First visible item index
	Cursor       int // Currently selected item index
	ItemCount    int // Total number of items
	VisibleCount int // How many items fit in view (calculated at render time)
}

// ScrollUp moves cursor up and adjusts offset
func (s *PanelScrollState) ScrollUp() {
	if s.Cursor > 0 {
		s.Cursor--
		s.EnsureCursorVisible()
	}
}

// ScrollDown moves cursor down and adjusts offset
func (s *PanelScrollState) ScrollDown() {
	if s.Cursor < s.ItemCount-1 {
		s.Cursor++
		s.EnsureCursorVisible()
	}
}

// JumpToFirst moves to first item
func (s *PanelScrollState) JumpToFirst() {
	s.Cursor = 0
	s.Offset = 0
}

// JumpToLast moves to last item with consistent padding
func (s *PanelScrollState) JumpToLast() {
	s.Cursor = s.ItemCount - 1
	// Only adjust offset if visible count is known
	if s.VisibleCount > 0 && s.ItemCount > s.VisibleCount {
		// Calculate offset so last item appears at the bottom, matching top padding
		// When at offset 0, first item is at line 0
		// When at last page, last item should be at line (VisibleCount-1)
		// So offset = ItemCount - VisibleCount
		s.Offset = s.ItemCount - s.VisibleCount
		// Ensure we don't go past the last item
		if s.Offset < 0 {
			s.Offset = 0
		}
	} else {
		// If all items fit, show from beginning
		s.Offset = 0
	}
}

// EnsureCursorVisible adjusts offset so cursor is in view
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

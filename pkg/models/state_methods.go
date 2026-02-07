package models

import "context"

// IncrementSpinnerFrame safely increments the spinner frame
func (s *State) IncrementSpinnerFrame() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.SpinnerFrame = (s.SpinnerFrame + 1) % 10
}

// GetSpinnerFrame safely gets the current spinner frame
func (s *State) GetSpinnerFrame() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.SpinnerFrame
}

// SetToolStartTime safely sets the start time for a tool
func (s *State) SetToolStartTime(tool string, time int64) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.ToolStartTimes[tool] = time
}

// GetCancelContext returns the cancellation context for running installations
func (s *State) GetCancelContext() context.Context {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.CancelCtx
}

// CancelInstallations cancels all running installation goroutines
func (s *State) CancelInstallations() {
	s.mu.RLock()
	cancel := s.CancelFunc
	s.mu.RUnlock()
	if cancel != nil {
		cancel()
	}
}

// ResetCancelContext creates a new cancellation context after abort
func (s *State) ResetCancelContext() {
	s.mu.Lock()
	defer s.mu.Unlock()
	ctx, cancel := context.WithCancel(context.Background())
	s.CancelCtx = ctx
	s.CancelFunc = cancel
}

// GetCurrentTool safely gets the tool currently being installed
func (s *State) GetCurrentTool() string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.CurrentTool
}

// SetCurrentTool safely sets the tool currently being installed
func (s *State) SetCurrentTool(tool string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.CurrentTool = tool
}

// GetSelectedMethod safely gets the selected installation method
func (s *State) GetSelectedMethod() string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.SelectedMethod
}

// SetSelectedMethod safely sets the selected installation method
func (s *State) SetSelectedMethod(method string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.SelectedMethod = method
}

// GetSelectedTools safely gets a copy of selected tools
func (s *State) GetSelectedTools() map[string]bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	tools := make(map[string]bool)
	for k, v := range s.SelectedTools {
		tools[k] = v
	}
	return tools
}

// SetSelectedTools safely sets the selected tools
func (s *State) SetSelectedTools(tools map[string]bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.SelectedTools = tools
}

// GetCurrentPage safely gets the current page
func (s *State) GetCurrentPage() Page {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.CurrentPage
}

// SetCurrentPage safely sets the current page
func (s *State) SetCurrentPage(page Page) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.CurrentPage = page
}

// GetActivePanel safely gets the active panel
func (s *State) GetActivePanel() Panel {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.ActivePanel
}

// SetActivePanel safely sets the active panel
func (s *State) SetActivePanel(panel Panel) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.ActivePanel = panel
}

// ClearToolStartTimes safely clears tool start times
func (s *State) ClearToolStartTimes() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.ToolStartTimes = make(map[string]int64)
}

// GetToolStartTime safely gets the start time for a tool
func (s *State) GetToolStartTime(tool string) int64 {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.ToolStartTimes[tool]
}

// GetLastEscapeTime safely gets the last escape time
func (s *State) GetLastEscapeTime() int64 {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.LastEscapeTime
}

// SetLastEscapeTime safely sets the last escape time
func (s *State) SetLastEscapeTime(time int64) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.LastEscapeTime = time
}

// GetSelectedAction safely gets the selected action type
func (s *State) GetSelectedAction() ActionType {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.SelectedAction
}

// Scroll state methods - thread-safe access to PanelScrollState

// GetPackageManagerScroll safely gets the package manager scroll state
func (s *State) GetPackageManagerScroll() PanelScrollState {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.PackageManagerScroll
}

// SetPackageManagerScroll safely sets the package manager scroll state
func (s *State) SetPackageManagerScroll(scroll PanelScrollState) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.PackageManagerScroll = scroll
}

// GetActionScroll safely gets the action scroll state
func (s *State) GetActionScroll() PanelScrollState {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.ActionScroll
}

// SetActionScroll safely sets the action scroll state
func (s *State) SetActionScroll(scroll PanelScrollState) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.ActionScroll = scroll
}

// GetToolsScroll safely gets the tools scroll state
func (s *State) GetToolsScroll() PanelScrollState {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.ToolsScroll
}

// SetToolsScroll safely sets the tools scroll state
func (s *State) SetToolsScroll(scroll PanelScrollState) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.ToolsScroll = scroll
}

// Search query accessors

// GetSearchQuery safely gets the current search query
func (s *State) GetSearchQuery() string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.SearchQuery
}

// SetSearchQuery safely sets the search query
func (s *State) SetSearchQuery(query string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.SearchQuery = query
}

// AppendSearchQuery safely appends a character to the search query
func (s *State) AppendSearchQuery(char rune) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.SearchQuery += string(char)
}

// BackspaceSearchQuery safely removes the last character from search query
func (s *State) BackspaceSearchQuery() {
	s.mu.Lock()
	defer s.mu.Unlock()
	if len(s.SearchQuery) > 0 {
		s.SearchQuery = s.SearchQuery[:len(s.SearchQuery)-1]
	}
}

// Filtered tools accessors

// GetFilteredTools safely gets the filtered tools list
func (s *State) GetFilteredTools() []string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.FilteredTools
}

// SetFilteredTools safely sets the filtered tools list
func (s *State) SetFilteredTools(tools []string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.FilteredTools = tools
}

// Search mode accessors

// GetIsSearchMode safely gets whether in search mode
func (s *State) GetIsSearchMode() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.IsSearchMode
}

// SetIsSearchMode safely sets search mode state
func (s *State) SetIsSearchMode(mode bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.IsSearchMode = mode
}

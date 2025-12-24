package models

import "context"

// AppendInstallOutput safely appends output to the installation output string
func (s *State) AppendInstallOutput(output string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.InstallOutput += output
}

// AddInstallResult safely adds an installation result
func (s *State) AddInstallResult(result InstallResult) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.InstallResults = append(s.InstallResults, result)
}

// IncrementInstallingIndex safely increments the installing index
func (s *State) IncrementInstallingIndex() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.InstallingIndex++
}

// SetInstallationDone safely sets the installation done flag
func (s *State) SetInstallationDone(done bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.InstallationDone = done
}

// GetInstallationDone safely gets the installation done flag
func (s *State) GetInstallationDone() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.InstallationDone
}

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

// GetInstallOutput safely gets the installation output
func (s *State) GetInstallOutput() string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.InstallOutput
}

// GetInstallResults safely gets a copy of installation results
func (s *State) GetInstallResults() []InstallResult {
	s.mu.RLock()
	defer s.mu.RUnlock()
	results := make([]InstallResult, len(s.InstallResults))
	copy(results, s.InstallResults)
	return results
}

// GetAbortInstallation safely gets the abort installation flag
func (s *State) GetAbortInstallation() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.AbortInstallation
}

// SetAbortInstallation safely sets the abort installation flag
func (s *State) SetAbortInstallation(abort bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.AbortInstallation = abort
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

// GetInstallingIndex safely gets the number of completed installations
func (s *State) GetInstallingIndex() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.InstallingIndex
}

// SetInstallingIndex safely sets the number of completed installations
func (s *State) SetInstallingIndex(index int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.InstallingIndex = index
}

// GetInstallStartTime safely gets the installation start time
func (s *State) GetInstallStartTime() int64 {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.InstallStartTime
}

// SetInstallStartTime safely sets the installation start time
func (s *State) SetInstallStartTime(time int64) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.InstallStartTime = time
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

// ClearInstallOutput safely clears the installation output
func (s *State) ClearInstallOutput() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.InstallOutput = ""
}

// ClearInstallResults safely clears all installation results
func (s *State) ClearInstallResults() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.InstallResults = []InstallResult{}
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

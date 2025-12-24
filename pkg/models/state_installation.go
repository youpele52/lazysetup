package models

// AppendInstallOutput safely appends output to the installation output string
func (s *State) AppendInstallOutput(output string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.InstallOutput += output
}

// ClearInstallOutput safely clears the installation output
func (s *State) ClearInstallOutput() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.InstallOutput = ""
}

// GetInstallOutput safely gets the installation output
func (s *State) GetInstallOutput() string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.InstallOutput
}

// AddInstallResult safely adds an installation result
func (s *State) AddInstallResult(result InstallResult) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.InstallResults = append(s.InstallResults, result)
}

// ClearInstallResults safely clears all installation results
func (s *State) ClearInstallResults() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.InstallResults = []InstallResult{}
}

// GetInstallResults safely gets a copy of installation results
func (s *State) GetInstallResults() []InstallResult {
	s.mu.RLock()
	defer s.mu.RUnlock()
	results := make([]InstallResult, len(s.InstallResults))
	copy(results, s.InstallResults)
	return results
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

// SetInstallingIndex safely sets the installing index
func (s *State) SetInstallingIndex(index int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.InstallingIndex = index
}

// IncrementInstallingIndex safely increments the installing index
func (s *State) IncrementInstallingIndex() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.InstallingIndex++
}

// GetInstallingIndex safely gets the installing index
func (s *State) GetInstallingIndex() int {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.InstallingIndex
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

// SetInstallStartTime safely sets the installation start time
func (s *State) SetInstallStartTime(time int64) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.InstallStartTime = time
}

// GetInstallStartTime safely gets the installation start time
func (s *State) GetInstallStartTime() int64 {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.InstallStartTime
}

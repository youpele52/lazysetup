package models

import (
	"context"
	"sync"

	"github.com/youpele52/lazysetup/pkg/config"
	"github.com/youpele52/lazysetup/pkg/tools"
)

// Page represents the current UI page being displayed
// Used to control which layout function to render
type Page string

const (
	PageMenu       Page = "menu"
	PageSelection  Page = "selection"
	PageTools      Page = "tools"
	PageInstalling Page = "installing"
	PageResults    Page = "results"
	PageMultiPanel Page = "multipanel"
)

// Panel represents the active panel in the multi-panel layout
// Only one panel can be active at a time for user interaction
type Panel int

const (
	PanelPackageManager Panel = 0
	PanelStatus         Panel = 1
	PanelAction         Panel = 2
	PanelTools          Panel = 3
)

// ActionType represents the type of action to perform on tools
type ActionType int

const (
	ActionCheck     ActionType = 0
	ActionInstall   ActionType = 1
	ActionUpdate    ActionType = 2
	ActionUninstall ActionType = 3
)

// InstallResult captures the outcome of a single tool installation
type InstallResult struct {
	Tool     string // Name of the tool that was installed
	Success  bool   // Whether installation completed successfully
	Error    string // Error message if installation failed
	Duration int64  // Time taken to install in seconds
	Retries  int    // Number of retry attempts made
}

// State holds all application state with thread-safe access
// Must be accessed through getter/setter methods to avoid race conditions
// as it is shared between UI thread and installation goroutines
type State struct {
	mu sync.RWMutex

	InstallMethods []string   // Available installation methods (Homebrew, APT, etc.)
	SelectedMethod string     // Confirmed method to use for installation
	CheckStatus    string     // Status of method availability check
	Error          string     // Current error message to display
	CurrentPage    Page       // Current page being rendered
	ActivePanel    Panel      // Active panel in multi-panel layout (0-3)
	SelectedAction ActionType // Currently selected action (Install, Update, Delete)

	// Panel scroll states for automatic scrolling
	PackageManagerScroll PanelScrollState // Scroll state for package manager panel
	ActionScroll         PanelScrollState // Scroll state for action panel
	ToolsScroll          PanelScrollState // Scroll state for tools panel

	Tools            []string         // Available tools to install
	SelectedTools    map[string]bool  // Tools user selected for installation
	InstallResults   []InstallResult  // Results of completed installations
	InstallOutput    string           // Accumulated output from installation commands
	CurrentTool      string           // Tool currently being installed
	InstallingIndex  int              // Number of tools completed installing
	InstallationDone bool             // Whether all installations are finished
	SpinnerFrame     int              // Current animation frame (0-9) for spinner
	InstallStartTime int64            // Unix timestamp when installation started
	ToolStartTimes   map[string]int64 // Start time for each tool installation

	LastEscapeTime    int64              // Unix timestamp of last Esc key press (for double-escape detection)
	AbortInstallation bool               // Flag to signal running installations to abort
	CancelCtx         context.Context    // Context for cancelling running installations
	CancelFunc        context.CancelFunc // Function to cancel the context

	// Popup state
	ShowSudoConfirm bool       // Whether to show sudo confirmation popup
	PendingAction   ActionType // Action waiting for sudo confirmation
	SudoPassword    string     // Temporary sudo password (cleared after action)
	PasswordInput   string     // Current password input buffer

	// Results display state
	LastRenderedResultCount int   // Track how many results were last rendered to avoid duplication
	ActionCompletionTime    int64 // Unix timestamp when last action completed (for auto-clear timeout)

	// Update state
	UpdateAvailable   bool   // Whether an update is available
	UpdateVersion     string // Latest version available
	UpdateMessage     string // Update status message to display
	UpdateDownloadURL string // URL to download the update
	UpdateMessageTime int64  // Unix timestamp when update message was shown
}

func NewState() *State {
	ctx, cancel := context.WithCancel(context.Background())
	return &State{
		InstallMethods:       config.InstallMethods,
		SelectedMethod:       config.InstallMethods[0],
		CurrentPage:          PageMultiPanel,
		ActivePanel:          PanelPackageManager,
		SelectedAction:       ActionCheck,
		PackageManagerScroll: PanelScrollState{ItemCount: len(config.InstallMethods)},
		ActionScroll:         PanelScrollState{ItemCount: len(config.Actions)},
		ToolsScroll:          PanelScrollState{ItemCount: len(tools.Tools)},
		SelectedTools:        make(map[string]bool),
		InstallResults:       []InstallResult{},
		ToolStartTimes:       make(map[string]int64),
		Tools:                tools.Tools,
		CancelCtx:            ctx,
		CancelFunc:           cancel,
	}
}

func (s *State) Reset() {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.SelectedMethod = ""
	s.PackageManagerScroll.JumpToFirst()
	s.CheckStatus = ""
	s.Error = ""
	s.CurrentPage = PageMultiPanel
	s.SelectedTools = make(map[string]bool)
	s.ToolsScroll.JumpToFirst()
	s.ActionScroll.JumpToFirst()
	s.SelectedAction = ActionCheck
	s.InstallResults = []InstallResult{}
}

// ResetActionState clears action-related state after an action completes
func (s *State) ResetActionState() {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.InstallResults = []InstallResult{}
	s.InstallOutput = ""
	s.CurrentTool = ""
	s.InstallingIndex = 0
	s.InstallationDone = false
	s.SpinnerFrame = 0
	s.InstallStartTime = 0
	s.ToolStartTimes = make(map[string]int64)
	s.SelectedTools = make(map[string]bool)
	s.ToolsScroll.JumpToFirst()
	s.ActionScroll.JumpToFirst()
	s.SelectedAction = ActionCheck
}

// GetShowSudoConfirm returns whether the sudo confirmation popup is visible
func (s *State) GetShowSudoConfirm() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.ShowSudoConfirm
}

// SetShowSudoConfirm sets the sudo confirmation popup visibility
func (s *State) SetShowSudoConfirm(show bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.ShowSudoConfirm = show
}

// GetPendingAction returns the action waiting for sudo confirmation
func (s *State) GetPendingAction() ActionType {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.PendingAction
}

// SetPendingAction sets the action waiting for sudo confirmation
func (s *State) SetPendingAction(action ActionType) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.PendingAction = action
}

// GetSudoPassword returns the temporary sudo password
func (s *State) GetSudoPassword() string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.SudoPassword
}

// SetSudoPassword sets the temporary sudo password
func (s *State) SetSudoPassword(password string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.SudoPassword = password
}

// ClearSudoPassword clears the sudo password from memory
func (s *State) ClearSudoPassword() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.SudoPassword = ""
	s.PasswordInput = ""
}

// GetPasswordInput returns the current password input buffer
func (s *State) GetPasswordInput() string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.PasswordInput
}

// SetPasswordInput sets the password input buffer
func (s *State) SetPasswordInput(input string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.PasswordInput = input
}

// AppendPasswordInput appends a character to the password input
func (s *State) AppendPasswordInput(char rune) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.PasswordInput += string(char)
}

// BackspacePasswordInput removes the last character from password input
func (s *State) BackspacePasswordInput() {
	s.mu.Lock()
	defer s.mu.Unlock()
	if len(s.PasswordInput) > 0 {
		s.PasswordInput = s.PasswordInput[:len(s.PasswordInput)-1]
	}
}

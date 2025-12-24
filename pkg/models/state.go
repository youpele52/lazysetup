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
	ActionInstall   ActionType = 0
	ActionUpdate    ActionType = 1
	ActionUninstall ActionType = 2
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
	SelectedIndex  int        // Currently selected method index in menu
	SelectedMethod string     // Confirmed method to use for installation
	CheckStatus    string     // Status of method availability check
	Error          string     // Current error message to display
	CurrentPage    Page       // Current page being rendered
	ActivePanel    Panel      // Active panel in multi-panel layout (0-3)
	SelectedAction ActionType // Currently selected action (Install, Update, Delete)
	ActionIndex    int        // Currently selected action index in action panel

	Tools            []string         // Available tools to install
	SelectedTools    map[string]bool  // Tools user selected for installation
	ToolsIndex       int              // Current tool index in selection list
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
}

func NewState() *State {
	ctx, cancel := context.WithCancel(context.Background())
	return &State{
		InstallMethods: config.InstallMethods,
		SelectedIndex:  0,
		SelectedMethod: config.InstallMethods[0],
		CurrentPage:    PageMultiPanel,
		ActivePanel:    PanelPackageManager,
		SelectedAction: ActionInstall,
		ActionIndex:    0,
		SelectedTools:  make(map[string]bool),
		ToolsIndex:     0,
		InstallResults: []InstallResult{},
		ToolStartTimes: make(map[string]int64),
		Tools:          tools.Tools,
		CancelCtx:      ctx,
		CancelFunc:     cancel,
	}
}

func (s *State) Reset() {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.SelectedMethod = ""
	s.SelectedIndex = 0
	s.CheckStatus = ""
	s.Error = ""
	s.CurrentPage = PageMultiPanel
	s.SelectedTools = make(map[string]bool)
	s.ToolsIndex = 0
	s.ActionIndex = 0
	s.SelectedAction = ActionInstall
	s.InstallResults = []InstallResult{}
}

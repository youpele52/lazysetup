package ui

import "github.com/youpele52/lazysetup/pkg/models"

// ProgressMessageParams groups parameters for BuildInstallationProgressMessage
type ProgressMessageParams struct {
	SelectedMethod   string
	CurrentTool      string
	InstallingIndex  int
	TotalTools       int
	InstallationDone bool
	SpinnerFrame     int
	InstallOutput    string
	Action           models.ActionType
}

// PanelRenderParams groups parameters for panel rendering functions
type PanelRenderParams struct {
	Gui             *interface{} // gocui.Gui - using interface{} to avoid import cycle
	State           *models.State
	ActivePanel     models.Panel
	LeftPanelWidth  int
	PackageManagerY int
	ActionHeight    int // Used by renderActionPanel
	ToolsStartY     int // Used by renderToolsPanel
	PanelHeight     int // Used by renderToolsPanel
}

package constants

const (
	ViewMenu       = "menu"
	ViewResult     = "result"
	ViewResults    = "results"
	ViewTools      = "tools"
	ViewInstalling = "installing"

	PanelTools          = "panel_tools"
	PanelPackageManager = "panel_package_manager"
	PanelProgress       = "panel_progress"
	PanelAction         = "panel_action"
	PanelStatusView     = "Status"

	TitlePackageManager = "Package Manager"
	TitleInstalling     = "Installing"
	TitleTools          = "Tools"
	TitleAction         = "Action"
	TitleStatus         = "Status"
	TitleSelection      = "Details"

	TextInstalling = "installing"
	TextTools      = "tools"

	MessageSelected       = "You have selected %s for the installation\n"
	ErrorNoInstallCommand = "No install command found"
	ErrorNoToolsSelected  = "Please select at least one tool for installation"

	// Installation status constants
	StatusSuccess              = "success"
	StatusFailed               = "failed"
	StatusAlreadyInstalled     = "Already installed"
	StatusNotInstalled         = "Not installed"
	ErrorInstallationTimedOut  = "Installation timed out after 15 minutes"
	ErrorInstallationCancelled = "Installation was cancelled"

	RadioSelected   = "●"
	RadioUnselected = "○"

	CheckboxSelected   = "☑"
	CheckboxUnselected = "☐"

	ResultsSummaryTitle = "Installation Summary"
	ResultsSeparator    = "===================="
	ResultsSuccess      = "✓ %s - Success (%ds)\n"
	ResultsFailed       = "✗ %s - Failed (%ds)\n"
	ResultsError        = "  Error: %s\n"
	ResultsTotal        = "Total: %d Success, %d Failed\n"
)

const Logo = ` _                          _               
| | __ _ _____   _ ___  ___| |_ _   _ _ __  
| |/ _` + "`" + `_  / | | / __|/ _ \ __| | | | '_ \ 
| | (_| |/ /| |_| \__ \  __/ |_| |_| | |_) |
|_|\__,_/___|\__, |___/\___|\__|\__,_| .__/ 
             |___/                   |_|    

Copyright 2025 Youpele Michael
`

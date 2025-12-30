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
	PopupConfirm        = "popup_confirm"

	TitlePackageManager = "Package Manager"
	TitleInstalling     = "Installing"
	TitleTools          = "Tools"
	TitleAction         = "Action"
	TitleStatus         = "Status"
	TitleSelection      = "Details"

	TextInstalling = "installing"
	TextTools      = "tools"

	ResultsSummaryTitle = "Installation Summary"
	ResultsSeparator    = "===================="
	ResultsSuccess      = "✓ %s - Success (%ds)\n"
	ResultsFailed       = "✗ %s - Failed (%ds)\n"
	ResultsError        = "  Error: %s\n"
	ResultsTotal        = "Total: %d Success, %d Failed\n"

	// Update messages
	UpdateAvailable    = "Update available: v%s → v%s"
	UpdateDownloading  = "Downloading update..."
	UpdateInstalling   = "Installing update..."
	UpdateSuccess      = "Update installed! Press 'r' to restart."
	UpdateFailed       = "Update failed: %s"
	UpdateCheckFailed  = "Failed to check for updates: %s"
	UpdateNotAvailable = "You're running the latest version (v%s)"
)

const Logo = ` _                          _               
| | __ _ _____   _ ___  ___| |_ _   _ _ __  
| |/ _` + "`" + `_  / | | / __|/ _ \ __| | | | '_ \ 
| | (_| |/ /| |_| \__ \  __/ |_| |_| | |_) |
|_|\__,_/___|\__, |___/\___|\__|\__,_| .__/ 
             |___/                   |_|    

Copyright 2025 P.E.L.E.
`

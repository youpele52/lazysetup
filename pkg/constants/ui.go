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

	// Nix package manager disclaimer
	NixDisclaimer = "⚠️  NIX NOTICE: Requires nixpkgs channel configured.\n    Some tools may not be available via Nix.\n    Setup guide: https://nixos.org/manual/nix/stable/installation/installing-binary\n    If errors occur, try another package manager."
)

const Logo = ` _                          _               
| | __ _ _____   _ ___  ___| |_ _   _ _ __  
| |/ _` + "`" + `_  / | | / __|/ _ \ __| | | | '_ \ 
| | (_| |/ /| |_| \__ \  __/ |_| |_| | |_) |
|_|\__,_/___|\__, |___/\___|\__|\__,_| .__/ 
             |___/                   |_|    

Copyright 2026 P.E.L.E. - https://youpele.com/
`

// ToolDisplayNames maps internal tool identifiers to user-friendly display names
// Used in the UI to show more readable tool names (e.g., "claude code" instead of "claude-code")
var ToolDisplayNames = map[string]string{
	"claude-code": "claude code",
	"opencode":    "opencode",
}

// GetToolDisplayName returns the display name for a tool
// Returns the tool name itself if no display name mapping exists
func GetToolDisplayName(tool string) string {
	if displayName, ok := ToolDisplayNames[tool]; ok {
		return displayName
	}
	return tool
}

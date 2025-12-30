package constants

const (
	GeneralError            = "An error occurred"
	UnknownMethodError      = "Unknown method"
	NoInstallCommandError   = "No install command found"
	NoUpdateCommandError    = "No update command found"
	NoUninstallCommandError = "No uninstall command found"
	NoToolCheckCommandError = "No check command found for %s"
	NoToolsSelectedError    = "Please select at least one tool for installation"
	NoToolsSelected         = "No tools selected"

	InstallationTimedOut  = "Installation timed out after 15 minutes"
	InstallationCancelled = "Installation was cancelled"
	UpdateTimedOut        = "Update timed out"
	UpdateCancelled       = "Update was cancelled"
	UninstallTimedOut     = "Uninstall timed out"
	UninstallCancelled    = "Uninstall was cancelled"
	CheckTimedOut         = "Check timed out"
	CheckCancelled        = "Check was cancelled"

	StatusSuccess          = "success"
	StatusFailed           = "failed"
	StatusAlreadyInstalled = "Already installed"
	StatusNotInstalled     = "Not installed"
	StatusPending          = "pending"

	ToolActionCheck     = "check"
	ToolActionInstall   = "install"
	ToolActionUpdate    = "update"
	ToolActionUninstall = "uninstall"

	ErrorHtopCurlNotSupported = "htop cannot be installed via Curl. Please use Homebrew or APT instead."

	SudoConfirmTitle   = "Sudo Password Required"
	SudoConfirmMessage = "Enter your sudo password:\n\nPassword: %s\n\nPress Enter to confirm or Esc to cancel."

	PasswordMask = "•"

	MessageSelected = "You have selected %s for the installation\n"

	RadioSelected   = "●"
	RadioUnselected = "○"

	CheckboxSelected   = "☑"
	CheckboxUnselected = "☐"
)

package commands

// CheckCommands maps installation methods to commands that check if the package manager is installed
// Returns version string if available, error if package manager is not found
var CheckCommands = map[string]string{
	"Homebrew":   "brew --version",
	"Curl":       "curl --version",
	"APT":        "apt --version",
	"YUM":        "yum --version",
	"Scoop":      "scoop --version",
	"Chocolatey": "choco --version",
}

// GetCheckCommand retrieves the version check command for a specific package manager
// Returns empty string if method is not found
func GetCheckCommand(method string) string {
	if cmd, ok := CheckCommands[method]; ok {
		return cmd
	}
	return ""
}

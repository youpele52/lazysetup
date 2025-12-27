package commands

// PackageManagerCheckCommands maps installation methods to commands that check if the package manager is installed
// Returns version string if available, error if package manager is not found
var PackageManagerCheckCommands = CheckCommandsType{
	"Homebrew":   "brew --version",
	"Curl":       "curl --version",
	"APT":        "apt --version",
	"YUM":        "yum --version",
	"Scoop":      "scoop --version",
	"Chocolatey": "choco --version",
}

var ToolCheckCommands = CheckCommandsType{
	"git":        "git --version",
	"docker":     "docker --version",
	"lazygit":    "lazygit --version",
	"lazydocker": "lazydocker --version",
	"htop":       "htop --version",
}

var CheckCommands = MergeMaps(ToolCheckCommands, PackageManagerCheckCommands)

// GetCheckCommand retrieves the version check command for a specific package manager
// Returns empty string if method is not found
func GetCheckCommand(method string) string {
	return GetCheckCommandBase(GetCheckCommandType{
		method:            method,
		CheckCommandsType: CheckCommands,
	})
}

// GetToolCheckCommand retrieves the version check command for a specific tool
// Returns empty string if tool is not found
func GetToolCheckCommand(tool string) string {
	return GetCheckCommandBase(GetCheckCommandType{
		method:            tool,
		CheckCommandsType: ToolCheckCommands,
	})
}

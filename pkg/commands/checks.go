package commands

import (
	"github.com/youpele52/lazysetup/pkg/tools"
)

// PackageManagerCheckCommands maps installation methods to commands that check if the package manager is installed
// Returns version string if available, error if package manager is not found
var PackageManagerCheckCommands = CheckCommandsType{
	"Homebrew":   "brew --version",
	"Curl":       "curl --version",
	"APT":        "apt --version",
	"YUM":        "yum --version",
	"Scoop":      "scoop --version",
	"Chocolatey": "choco --version",
	"Pacman":     "pacman --version",
	"DNF":        "dnf --version",
	"Nix":        "nix --version",
}

// ToolCheckCommands maps tool names to their version check commands
// Auto-generated using helper functions for all tools in the tools package
var ToolCheckCommands = buildToolCheckCommands()

// buildToolCheckCommands creates check commands for all tools
func buildToolCheckCommands() CheckCommandsType {
	result := make(CheckCommandsType)

	for _, tool := range tools.Tools {
		result[tool] = GenerateCheckCommand(tool)
	}

	return result
}

// CheckCommands merges tool and package manager check commands
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

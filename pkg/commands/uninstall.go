package commands

import (
	"fmt"

	"github.com/youpele52/lazysetup/pkg/tools"
)

// PackageManagerUninstallCommands maps package managers to tools and their uninstall commands
// Standard package manager commands are auto-generated using helper functions.
// Curl commands use direct binary removal since Curl-based installs don't track packages.
var PackageManagerUninstallCommands = LifecycleCommandsType{
	"Homebrew":   buildToolMap("Homebrew", "uninstall"),
	"APT":        buildToolMap("APT", "uninstall"),
	"YUM":        buildToolMap("YUM", "uninstall"),
	"Scoop":      buildToolMap("Scoop", "uninstall"),
	"Chocolatey": buildToolMap("Chocolatey", "uninstall"),
	"Pacman":     buildToolMap("Pacman", "uninstall"),
	"Curl":       buildCurlUninstallMap(),
}

// buildCurlUninstallMap creates uninstall commands for Curl-based installs
// These remove binaries from common installation locations
func buildCurlUninstallMap() map[string]string {
	uninstallMap := make(map[string]string)

	for _, tool := range tools.Tools {
		switch tool {
		case "nvim":
			uninstallMap[tool] = "rm -rf /usr/local/bin/nvim /usr/local/share/nvim /usr/local/lib/nvim"
		case "zsh":
			uninstallMap[tool] = "rm -rf /usr/local/bin/zsh /usr/local/share/zsh ~/.oh-my-zsh"
		case "fzf":
			uninstallMap[tool] = "rm -rf ~/.fzf /usr/local/bin/fzf"
		case "zoxide":
			uninstallMap[tool] = "rm -rf ~/.local/bin/zoxide /usr/local/bin/zoxide"
		case "starship":
			uninstallMap[tool] = "rm -f /usr/local/bin/starship ~/.local/bin/starship"
		case "node":
			uninstallMap[tool] = "rm -rf /usr/local/bin/node /usr/local/bin/npm /usr/local/lib/node_modules /usr/local/include/node"
		case "python3":
			uninstallMap[tool] = "rm -rf /usr/local/bin/python3 /usr/local/bin/pip3 /usr/local/lib/python3*"
		case "gh":
			uninstallMap[tool] = "rm -f /usr/local/bin/gh"
		case "eza":
			uninstallMap[tool] = "rm -f /usr/local/bin/eza"
		case "delta":
			uninstallMap[tool] = "rm -f /usr/local/bin/delta"
		case "btop":
			uninstallMap[tool] = "rm -rf /usr/local/bin/btop ~/.config/btop"
		case "httpie":
			uninstallMap[tool] = "rm -f /usr/local/bin/http /usr/local/bin/https"
		case "lazysql":
			uninstallMap[tool] = "rm -f /usr/local/bin/lazysql"
		case "claude-code":
			uninstallMap[tool] = "rm -f /usr/local/bin/claude"
		case "opencode":
			uninstallMap[tool] = "rm -f /usr/local/bin/opencode"
		case "bun":
			uninstallMap[tool] = "rm -rf ~/.bun"
		case "uv":
			uninstallMap[tool] = "rm ~/.local/bin/uv ~/.local/bin/uvx"
		case "rsync":
			uninstallMap[tool] = "rm -f /usr/local/bin/rsync"
		case "kubectl":
			uninstallMap[tool] = "rm -f /usr/local/bin/kubectl ~/.kube/config"
		case "k9s":
			uninstallMap[tool] = "rm -rf /usr/local/bin/k9s ~/.config/k9s ~/.k9s"
		case "terraform":
			uninstallMap[tool] = "rm -f /usr/local/bin/terraform"
		case "helm":
			uninstallMap[tool] = "rm -f /usr/local/bin/helm ~/.config/helm"
		case "pnpm":
			uninstallMap[tool] = "rm -rf ~/.local/share/pnpm ~/.pnpm-store"
		default:
			uninstallMap[tool] = fmt.Sprintf("rm -f /usr/local/bin/%s", tool)
		}
	}

	return uninstallMap
}

// GetUninstallCommand retrieves the uninstall command for a specific tool using a specific method
// Returns empty string if method or tool is not found in the commands map
func GetUninstallCommand(method, tool string) string {
	return GetLifecycleCommand(GetLifecycleCommandType{
		method:                method,
		tool:                  tool,
		LifecycleCommandsType: PackageManagerUninstallCommands,
	})
}

package commands

import "fmt"

type LifecycleCommandsType map[string]map[string]string
type CheckCommandsType map[string]string

func MergeMaps[K comparable, V any](maps ...map[K]V) map[K]V {
	m := make(map[K]V)
	for _, src := range maps {
		for k, v := range src {
			m[k] = v
		}
	}
	return m
}

type GetLifecycleCommandType struct {
	LifecycleCommandsType
	method, tool string
}

func GetLifecycleCommand(params GetLifecycleCommandType) string {
	if methodCmds, ok := params.LifecycleCommandsType[params.method]; ok {
		if cmd, ok := methodCmds[params.tool]; ok {
			return cmd
		}
	}
	return ""
}

type GetCheckCommandType struct {
	CheckCommandsType
	method string
}

func GetCheckCommandBase(params GetCheckCommandType) string {
	if cmd, ok := params.CheckCommandsType[params.method]; ok {
		return cmd
	}
	return ""
}

// PackageNameMappings maps tool names to package manager-specific names
// Some tools have different package names across different package managers
var PackageNameMappings = map[string]map[string]string{
	"nvim": {
		"Homebrew":   "nvim",
		"APT":        "neovim",
		"YUM":        "neovim",
		"Scoop":      "neovim",
		"Chocolatey": "neovim",
		"Pacman":     "neovim",
	},
	"fd": {
		"Homebrew":   "fd",
		"APT":        "fd-find",
		"YUM":        "fd-find",
		"Scoop":      "fd",
		"Chocolatey": "fd",
		"Pacman":     "fd",
	},
	"docker": {
		"Homebrew":   "docker",
		"APT":        "docker.io",
		"YUM":        "docker",
		"Scoop":      "docker",
		"Chocolatey": "docker-desktop",
		"Pacman":     "docker",
	},
	"node": {
		"APT":    "nodejs",
		"YUM":    "nodejs",
		"Pacman": "nodejs",
	},
	"python3": {
		"Homebrew":   "python3",
		"APT":        "python3",
		"YUM":        "python3",
		"Scoop":      "python",
		"Chocolatey": "python",
		"Pacman":     "python",
	},
	"delta": {
		"APT":    "git-delta",
		"YUM":    "git-delta",
		"Pacman": "git-delta",
	},
	"httpie": {
		"APT": "httpie",
		"YUM": "httpie",
	},
	"claude-code": {
		"Homebrew": "claude-code",
	},
	"opencode": {
		"Homebrew": "opencode",
	},
}

// GetPackageName returns the package name for a tool on a specific package manager
// Returns the tool name itself if no mapping exists
func GetPackageName(tool, method string) string {
	if methodMap, ok := PackageNameMappings[tool]; ok {
		if pkgName, ok := methodMap[method]; ok {
			return pkgName
		}
	}
	return tool
}

// GenerateInstallCommand creates an install command for a tool using the specified method
func GenerateInstallCommand(method, tool string) string {
	pkgName := GetPackageName(tool, method)

	switch method {
	case "Homebrew":
		return fmt.Sprintf("brew install %s", pkgName)
	case "APT":
		return fmt.Sprintf("apt-get install -y %s", pkgName)
	case "YUM":
		return fmt.Sprintf("yum install -y %s", pkgName)
	case "Scoop":
		return fmt.Sprintf("scoop install %s", pkgName)
	case "Chocolatey":
		return fmt.Sprintf("choco install %s -y", pkgName)
	case "Pacman":
		return fmt.Sprintf("pacman -S --noconfirm %s", pkgName)
	case "DNF":
		return fmt.Sprintf("dnf install -y %s", pkgName)
	case "Nix":
		return fmt.Sprintf("nix-env -iA nixpkgs.%s", pkgName)
	case "Curl":
		// For Curl, return empty string - these need to be handled specially
		return ""
	default:
		return ""
	}
}

// GenerateUpdateCommand creates an update command for a tool using the specified method
func GenerateUpdateCommand(method, tool string) string {
	pkgName := GetPackageName(tool, method)

	switch method {
	case "Homebrew":
		return fmt.Sprintf("brew upgrade %s", pkgName)
	case "APT":
		return fmt.Sprintf("apt-get update && apt-get upgrade -y %s", pkgName)
	case "YUM":
		return fmt.Sprintf("yum update -y %s", pkgName)
	case "Scoop":
		return fmt.Sprintf("scoop update %s", pkgName)
	case "Chocolatey":
		return fmt.Sprintf("choco upgrade %s -y", pkgName)
	case "Pacman":
		return fmt.Sprintf("pacman -Syu --noconfirm %s", pkgName)
	case "DNF":
		return fmt.Sprintf("dnf upgrade -y %s", pkgName)
	case "Nix":
		return fmt.Sprintf("nix-env -u %s", pkgName)
	case "Curl":
		// For Curl, reinstall is the same as install
		return ""
	default:
		return ""
	}
}

// GenerateUninstallCommand creates an uninstall command for a tool using the specified method
func GenerateUninstallCommand(method, tool string) string {
	pkgName := GetPackageName(tool, method)

	switch method {
	case "Homebrew":
		return fmt.Sprintf("brew uninstall %s", pkgName)
	case "APT":
		return fmt.Sprintf("apt-get remove -y %s", pkgName)
	case "YUM":
		return fmt.Sprintf("yum remove -y %s", pkgName)
	case "Scoop":
		return fmt.Sprintf("scoop uninstall %s", pkgName)
	case "Chocolatey":
		return fmt.Sprintf("choco uninstall %s -y", pkgName)
	case "Pacman":
		return fmt.Sprintf("pacman -R --noconfirm %s", pkgName)
	case "DNF":
		return fmt.Sprintf("dnf remove -y %s", pkgName)
	case "Nix":
		return fmt.Sprintf("nix-env -e %s", pkgName)
	case "Curl":
		return fmt.Sprintf("rm -f /usr/local/bin/%s", tool)
	default:
		return ""
	}
}

// GenerateCheckCommand creates a version check command for a tool
func GenerateCheckCommand(tool string) string {
	switch tool {
	case "tmux":
		return "tmux -V"
	case "ripgrep":
		return "rg --version"
	case "claude-code":
		return "claude --version"
	default:
		return fmt.Sprintf("%s --version", tool)
	}
}

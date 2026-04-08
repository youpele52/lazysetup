package commands

import "fmt"

type LifecycleCommandsType map[string]map[string]string
type CheckCommandsType map[string]string

const (
	openclawUnixInstallScript    = "curl -fsSL https://openclaw.ai/install.sh | bash"
	openclawWindowsInstallScript = "powershell -c \"irm https://openclaw.ai/install.ps1 | iex\""
)

var supportedMethodsByTool = map[string]map[string]bool{
	"git": {
		"Homebrew": true, "APT": true, "YUM": true, "DNF": true, "Pacman": true, "Nix": true, "Scoop": true, "Chocolatey": true, "Curl": true,
	},
	"docker": {
		"Homebrew": true, "APT": true, "YUM": true, "DNF": true, "Pacman": true, "Nix": true, "Scoop": true, "Chocolatey": true, "Curl": true,
	},
	"lazygit": {
		"Homebrew": true, "APT": true, "YUM": true, "DNF": true, "Pacman": true, "Nix": true, "Scoop": true, "Chocolatey": true, "Curl": true,
	},
	"lazydocker": {
		"Homebrew": true, "APT": true, "YUM": true, "DNF": true, "Pacman": true, "Nix": true, "Scoop": true, "Chocolatey": true, "Curl": true,
	},
	"gh": {
		"Homebrew": true, "APT": true, "YUM": true, "DNF": true, "Pacman": true, "Nix": true, "Scoop": true, "Chocolatey": true, "Curl": true,
	},
	"make": {
		"Homebrew": true, "APT": true, "YUM": true, "DNF": true, "Pacman": true, "Nix": true, "Scoop": true, "Chocolatey": true, "Curl": true,
	},
	"just": {
		"Homebrew": true, "APT": true, "YUM": true, "DNF": true, "Pacman": true, "Nix": true, "Scoop": true, "Chocolatey": true, "Curl": true,
	},
	"zsh": {
		"Homebrew": true, "APT": true, "YUM": true, "DNF": true, "Pacman": true, "Nix": true, "Scoop": true, "Chocolatey": true, "Curl": true,
	},
	"tmux": {
		"Homebrew": true, "APT": true, "YUM": true, "DNF": true, "Pacman": true, "Nix": true, "Scoop": true, "Chocolatey": true, "Curl": true,
	},
	"nvim": {
		"Homebrew": true, "APT": true, "YUM": true, "DNF": true, "Pacman": true, "Nix": true, "Scoop": true, "Chocolatey": true, "Curl": true,
	},
	"node": {
		"Homebrew": true, "APT": true, "YUM": true, "DNF": true, "Pacman": true, "Nix": true, "Scoop": true, "Chocolatey": true, "Curl": true,
	},
	"python3": {
		"Homebrew": true, "APT": true, "YUM": true, "DNF": true, "Pacman": true, "Nix": true, "Scoop": true, "Chocolatey": true, "Curl": true,
	},
	"bun": {
		"Homebrew": true, "APT": true, "YUM": true, "DNF": true, "Pacman": true, "Nix": true, "Scoop": true, "Chocolatey": true, "Curl": true,
	},
	"pnpm": {
		"Homebrew": true, "APT": true, "YUM": true, "DNF": true, "Pacman": true, "Nix": true, "Scoop": true, "Chocolatey": true, "Curl": true,
	},
	"uv": {
		"Homebrew": true, "APT": true, "YUM": true, "DNF": true, "Pacman": true, "Nix": true, "Scoop": true, "Chocolatey": true, "Curl": true,
	},
	"claude-code": {
		"Homebrew": true, "Curl": true,
	},
	"opencode": {
		"Homebrew": true, "Curl": true,
	},
	"codex": {
		"Homebrew": true, "Curl": true,
	},
	"openclaw": {
		"Homebrew": true, "APT": true, "YUM": true, "DNF": true, "Pacman": true, "Nix": true, "Curl": true,
	},
	"ollama": {
		"Homebrew": true, "Curl": true,
	},
	"btop": {
		"Homebrew": true, "APT": true, "YUM": true, "DNF": true, "Pacman": true, "Nix": true, "Scoop": true, "Chocolatey": true, "Curl": true,
	},
	"fzf": {
		"Homebrew": true, "APT": true, "YUM": true, "DNF": true, "Pacman": true, "Nix": true, "Scoop": true, "Chocolatey": true, "Curl": true,
	},
	"ripgrep": {
		"Homebrew": true, "APT": true, "YUM": true, "DNF": true, "Pacman": true, "Nix": true, "Scoop": true, "Chocolatey": true, "Curl": true,
	},
	"fd": {
		"Homebrew": true, "APT": true, "YUM": true, "DNF": true, "Pacman": true, "Nix": true, "Scoop": true, "Chocolatey": true, "Curl": true,
	},
	"bat": {
		"Homebrew": true, "APT": true, "YUM": true, "DNF": true, "Pacman": true, "Nix": true, "Scoop": true, "Chocolatey": true, "Curl": true,
	},
	"eza": {
		"Homebrew": true, "APT": true, "YUM": true, "DNF": true, "Pacman": true, "Nix": true, "Scoop": true, "Chocolatey": true, "Curl": true,
	},
	"zoxide": {
		"Homebrew": true, "APT": true, "YUM": true, "DNF": true, "Pacman": true, "Nix": true, "Scoop": true, "Chocolatey": true, "Curl": true,
	},
	"tree": {
		"Homebrew": true, "APT": true, "YUM": true, "DNF": true, "Pacman": true, "Nix": true, "Scoop": true, "Chocolatey": true, "Curl": true,
	},
	"rsync": {
		"Homebrew": true, "APT": true, "YUM": true, "DNF": true, "Pacman": true, "Nix": true, "Scoop": true, "Chocolatey": true, "Curl": true,
	},
	"starship": {
		"Homebrew": true, "APT": true, "YUM": true, "DNF": true, "Pacman": true, "Nix": true, "Scoop": true, "Chocolatey": true, "Curl": true,
	},
	"delta": {
		"Homebrew": true, "APT": true, "YUM": true, "DNF": true, "Pacman": true, "Nix": true, "Scoop": true, "Chocolatey": true, "Curl": true,
	},
	"jq": {
		"Homebrew": true, "APT": true, "YUM": true, "DNF": true, "Pacman": true, "Nix": true, "Scoop": true, "Chocolatey": true, "Curl": true,
	},
	"yq": {
		"Homebrew": true, "APT": true, "YUM": true, "DNF": true, "Pacman": true, "Nix": true, "Scoop": true, "Chocolatey": true,
	},
	"kubectl": {
		"Homebrew": true, "APT": true, "YUM": true, "DNF": true, "Pacman": true, "Nix": true, "Scoop": true, "Chocolatey": true, "Curl": true,
	},
	"k9s": {
		"Homebrew": true, "APT": true, "YUM": true, "DNF": true, "Pacman": true, "Nix": true, "Scoop": true, "Chocolatey": true, "Curl": true,
	},
	"terraform": {
		"Homebrew": true, "APT": true, "YUM": true, "DNF": true, "Pacman": true, "Nix": true, "Scoop": true, "Chocolatey": true, "Curl": true,
	},
	"helm": {
		"Homebrew": true, "APT": true, "YUM": true, "DNF": true, "Pacman": true, "Nix": true, "Scoop": true, "Chocolatey": true, "Curl": true,
	},
	"podman": {
		"Homebrew": true, "APT": true, "YUM": true, "DNF": true, "Pacman": true, "Nix": true, "Scoop": true, "Chocolatey": true,
	},
	"httpie": {
		"Homebrew": true, "APT": true, "YUM": true, "DNF": true, "Pacman": true, "Nix": true, "Scoop": true, "Chocolatey": true, "Curl": true,
	},
	"wget": {
		"Homebrew": true, "APT": true, "YUM": true, "DNF": true, "Pacman": true, "Nix": true, "Scoop": true, "Chocolatey": true, "Curl": true,
	},
	"tldr": {
		"Homebrew": true, "APT": true, "YUM": true, "DNF": true, "Pacman": true, "Nix": true, "Scoop": true, "Chocolatey": true, "Curl": true,
	},
	"lazysql": {
		"Homebrew": true, "APT": true, "YUM": true, "DNF": true, "Pacman": true, "Nix": true, "Scoop": true, "Chocolatey": true, "Curl": true,
	},
	"direnv": {
		"Homebrew": true, "APT": true, "YUM": true, "DNF": true, "Pacman": true, "Nix": true, "Scoop": true, "Chocolatey": true,
	},
	"mise": {
		"Homebrew": true, "APT": true, "YUM": true, "DNF": true, "Pacman": true, "Nix": true, "Scoop": true, "Chocolatey": true,
	},
	"shellcheck": {
		"Homebrew": true, "APT": true, "YUM": true, "DNF": true, "Pacman": true, "Nix": true, "Scoop": true, "Chocolatey": true,
	},
	"pre-commit": {
		"Homebrew": true, "APT": true, "YUM": true, "DNF": true, "Pacman": true, "Nix": true, "Scoop": true, "Chocolatey": true,
	},
}

func MergeMaps[K comparable, V any](maps ...map[K]V) map[K]V {
	m := make(map[K]V)
	for _, src := range maps {
		for k, v := range src {
			m[k] = v
		}
	}
	return m
}

// IsToolSupportedByMethod returns whether a tool is available for a package manager.
func IsToolSupportedByMethod(tool, method string) bool {
	methodMap, ok := supportedMethodsByTool[tool]
	if !ok {
		return false
	}

	return methodMap[method]
}

// GetSupportedToolsForMethod returns all tools supported by the selected package manager.
func GetSupportedToolsForMethod(method string) []string {
	result := make([]string, 0)
	for tool, methodMap := range supportedMethodsByTool {
		if methodMap[method] {
			result = append(result, tool)
		}
	}
	return result
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
	"codex": {
		"Homebrew": "codex",
	},
	"openclaw": {
		"Homebrew": "openclaw",
	},
	"bun": {
		"Homebrew": "oven-sh/bun/bun",
	},
	"ollama": {
		"Nix": "ollama",
	},
	"yq": {
		"Scoop":      "yq",
		"Chocolatey": "yq",
	},
	"kubectl": {
		"Chocolatey": "kubernetes-cli",
	},
	"helm": {
		"Chocolatey": "kubernetes-helm",
	},
	"podman": {
		"Chocolatey": "podman-desktop",
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
	if !IsToolSupportedByMethod(tool, method) {
		return ""
	}

	pkgName := GetPackageName(tool, method)

	switch method {
	case "Homebrew":
		if tool == "openclaw" {
			return openclawUnixInstallScript
		}
		if tool == "codex" || tool == "openclaw" {
			return fmt.Sprintf("brew install --cask %s", pkgName)
		}
		return fmt.Sprintf("brew install %s", pkgName)
	case "APT":
		if tool == "openclaw" {
			return openclawUnixInstallScript
		}
		return fmt.Sprintf("apt-get install -y %s", pkgName)
	case "YUM":
		if tool == "openclaw" {
			return openclawUnixInstallScript
		}
		return fmt.Sprintf("yum install -y %s", pkgName)
	case "Scoop":
		return fmt.Sprintf("scoop install %s", pkgName)
	case "Chocolatey":
		return fmt.Sprintf("choco install %s -y", pkgName)
	case "Pacman":
		if tool == "openclaw" {
			return openclawUnixInstallScript
		}
		return fmt.Sprintf("pacman -S --noconfirm %s", pkgName)
	case "DNF":
		if tool == "openclaw" {
			return openclawUnixInstallScript
		}
		return fmt.Sprintf("dnf install -y %s", pkgName)
	case "Nix":
		if tool == "openclaw" {
			return openclawUnixInstallScript
		}
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
	if !IsToolSupportedByMethod(tool, method) {
		return ""
	}

	pkgName := GetPackageName(tool, method)

	switch method {
	case "Homebrew":
		if tool == "openclaw" {
			return openclawUnixInstallScript
		}
		if tool == "codex" || tool == "openclaw" {
			return fmt.Sprintf("brew upgrade --cask %s", pkgName)
		}
		return fmt.Sprintf("brew upgrade %s", pkgName)
	case "APT":
		if tool == "openclaw" {
			return openclawUnixInstallScript
		}
		return fmt.Sprintf("apt-get update && apt-get upgrade -y %s", pkgName)
	case "YUM":
		if tool == "openclaw" {
			return openclawUnixInstallScript
		}
		return fmt.Sprintf("yum update -y %s", pkgName)
	case "Scoop":
		return fmt.Sprintf("scoop update %s", pkgName)
	case "Chocolatey":
		return fmt.Sprintf("choco upgrade %s -y", pkgName)
	case "Pacman":
		if tool == "openclaw" {
			return openclawUnixInstallScript
		}
		return fmt.Sprintf("pacman -Syu --noconfirm %s", pkgName)
	case "DNF":
		if tool == "openclaw" {
			return openclawUnixInstallScript
		}
		return fmt.Sprintf("dnf upgrade -y %s", pkgName)
	case "Nix":
		if tool == "openclaw" {
			return openclawUnixInstallScript
		}
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
	if !IsToolSupportedByMethod(tool, method) {
		return ""
	}

	pkgName := GetPackageName(tool, method)

	switch method {
	case "Homebrew":
		if tool == "openclaw" {
			return getCurlUninstallCommand(tool)
		}
		if tool == "codex" || tool == "openclaw" {
			return fmt.Sprintf("brew uninstall --cask %s", pkgName)
		}
		return fmt.Sprintf("brew uninstall %s", pkgName)
	case "APT":
		if tool == "openclaw" {
			return getCurlUninstallCommand(tool)
		}
		return fmt.Sprintf("apt-get remove -y %s", pkgName)
	case "YUM":
		if tool == "openclaw" {
			return getCurlUninstallCommand(tool)
		}
		return fmt.Sprintf("yum remove -y %s", pkgName)
	case "Scoop":
		return fmt.Sprintf("scoop uninstall %s", pkgName)
	case "Chocolatey":
		return fmt.Sprintf("choco uninstall %s -y", pkgName)
	case "Pacman":
		if tool == "openclaw" {
			return getCurlUninstallCommand(tool)
		}
		return fmt.Sprintf("pacman -R --noconfirm %s", pkgName)
	case "DNF":
		if tool == "openclaw" {
			return getCurlUninstallCommand(tool)
		}
		return fmt.Sprintf("dnf remove -y %s", pkgName)
	case "Nix":
		if tool == "openclaw" {
			return getCurlUninstallCommand(tool)
		}
		return fmt.Sprintf("nix-env -e %s", pkgName)
	case "Curl":
		return getCurlUninstallCommand(tool)
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
	case "codex":
		return "codex --version"
	case "ollama":
		return "ollama -v"
	case "shellcheck":
		return "shellcheck --version"
	default:
		return fmt.Sprintf("%s --version", tool)
	}
}

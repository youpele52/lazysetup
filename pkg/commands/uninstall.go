package commands

// PackageManagerUninstallCommands maps package managers to tools and their uninstall commands
var PackageManagerUninstallCommands = LifecycleCommandsType{
	"Homebrew": {
		"git":        "brew uninstall git",
		"docker":     "brew uninstall docker",
		"lazygit":    "brew uninstall lazygit",
		"lazydocker": "brew uninstall lazydocker",
		"htop":       "brew uninstall htop",
	},
	"Curl": {
		"git":        "rm -f /usr/local/bin/git",
		"docker":     "rm -f /usr/local/bin/docker",
		"lazygit":    "rm -f /usr/local/bin/lazygit",
		"lazydocker": "rm -f /usr/local/bin/lazydocker",
		"htop":       "rm -f /usr/local/bin/htop",
	},
	"APT": {
		"git":        "apt-get remove -y git",
		"docker":     "apt-get remove -y docker.io",
		"lazygit":    "apt-get remove -y lazygit",
		"lazydocker": "apt-get remove -y lazydocker",
		"htop":       "apt-get remove -y htop",
	},
	"YUM": {
		"git":        "yum remove -y git",
		"docker":     "yum remove -y docker",
		"lazygit":    "yum remove -y lazygit",
		"lazydocker": "yum remove -y lazydocker",
		"htop":       "yum remove -y htop",
	},
	"Scoop": {
		"git":        "scoop uninstall git",
		"docker":     "scoop uninstall docker",
		"lazygit":    "scoop uninstall lazygit",
		"lazydocker": "scoop uninstall lazydocker",
		"htop":       "scoop uninstall htop",
	},
	"Chocolatey": {
		"git":        "choco uninstall git -y",
		"docker":     "choco uninstall docker-desktop -y",
		"lazygit":    "choco uninstall lazygit -y",
		"lazydocker": "choco uninstall lazydocker -y",
		"htop":       "choco uninstall htop -y",
	},
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

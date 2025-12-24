package commands

// InstallCommands maps installation methods to tools and their installation commands
// Outer key: package manager (Homebrew, Curl, APT, YUM, Scoop, Chocolatey)
// Inner key: tool name (git, docker, lazygit, lazydocker)
// Value: shell command to install the tool
var InstallCommands = map[string]map[string]string{
	"Homebrew": {
		"git":        "brew install git",
		"docker":     "brew install docker",
		"lazygit":    "brew install lazygit",
		"lazydocker": "brew install lazydocker",
	},
	"Curl": {
		"git":        "curl -fsSL https://git-scm.com/download/linux -o /tmp/git-installer.sh && chmod +x /tmp/git-installer.sh && /tmp/git-installer.sh",
		"docker":     "curl -fsSL https://get.docker.com -o /tmp/get-docker.sh && chmod +x /tmp/get-docker.sh && sh /tmp/get-docker.sh",
		"lazygit":    "curl -fsSL https://raw.githubusercontent.com/jesseduffield/lazygit/master/pkg/installer/install.sh -o /tmp/lazygit-install.sh && chmod +x /tmp/lazygit-install.sh && /tmp/lazygit-install.sh",
		"lazydocker": "curl -fsSL https://raw.githubusercontent.com/jesseduffield/lazydocker/master/scripts/install.sh -o /tmp/lazydocker-install.sh && chmod +x /tmp/lazydocker-install.sh && /tmp/lazydocker-install.sh",
	},
	"APT": {
		"git":        "apt-get install -y git",
		"docker":     "apt-get install -y docker.io",
		"lazygit":    "apt-get install -y lazygit",
		"lazydocker": "apt-get install -y lazydocker",
	},
	"YUM": {
		"git":        "yum install -y git",
		"docker":     "yum install -y docker",
		"lazygit":    "yum install -y lazygit",
		"lazydocker": "yum install -y lazydocker",
	},
	"Scoop": {
		"git":        "scoop install git",
		"docker":     "scoop install docker",
		"lazygit":    "scoop install lazygit",
		"lazydocker": "scoop install lazydocker",
	},
	"Chocolatey": {
		"git":        "choco install git -y",
		"docker":     "choco install docker-desktop -y",
		"lazygit":    "choco install lazygit -y",
		"lazydocker": "choco install lazydocker -y",
	},
}

// GetInstallCommand retrieves the installation command for a specific tool using a specific method
// Returns empty string if method or tool is not found in the commands map
func GetInstallCommand(method, tool string) string {
	if methodCmds, ok := InstallCommands[method]; ok {
		if cmd, ok := methodCmds[tool]; ok {
			return cmd
		}
	}
	return ""
}

// UpdateCommands maps package managers to tools and their update commands
var UpdateCommands = map[string]map[string]string{
	"Homebrew": {
		"git":        "brew upgrade git",
		"docker":     "brew upgrade docker",
		"lazygit":    "brew upgrade lazygit",
		"lazydocker": "brew upgrade lazydocker",
	},
	"Curl": {
		"git":        "curl -fsSL https://git-scm.com/download/linux -o /tmp/git-installer.sh && chmod +x /tmp/git-installer.sh && /tmp/git-installer.sh",
		"docker":     "curl -fsSL https://get.docker.com -o /tmp/get-docker.sh && chmod +x /tmp/get-docker.sh && sh /tmp/get-docker.sh",
		"lazygit":    "curl -fsSL https://raw.githubusercontent.com/jesseduffield/lazygit/master/pkg/installer/install.sh -o /tmp/lazygit-install.sh && chmod +x /tmp/lazygit-install.sh && /tmp/lazygit-install.sh",
		"lazydocker": "curl -fsSL https://raw.githubusercontent.com/jesseduffield/lazydocker/master/scripts/install.sh -o /tmp/lazydocker-install.sh && chmod +x /tmp/lazydocker-install.sh && /tmp/lazydocker-install.sh",
	},
	"APT": {
		"git":        "apt-get update && apt-get upgrade -y git",
		"docker":     "apt-get update && apt-get upgrade -y docker.io",
		"lazygit":    "apt-get update && apt-get upgrade -y lazygit",
		"lazydocker": "apt-get update && apt-get upgrade -y lazydocker",
	},
	"YUM": {
		"git":        "yum update -y git",
		"docker":     "yum update -y docker",
		"lazygit":    "yum update -y lazygit",
		"lazydocker": "yum update -y lazydocker",
	},
	"Scoop": {
		"git":        "scoop update git",
		"docker":     "scoop update docker",
		"lazygit":    "scoop update lazygit",
		"lazydocker": "scoop update lazydocker",
	},
	"Chocolatey": {
		"git":        "choco upgrade git -y",
		"docker":     "choco upgrade docker-desktop -y",
		"lazygit":    "choco upgrade lazygit -y",
		"lazydocker": "choco upgrade lazydocker -y",
	},
}

// DeleteCommands maps package managers to tools and their uninstall commands
var DeleteCommands = map[string]map[string]string{
	"Homebrew": {
		"git":        "brew uninstall git",
		"docker":     "brew uninstall docker",
		"lazygit":    "brew uninstall lazygit",
		"lazydocker": "brew uninstall lazydocker",
	},
	"Curl": {
		"git":        "rm -f /usr/local/bin/git",
		"docker":     "rm -f /usr/local/bin/docker",
		"lazygit":    "rm -f /usr/local/bin/lazygit",
		"lazydocker": "rm -f /usr/local/bin/lazydocker",
	},
	"APT": {
		"git":        "apt-get remove -y git",
		"docker":     "apt-get remove -y docker.io",
		"lazygit":    "apt-get remove -y lazygit",
		"lazydocker": "apt-get remove -y lazydocker",
	},
	"YUM": {
		"git":        "yum remove -y git",
		"docker":     "yum remove -y docker",
		"lazygit":    "yum remove -y lazygit",
		"lazydocker": "yum remove -y lazydocker",
	},
	"Scoop": {
		"git":        "scoop uninstall git",
		"docker":     "scoop uninstall docker",
		"lazygit":    "scoop uninstall lazygit",
		"lazydocker": "scoop uninstall lazydocker",
	},
	"Chocolatey": {
		"git":        "choco uninstall git -y",
		"docker":     "choco uninstall docker-desktop -y",
		"lazygit":    "choco uninstall lazygit -y",
		"lazydocker": "choco uninstall lazydocker -y",
	},
}

// GetUpdateCommand retrieves the update command for a specific tool using a specific method
// Returns empty string if method or tool is not found in the commands map
func GetUpdateCommand(method, tool string) string {
	if methodCmds, ok := UpdateCommands[method]; ok {
		if cmd, ok := methodCmds[tool]; ok {
			return cmd
		}
	}
	return ""
}

// GetDeleteCommand retrieves the uninstall/delete command for a specific tool using a specific method
// Returns empty string if method or tool is not found in the commands map
func GetDeleteCommand(method, tool string) string {
	if methodCmds, ok := DeleteCommands[method]; ok {
		if cmd, ok := methodCmds[tool]; ok {
			return cmd
		}
	}
	return ""
}

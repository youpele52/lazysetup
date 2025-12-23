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

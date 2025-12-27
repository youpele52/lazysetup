package commands

// PackageManagerInstallCommands maps installation methods to tools and their installation commands
// Outer key: package manager (Homebrew, Curl, APT, YUM, Scoop, Chocolatey)
// Inner key: tool name (git, docker, lazygit, lazydocker)
// Value: shell command to install the tool
var PackageManagerInstallCommands = LifecycleCommandsType{
	"Homebrew": {
		"git":        "brew install git",
		"docker":     "brew install docker",
		"lazygit":    "brew install lazygit",
		"lazydocker": "brew install lazydocker",
		"htop":       "brew install htop",
	},
	"Curl": {
		"git":        "curl -fsSL https://git-scm.com/download/linux -o /tmp/git-installer.sh && chmod +x /tmp/git-installer.sh && /tmp/git-installer.sh",
		"docker":     "curl -fsSL https://get.docker.com -o /tmp/get-docker.sh && chmod +x /tmp/get-docker.sh && sh /tmp/get-docker.sh",
		"lazygit":    "curl -fsSL https://raw.githubusercontent.com/jesseduffield/lazygit/master/pkg/installer/install.sh -o /tmp/lazygit-install.sh && chmod +x /tmp/lazygit-install.sh && /tmp/lazygit-install.sh",
		"lazydocker": "curl -fsSL https://raw.githubusercontent.com/jesseduffield/lazydocker/master/scripts/install.sh -o /tmp/lazydocker-install.sh && chmod +x /tmp/lazydocker-install.sh && /tmp/lazydocker-install.sh",
		"htop":       "curl -fsSL https://github.com/htop-dev/htop/releases/download/3.3.0/htop-3.3.0.tar.xz -o /tmp/htop.tar.xz && cd /tmp && tar -xf htop.tar.xz && cd htop-3.3.0 && ./configure && make && sudo make install",
	},
	"APT": {
		"git":        "apt-get install -y git",
		"docker":     "apt-get install -y docker.io",
		"lazygit":    "apt-get install -y lazygit",
		"lazydocker": "apt-get install -y lazydocker",
		"htop":       "apt-get install -y htop",
	},
	"YUM": {
		"git":        "yum install -y git",
		"docker":     "yum install -y docker",
		"lazygit":    "yum install -y lazygit",
		"lazydocker": "yum install -y lazydocker",
		"htop":       "yum install -y htop",
	},
	"Scoop": {
		"git":        "scoop install git",
		"docker":     "scoop install docker",
		"lazygit":    "scoop install lazygit",
		"lazydocker": "scoop install lazydocker",
		"htop":       "scoop install htop",
	},
	"Chocolatey": {
		"git":        "choco install git -y",
		"docker":     "choco install docker-desktop -y",
		"lazygit":    "choco install lazygit -y",
		"lazydocker": "choco install lazydocker -y",
		"htop":       "choco install htop -y",
	},
}

// GetInstallCommand retrieves the installation command for a specific tool using a specific method
// Returns empty string if method or tool is not found in the commands map
func GetInstallCommand(method, tool string) string {
	return GetLifecycleCommand(GetLifecycleCommandType{
		method:                method,
		tool:                  tool,
		LifecycleCommandsType: PackageManagerInstallCommands,
	})
}

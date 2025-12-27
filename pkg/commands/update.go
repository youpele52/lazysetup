package commands

// PackageManagerUpdateCommands maps package managers to tools and their update commands
var PackageManagerUpdateCommands = LifecycleCommandsType{
	"Homebrew": {
		"git":        "brew upgrade git",
		"docker":     "brew upgrade docker",
		"lazygit":    "brew upgrade lazygit",
		"lazydocker": "brew upgrade lazydocker",
		"htop":       "brew upgrade htop",
	},
	"Curl": {
		"git":        "curl -fsSL https://git-scm.com/download/linux -o /tmp/git-installer.sh && chmod +x /tmp/git-installer.sh && /tmp/git-installer.sh",
		"docker":     "curl -fsSL https://get.docker.com -o /tmp/get-docker.sh && chmod +x /tmp/get-docker.sh && sh /tmp/get-docker.sh",
		"lazygit":    "curl -fsSL https://raw.githubusercontent.com/jesseduffield/lazygit/master/pkg/installer/install.sh -o /tmp/lazygit-install.sh && chmod +x /tmp/lazygit-install.sh && /tmp/lazygit-install.sh",
		"lazydocker": "curl -fsSL https://raw.githubusercontent.com/jesseduffield/lazydocker/master/scripts/install.sh -o /tmp/lazydocker-install.sh && chmod +x /tmp/lazydocker-install.sh && /tmp/lazydocker-install.sh",
		"htop":       "curl -fsSL https://github.com/htop-dev/htop/releases/download/3.3.0/htop-3.3.0.tar.xz -o /tmp/htop.tar.xz && cd /tmp && tar -xf htop.tar.xz && cd htop-3.3.0 && ./configure && make && sudo make install",
	},
	"APT": {
		"git":        "apt-get update && apt-get upgrade -y git",
		"docker":     "apt-get update && apt-get upgrade -y docker.io",
		"lazygit":    "apt-get update && apt-get upgrade -y lazygit",
		"lazydocker": "apt-get update && apt-get upgrade -y lazydocker",
		"htop":       "apt-get update && apt-get upgrade -y htop",
	},
	"YUM": {
		"git":        "yum update -y git",
		"docker":     "yum update -y docker",
		"lazygit":    "yum update -y lazygit",
		"lazydocker": "yum update -y lazydocker",
		"htop":       "yum update -y htop",
	},
	"Scoop": {
		"git":        "scoop update git",
		"docker":     "scoop update docker",
		"lazygit":    "scoop update lazygit",
		"lazydocker": "scoop update lazydocker",
		"htop":       "scoop update htop",
	},
	"Chocolatey": {
		"git":        "choco upgrade git -y",
		"docker":     "choco upgrade docker-desktop -y",
		"lazygit":    "choco upgrade lazygit -y",
		"lazydocker": "choco upgrade lazydocker -y",
		"htop":       "choco upgrade htop -y",
	},
}

// GetUpdateCommand retrieves the update command for a specific tool using a specific method
// Returns empty string if method or tool is not found in the commands map
func GetUpdateCommand(method, tool string) string {
	return GetLifecycleCommand(GetLifecycleCommandType{
		method:                method,
		tool:                  tool,
		LifecycleCommandsType: PackageManagerUpdateCommands,
	})
}

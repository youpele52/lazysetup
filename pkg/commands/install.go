package commands

var InstallCommands = map[string]map[string]string{
	"Homebrew": {
		"git":        "brew install git",
		"docker":     "brew install docker",
		"lazygit":    "brew install lazygit",
		"lazydocker": "brew install lazydocker",
	},
	"Curl": {
		"git":        "curl https://git-scm.com/download/linux -o git-installer.sh && bash git-installer.sh",
		"docker":     "curl https://get.docker.com -o get-docker.sh && bash get-docker.sh",
		"lazygit":    "curl https://raw.githubusercontent.com/jesseduffield/lazygit/master/pkg/installer/install.sh | bash",
		"lazydocker": "curl https://raw.githubusercontent.com/jesseduffield/lazydocker/master/scripts/install.sh | bash",
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

func GetInstallCommand(method, tool string) string {
	if methodCmds, ok := InstallCommands[method]; ok {
		if cmd, ok := methodCmds[tool]; ok {
			return cmd
		}
	}
	return ""
}

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
		"nvim":       "brew install nvim",
		"zsh":        "brew install zsh",
		"tmux":       "brew install tmux",
		"fzf":        "brew install fzf",
		"ripgrep":    "brew install ripgrep",
		"fd":         "brew install fd",
		"bat":        "brew install bat",
		"jq":         "brew install jq",
	},
	"Curl": {
		"git":        "curl -fsSL https://git-scm.com/download/linux -o /tmp/git-installer.sh && chmod +x /tmp/git-installer.sh && /tmp/git-installer.sh",
		"docker":     "curl -fsSL https://get.docker.com -o /tmp/get-docker.sh && chmod +x /tmp/get-docker.sh && sh /tmp/get-docker.sh",
		"lazygit":    "curl -fsSL https://raw.githubusercontent.com/jesseduffield/lazygit/master/pkg/installer/install.sh -o /tmp/lazygit-install.sh && chmod +x /tmp/lazygit-install.sh && /tmp/lazygit-install.sh",
		"lazydocker": "curl -fsSL https://raw.githubusercontent.com/jesseduffield/lazydocker/master/scripts/install.sh -o /tmp/lazydocker-install.sh && chmod +x /tmp/lazydocker-install.sh && /tmp/lazydocker-install.sh",
		"htop":       "curl -fsSL https://github.com/htop-dev/htop/releases/download/3.3.0/htop-3.3.0.tar.xz -o /tmp/htop.tar.xz && cd /tmp && tar -xf htop.tar.xz && cd htop-3.3.0 && ./configure && make && sudo make install",
		"nvim":       "curl -fsSL https://github.com/neovim/neovim/releases/latest/download/nvim-linux64.tar.gz -o /tmp/nvim.tar.gz && cd /tmp && tar -xzf nvim.tar.gz && sudo cp -r nvim-linux64/* /usr/local/",
		"zsh":        "curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh -o /tmp/zsh-install.sh && chmod +x /tmp/zsh-install.sh && sh /tmp/zsh-install.sh",
		"tmux":       "curl -fsSL https://github.com/tmux/tmux/releases/latest/download/tmux-3.4.tar.gz -o /tmp/tmux.tar.gz && cd /tmp && tar -xzf tmux.tar.gz && cd tmux-3.4 && ./configure && make && sudo make install",
		"fzf":        "curl -fsSL https://raw.githubusercontent.com/junegunn/fzf/master/install -o /tmp/fzf-install.sh && chmod +x /tmp/fzf-install.sh && /tmp/fzf-install.sh",
		"ripgrep":    "curl -fsSL https://github.com/BurntSushi/ripgrep/releases/latest/download/ripgrep_$(uname -m)-unknown-linux-gnu.tar.gz -o /tmp/rg.tar.gz && cd /tmp && tar -xzf rg.tar.gz && sudo cp ripgrep*/rg /usr/local/bin/",
		"fd":         "curl -fsSL https://github.com/sharkdp/fd/releases/latest/download/fd-$(uname -m)-unknown-linux-gnu.tar.gz -o /tmp/fd.tar.gz && cd /tmp && tar -xzf fd.tar.gz && sudo cp fd-*/fd /usr/local/bin/",
		"bat":        "curl -fsSL https://github.com/sharkdp/bat/releases/latest/download/bat-$(uname -m)-unknown-linux-gnu.tar.gz -o /tmp/bat.tar.gz && cd /tmp && tar -xzf bat.tar.gz && sudo cp bat-*/bat /usr/local/bin/",
		"jq":         "curl -fsSL https://github.com/stedolan/jq/releases/latest/download/jq-linux64 -o /tmp/jq && chmod +x /tmp/jq && sudo mv /tmp/jq /usr/local/bin/",
	},
	"APT": {
		"git":        "apt-get install -y git",
		"docker":     "apt-get install -y docker.io",
		"lazygit":    "apt-get install -y lazygit",
		"lazydocker": "apt-get install -y lazydocker",
		"htop":       "apt-get install -y htop",
		"nvim":       "apt-get install -y neovim",
		"zsh":        "apt-get install -y zsh",
		"tmux":       "apt-get install -y tmux",
		"fzf":        "apt-get install -y fzf",
		"ripgrep":    "apt-get install -y ripgrep",
		"fd":         "apt-get install -y fd-find",
		"bat":        "apt-get install -y bat",
		"jq":         "apt-get install -y jq",
	},
	"YUM": {
		"git":        "yum install -y git",
		"docker":     "yum install -y docker",
		"lazygit":    "yum install -y lazygit",
		"lazydocker": "yum install -y lazydocker",
		"htop":       "yum install -y htop",
		"nvim":       "yum install -y neovim",
		"zsh":        "yum install -y zsh",
		"tmux":       "yum install -y tmux",
		"fzf":        "yum install -y fzf",
		"ripgrep":    "yum install -y ripgrep",
		"fd":         "yum install -y fd-find",
		"bat":        "yum install -y bat",
		"jq":         "yum install -y jq",
	},
	"Scoop": {
		"git":        "scoop install git",
		"docker":     "scoop install docker",
		"lazygit":    "scoop install lazygit",
		"lazydocker": "scoop install lazydocker",
		"htop":       "scoop install htop",
		"nvim":       "scoop install neovim",
		"tmux":       "scoop install tmux",
		"fzf":        "scoop install fzf",
		"ripgrep":    "scoop install ripgrep",
		"fd":         "scoop install fd",
		"bat":        "scoop install bat",
		"jq":         "scoop install jq",
	},
	"Chocolatey": {
		"git":        "choco install git -y",
		"docker":     "choco install docker-desktop -y",
		"lazygit":    "choco install lazygit -y",
		"lazydocker": "choco install lazydocker -y",
		"htop":       "choco install htop -y",
		"nvim":       "choco install neovim -y",
		"fzf":        "choco install fzf -y",
		"ripgrep":    "choco install ripgrep -y",
		"fd":         "choco install fd -y",
		"bat":        "choco install bat -y",
		"jq":         "choco install jq -y",
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

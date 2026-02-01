package commands

// PackageManagerUpdateCommands maps package managers to tools and their update commands
// Standard package manager commands are auto-generated using helper functions.
// Curl commands remain hardcoded due to their complex nature (reinstall via download).
var PackageManagerUpdateCommands = LifecycleCommandsType{
	"Homebrew":   buildToolMap("Homebrew", "update"),
	"APT":        buildToolMap("APT", "update"),
	"YUM":        buildToolMap("YUM", "update"),
	"Scoop":      buildToolMap("Scoop", "update"),
	"Chocolatey": buildToolMap("Chocolatey", "update"),
	"Pacman":     buildToolMap("Pacman", "update"),
	"Curl": {
		"git":         "curl -fsSL https://git-scm.com/download/linux -o /tmp/git-installer.sh && chmod +x /tmp/git-installer.sh && /tmp/git-installer.sh",
		"docker":      "curl -fsSL https://get.docker.com -o /tmp/get-docker.sh && chmod +x /tmp/get-docker.sh && sh /tmp/get-docker.sh",
		"lazygit":     "curl -fsSL https://raw.githubusercontent.com/jesseduffield/lazygit/master/pkg/installer/install.sh -o /tmp/lazygit-install.sh && chmod +x /tmp/lazygit-install.sh && /tmp/lazygit-install.sh",
		"lazydocker":  "curl -fsSL https://raw.githubusercontent.com/jesseduffield/lazydocker/master/scripts/install.sh -o /tmp/lazydocker-install.sh && chmod +x /tmp/lazydocker-install.sh && /tmp/lazydocker-install.sh",
		"htop":        "curl -fsSL https://github.com/htop-dev/htop/releases/latest/download/htop.tar.xz -o /tmp/htop.tar.xz && cd /tmp && tar -xf htop.tar.xz && cd htop-* && ./autogen.sh && ./configure && make && sudo make install",
		"nvim":        "curl -fsSL https://github.com/neovim/neovim/releases/latest/download/nvim-linux64.tar.gz -o /tmp/nvim.tar.gz && cd /tmp && tar -xzf nvim.tar.gz && sudo cp -r nvim-linux64/* /usr/local/",
		"zsh":         "curl -fsSL https://raw.githubusercontent.com/ohmyzsh/ohmyzsh/master/tools/install.sh -o /tmp/zsh-install.sh && chmod +x /tmp/zsh-install.sh && sh /tmp/zsh-install.sh",
		"tmux":        "curl -fsSL https://github.com/tmux/tmux/releases/latest/download/tmux.tar.gz -o /tmp/tmux.tar.gz && cd /tmp && tar -xzf tmux.tar.gz && cd tmux-* && ./configure && make && sudo make install",
		"fzf":         "curl -fsSL https://raw.githubusercontent.com/junegunn/fzf/master/install -o /tmp/fzf-install.sh && chmod +x /tmp/fzf-install.sh && /tmp/fzf-install.sh",
		"ripgrep":     "curl -fsSL https://github.com/BurntSushi/ripgrep/releases/latest/download/ripgrep_$(uname -m)-unknown-linux-gnu.tar.gz -o /tmp/rg.tar.gz && cd /tmp && tar -xzf rg.tar.gz && sudo cp ripgrep*/rg /usr/local/bin/",
		"fd":          "curl -fsSL https://github.com/sharkdp/fd/releases/latest/download/fd-$(uname -m)-unknown-linux-gnu.tar.gz -o /tmp/fd.tar.gz && cd /tmp && tar -xzf fd.tar.gz && sudo cp fd-*/fd /usr/local/bin/",
		"bat":         "curl -fsSL https://github.com/sharkdp/bat/releases/latest/download/bat-$(uname -m)-unknown-linux-gnu.tar.gz -o /tmp/bat.tar.gz && cd /tmp && tar -xzf bat.tar.gz && sudo cp bat-*/bat /usr/local/bin/",
		"jq":          "curl -fsSL https://github.com/stedolan/jq/releases/latest/download/jq-linux64 -o /tmp/jq && chmod +x /tmp/jq && sudo mv /tmp/jq /usr/local/bin/",
		"node":        "curl -fsSL https://nodejs.org/dist/latest-v20.x/node-latest-v20.x-linux-x64.tar.xz -o /tmp/node.tar.xz && cd /tmp && tar -xf node.tar.xz && sudo cp -r node-*/* /usr/local/",
		"gh":          "curl -fsSL https://github.com/cli/cli/releases/latest/download/gh_$(uname -s)_$(uname -m).tar.gz -o /tmp/gh.tar.gz && cd /tmp && tar -xzf gh.tar.gz && sudo cp gh_*/bin/gh /usr/local/bin/",
		"eza":         "curl -fsSL https://github.com/eza-community/eza/releases/latest/download/eza_$(uname -m)-unknown-linux-gnu.tar.gz -o /tmp/eza.tar.gz && cd /tmp && tar -xzf eza.tar.gz && sudo cp eza /usr/local/bin/",
		"zoxide":      "curl -fsSL https://raw.githubusercontent.com/ajeetdsouza/zoxide/main/install.sh | sh",
		"starship":    "curl -fsSL https://starship.rs/install.sh | sh -s -- -y",
		"python3":     "curl -fsSL https://www.python.org/ftp/python/3.13.1/Python-3.13.1.tgz -o /tmp/python.tgz && cd /tmp && tar -xzf python.tgz && cd Python-* && ./configure --enable-optimizations && make && sudo make install",
		"delta":       "curl -fsSL https://github.com/dandavison/delta/releases/latest/download/delta-$(uname -m)-unknown-linux-gnu.tar.gz -o /tmp/delta.tar.gz && cd /tmp && tar -xzf delta.tar.gz && sudo cp delta-*/delta /usr/local/bin/",
		"btop":        "curl -fsSL https://github.com/aristocratos/btop/releases/latest/download/btop-$(uname -m)-linux-musl.tbz -o /tmp/btop.tbz && cd /tmp && tar -xjf btop.tbz && sudo cp btop/bin/btop /usr/local/bin/",
		"httpie":      "curl -fsSL https://github.com/httpie/cli/releases/latest/download/httpie-$(uname -m)-unknown-linux-musl.tar.gz -o /tmp/httpie.tar.gz && cd /tmp && tar -xzf httpie.tar.gz && sudo cp httpie*/bin/http /usr/local/bin/ && sudo cp httpie*/bin/https /usr/local/bin/",
		"lazysql":     "curl -fsSL https://github.com/jorgerojas26/lazysql/releases/latest/download/lazysql_$(uname -s)_$(uname -m).tar.gz -o /tmp/lazysql.tar.gz && cd /tmp && tar -xzf lazysql.tar.gz && sudo cp lazysql /usr/local/bin/",
		"tree":        "curl -fsSL http://mama.indstate.edu/users/ice/tree/src/tree-2.1.3.tgz -o /tmp/tree.tgz && cd /tmp && tar -xzf tree.tgz && cd tree-* && make && sudo make install",
		"make":        "curl -fsSL https://ftp.gnu.org/gnu/make/make-4.4.1.tar.gz -o /tmp/make.tar.gz && cd /tmp && tar -xzf make.tar.gz && cd make-* && ./configure && make && sudo make install",
		"wget":        "curl -fsSL https://ftp.gnu.org/gnu/wget/wget-latest.tar.gz -o /tmp/wget.tar.gz && cd /tmp && tar -xzf wget.tar.gz && cd wget-* && ./configure && make && sudo make install",
		"tldr":        "curl -fsSL https://github.com/tldr-pages/tlrc/releases/latest/download/tlrc-$(uname -m)-unknown-linux-musl -o /tmp/tldr && chmod +x /tmp/tldr && sudo mv /tmp/tldr /usr/local/bin/",
		"claude-code": "curl -fsSL https://claude.ai/install.sh | bash",
		"opencode":    "curl -fsSL https://opencode.ai/install | bash",
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

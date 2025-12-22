package commands

var CheckCommands = map[string]string{
	"Homebrew":   "brew --version",
	"Curl":       "curl --version",
	"APT":        "apt --version",
	"YUM":        "yum --version",
	"Scoop":      "scoop --version",
	"Chocolatey": "choco --version",
}

func GetCheckCommand(method string) string {
	if cmd, ok := CheckCommands[method]; ok {
		return cmd
	}
	return ""
}

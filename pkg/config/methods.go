package config

var InstallMethods = []string{
	"Homebrew",
	"APT",
	"Curl",
	"YUM",
	"Scoop",
	"Chocolatey",
	"Pacman",
	"DNF",
	"Nix",
}

var Actions = []string{
	"Check",
	"Install",
	"Update",
	"Uninstall",
}

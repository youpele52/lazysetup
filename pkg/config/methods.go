package config

var InstallMethods = []string{
	"Homebrew",
	"APT",
	"Curl",
	"YUM",
	"Scoop",
	"Chocolatey",
}

var Actions = []string{
	"Check",
	"Install",
	"Update",
	"Uninstall",
}

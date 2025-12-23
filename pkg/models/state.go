package models

import "github.com/youpele52/lazysetup/pkg/config"

type Page string

const (
	PageMenu       Page = "menu"
	PageSelection  Page = "selection"
	PageTools      Page = "tools"
	PageInstalling Page = "installing"
	PageResults    Page = "results"
)

type InstallResult struct {
	Tool     string
	Success  bool
	Error    string
	Duration int64
	Retries  int
}

type State struct {
	InstallMethods []string
	SelectedIndex  int
	SelectedMethod string
	CheckStatus    string
	Error          string
	CurrentPage    Page

	Tools            []string
	SelectedTools    map[string]bool
	ToolsIndex       int
	InstallResults   []InstallResult
	InstallOutput    string
	CurrentTool      string
	InstallingIndex  int
	InstallationDone bool
	SpinnerFrame     int
	InstallStartTime int64
	ToolStartTimes   map[string]int64
}

func NewState() *State {
	return &State{
		InstallMethods: config.InstallMethods,
		SelectedIndex:  0,
		SelectedMethod: "",
		CurrentPage:    PageMenu,
		SelectedTools:  make(map[string]bool),
		ToolsIndex:     0,
		InstallResults: []InstallResult{},
		ToolStartTimes: make(map[string]int64),
	}
}

func (s *State) Reset() {
	s.SelectedMethod = ""
	s.SelectedIndex = 0
	s.CheckStatus = ""
	s.Error = ""
	s.CurrentPage = PageMenu
	s.SelectedTools = make(map[string]bool)
	s.ToolsIndex = 0
	s.InstallResults = []InstallResult{}
}

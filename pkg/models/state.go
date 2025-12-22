package models

import "github.com/youpele52/lazysetup/pkg/config"

type Page string

const (
	PageMenu      Page = "menu"
	PageSelection Page = "selection"
)

type State struct {
	InstallMethods []string
	SelectedIndex  int
	SelectedMethod string
	CheckStatus    string
	Error          string
	CurrentPage    Page
}

func NewState() *State {
	return &State{
		InstallMethods: config.InstallMethods,
		SelectedIndex:  0,
		SelectedMethod: "",
		CurrentPage:    PageMenu,
	}
}

func (s *State) Reset() {
	s.SelectedMethod = ""
	s.SelectedIndex = 0
	s.CheckStatus = ""
	s.Error = ""
	s.CurrentPage = PageMenu
}

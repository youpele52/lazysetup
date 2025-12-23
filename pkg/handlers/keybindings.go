package handlers

import (
	"os/exec"
	"strings"
	"time"

	"github.com/jesseduffield/gocui"
	"github.com/youpele52/lazysetup/pkg/commands"
	"github.com/youpele52/lazysetup/pkg/constants"
	"github.com/youpele52/lazysetup/pkg/models"
	"github.com/youpele52/lazysetup/pkg/tools"
)

func CursorUp(state *models.State) func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		if state.SelectedIndex > 0 {
			state.SelectedIndex--
		}
		return nil
	}
}

func CursorDown(state *models.State) func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		if state.SelectedIndex < len(state.InstallMethods)-1 {
			state.SelectedIndex++
		}
		return nil
	}
}

func SelectMethod(state *models.State) func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		state.SelectedMethod = state.InstallMethods[state.SelectedIndex]
		state.CheckStatus, state.Error = checkInstallation(state.SelectedMethod)
		if state.Error == "" && state.CheckStatus == "Already installed" {
			state.Tools = tools.Tools
			state.SelectedTools = make(map[string]bool)
			state.ToolsIndex = 0
			state.CurrentPage = models.PageTools
		} else {
			state.CurrentPage = models.PageSelection
		}
		return nil
	}
}

func ToolsCursorUp(state *models.State) func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		if state.ToolsIndex > 0 {
			state.ToolsIndex--
		}
		return nil
	}
}

func ToolsCursorDown(state *models.State) func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		if state.ToolsIndex < len(state.Tools)-1 {
			state.ToolsIndex++
		}
		return nil
	}
}

func ToggleTool(state *models.State) func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		tool := state.Tools[state.ToolsIndex]
		state.SelectedTools[tool] = !state.SelectedTools[tool]
		return nil
	}
}

func StartInstallation(state *models.State) func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		state.InstallResults = []models.InstallResult{}
		state.CurrentPage = models.PageInstalling
		state.InstallingIndex = 0
		state.InstallOutput = ""
		state.InstallationDone = false
		state.SpinnerFrame = 0

		go func() {
			spinnerTicker := time.NewTicker(100 * time.Millisecond)
			defer spinnerTicker.Stop()

			spinnerDone := make(chan bool)

			go func() {
				for {
					select {
					case <-spinnerDone:
						return
					case <-spinnerTicker.C:
						if !state.InstallationDone {
							state.SpinnerFrame = (state.SpinnerFrame + 1) % 10
						}
					}
				}
			}()

			for _, tool := range state.Tools {
				if state.SelectedTools[tool] {
					state.CurrentTool = tool
					state.InstallOutput = "Installing " + tool + "...\n\n"
					status, errMsg, output := installToolWithOutput(state.SelectedMethod, tool)
					state.InstallOutput += output

					result := models.InstallResult{
						Tool:    tool,
						Success: status == "success",
						Error:   errMsg,
					}
					state.InstallResults = append(state.InstallResults, result)
					state.InstallingIndex++
				}
			}
			state.InstallationDone = true
			spinnerDone <- true
			time.Sleep(1 * time.Second)
			state.CurrentPage = models.PageResults
		}()

		return nil
	}
}

func installToolWithOutput(method, tool string) (string, string, string) {
	cmd := commands.GetInstallCommand(method, tool)
	if cmd == "" {
		return "failed", constants.ErrorNoInstallCommand, ""
	}

	parts := strings.Fields(cmd)
	execCmd := exec.Command(parts[0], parts[1:]...)
	output, err := execCmd.CombinedOutput()
	outputStr := string(output)

	if err != nil {
		errMsg := err.Error()
		return "failed", errMsg, outputStr
	}

	return "success", "", outputStr
}

func checkInstallation(method string) (string, string) {
	cmd := commands.GetCheckCommand(method)
	if cmd == "" {
		return "", "Unknown method"
	}

	parts := strings.Fields(cmd)
	execCmd := exec.Command(parts[0], parts[1:]...)
	output, err := execCmd.CombinedOutput()

	if err != nil {
		errMsg := err.Error()
		if len(output) > 0 {
			errMsg = string(output)
		}
		return "Not installed", errMsg
	}

	if len(output) > 0 {
		return "Already installed", ""
	}
	return "Not installed", ""
}

func GoBack(state *models.State) func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		state.Reset()
		return nil
	}
}

func Quit(g *gocui.Gui, v *gocui.View) error {
	return gocui.ErrQuit
}

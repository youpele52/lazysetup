package handlers

import (
	"math/rand"
	"os/exec"
	"strings"
	"sync"
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
		state.InstallStartTime = time.Now().Unix()
		state.ToolStartTimes = make(map[string]int64)

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

			var wg sync.WaitGroup
			var mu sync.Mutex
			resultsChan := make(chan models.InstallResult, len(state.Tools))

			for _, tool := range state.Tools {
				if state.SelectedTools[tool] {
					wg.Add(1)
					go func(toolName string) {
						defer wg.Done()
						state.ToolStartTimes[toolName] = time.Now().Unix()
						status, errMsg, output := installToolWithRetry(state.SelectedMethod, toolName)

						mu.Lock()
						state.InstallOutput += "Tool: " + toolName + "\n" + output + "\n"
						mu.Unlock()

						duration := time.Now().Unix() - state.ToolStartTimes[toolName]
						result := models.InstallResult{
							Tool:     toolName,
							Success:  status == "success",
							Error:    errMsg,
							Duration: duration,
							Retries:  0,
						}
						resultsChan <- result
					}(tool)
				}
			}

			go func() {
				wg.Wait()
				close(resultsChan)
			}()

			for result := range resultsChan {
				state.InstallResults = append(state.InstallResults, result)
				state.InstallingIndex++
			}

			state.InstallationDone = true
			spinnerDone <- true
			time.Sleep(1 * time.Second)
			state.CurrentPage = models.PageResults
		}()

		return nil
	}
}

func installToolWithRetry(method, tool string) (string, string, string) {
	maxRetries := 2
	var lastErr string
	var lastOutput string
	retryCount := 0

	for attempt := 0; attempt <= maxRetries; attempt++ {
		if attempt > 0 {
			randomDelay := time.Duration(rand.Intn(40)) * time.Second
			time.Sleep(randomDelay)
			retryCount++
		}

		status, errMsg, output := installToolWithOutput(method, tool)
		lastOutput = output
		lastErr = errMsg

		if status == "success" {
			return "success", "", output
		}
	}

	return "failed", lastErr, lastOutput
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

func NextPanel(state *models.State) func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		if state.CurrentPage == models.PageMultiPanel {
			state.ActivePanel = (state.ActivePanel + 1) % 3
		}
		return nil
	}
}

func PrevPanel(state *models.State) func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		if state.CurrentPage == models.PageMultiPanel {
			state.ActivePanel = (state.ActivePanel + 2) % 3
		}
		return nil
	}
}

func SwitchToPanel(state *models.State, panel models.Panel) func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		if state.CurrentPage == models.PageMultiPanel {
			state.ActivePanel = panel
		}
		return nil
	}
}

func MultiPanelCursorUp(state *models.State) func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		if state.CurrentPage == models.PageMultiPanel {
			switch state.ActivePanel {
			case models.PanelInstallation:
				if state.SelectedIndex > 0 {
					state.SelectedIndex--
				}
			case models.PanelTools:
				if state.ToolsIndex > 0 {
					state.ToolsIndex--
				}
			}
		}
		return nil
	}
}

func MultiPanelCursorDown(state *models.State) func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		if state.CurrentPage == models.PageMultiPanel {
			switch state.ActivePanel {
			case models.PanelInstallation:
				if state.SelectedIndex < len(state.InstallMethods)-1 {
					state.SelectedIndex++
				}
			case models.PanelTools:
				if state.ToolsIndex < len(state.Tools)-1 {
					state.ToolsIndex++
				}
			}
		}
		return nil
	}
}

func MultiPanelToggleTool(state *models.State) func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		if state.CurrentPage == models.PageMultiPanel && state.ActivePanel == models.PanelTools {
			tool := state.Tools[state.ToolsIndex]
			state.SelectedTools[tool] = !state.SelectedTools[tool]
		}
		return nil
	}
}

func MultiPanelSelectMethod(state *models.State) func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		if state.CurrentPage == models.PageMultiPanel && state.ActivePanel == models.PanelInstallation {
			state.SelectedMethod = state.InstallMethods[state.SelectedIndex]
			state.CheckStatus, state.Error = checkInstallation(state.SelectedMethod)

			// Only proceed if check passed (no error)
			if state.Error == "" {
				state.Tools = tools.Tools
				state.SelectedTools = make(map[string]bool)
				state.ToolsIndex = 0
				state.ActivePanel = models.PanelTools
			}
		}
		return nil
	}
}

func MultiPanelStartInstallation(state *models.State) func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		if state.CurrentPage == models.PageMultiPanel && state.ActivePanel == models.PanelTools {
			// Check if at least one tool is selected
			selectedCount := 0
			for _, selected := range state.SelectedTools {
				if selected {
					selectedCount++
				}
			}

			if selectedCount == 0 {
				state.Error = constants.ErrorNoToolsSelected
				return nil
			}

			state.InstallResults = []models.InstallResult{}
			state.InstallingIndex = 0
			state.InstallOutput = ""
			state.InstallationDone = false
			state.SpinnerFrame = 0
			state.InstallStartTime = time.Now().Unix()
			state.ToolStartTimes = make(map[string]int64)

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

				var wg sync.WaitGroup
				var mu sync.Mutex
				resultsChan := make(chan models.InstallResult, len(state.Tools))

				for _, tool := range state.Tools {
					if state.SelectedTools[tool] {
						wg.Add(1)
						go func(toolName string) {
							defer wg.Done()
							state.ToolStartTimes[toolName] = time.Now().Unix()
							status, errMsg, output := installToolWithRetry(state.SelectedMethod, toolName)

							mu.Lock()
							state.InstallOutput += "Tool: " + toolName + "\n" + output + "\n"
							mu.Unlock()

							duration := time.Now().Unix() - state.ToolStartTimes[toolName]
							result := models.InstallResult{
								Tool:     toolName,
								Success:  status == "success",
								Error:    errMsg,
								Duration: duration,
								Retries:  0,
							}
							resultsChan <- result
						}(tool)
					}
				}

				go func() {
					wg.Wait()
					close(resultsChan)
				}()

				for result := range resultsChan {
					state.InstallResults = append(state.InstallResults, result)
					state.InstallingIndex++
				}

				state.InstallationDone = true
				spinnerDone <- true
				time.Sleep(1 * time.Second)
			}()
		}
		return nil
	}
}

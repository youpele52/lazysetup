package handlers

import (
	"context"
	"sync"
	"time"

	"github.com/jesseduffield/gocui"
	"github.com/youpele52/lazysetup/pkg/commands"
	"github.com/youpele52/lazysetup/pkg/constants"
	"github.com/youpele52/lazysetup/pkg/executor"
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

// StartInstallation initiates parallel installation of selected tools
// It launches goroutines for each tool, collects results in a channel,
// and updates the UI in real-time with spinner animation
// NOTE: This function is for old single-page layout, use MultiPanelStartInstallation for multi-panel mode
func StartInstallation(state *models.State) func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		// Initialize installation state using thread-safe methods
		state.ClearInstallResults()
		state.ClearToolStartTimes()
		state.ClearInstallOutput()
		state.SetCurrentPage(models.PageInstalling)
		state.SetInstallingIndex(0)
		state.SetInstallationDone(false)
		state.SetInstallStartTime(time.Now().Unix())
		state.SetCurrentTool("")

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
						if !state.GetInstallationDone() {
							state.IncrementSpinnerFrame()
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
						startTime := time.Now().Unix()
						state.SetToolStartTime(toolName, startTime)
						status, errMsg, output := installToolWithRetry(state, state.SelectedMethod, toolName)

						mu.Lock()
						state.AppendInstallOutput("Tool: " + toolName + "\n" + output + "\n")
						mu.Unlock()

						duration := time.Now().Unix() - state.GetToolStartTime(toolName)
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
				state.AddInstallResult(result)
				state.IncrementInstallingIndex()
			}

			state.SetInstallationDone(true)
			spinnerDone <- true
			time.Sleep(1 * time.Second)
			state.CurrentPage = models.PageResults
		}()

		return nil
	}
}

// installToolWithRetry attempts to install a tool with automatic retry logic
// on failure. It will retry up to maxRetries times with exponential backoff.
// Returns: (status, errorMsg, output) where status is "success" or "failed"
func installToolWithRetry(state *models.State, method, tool string) (string, string, string) {
	maxRetries := 2
	var lastErr string
	var lastOutput string

	for attempt := 0; attempt <= maxRetries; attempt++ {
		if attempt > 0 {
			// Exponential backoff: 1s, 2s, 4s
			delay := time.Duration(1<<uint(attempt-1)) * time.Second
			time.Sleep(delay)
		}

		status, errMsg, output := installToolWithOutput(state, method, tool)
		lastOutput = output
		lastErr = errMsg

		if status == "success" {
			return "success", "", output
		}
	}

	return "failed", lastErr, lastOutput
}

// installToolWithOutput executes installation command with cancellation support
// Uses state's cancel context to allow aborting running installations
// Returns: (status, errorMsg, output) where status is "success" or "failed"
func installToolWithOutput(state *models.State, method, tool string) (string, string, string) {
	cmd := commands.GetInstallCommand(method, tool)
	if cmd == "" {
		return "failed", constants.ErrorNoInstallCommand, ""
	}

	ctx := state.GetCancelContext()
	result := executor.ExecuteWithTimeout(ctx, cmd, 15*time.Minute)

	if result.TimedOut {
		return "failed", "Installation timed out after 15 minutes", result.Output
	}
	if result.Cancelled {
		return "failed", "Installation was cancelled", result.Output
	}
	if result.ExitCode != 0 {
		return "failed", result.GetErrorMessage(), result.Output
	}

	return "success", "", result.Output
}

// checkInstallation verifies if a package manager is installed and available
// Returns: (status, errorMsg) where status is "Already installed" or "Not installed"
func checkInstallation(method string) (string, string) {
	cmd := commands.GetCheckCommand(method)
	if cmd == "" {
		return "", "Unknown method"
	}

	ctx := context.Background()
	result := executor.ExecuteWithTimeout(ctx, cmd, 10*time.Second)

	if result.ExitCode != 0 {
		errMsg := result.GetErrorMessage()
		if result.Output != "" {
			errMsg = result.Output
		}
		return "Not installed", errMsg
	}

	return "Already installed", ""
}

// GoBack implements double-escape key handling for aborting installations
// First Esc: records timestamp; second Esc within 500ms: resets state and aborts
// This pattern prevents accidental aborts while allowing quick exit
func GoBack(state *models.State) func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		now := time.Now().UnixMilli()
		lastEscape := state.GetLastEscapeTime()

		// Check if this is a double escape (within 500ms)
		if lastEscape > 0 && (now-lastEscape) < 500 {
			// Double escape: cancel running installations and reset state
			state.CancelInstallations()
			state.SetAbortInstallation(true)
			state.Reset()
			state.SetLastEscapeTime(0)
			state.ResetCancelContext()
		} else {
			// First escape: mark the time
			state.SetLastEscapeTime(now)

			// Reset the timestamp after 600ms if escape is not pressed again
			go func() {
				time.Sleep(600 * time.Millisecond)
				if state.GetLastEscapeTime() == now {
					state.SetLastEscapeTime(0)
				}
			}()
		}
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

// MultiPanelStartInstallation initiates parallel installation in multi-panel mode
// Validates tool selection, launches goroutines, collects results, and handles abort requests
func MultiPanelStartInstallation(state *models.State) func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		if state.GetCurrentPage() == models.PageMultiPanel && state.ActivePanel == models.PanelTools {
			// Check if at least one tool is selected
			selectedTools := state.GetSelectedTools()
			selectedCount := 0
			for _, selected := range selectedTools {
				if selected {
					selectedCount++
				}
			}

			if selectedCount == 0 {
				state.Error = constants.ErrorNoToolsSelected
				return nil
			}

			// Initialize installation state using thread-safe methods
			state.ClearInstallResults()
			state.ClearToolStartTimes()
			state.ClearInstallOutput()
			state.SetInstallingIndex(0)
			state.SetInstallationDone(false)
			state.SetInstallStartTime(time.Now().Unix())

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
							if !state.GetInstallationDone() {
								state.IncrementSpinnerFrame()
							}
						}
					}
				}()

				var wg sync.WaitGroup
				var mu sync.Mutex
				resultsChan := make(chan models.InstallResult, len(state.Tools))

				for _, tool := range state.Tools {
					if state.SelectedTools[tool] {
						// Check if abort was requested
						if state.GetAbortInstallation() {
							break
						}

						wg.Add(1)
						go func(toolName string) {
							defer wg.Done()

							// Check abort flag before starting installation
							if state.GetAbortInstallation() {
								return
							}

							startTime := time.Now().Unix()
							state.SetToolStartTime(toolName, startTime)
							status, errMsg, output := installToolWithRetry(state, state.SelectedMethod, toolName)

							mu.Lock()
							state.AppendInstallOutput("Tool: " + toolName + "\n" + output + "\n")
							mu.Unlock()

							duration := time.Now().Unix() - state.GetToolStartTime(toolName)
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
					state.AddInstallResult(result)
					state.IncrementInstallingIndex()
				}

				state.SetInstallationDone(true)
				spinnerDone <- true
				time.Sleep(1 * time.Second)
			}()
		}
		return nil
	}
}

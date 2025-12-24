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
		if state.Error == "" && state.CheckStatus == constants.StatusAlreadyInstalled {
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

// installToolWithRetry attempts to install a tool with automatic retry logic
// on failure. It will retry up to maxRetries times with exponential backoff.
// Returns: (status, errorMsg, output) where status is StatusSuccess or StatusFailed
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

		if status == constants.StatusSuccess {
			return constants.StatusSuccess, "", output
		}
	}

	return constants.StatusFailed, lastErr, lastOutput
}

// installToolWithOutput executes installation command with cancellation support
// Uses state's cancel context to allow aborting running installations
// Returns: (status, errorMsg, output) where status is StatusSuccess or StatusFailed
// errorMsg contains the actual error from command output when possible
func installToolWithOutput(state *models.State, method, tool string) (string, string, string) {
	cmd := commands.GetInstallCommand(method, tool)
	if cmd == "" {
		return constants.StatusFailed, constants.ErrorNoInstallCommand, ""
	}

	ctx := state.GetCancelContext()
	result := executor.ExecuteWithTimeout(ctx, cmd, 15*time.Minute)

	if result.TimedOut {
		return constants.StatusFailed, constants.ErrorInstallationTimedOut, result.Output
	}
	if result.Cancelled {
		return constants.StatusFailed, constants.ErrorInstallationCancelled, result.Output
	}
	if result.ExitCode != 0 {
		// Use actual command output as error message if available, otherwise use generic message
		errMsg := result.GetErrorMessage()
		if result.Output != "" {
			// Trim the output to first 200 chars to avoid cluttering the UI
			trimmedOutput := result.Output
			if len(trimmedOutput) > 200 {
				trimmedOutput = trimmedOutput[:200] + "..."
			}
			errMsg = trimmedOutput
		}
		return constants.StatusFailed, errMsg, result.Output
	}

	return constants.StatusSuccess, "", result.Output
}

// checkInstallation verifies if a package manager is installed and available
// Returns: (status, errorMsg) where status is StatusAlreadyInstalled or StatusNotInstalled
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
		return constants.StatusNotInstalled, errMsg
	}

	return constants.StatusAlreadyInstalled, ""
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

// NextPanel switches to the next panel (Tab key)
// Cycles through: PackageManager(0) -> Action(2) -> Tools(3) -> PackageManager(0)
// Status panel (1) is skipped as it's read-only
func NextPanel(state *models.State) func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		currentPage := state.GetCurrentPage()
		if currentPage == models.PageMultiPanel {
			activePanel := state.GetActivePanel()
			var newPanel models.Panel
			switch activePanel {
			case models.PanelPackageManager:
				newPanel = models.PanelAction
			case models.PanelAction:
				newPanel = models.PanelTools
			case models.PanelTools:
				newPanel = models.PanelPackageManager
			default:
				newPanel = models.PanelPackageManager
			}
			state.SetActivePanel(newPanel)
		}
		return nil
	}
}

// PrevPanel switches to the previous panel (Shift+Tab)
// Cycles through: PackageManager(0) <- Action(2) <- Tools(3) <- PackageManager(0)
// Status panel (1) is skipped as it's read-only
func PrevPanel(state *models.State) func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		currentPage := state.GetCurrentPage()
		if currentPage == models.PageMultiPanel {
			activePanel := state.GetActivePanel()
			var newPanel models.Panel
			switch activePanel {
			case models.PanelPackageManager:
				newPanel = models.PanelTools
			case models.PanelAction:
				newPanel = models.PanelPackageManager
			case models.PanelTools:
				newPanel = models.PanelAction
			default:
				newPanel = models.PanelPackageManager
			}
			state.SetActivePanel(newPanel)
		}
		return nil
	}
}

// SwitchToPanel switches to a specific panel (0, 1, 2 keys)
func SwitchToPanel(state *models.State, panel models.Panel) func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		currentPage := state.GetCurrentPage()
		if currentPage == models.PageMultiPanel {
			state.SetActivePanel(panel)
		}
		return nil
	}
}

func MultiPanelCursorUp(state *models.State) func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		currentPage := state.GetCurrentPage()
		if currentPage == models.PageMultiPanel {
			activePanel := state.GetActivePanel()
			switch activePanel {
			case models.PanelPackageManager:
				if state.SelectedIndex > 0 {
					state.SelectedIndex--
				}
			case models.PanelAction:
				if state.ActionIndex > 0 {
					state.ActionIndex--
					state.SelectedAction = models.ActionType(state.ActionIndex)
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
		currentPage := state.GetCurrentPage()
		if currentPage == models.PageMultiPanel {
			activePanel := state.GetActivePanel()
			switch activePanel {
			case models.PanelPackageManager:
				if state.SelectedIndex < len(state.InstallMethods)-1 {
					state.SelectedIndex++
				}
			case models.PanelAction:
				if state.ActionIndex < 2 { // 3 actions: Install, Update, Delete
					state.ActionIndex++
					state.SelectedAction = models.ActionType(state.ActionIndex)
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
		if state.GetCurrentPage() == models.PageMultiPanel && state.GetActivePanel() == models.PanelTools {
			tool := state.Tools[state.ToolsIndex]
			state.SelectedTools[tool] = !state.SelectedTools[tool]
		}
		return nil
	}
}

func MultiPanelSelectMethod(state *models.State) func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		if state.GetCurrentPage() == models.PageMultiPanel && state.GetActivePanel() == models.PanelPackageManager {
			state.SelectedMethod = state.InstallMethods[state.SelectedIndex]
			state.CheckStatus, state.Error = checkInstallation(state.SelectedMethod)

			// Only proceed if check passed (no error)
			if state.Error == "" {
				state.Tools = tools.Tools
				state.SelectedTools = make(map[string]bool)
				state.ToolsIndex = 0
				state.SetActivePanel(models.PanelAction)
			}
		}
		return nil
	}
}

// MultiPanelSelectAction handles Enter key in the Action panel
func MultiPanelSelectAction(state *models.State) func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		if state.GetCurrentPage() == models.PageMultiPanel && state.GetActivePanel() == models.PanelAction {
			state.SelectedAction = models.ActionType(state.ActionIndex)
			state.SetActivePanel(models.PanelTools)
		}
		return nil
	}
}

// MultiPanelExecuteAction handles Enter key in the Tools panel
// Executes the selected action (Install, Update, Delete) on selected tools
func MultiPanelExecuteAction(state *models.State) func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		if state.GetCurrentPage() != models.PageMultiPanel || state.GetActivePanel() != models.PanelTools {
			return nil
		}

		// Check if at least one tool is selected
		selectedTools := state.GetSelectedTools()
		selectedCount := 0
		for _, selected := range selectedTools {
			if selected {
				selectedCount++
			}
		}

		if selectedCount == 0 {
			state.Error = "No tools selected"
			return nil
		}

		// Execute based on selected action
		switch state.SelectedAction {
		case models.ActionInstall:
			return executeInstallAction(state)
		case models.ActionUpdate:
			return executeUpdateAction(state)
		case models.ActionDelete:
			return executeDeleteAction(state)
		}

		return nil
	}
}

// executeInstallAction runs installation for selected tools
func executeInstallAction(state *models.State) error {
	state.ClearInstallResults()
	state.ClearToolStartTimes()
	state.ClearInstallOutput()
	state.SetInstallingIndex(0)
	state.SetInstallationDone(false)
	state.SetInstallStartTime(time.Now().Unix())

	go runToolAction(state, "install")
	return nil
}

// executeUpdateAction runs update for selected tools
func executeUpdateAction(state *models.State) error {
	state.ClearInstallResults()
	state.ClearToolStartTimes()
	state.ClearInstallOutput()
	state.SetInstallingIndex(0)
	state.SetInstallationDone(false)
	state.SetInstallStartTime(time.Now().Unix())

	go runToolAction(state, "update")
	return nil
}

// executeDeleteAction runs uninstall/delete for selected tools
func executeDeleteAction(state *models.State) error {
	state.ClearInstallResults()
	state.ClearToolStartTimes()
	state.ClearInstallOutput()
	state.SetInstallingIndex(0)
	state.SetInstallationDone(false)
	state.SetInstallStartTime(time.Now().Unix())

	go runToolAction(state, "delete")
	return nil
}

// runToolAction executes the specified action (install/update/delete) on selected tools
func runToolAction(state *models.State, action string) {
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
			if state.GetAbortInstallation() {
				break
			}

			wg.Add(1)
			go func(toolName string) {
				defer wg.Done()

				if state.GetAbortInstallation() {
					return
				}

				startTime := time.Now().Unix()
				state.SetToolStartTime(toolName, startTime)

				var status, errMsg, output string
				switch action {
				case "install":
					status, errMsg, output = installToolWithRetry(state, state.SelectedMethod, toolName)
				case "update":
					status, errMsg, output = updateToolWithOutput(state, state.SelectedMethod, toolName)
				case "delete":
					status, errMsg, output = deleteToolWithOutput(state, state.SelectedMethod, toolName)
				}

				mu.Lock()
				state.AppendInstallOutput("Tool: " + toolName + " (" + action + ")\n" + output + "\n")
				mu.Unlock()

				duration := time.Now().Unix() - state.GetToolStartTime(toolName)
				result := models.InstallResult{
					Tool:     toolName,
					Success:  status == constants.StatusSuccess,
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
}

// updateToolWithOutput executes update command for a tool
func updateToolWithOutput(state *models.State, method, tool string) (string, string, string) {
	cmd := commands.GetUpdateCommand(method, tool)
	if cmd == "" {
		return constants.StatusFailed, "No update command found for " + tool, ""
	}

	ctx := state.GetCancelContext()
	result := executor.ExecuteWithTimeout(ctx, cmd, 15*time.Minute)

	if result.TimedOut {
		return constants.StatusFailed, constants.ErrorInstallationTimedOut, result.Output
	}
	if result.Cancelled {
		return constants.StatusFailed, constants.ErrorInstallationCancelled, result.Output
	}
	if result.ExitCode != 0 {
		errMsg := result.GetErrorMessage()
		if result.Output != "" {
			trimmedOutput := result.Output
			if len(trimmedOutput) > 200 {
				trimmedOutput = trimmedOutput[:200] + "..."
			}
			errMsg = trimmedOutput
		}
		return constants.StatusFailed, errMsg, result.Output
	}

	return constants.StatusSuccess, "", result.Output
}

// deleteToolWithOutput executes uninstall/delete command for a tool
func deleteToolWithOutput(state *models.State, method, tool string) (string, string, string) {
	cmd := commands.GetDeleteCommand(method, tool)
	if cmd == "" {
		return constants.StatusFailed, "No delete command found for " + tool, ""
	}

	ctx := state.GetCancelContext()
	result := executor.ExecuteWithTimeout(ctx, cmd, 15*time.Minute)

	if result.TimedOut {
		return constants.StatusFailed, constants.ErrorInstallationTimedOut, result.Output
	}
	if result.Cancelled {
		return constants.StatusFailed, constants.ErrorInstallationCancelled, result.Output
	}
	if result.ExitCode != 0 {
		errMsg := result.GetErrorMessage()
		if result.Output != "" {
			trimmedOutput := result.Output
			if len(trimmedOutput) > 200 {
				trimmedOutput = trimmedOutput[:200] + "..."
			}
			errMsg = trimmedOutput
		}
		return constants.StatusFailed, errMsg, result.Output
	}

	return constants.StatusSuccess, "", result.Output
}

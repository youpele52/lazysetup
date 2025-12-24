package handlers

import (
	"sync"
	"time"

	"github.com/youpele52/lazysetup/pkg/commands"
	"github.com/youpele52/lazysetup/pkg/constants"
	"github.com/youpele52/lazysetup/pkg/executor"
	"github.com/youpele52/lazysetup/pkg/models"
)

// runToolAction executes the specified action on all selected tools concurrently
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
				params := ToolActionParams{
					State:  state,
					Method: state.SelectedMethod,
					Tool:   toolName,
				}
				switch action {
				case "install":
					status, errMsg, output = installToolWithRetry(state, state.SelectedMethod, toolName)
				case "update":
					status, errMsg, output = updateToolWithOutput(params)
				case "uninstall":
					status, errMsg, output = uninstallToolWithOutput(params)
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
func updateToolWithOutput(params ToolActionParams) (string, string, string) {
	cmd := commands.GetUpdateCommand(params.Method, params.Tool)
	if cmd == "" {
		return constants.StatusFailed, "No update command found for " + params.Tool, ""
	}

	ctx := params.State.GetCancelContext()
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

// uninstallToolWithOutput executes uninstall command for a tool
func uninstallToolWithOutput(params ToolActionParams) (string, string, string) {
	cmd := commands.GetUninstallCommand(params.Method, params.Tool)
	if cmd == "" {
		return constants.StatusFailed, "No uninstall command found for " + params.Tool, ""
	}

	ctx := params.State.GetCancelContext()
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

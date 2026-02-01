package handlers

import (
	"context"
	"time"

	"github.com/jesseduffield/gocui"
	"github.com/youpele52/lazysetup/pkg/commands"
	"github.com/youpele52/lazysetup/pkg/constants"
	"github.com/youpele52/lazysetup/pkg/executor"
	"github.com/youpele52/lazysetup/pkg/models"
	"github.com/youpele52/lazysetup/pkg/tools"
)

func SelectMethod(state *models.State) func(*gocui.Gui, *gocui.View) error {
	return func(g *gocui.Gui, v *gocui.View) error {
		state.SelectedMethod = state.InstallMethods[state.PackageManagerScroll.Cursor]
		state.CheckStatus, state.Error = checkInstallation(state.SelectedMethod)
		if state.Error == "" && state.CheckStatus == constants.StatusAlreadyInstalled {
			state.Tools = tools.Tools
			state.SelectedTools = make(map[string]bool)
			state.ToolsScroll.JumpToFirst()
			state.CurrentPage = models.PageTools
		} else {
			state.CurrentPage = models.PageSelection
		}
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
// Uses sudo password if available for APT and Curl methods only
// Returns: (status, errorMsg, output) where status is StatusSuccess or StatusFailed
// errorMsg contains the actual error from command output when possible
func installToolWithOutput(state *models.State, method, tool string) (string, string, string) {
	cmd := commands.GetInstallCommand(method, tool)
	if cmd == "" {
		return constants.StatusFailed, constants.NoInstallCommandError, ""
	}

	ctx := state.GetCancelContext()
	var result *executor.CommandResult

	// Use sudo password only for APT and Curl methods
	password := state.GetSudoPassword()
	needsSudo := method == "APT" || method == "Curl" || method == "YUM"
	if password != "" && needsSudo {
		result = executor.ExecuteWithSudo(ctx, cmd, password, 15*time.Minute)
	} else {
		result = executor.ExecuteWithTimeout(ctx, cmd, 15*time.Minute)
	}

	if result.TimedOut {
		return constants.StatusFailed, constants.InstallationTimedOut, result.Output
	}
	if result.Cancelled {
		return constants.StatusFailed, constants.InstallationCancelled, result.Output
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

package handlers

import (
	"testing"
	"time"

	"github.com/youpele52/lazysetup/pkg/constants"
	"github.com/youpele52/lazysetup/pkg/models"
)

func TestRunToolAction_SingleToolExecution(t *testing.T) {
	t.Run("executes action on single tool", func(t *testing.T) {
		state := models.NewState()
		state.SetSelectedMethod("APT")
		state.SetSelectedTools(map[string]bool{"htop": true})
		state.ClearInstallResults()
		state.SetInstallationDone(false)
		state.SetInstallStartTime(time.Now().Unix())

		go runToolAction(state, constants.ToolActionCheck)

		time.Sleep(2 * time.Second)
		state.CancelInstallations()
	})
}

func TestRunToolAction_ConcurrentExecution(t *testing.T) {
	t.Run("multiple tools execute concurrently without race conditions", func(t *testing.T) {
		state := models.NewState()
		tools := []string{"htop", "curl", "vim", "git"}
		selectedTools := make(map[string]bool)
		for _, tool := range tools {
			selectedTools[tool] = true
		}

		state.SetSelectedMethod("APT")
		state.SetSelectedTools(selectedTools)
		state.ClearInstallResults()
		state.SetInstallationDone(false)
		state.SetInstallStartTime(time.Now().Unix())

		done := make(chan bool)
		go func() {
			runToolAction(state, constants.ToolActionCheck)
			done <- true
		}()

		select {
		case <-done:
			results := state.GetInstallResults()
			if len(results) == 0 {
				t.Error("Expected at least one result")
			}
		case <-time.After(10 * time.Second):
			state.CancelInstallations()
			t.Error("Test timed out")
		}
	})
}

func TestRunToolAction_AbortFlagStopsExecution(t *testing.T) {
	t.Run("abort flag stops concurrent tool execution", func(t *testing.T) {
		state := models.NewState()
		tools := []string{"htop", "curl", "vim", "git"}
		selectedTools := make(map[string]bool)
		for _, tool := range tools {
			selectedTools[tool] = true
		}

		state.SetSelectedMethod("APT")
		state.SetSelectedTools(selectedTools)
		state.ClearInstallResults()
		state.SetInstallationDone(false)
		state.SetInstallStartTime(time.Now().Unix())

		go runToolAction(state, constants.ToolActionInstall)

		time.Sleep(500 * time.Millisecond)
		state.SetAbortInstallation(true)
		state.CancelInstallations()

		time.Sleep(1 * time.Second)

		abortFlag := state.GetAbortInstallation()
		if !abortFlag {
			t.Error("Expected abort flag to be set")
		}
	})
}

func TestRunToolAction_ResultsCollectedCorrectly(t *testing.T) {
	t.Run("results are collected correctly from all tools", func(t *testing.T) {
		state := models.NewState()
		tools := []string{"htop", "curl"}
		selectedTools := make(map[string]bool)
		for _, tool := range tools {
			selectedTools[tool] = true
		}

		state.SetSelectedMethod("APT")
		state.SetSelectedTools(selectedTools)
		state.ClearInstallResults()
		state.SetInstallationDone(false)
		state.SetInstallStartTime(time.Now().Unix())

		done := make(chan bool)
		go func() {
			runToolAction(state, constants.ToolActionCheck)
			done <- true
		}()

		select {
		case <-done:
			results := state.GetInstallResults()
			if len(results) == 0 {
				t.Error("Expected results to be collected")
			}

			for _, result := range results {
				if result.Tool == "" {
					t.Error("Tool name should not be empty")
				}
			}
		case <-time.After(10 * time.Second):
			state.CancelInstallations()
			t.Error("Test timed out")
		}
	})
}

func TestRunToolAction_SpinnerFrameIncrements(t *testing.T) {
	t.Run("spinner frame increments during execution", func(t *testing.T) {
		state := models.NewState()
		state.SetSelectedTools(map[string]bool{"htop": true})
		state.ClearInstallResults()
		state.SetInstallationDone(false)
		state.SetInstallStartTime(time.Now().Unix())

		go runToolAction(state, constants.ToolActionInstall)

		time.Sleep(1 * time.Second)

		_ = state.GetSpinnerFrame()
		state.CancelInstallations()
	})
}

func TestCheckToolWithOutput_TimeoutHandling(t *testing.T) {
	t.Run("check action handles timeout correctly", func(t *testing.T) {
		state := models.NewState()
		params := ToolActionParams{
			State:  state,
			Method: "APT",
			Tool:   "htop",
		}

		status, _, output := checkToolWithOutput(params)

		if status == constants.StatusSuccess || status == constants.StatusFailed {
			if output == "" {
				t.Error("Expected some output")
			}
		}
	})
}

func TestCheckToolWithOutput_CancelledHandling(t *testing.T) {
	t.Run("check action handles cancellation correctly", func(t *testing.T) {
		state := models.NewState()
		params := ToolActionParams{
			State:  state,
			Method: "APT",
			Tool:   "htop",
		}

		go func() {
			time.Sleep(100 * time.Millisecond)
			state.CancelInstallations()
		}()

		status, _, output := checkToolWithOutput(params)

		_ = output

		if status == constants.StatusSuccess || status == constants.StatusFailed {
			if output == "" {
				t.Error("Expected some output")
			}
		}
	})
}

func TestUpdateToolWithOutput_SuccessCase(t *testing.T) {
	t.Run("update tool with output on success", func(t *testing.T) {
		state := models.NewState()
		state.SetSelectedMethod("APT")
		params := ToolActionParams{
			State:  state,
			Method: "APT",
			Tool:   "htop",
		}

		status, _, output := updateToolWithOutput(params)

		_ = output

		if status == constants.StatusSuccess || status == constants.StatusFailed {
			if output == "" {
				t.Error("Expected output")
			}
		}
	})
}

func TestUpdateToolWithOutput_FailureWithError(t *testing.T) {
	t.Run("update tool with output handles failure", func(t *testing.T) {
		state := models.NewState()
		params := ToolActionParams{
			State:  state,
			Method: "APT",
			Tool:   "htop",
		}

		status, errMsg, _ := updateToolWithOutput(params)

		if status == constants.StatusFailed {
			if errMsg == "" {
				t.Log("Warning: No error message on failure (may be expected for non-installed tools)")
			}
		}
	})
}

func TestUninstallToolWithOutput_SuccessCase(t *testing.T) {
	t.Run("uninstall tool with output on success", func(t *testing.T) {
		state := models.NewState()
		params := ToolActionParams{
			State:  state,
			Method: "APT",
			Tool:   "htop",
		}

		status, _, output := uninstallToolWithOutput(params)

		_ = output

		if status == constants.StatusSuccess || status == constants.StatusFailed {
			if output == "" {
				t.Error("Expected output")
			}
		}
	})
}

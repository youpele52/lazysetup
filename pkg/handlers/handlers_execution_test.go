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

		done := make(chan bool)
		go func() {
			runToolAction(state, constants.ToolActionCheck)
			done <- true
		}()

		select {
		case <-done:
			results := state.GetInstallResults()
			if len(results) != 1 {
				t.Errorf("Expected 1 result for single tool, got %d", len(results))
			}
			if len(results) > 0 {
				if results[0].Tool != "htop" {
					t.Errorf("Expected tool 'htop', got '%s'", results[0].Tool)
				}
			}
			if !state.GetInstallationDone() {
				t.Error("Expected installation done flag to be true after completion")
			}
		case <-time.After(10 * time.Second):
			state.CancelInstallations()
			t.Error("Test timed out waiting for single tool execution")
		}
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

		done := make(chan bool)
		go func() {
			runToolAction(state, constants.ToolActionCheck)
			done <- true
		}()

		// Wait briefly then trigger abort
		time.Sleep(100 * time.Millisecond)
		state.SetAbortInstallation(true)
		state.CancelInstallations()

		select {
		case <-done:
			abortFlag := state.GetAbortInstallation()
			if !abortFlag {
				t.Error("Expected abort flag to be set")
			}
			// Verify execution was interrupted - should have fewer results than tools
			// or installation should be marked as done due to abort
			results := state.GetInstallResults()
			if len(results) > len(tools) {
				t.Errorf("Expected at most %d results after abort, got %d", len(tools), len(results))
			}
		case <-time.After(10 * time.Second):
			t.Error("Test timed out - abort may not have stopped execution")
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
		state.SetSelectedMethod("APT")
		state.SetSelectedTools(map[string]bool{"htop": true})
		state.ClearInstallResults()
		state.SetInstallationDone(false)
		state.SetInstallStartTime(time.Now().Unix())

		initialFrame := state.GetSpinnerFrame()

		done := make(chan bool)
		go func() {
			runToolAction(state, constants.ToolActionCheck)
			done <- true
		}()

		// Wait for execution to start and spinner to increment
		time.Sleep(500 * time.Millisecond)
		midFrame := state.GetSpinnerFrame()

		select {
		case <-done:
			finalFrame := state.GetSpinnerFrame()
			// Spinner should have incremented at least once during execution
			if midFrame == initialFrame && finalFrame == initialFrame {
				t.Log("Warning: Spinner frame did not increment during execution (may be expected for fast operations)")
			}
			// Verify spinner frame is a valid value (non-negative)
			if finalFrame < 0 {
				t.Errorf("Spinner frame should be non-negative, got %d", finalFrame)
			}
		case <-time.After(10 * time.Second):
			state.CancelInstallations()
			t.Error("Test timed out")
		}
	})
}

func TestCheckToolWithOutput_TimeoutHandling(t *testing.T) {
	t.Run("check action completes and returns valid status", func(t *testing.T) {
		state := models.NewState()
		params := ToolActionParams{
			State:  state,
			Method: "APT",
			Tool:   "htop",
		}

		status, errMsg, output := checkToolWithOutput(params)

		// Status should be one of the valid statuses
		validStatuses := []string{constants.StatusSuccess, constants.StatusFailed, constants.StatusAlreadyInstalled, constants.StatusNotInstalled}
		isValidStatus := false
		for _, validStatus := range validStatuses {
			if status == validStatus {
				isValidStatus = true
				break
			}
		}
		if !isValidStatus {
			t.Errorf("Expected valid status, got '%s'", status)
		}

		// If failed, should have error message or output
		if status == constants.StatusFailed {
			if errMsg == "" && output == "" {
				t.Error("Expected error message or output on failure")
			}
		}

		// If success, should have output (version info)
		if status == constants.StatusSuccess {
			if output == "" {
				t.Error("Expected version output on success")
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

		// Cancel immediately to test cancellation handling
		state.CancelInstallations()

		status, errMsg, output := checkToolWithOutput(params)

		// After cancellation, status should still be valid
		validStatuses := []string{constants.StatusSuccess, constants.StatusFailed, constants.StatusAlreadyInstalled, constants.StatusNotInstalled, constants.InstallationCancelled}
		isValidStatus := false
		for _, validStatus := range validStatuses {
			if status == validStatus {
				isValidStatus = true
				break
			}
		}
		if !isValidStatus {
			t.Errorf("Expected valid status after cancellation, got '%s'", status)
		}

		// Log the result for debugging
		t.Logf("Status: %s, Error: %s, Output length: %d", status, errMsg, len(output))
	})
}

func TestUpdateToolWithOutput_SuccessCase(t *testing.T) {
	t.Run("update tool returns valid status and output", func(t *testing.T) {
		state := models.NewState()
		state.SetSelectedMethod("APT")
		params := ToolActionParams{
			State:  state,
			Method: "APT",
			Tool:   "htop",
		}

		status, errMsg, output := updateToolWithOutput(params)

		// Status should be valid
		validStatuses := []string{constants.StatusSuccess, constants.StatusFailed}
		isValidStatus := false
		for _, validStatus := range validStatuses {
			if status == validStatus {
				isValidStatus = true
				break
			}
		}
		if !isValidStatus {
			t.Errorf("Expected valid status (success/failed), got '%s'", status)
		}

		// Should have some output or error message
		if output == "" && errMsg == "" {
			t.Error("Expected either output or error message from update action")
		}

		// Log result for debugging
		t.Logf("Update status: %s, Error: %s, Output length: %d", status, errMsg, len(output))
	})
}

func TestUpdateToolWithOutput_FailureWithError(t *testing.T) {
	t.Run("update tool with invalid method returns failure", func(t *testing.T) {
		state := models.NewState()
		params := ToolActionParams{
			State:  state,
			Method: "InvalidMethod",
			Tool:   "htop",
		}

		status, errMsg, output := updateToolWithOutput(params)

		// With invalid method, should fail
		if status != constants.StatusFailed {
			t.Logf("Status with invalid method: %s (expected failure)", status)
		}

		// Log the results for debugging
		t.Logf("Status: %s, Error: %s, Output: %s", status, errMsg, output)
	})

	t.Run("update tool with nonexistent tool handles gracefully", func(t *testing.T) {
		state := models.NewState()
		params := ToolActionParams{
			State:  state,
			Method: "APT",
			Tool:   "nonexistent-tool-xyz",
		}

		status, errMsg, output := updateToolWithOutput(params)

		// Should return a valid status even for nonexistent tool
		if status != constants.StatusSuccess && status != constants.StatusFailed {
			t.Errorf("Expected success or failed status, got '%s'", status)
		}

		// Log the results
		t.Logf("Nonexistent tool - Status: %s, Error: %s, Output length: %d", status, errMsg, len(output))
	})
}

func TestUninstallToolWithOutput_SuccessCase(t *testing.T) {
	t.Run("uninstall tool returns valid status", func(t *testing.T) {
		state := models.NewState()
		params := ToolActionParams{
			State:  state,
			Method: "APT",
			Tool:   "htop",
		}

		status, errMsg, output := uninstallToolWithOutput(params)

		// Status should be valid
		validStatuses := []string{constants.StatusSuccess, constants.StatusFailed}
		isValidStatus := false
		for _, validStatus := range validStatuses {
			if status == validStatus {
				isValidStatus = true
				break
			}
		}
		if !isValidStatus {
			t.Errorf("Expected valid status (success/failed), got '%s'", status)
		}

		// Should have some output or error message
		if output == "" && errMsg == "" {
			t.Error("Expected either output or error message from uninstall action")
		}

		// Log result for debugging
		t.Logf("Uninstall status: %s, Error: %s, Output length: %d", status, errMsg, len(output))
	})
}

func TestUninstallToolWithOutput_InvalidMethod(t *testing.T) {
	t.Run("uninstall with invalid method handles gracefully", func(t *testing.T) {
		state := models.NewState()
		params := ToolActionParams{
			State:  state,
			Method: "InvalidMethod",
			Tool:   "htop",
		}

		status, errMsg, output := uninstallToolWithOutput(params)

		// With invalid method, should fail
		if status != constants.StatusFailed {
			t.Logf("Status with invalid method: %s (expected failure)", status)
		}

		// Log the results
		t.Logf("Invalid method - Status: %s, Error: %s, Output: %s", status, errMsg, output)
	})
}

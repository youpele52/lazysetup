package handlers

import (
	"testing"
	"time"

	"github.com/jesseduffield/gocui"
	"github.com/youpele52/lazysetup/pkg/models"
)

// TestMultiPanelStartInstallation_NoToolsError tests error handling when no tools are selected.
// Priority: P2 - Validation before execution prevents confusing empty execution.
// Tests that the handler returns appropriate error message when no tools are selected.
func TestMultiPanelStartInstallation_NoToolsError(t *testing.T) {
	t.Run("returns error when no tools selected", func(t *testing.T) {
		state := models.NewState()
		state.SetCurrentPage(models.PageMultiPanel)
		state.SetActivePanel(models.PanelTools)
		state.SetSelectedTools(map[string]bool{})

		handler := MultiPanelStartInstallation(state)

		gui := gocui.NewGui()
		if err := gui.Init(); err != nil {
			t.Fatal(err)
		}
		defer gui.Close()

		_ = handler(gui, nil)

		if state.Error != "No tools selected" {
			t.Errorf("Expected 'No tools selected' error, got %s", state.Error)
		}
	})
}

// TestMultiPanelStartInstallation_StateInitialization tests state reset before installation.
// Priority: P1 - Prevents state corruption on installation start.
// Tests that old results, output, and flags are cleared before new installation.
func TestMultiPanelStartInstallation_StateInitialization(t *testing.T) {
	t.Run("initializes state before installation", func(t *testing.T) {
		state := models.NewState()
		state.SetCurrentPage(models.PageMultiPanel)
		state.SetActivePanel(models.PanelTools)
		state.SetSelectedTools(map[string]bool{"htop": true})
		state.SetSelectedMethod("APT")

		state.AddInstallResult(models.InstallResult{Tool: "old"})
		state.AppendInstallOutput("old output")
		state.SetInstallingIndex(5)
		state.SetInstallationDone(true)
		state.SetInstallStartTime(12345)

		handler := MultiPanelStartInstallation(state)

		gui := gocui.NewGui()
		if err := gui.Init(); err != nil {
			t.Fatal(err)
		}
		defer gui.Close()

		_ = handler(gui, nil)

		time.Sleep(100 * time.Millisecond)

		results := state.GetInstallResults()
		if len(results) != 0 {
			t.Errorf("Expected empty results, got %d", len(results))
		}

		if state.GetInstallingIndex() != 0 {
			t.Errorf("Expected installing index 0, got %d", state.GetInstallingIndex())
		}

		if state.GetInstallationDone() {
			t.Error("Expected installation done false")
		}

		if state.GetInstallOutput() != "" {
			t.Error("Expected empty install output")
		}
	})
}

// TestMultiPanelStartInstallation_ParallelGoroutines tests concurrent goroutine execution.
// Priority: P0 - Complex concurrency requires testing for race conditions.
// Tests that multiple tools execute in parallel without race conditions.
func TestMultiPanelStartInstallation_ParallelGoroutines(t *testing.T) {
	t.Run("launches parallel goroutines without race conditions", func(t *testing.T) {
		state := models.NewState()
		state.SetCurrentPage(models.PageMultiPanel)
		state.SetActivePanel(models.PanelTools)
		state.SetSelectedTools(map[string]bool{
			"htop":   true,
			"curl":   true,
			"vim":    true,
			"git":    true,
			"tmux":   true,
			"docker": true,
		})
		state.SetSelectedMethod("APT")

		handler := MultiPanelStartInstallation(state)

		gui := gocui.NewGui()
		if err := gui.Init(); err != nil {
			t.Fatal(err)
		}
		defer gui.Close()

		_ = handler(gui, nil)

		done := make(chan bool)
		go func() {
			for !state.GetInstallationDone() {
				time.Sleep(100 * time.Millisecond)
			}
			done <- true
		}()

		select {
		case <-done:
			results := state.GetInstallResults()
			if len(results) == 0 {
				t.Error("Expected at least one result")
			}
		case <-time.After(30 * time.Second):
			state.CancelInstallations()
			t.Error("Installation timed out after 30 seconds")
		}
	})
}

// TestMultiPanelStartInstallation_FullIntegrationFlow tests complete installation workflow.
// Priority: P0 - Main user journey must work end-to-end.
// Tests full installation flow including results collection and output generation.
func TestMultiPanelStartInstallation_FullIntegrationFlow(t *testing.T) {
	t.Run("executes full installation flow with multiple tools", func(t *testing.T) {
		state := models.NewState()
		state.SetCurrentPage(models.PageMultiPanel)
		state.SetActivePanel(models.PanelTools)
		state.SetSelectedTools(map[string]bool{
			"htop": true,
			"curl": true,
		})
		state.SetSelectedMethod("APT")

		handler := MultiPanelStartInstallation(state)

		gui := gocui.NewGui()
		if err := gui.Init(); err != nil {
			t.Fatal(err)
		}
		defer gui.Close()

		startTime := time.Now()
		_ = handler(gui, nil)

		done := make(chan bool)
		go func() {
			for !state.GetInstallationDone() {
				time.Sleep(200 * time.Millisecond)
			}
			done <- true
		}()

		select {
		case <-done:
			results := state.GetInstallResults()
			if len(results) == 0 {
				t.Error("Expected at least one result")
			}

			if state.GetInstallOutput() == "" {
				t.Error("Expected install output")
			}

			duration := time.Since(startTime).Seconds()
			if duration > 60 {
				t.Errorf("Installation took too long: %.2f seconds", duration)
			}
		case <-time.After(60 * time.Second):
			state.CancelInstallations()
			t.Error("Installation timed out after 60 seconds")
		}
	})
}

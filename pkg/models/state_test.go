package models

import (
	"testing"
	"time"
)

// TestNewState tests initial state creation and default values.
// Priority: P2 - Prevents nil pointer panics on startup.
// Tests that NewState returns non-nil state with correct default values.
func TestNewState(t *testing.T) {
	t.Run("creates valid initial state", func(t *testing.T) {
		state := NewState()

		if state == nil {
			t.Fatal("NewState() returned nil")
		}

		if len(state.InstallMethods) == 0 {
			t.Error("InstallMethods should not be empty")
		}

		if state.CurrentPage != PageMultiPanel {
			t.Errorf("Expected PageMultiPanel, got %v", state.CurrentPage)
		}

		if state.ActivePanel != PanelPackageManager {
			t.Errorf("Expected PanelPackageManager, got %v", state.ActivePanel)
		}
	})
}

// TestState_ConcurrentAccess tests thread-safety of state with multiple goroutines.
// Priority: P0 - Race conditions can cause crashes or data corruption.
// Stress tests with 100 goroutines performing concurrent read/write operations.
func TestState_ConcurrentAccess(t *testing.T) {
	t.Run("multiple goroutines can safely read and write", func(t *testing.T) {
		state := NewState()
		numGoroutines := 100
		done := make(chan bool, numGoroutines)

		for i := 0; i < numGoroutines; i++ {
			go func(id int) {
				defer func() { done <- true }()

				for j := 0; j < 100; j++ {
					state.SetCurrentTool(toolName(id, j))
					_ = state.GetCurrentTool()

					state.SetInstallingIndex(id)
					_ = state.GetInstallingIndex()

					state.IncrementSpinnerFrame()
					_ = state.GetSpinnerFrame()
				}
			}(i)
		}

		for i := 0; i < numGoroutines; i++ {
			<-done
		}
	})
}

// TestState_PasswordMethods_ThreadSafety tests thread-safety of password operations.
// Priority: P1 - Password security must be maintained under concurrent access.
// Tests Set, Get, and Clear password operations with 50 concurrent goroutines.
func TestState_PasswordMethods_ThreadSafety(t *testing.T) {
	t.Run("concurrent password access is safe", func(t *testing.T) {
		state := NewState()
		numGoroutines := 50
		done := make(chan bool, numGoroutines)

		for i := 0; i < numGoroutines; i++ {
			go func(id int) {
				defer func() { done <- true }()

				for j := 0; j < 50; j++ {
					password := testPassword(id, j)
					state.SetSudoPassword(password)
					_ = state.GetSudoPassword()
					state.ClearSudoPassword()
				}
			}(i)
		}

		for i := 0; i < numGoroutines; i++ {
			<-done
		}
	})
}

// TestState_PasswordInputMethods_ThreadSafety tests thread-safety of password input buffer.
// Priority: P1 - Input buffer corruption causes incorrect passwords.
// Tests Set, Get, Append, and Backspace operations with concurrent goroutines.
func TestState_PasswordInputMethods_ThreadSafety(t *testing.T) {
	t.Run("concurrent password input operations are safe", func(t *testing.T) {
		state := NewState()
		numGoroutines := 50
		done := make(chan bool, numGoroutines)

		for i := 0; i < numGoroutines; i++ {
			go func(id int) {
				defer func() { done <- true }()

				for j := 0; j < 50; j++ {
					input := testInput(id, j)
					state.SetPasswordInput(input)
					_ = state.GetPasswordInput()
					state.AppendPasswordInput('x')
					state.BackspacePasswordInput()
				}
			}(i)
		}

		for i := 0; i < numGoroutines; i++ {
			<-done
		}

		finalInput := state.GetPasswordInput()
		if len(finalInput) > 100 {
			t.Errorf("Unexpectedly long input: %d chars", len(finalInput))
		}
	})
}

// TestState_SelectedTools_ConcurrentOperations tests thread-safety of tool selection.
// Priority: P1 - Tool selection corruption causes wrong installations.
// Tests concurrent Get and Set operations on selected tools map.
func TestState_SelectedTools_ConcurrentOperations(t *testing.T) {
	t.Run("concurrent selected tools operations are safe", func(t *testing.T) {
		state := NewState()
		numGoroutines := 20
		done := make(chan bool, numGoroutines)

		for i := 0; i < numGoroutines; i++ {
			go func(id int) {
				defer func() { done <- true }()

				for j := 0; j < 20; j++ {
					toolName := testTool(id, j)
					selectedTools := state.GetSelectedTools()
					selectedTools[toolName] = true
					state.SetSelectedTools(selectedTools)
					_ = state.GetSelectedTools()
				}
			}(i)
		}

		for i := 0; i < numGoroutines; i++ {
			<-done
		}
	})
}

// TestState_InstallOutputMethods_ThreadSafety tests thread-safety of install output.
// Priority: P1 - Output display corruption breaks user feedback.
// Tests Append, Get, and Clear output operations with concurrent goroutines.
func TestState_InstallOutputMethods_ThreadSafety(t *testing.T) {
	t.Run("concurrent install output operations are safe", func(t *testing.T) {
		state := NewState()
		numGoroutines := 30
		done := make(chan bool, numGoroutines)

		for i := 0; i < numGoroutines; i++ {
			go func(id int) {
				defer func() { done <- true }()

				for j := 0; j < 30; j++ {
					output := testOutput(id, j)
					state.AppendInstallOutput(output)
					_ = state.GetInstallOutput()
					state.ClearInstallOutput()
				}
			}(i)
		}

		for i := 0; i < numGoroutines; i++ {
			<-done
		}
	})
}

// TestState_InstallResults_ConcurrentAccess tests thread-safety of install results.
// Priority: P1 - Results tracking corruption causes incorrect status display.
// Tests AddInstallResult and GetInstallResults with exact count verification.
func TestState_InstallResults_ConcurrentAccess(t *testing.T) {
	t.Run("concurrent install results operations are safe", func(t *testing.T) {
		state := NewState()
		numGoroutines := 30
		done := make(chan bool, numGoroutines)

		for i := 0; i < numGoroutines; i++ {
			go func(id int) {
				defer func() { done <- true }()

				for j := 0; j < 20; j++ {
					result := InstallResult{
						Tool:     toolName(id, j),
						Success:  j%2 == 0,
						Error:    testError(id, j),
						Duration: int64(j),
						Retries:  0,
					}
					state.AddInstallResult(result)
					_ = state.GetInstallResults()
				}
			}(i)
		}

		for i := 0; i < numGoroutines; i++ {
			<-done
		}

		results := state.GetInstallResults()
		if len(results) != numGoroutines*20 {
			t.Errorf("Expected %d results, got %d", numGoroutines*20, len(results))
		}
	})
}

// TestState_InstallationDoneFlag_ThreadSafety tests thread-safety of installation done flag.
// Priority: P1 - Premature completion flag blocks UI updates.
// Tests concurrent Set and Get operations on installation done flag.
func TestState_InstallationDoneFlag_ThreadSafety(t *testing.T) {
	t.Run("concurrent installation done flag operations are safe", func(t *testing.T) {
		state := NewState()
		numGoroutines := 50
		done := make(chan bool, numGoroutines)

		for i := 0; i < numGoroutines; i++ {
			go func(id int) {
				defer func() { done <- true }()

				for j := 0; j < 50; j++ {
					state.SetInstallationDone(j%2 == 0)
					_ = state.GetInstallationDone()
				}
			}(i)
		}

		for i := 0; i < numGoroutines; i++ {
			<-done
		}
	})
}

// TestState_PagePanelMethods_ThreadSafety tests thread-safety of page and panel navigation.
// Priority: P1 - UI navigation breaks with incorrect state.
// Tests concurrent page and panel switching operations.
func TestState_PagePanelMethods_ThreadSafety(t *testing.T) {
	t.Run("concurrent page and panel operations are safe", func(t *testing.T) {
		state := NewState()
		numGoroutines := 40
		done := make(chan bool, numGoroutines)
		pages := []Page{PageMenu, PageSelection, PageTools, PageInstalling, PageResults, PageMultiPanel}
		panels := []Panel{PanelPackageManager, PanelStatus, PanelAction, PanelTools}

		for i := 0; i < numGoroutines; i++ {
			go func(id int) {
				defer func() { done <- true }()

				for j := 0; j < 40; j++ {
					state.SetCurrentPage(pages[j%len(pages)])
					_ = state.GetCurrentPage()
					state.SetActivePanel(panels[j%len(panels)])
					_ = state.GetActivePanel()
				}
			}(i)
		}

		for i := 0; i < numGoroutines; i++ {
			<-done
		}
	})
}

// TestState_ToolStartTime_ConcurrentAccess tests thread-safety of tool start time tracking.
// Priority: P2 - Duration display depends on correct timestamps.
// Tests concurrent Set and Get operations on tool start times.
func TestState_ToolStartTime_ConcurrentAccess(t *testing.T) {
	t.Run("concurrent tool start time operations are safe", func(t *testing.T) {
		state := NewState()
		numGoroutines := 30
		done := make(chan bool, numGoroutines)

		for i := 0; i < numGoroutines; i++ {
			go func(id int) {
				defer func() { done <- true }()

				for j := 0; j < 30; j++ {
					tool := testTool(id, j)
					timestamp := time.Now().Unix()
					state.SetToolStartTime(tool, timestamp)
					_ = state.GetToolStartTime(tool)
				}
			}(i)
		}

		for i := 0; i < numGoroutines; i++ {
			<-done
		}
	})
}

// TestState_CancelContext_ThreadSafety tests thread-safety of cancel context operations.
// Priority: P1 - Cancellation must work correctly for user abort functionality.
// Tests concurrent GetCancelContext, CancelInstallations, and ResetCancelContext.
func TestState_CancelContext_ThreadSafety(t *testing.T) {
	t.Run("concurrent cancel context operations are safe", func(t *testing.T) {
		state := NewState()
		numGoroutines := 30
		done := make(chan bool, numGoroutines)

		for i := 0; i < numGoroutines; i++ {
			go func(id int) {
				defer func() { done <- true }()

				for j := 0; j < 30; j++ {
					_ = state.GetCancelContext()
					state.CancelInstallations()
					time.Sleep(1 * time.Millisecond)
					state.ResetCancelContext()
				}
			}(i)
		}

		for i := 0; i < numGoroutines; i++ {
			<-done
		}
	})
}

// TestState_AbortFlag_ThreadSafety tests thread-safety of abort flag operations.
// Priority: P1 - Abort flag must be thread-safe for user control.
// Tests concurrent Set and Get operations on abort installation flag.
func TestState_AbortFlag_ThreadSafety(t *testing.T) {
	t.Run("concurrent abort flag operations are safe", func(t *testing.T) {
		state := NewState()
		numGoroutines := 40
		done := make(chan bool, numGoroutines)

		for i := 0; i < numGoroutines; i++ {
			go func(id int) {
				defer func() { done <- true }()

				for j := 0; j < 40; j++ {
					state.SetAbortInstallation(j%2 == 0)
					_ = state.GetAbortInstallation()
				}
			}(i)
		}

		for i := 0; i < numGoroutines; i++ {
			<-done
		}
	})
}

// TestState_Reset tests that Reset clears all state fields.
// Priority: P2 - Prevents state corruption between operations.
// Tests that Reset clears selected method, tools, and resets page to default.
func TestState_Reset(t *testing.T) {
	t.Run("reset clears all fields", func(t *testing.T) {
		state := NewState()

		state.SetSelectedMethod("APT")
		state.SetSelectedTools(map[string]bool{"htop": true})
		state.SetCurrentPage(PageResults)
		state.SetActivePanel(PanelTools)

		state.Reset()

		if state.GetSelectedMethod() != "" {
			t.Errorf("Expected empty method, got %s", state.GetSelectedMethod())
		}

		selectedTools := state.GetSelectedTools()
		if len(selectedTools) != 0 {
			t.Errorf("Expected empty tools, got %v", selectedTools)
		}

		if state.GetCurrentPage() != PageMultiPanel {
			t.Errorf("Expected PageMultiPanel, got %v", state.GetCurrentPage())
		}
	})
}

func testTool(id, j int) string {
	return testToolName(id, j)
}

func toolName(id, j int) string {
	return testToolName(id, j)
}

func testToolName(id, j int) string {
	return "test-tool"
}

func testPassword(id, j int) string {
	return "password"
}

func testInput(id, j int) string {
	return "input"
}

func testOutput(id, j int) string {
	return "output"
}

func testError(id, j int) string {
	return "error"
}

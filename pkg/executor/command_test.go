package executor

import (
	"context"
	"testing"
	"time"
)

// TestExecute_SuccessfulCommand tests the Execute function for successful command execution.
// Priority: P1 - Core execution path for all installations.
// Tests that a simple command returns exit code 0, captures stdout/stderr, and sets correct flags.
func TestExecute_SuccessfulCommand(t *testing.T) {
	t.Run("successful command returns exit code 0", func(t *testing.T) {
		ctx := context.Background()
		result := Execute(ctx, "echo hello")

		if result.ExitCode != 0 {
			t.Errorf("Expected exit code 0, got %d", result.ExitCode)
		}
		if result.Error != nil {
			t.Errorf("Expected no error, got %v", result.Error)
		}
		if !result.IsSuccess() {
			t.Error("Expected IsSuccess() to return true")
		}
		if result.TimedOut {
			t.Error("Expected TimedOut to be false")
		}
		if result.Cancelled {
			t.Error("Expected Cancelled to be false")
		}
	})

	t.Run("captures stdout output", func(t *testing.T) {
		ctx := context.Background()
		result := Execute(ctx, "echo 'test output'")

		if result.ExitCode != 0 {
			t.Errorf("Expected exit code 0, got %d", result.ExitCode)
		}
		if result.Output == "" {
			t.Error("Expected output to contain 'test output'")
		}
	})

	t.Run("captures stderr output", func(t *testing.T) {
		ctx := context.Background()
		result := Execute(ctx, "echo 'error output' >&2")

		if result.Output == "" {
			t.Error("Expected output to contain stderr")
		}
	})
}

// TestExecuteWithTimeout_TimeoutBehavior tests timeout behavior of ExecuteWithTimeout.
// Priority: P1 - Prevents hanging installations indefinitely.
// Tests that commands complete before timeout and that slow commands are properly terminated.
func TestExecuteWithTimeout_TimeoutBehavior(t *testing.T) {
	t.Run("command completes before timeout", func(t *testing.T) {
		ctx := context.Background()
		result := ExecuteWithTimeout(ctx, "echo fast", 5*time.Second)

		if result.TimedOut {
			t.Error("Expected TimedOut to be false for fast command")
		}
		if result.ExitCode != 0 {
			t.Errorf("Expected exit code 0, got %d", result.ExitCode)
		}
	})

	t.Run("command times out correctly", func(t *testing.T) {
		ctx := context.Background()
		start := time.Now()
		result := ExecuteWithTimeout(ctx, "sleep 10", 500*time.Millisecond)
		elapsed := time.Since(start)

		if !result.TimedOut {
			t.Error("Expected TimedOut to be true")
		}
		if result.ExitCode != -1 {
			t.Errorf("Expected exit code -1 for timeout, got %d", result.ExitCode)
		}
		if result.Error == nil {
			t.Error("Expected error for timeout")
		}
		// Verify timeout happened within reasonable time
		if elapsed > 2*time.Second {
			t.Errorf("Timeout took too long: %v", elapsed)
		}
	})
}

// TestExecuteWithTimeout_CancellationBehavior tests context cancellation handling.
// Priority: P1 - User abort functionality depends on this.
// Tests that cancelling the context stops command execution and sets the Cancelled flag.
func TestExecuteWithTimeout_CancellationBehavior(t *testing.T) {
	t.Run("cancelled context stops execution", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())

		done := make(chan *CommandResult)
		go func() {
			result := ExecuteWithTimeout(ctx, "sleep 10", 30*time.Second)
			done <- result
		}()

		// Cancel after brief delay
		time.Sleep(100 * time.Millisecond)
		cancel()

		select {
		case result := <-done:
			if !result.Cancelled {
				t.Error("Expected Cancelled to be true")
			}
			if result.ExitCode != -1 {
				t.Errorf("Expected exit code -1 for cancellation, got %d", result.ExitCode)
			}
		case <-time.After(5 * time.Second):
			t.Error("Test timed out waiting for cancellation")
		}
	})
}

// TestExecuteWithTimeout_ExitCodeHandling tests exit code capture from commands.
// Priority: P1 - Incorrect exit codes cause false success/failure reporting.
// Tests that non-zero exit codes are captured correctly and errors are set appropriately.
func TestExecuteWithTimeout_ExitCodeHandling(t *testing.T) {
	t.Run("non-zero exit code captured correctly", func(t *testing.T) {
		ctx := context.Background()
		result := ExecuteWithTimeout(ctx, "exit 1", 5*time.Second)

		if result.ExitCode != 1 {
			t.Errorf("Expected exit code 1, got %d", result.ExitCode)
		}
		if result.Error == nil {
			t.Error("Expected error for non-zero exit code")
		}
		if result.IsSuccess() {
			t.Error("Expected IsSuccess() to return false")
		}
	})

	t.Run("exit code 2 captured correctly", func(t *testing.T) {
		ctx := context.Background()
		result := ExecuteWithTimeout(ctx, "exit 2", 5*time.Second)

		if result.ExitCode != 2 {
			t.Errorf("Expected exit code 2, got %d", result.ExitCode)
		}
	})

	t.Run("command not found returns error", func(t *testing.T) {
		ctx := context.Background()
		result := ExecuteWithTimeout(ctx, "nonexistent_command_xyz", 5*time.Second)

		if result.Error == nil {
			t.Error("Expected error for nonexistent command")
		}
		if result.IsSuccess() {
			t.Error("Expected IsSuccess() to return false")
		}
	})
}

// TestExecuteWithTimeout_ShellOperators tests shell operator handling (&&, |, >).
// Priority: P2 - Ensures complex installation commands work correctly.
// Tests that shell operators are properly interpreted by using sh -c.
func TestExecuteWithTimeout_ShellOperators(t *testing.T) {
	t.Run("handles && operator", func(t *testing.T) {
		ctx := context.Background()
		result := ExecuteWithTimeout(ctx, "echo first && echo second", 5*time.Second)

		if result.ExitCode != 0 {
			t.Errorf("Expected exit code 0, got %d", result.ExitCode)
		}
		if result.Output == "" {
			t.Error("Expected output from both commands")
		}
	})

	t.Run("handles pipe operator", func(t *testing.T) {
		ctx := context.Background()
		result := ExecuteWithTimeout(ctx, "echo hello | cat", 5*time.Second)

		if result.ExitCode != 0 {
			t.Errorf("Expected exit code 0, got %d", result.ExitCode)
		}
	})
}

// TestExecuteWithSudo_PasswordHandling tests sudo command execution with password.
// Priority: P1 - APT/YUM installations require sudo to work.
// Tests password escaping, timeout, and cancellation behavior with sudo commands.
// Note: These tests don't actually use sudo (would require real password).
func TestExecuteWithSudo_PasswordHandling(t *testing.T) {

	t.Run("handles special characters in password", func(t *testing.T) {
		ctx := context.Background()
		// Test with special characters - command will fail but shouldn't crash
		result := ExecuteWithSudo(ctx, "echo test", "pass'word\"with$pecial", 2*time.Second)

		// Should complete without panic
		if result == nil {
			t.Error("Expected non-nil result")
		}
	})

	t.Run("timeout works with sudo command", func(t *testing.T) {
		ctx := context.Background()
		start := time.Now()
		result := ExecuteWithSudo(ctx, "sleep 10", "testpass", 1*time.Second)
		elapsed := time.Since(start)

		if !result.TimedOut {
			t.Error("Expected TimedOut to be true")
		}
		// Allow more time for sudo command overhead
		if elapsed > 5*time.Second {
			t.Errorf("Timeout took too long: %v", elapsed)
		}
	})

	t.Run("cancellation works with sudo command", func(t *testing.T) {
		ctx, cancel := context.WithCancel(context.Background())

		done := make(chan *CommandResult)
		go func() {
			result := ExecuteWithSudo(ctx, "sleep 10", "testpass", 30*time.Second)
			done <- result
		}()

		time.Sleep(100 * time.Millisecond)
		cancel()

		select {
		case result := <-done:
			if !result.Cancelled {
				t.Error("Expected Cancelled to be true")
			}
		case <-time.After(5 * time.Second):
			t.Error("Test timed out waiting for cancellation")
		}
	})
}

// TestCommandResult_GetErrorMessage tests error message generation from CommandResult.
// Priority: P2 - User-facing error messages must be clear and accurate.
// Tests that timeout, cancellation, and success states return appropriate messages.
func TestCommandResult_GetErrorMessage(t *testing.T) {
	t.Run("returns timeout message", func(t *testing.T) {
		result := &CommandResult{TimedOut: true}
		msg := result.GetErrorMessage()
		if msg != "Installation timed out" {
			t.Errorf("Expected 'Installation timed out', got '%s'", msg)
		}
	})

	t.Run("returns cancelled message", func(t *testing.T) {
		result := &CommandResult{Cancelled: true}
		msg := result.GetErrorMessage()
		if msg != "Installation was cancelled" {
			t.Errorf("Expected 'Installation was cancelled', got '%s'", msg)
		}
	})

	t.Run("returns empty for success", func(t *testing.T) {
		result := &CommandResult{}
		msg := result.GetErrorMessage()
		if msg != "" {
			t.Errorf("Expected empty message, got '%s'", msg)
		}
	})
}

// TestCommandResult_Duration tests that command execution duration is recorded.
// Priority: P2 - Duration tracking is used for performance monitoring.
// Tests that Duration field is populated with a positive value after execution.
func TestCommandResult_Duration(t *testing.T) {
	t.Run("duration is recorded", func(t *testing.T) {
		ctx := context.Background()
		result := ExecuteWithTimeout(ctx, "sleep 0.1", 5*time.Second)

		if result.Duration <= 0 {
			t.Errorf("Expected positive duration, got %d", result.Duration)
		}
	})
}

package executor

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"time"
)

// CommandResult contains the outcome of a command execution
type CommandResult struct {
	Output    string // Combined stdout and stderr
	Error     error  // Execution error (nil on success)
	ExitCode  int    // Process exit code
	Duration  int64  // Execution time in milliseconds
	TimedOut  bool   // Whether command timed out
	Cancelled bool   // Whether command was cancelled
}

// Execute runs a shell command with timeout and cancellation support
// Uses sh -c to properly handle shell operators like &&, |, >
// Default timeout is 10 minutes for installations
func Execute(ctx context.Context, command string) *CommandResult {
	return ExecuteWithTimeout(ctx, command, 10*time.Minute)
}

// ExecuteWithTimeout runs a shell command with specified timeout and cancellation support
// Uses sh -c to properly handle shell operators like &&, |, >
// Returns CommandResult with output, error, exit code, duration, and status flags
func ExecuteWithTimeout(ctx context.Context, command string, timeout time.Duration) *CommandResult {
	startTime := time.Now()
	result := &CommandResult{}

	// Create context with timeout
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	// Use sh -c to properly handle shell operators
	cmd := exec.CommandContext(ctx, "sh", "-c", command)

	// Capture both stdout and stderr
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	// Run the command
	err := cmd.Run()

	// Calculate duration
	result.Duration = time.Since(startTime).Milliseconds()

	// Check for timeout
	if ctx.Err() == context.DeadlineExceeded {
		result.TimedOut = true
		result.Error = fmt.Errorf("command timed out after %v", timeout)
		result.ExitCode = -1
		result.Output = stdout.String() + stderr.String()
		return result
	}

	// Check for cancellation
	if ctx.Err() == context.Canceled {
		result.Cancelled = true
		result.Error = fmt.Errorf("command was cancelled")
		result.ExitCode = -1
		result.Output = stdout.String() + stderr.String()
		return result
	}

	// Combine output
	result.Output = stdout.String() + stderr.String()

	// Handle execution errors
	if err != nil {
		result.Error = err
		if exitErr, ok := err.(*exec.ExitError); ok {
			result.ExitCode = exitErr.ExitCode()
		} else {
			result.ExitCode = -1
		}
	} else {
		result.ExitCode = 0
	}

	return result
}

// IsSuccess checks if command executed successfully (exit code 0, no error)
func (r *CommandResult) IsSuccess() bool {
	return r.Error == nil && r.ExitCode == 0
}

// GetErrorMessage returns a human-readable error message
func (r *CommandResult) GetErrorMessage() string {
	if r.TimedOut {
		return "Installation timed out"
	}
	if r.Cancelled {
		return "Installation was cancelled"
	}
	if r.Error != nil {
		return r.Error.Error()
	}
	return ""
}

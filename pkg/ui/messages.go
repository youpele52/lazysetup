package ui

import (
	"fmt"
	"strings"

	"github.com/youpele52/lazysetup/pkg/colors"
	"github.com/youpele52/lazysetup/pkg/constants"
	"github.com/youpele52/lazysetup/pkg/models"
)

// Action text constants
const (
	actionTextInstall   = "Installation"
	actionTextUpdate    = "Update"
	actionTextUninstall = "Uninstall"
	actionTextDefault   = "Action"

	actionVerbInstall   = "Installing"
	actionVerbUpdate    = "Updating"
	actionVerbUninstall = "Uninstalling"
	actionVerbDefault   = "Processing"

	actionNameInstall   = "install"
	actionNameUpdate    = "update"
	actionNameUninstall = "uninstall"
	actionNameDefault   = "action"

	resultSuccessful = "successful"
	resultFailed     = "failed"
)

// MessageBuilder provides a fluent interface for building multi-line UI messages
// It allows chaining method calls to construct formatted messages with lines, separators, and spacing
type MessageBuilder struct {
	lines []string
}

func NewMessageBuilder() *MessageBuilder {
	return &MessageBuilder{
		lines: []string{},
	}
}

func (mb *MessageBuilder) AddLine(text string) *MessageBuilder {
	mb.lines = append(mb.lines, text)
	return mb
}

func (mb *MessageBuilder) AddBlankLine() *MessageBuilder {
	mb.lines = append(mb.lines, "")
	return mb
}

func (mb *MessageBuilder) AddSeparator() *MessageBuilder {
	mb.lines = append(mb.lines, constants.ResultsSeparator)
	return mb
}

func (mb *MessageBuilder) Build() string {
	return strings.Join(mb.lines, "\n")
}

func BuildInstallationProgressMessage(params ProgressMessageParams) string {
	mb := NewMessageBuilder()

	// Get action-specific text
	actionText := getActionText(params.Action)
	actionVerb := getActionVerb(params.Action)

	mb.AddLine(fmt.Sprintf("%s Method: %s", actionText, params.SelectedMethod))
	mb.AddBlankLine()
	if params.CurrentTool != "" {
		mb.AddLine(fmt.Sprintf("Current Tool: %s", params.CurrentTool))
	}
	mb.AddLine(fmt.Sprintf("Status: %d/%d", params.InstallingIndex, params.TotalTools))
	mb.AddBlankLine()
	mb.AddSeparator()

	if !params.InstallationDone {
		spinner := getSpinner(params.SpinnerFrame)
		mb.AddLine(fmt.Sprintf("%s %s...", spinner, actionVerb))
	}

	// Only show install output for non-check actions during progress
	if params.InstallOutput != "" && params.Action != models.ActionCheck {
		mb.AddLine(params.InstallOutput)
	}

	return mb.Build()
}

// getActionText returns the display text for the action type header
func getActionText(action models.ActionType) string {
	switch action {
	case models.ActionInstall:
		return actionTextInstall
	case models.ActionUpdate:
		return actionTextUpdate
	case models.ActionUninstall:
		return actionTextUninstall
	default:
		return actionTextDefault
	}
}

// getActionVerb returns the verb form for the action type
func getActionVerb(action models.ActionType) string {
	switch action {
	case models.ActionInstall:
		return actionVerbInstall
	case models.ActionUpdate:
		return actionVerbUpdate
	case models.ActionUninstall:
		return actionVerbUninstall
	default:
		return actionVerbDefault
	}
}

// BuildInstallationResultsMessage creates a formatted summary of action results
// Shows success/failure for each tool with color coding, error details
// Results are displayed in reverse order (latest at top) like rolling credits
func BuildInstallationResultsMessage(results []models.InstallResult, action models.ActionType) string {
	mb := NewMessageBuilder()

	successCount := 0
	failureCount := 0

	// Get action-specific result text
	successVerb := getSuccessVerb(action)
	failVerb := getFailVerb(action)

	// For check action, display version info instead of generic success message
	isCheckAction := action == models.ActionCheck

	// Display results in reverse order (latest at top, like rolling credits)
	for i := len(results) - 1; i >= 0; i-- {
		result := results[i]
		if result.Success {
			if isCheckAction && result.Error != "" {
				// For check action, show the version output with proper formatting
				mb.AddLine(fmt.Sprintf("%s✓ %s%s", colors.ANSIGreen, result.Tool, colors.ANSIReset))
				// Split version output by newlines and indent each line
				versionLines := strings.Split(strings.TrimSpace(result.Error), "\n")
				for _, versionLine := range versionLines {
					if strings.TrimSpace(versionLine) != "" {
						mb.AddLine(fmt.Sprintf("  %s", strings.TrimSpace(versionLine)))
					}
				}
			} else {
				successLine := fmt.Sprintf("%s✓ %s - %s successful (%ds)%s", colors.ANSIGreen, result.Tool, successVerb, result.Duration, colors.ANSIReset)
				mb.AddLine(successLine)
			}
			successCount++
		} else {
			failedLine := fmt.Sprintf("%s✗ %s - %s failed (%ds)%s", colors.ANSIRed, result.Tool, failVerb, result.Duration, colors.ANSIReset)
			mb.AddLine(failedLine)
			if result.Error != "" {
				// Display error message, split by newlines for readability
				errorLines := strings.Split(result.Error, "\n")
				for _, errLine := range errorLines {
					if strings.TrimSpace(errLine) != "" {
						displayLine := fmt.Sprintf("%s  %s%s", colors.ANSIRed, strings.TrimSpace(errLine), colors.ANSIReset)
						mb.AddLine(displayLine)
						break // Only show first line to avoid cluttering UI
					}
				}
			}
			failureCount++
		}
		mb.AddBlankLine()
	}

	mb.AddSeparator()
	mb.AddLine(fmt.Sprintf("Total: %d Success, %d Failed", successCount, failureCount))
	mb.AddSeparator()

	return mb.Build()
}

// getSuccessVerb returns the action verb for success messages
func getSuccessVerb(action models.ActionType) string {
	switch action {
	case models.ActionInstall:
		return actionNameInstall
	case models.ActionUpdate:
		return actionNameUpdate
	case models.ActionUninstall:
		return actionNameUninstall
	default:
		return actionNameDefault
	}
}

// getFailVerb returns the action verb for failure messages
func getFailVerb(action models.ActionType) string {
	switch action {
	case models.ActionInstall:
		return actionNameInstall
	case models.ActionUpdate:
		return actionNameUpdate
	case models.ActionUninstall:
		return actionNameUninstall
	default:
		return actionNameDefault
	}
}

// BuildNewResultsMessage builds a message for only new results (not previously rendered)
// Used for rolling credits display to avoid duplication
func BuildNewResultsMessage(results []models.InstallResult, action models.ActionType) string {
	mb := NewMessageBuilder()

	successCount := 0
	failureCount := 0

	// Get action-specific result text
	successVerb := getSuccessVerb(action)
	failVerb := getFailVerb(action)

	// For check action, display version info instead of generic success message
	isCheckAction := action == models.ActionCheck

	// Display results in reverse order (latest at top, like rolling credits)
	for i := len(results) - 1; i >= 0; i-- {
		result := results[i]
		if result.Success {
			if isCheckAction && result.Error != "" {
				// For check action, show the version output with proper formatting
				mb.AddLine(fmt.Sprintf("%s✓ %s%s", colors.ANSIGreen, result.Tool, colors.ANSIReset))
				// Split version output by newlines and indent each line
				versionLines := strings.Split(strings.TrimSpace(result.Error), "\n")
				for _, versionLine := range versionLines {
					if strings.TrimSpace(versionLine) != "" {
						mb.AddLine(fmt.Sprintf("  %s", strings.TrimSpace(versionLine)))
					}
				}
			} else {
				successLine := fmt.Sprintf("%s✓ %s - %s successful (%ds)%s", colors.ANSIGreen, result.Tool, successVerb, result.Duration, colors.ANSIReset)
				mb.AddLine(successLine)
			}
			successCount++
		} else {
			failedLine := fmt.Sprintf("%s✗ %s - %s failed (%ds)%s", colors.ANSIRed, result.Tool, failVerb, result.Duration, colors.ANSIReset)
			mb.AddLine(failedLine)
			if result.Error != "" {
				// Display error message, split by newlines for readability
				errorLines := strings.Split(result.Error, "\n")
				for _, errLine := range errorLines {
					if strings.TrimSpace(errLine) != "" {
						displayLine := fmt.Sprintf("%s  %s%s", colors.ANSIRed, strings.TrimSpace(errLine), colors.ANSIReset)
						mb.AddLine(displayLine)
						break // Only show first line to avoid cluttering UI
					}
				}
			}
			failureCount++
		}
		mb.AddBlankLine()
	}

	mb.AddSeparator()
	mb.AddLine(fmt.Sprintf("Total: %d Success, %d Failed", successCount, failureCount))
	mb.AddSeparator()

	return mb.Build()
}

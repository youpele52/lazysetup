package ui

import (
	"fmt"
	"strings"

	"github.com/youpele52/lazysetup/pkg/colors"
	"github.com/youpele52/lazysetup/pkg/constants"
	"github.com/youpele52/lazysetup/pkg/models"
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
	mb.AddLine(fmt.Sprintf("Current Tool: %s", params.CurrentTool))
	mb.AddLine(fmt.Sprintf("Status: %d/%d", params.InstallingIndex, params.TotalTools))
	mb.AddBlankLine()
	mb.AddSeparator()

	if !params.InstallationDone {
		spinner := getSpinner(params.SpinnerFrame)
		mb.AddLine(fmt.Sprintf("%s %s...", spinner, actionVerb))
	}

	if params.InstallOutput != "" {
		mb.AddLine(params.InstallOutput)
	}

	return mb.Build()
}

// getActionText returns the display text for the action type header
func getActionText(action models.ActionType) string {
	switch action {
	case models.ActionInstall:
		return "Installation"
	case models.ActionUpdate:
		return "Update"
	case models.ActionUninstall:
		return "Uninstall"
	default:
		return "Action"
	}
}

// getActionVerb returns the verb form for the action type
func getActionVerb(action models.ActionType) string {
	switch action {
	case models.ActionInstall:
		return "Installing"
	case models.ActionUpdate:
		return "Updating"
	case models.ActionUninstall:
		return "Uninstalling"
	default:
		return "Processing"
	}
}

// BuildInstallationResultsMessage creates a formatted summary of action results
// Shows success/failure for each tool with color coding, error details, and totals
func BuildInstallationResultsMessage(results []models.InstallResult, action models.ActionType) string {
	mb := NewMessageBuilder()

	mb.AddSeparator()
	mb.AddBlankLine()

	successCount := 0
	failureCount := 0

	// Get action-specific result text
	successVerb := getSuccessVerb(action)
	failVerb := getFailVerb(action)

	for _, result := range results {
		if result.Success {
			successLine := fmt.Sprintf("%s✓ %s - %s (%ds)%s", colors.ANSIGreen, result.Tool, successVerb, result.Duration, colors.ANSIReset)
			mb.AddLine(successLine)
			successCount++
		} else {
			failedLine := fmt.Sprintf("%s✗ %s - %s (%ds)%s", colors.ANSIRed, result.Tool, failVerb, result.Duration, colors.ANSIReset)
			mb.AddLine(failedLine)
			if result.Error != "" {
				// Display error message, split by newlines for readability
				errorLines := strings.Split(result.Error, "\n")
				for _, errLine := range errorLines {
					if strings.TrimSpace(errLine) != "" {
						displayLine := fmt.Sprintf("%s  Error: %s%s", colors.ANSIRed, strings.TrimSpace(errLine), colors.ANSIReset)
						mb.AddLine(displayLine)
						break // Only show first line to avoid cluttering UI
					}
				}
			}
			failureCount++
		}
	}

	mb.AddBlankLine()
	mb.AddSeparator()
	mb.AddLine(fmt.Sprintf("Total: %d Success, %d Failed", successCount, failureCount))

	return mb.Build()
}

// getSuccessVerb returns the past tense success verb for the action type
func getSuccessVerb(action models.ActionType) string {
	switch action {
	case models.ActionInstall:
		return "Successfully installed"
	case models.ActionUpdate:
		return "Successfully updated"
	case models.ActionUninstall:
		return "Successfully uninstalled"
	default:
		return "Success"
	}
}

// getFailVerb returns the failure verb for the action type
func getFailVerb(action models.ActionType) string {
	switch action {
	case models.ActionInstall:
		return "Failed to install"
	case models.ActionUpdate:
		return "Failed to update"
	case models.ActionUninstall:
		return "Failed to uninstall"
	default:
		return "Failed"
	}
}

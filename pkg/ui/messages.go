package ui

import (
	"fmt"
	"strings"

	"github.com/youpele52/lazysetup/pkg/colors"
	"github.com/youpele52/lazysetup/pkg/constants"
	"github.com/youpele52/lazysetup/pkg/models"
)

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

func BuildInstallationProgressMessage(currentTool string, installingIndex, totalTools int, installationDone bool, spinnerFrame int, installOutput string) string {
	mb := NewMessageBuilder()

	mb.AddLine(fmt.Sprintf("Current Tool: %s", currentTool))
	mb.AddLine(fmt.Sprintf("Progress: %d/%d", installingIndex, totalTools))
	mb.AddBlankLine()
	mb.AddSeparator()

	if !installationDone {
		spinner := getSpinner(spinnerFrame)
		mb.AddLine(fmt.Sprintf("%s Installing...", spinner))
	}

	if installOutput != "" {
		mb.AddLine(installOutput)
	}

	return mb.Build()
}

func BuildInstallationResultsMessage(results []models.InstallResult) string {
	mb := NewMessageBuilder()

	mb.AddSeparator()
	mb.AddBlankLine()

	successCount := 0
	failureCount := 0

	for _, result := range results {
		if result.Success {
			successLine := fmt.Sprintf("%s✓ %s - Success (%ds)%s", colors.ANSIGreen, result.Tool, result.Duration, colors.ANSIReset)
			mb.AddLine(successLine)
			successCount++
		} else {
			failedLine := fmt.Sprintf("%s✗ %s - Failed (%ds)%s", colors.ANSIRed, result.Tool, result.Duration, colors.ANSIReset)
			mb.AddLine(failedLine)
			if result.Error != "" {
				errorLine := fmt.Sprintf("%s  Error: %s%s", colors.ANSIRed, result.Error, colors.ANSIReset)
				mb.AddLine(errorLine)
			}
			failureCount++
		}
	}

	mb.AddBlankLine()
	mb.AddSeparator()
	mb.AddLine(fmt.Sprintf("Total: %d Success, %d Failed", successCount, failureCount))

	return mb.Build()
}

package ui

import (
	"fmt"
	"strings"

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
			mb.AddLine(fmt.Sprintf("✓ %s - Success (%ds)", result.Tool, result.Duration))
			successCount++
		} else {
			mb.AddLine(fmt.Sprintf("✗ %s - Failed (%ds)", result.Tool, result.Duration))
			if result.Error != "" {
				mb.AddLine(fmt.Sprintf("  Error: %s", result.Error))
			}
			failureCount++
		}
	}

	mb.AddBlankLine()
	mb.AddSeparator()
	mb.AddLine(fmt.Sprintf("Total: %d Success, %d Failed", successCount, failureCount))

	return mb.Build()
}

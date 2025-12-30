package messages

import (
	"fmt"
	"strings"

	"github.com/youpele52/lazysetup/pkg/colors"
	"github.com/youpele52/lazysetup/pkg/constants"
	"github.com/youpele52/lazysetup/pkg/models"
)

const (
	ANSIBlue = "\033[1;34m"
)

type ProgressMessageParams struct {
	SelectedMethod   string
	CurrentTool      string
	InstallingIndex  int
	TotalTools       int
	InstallationDone bool
	SpinnerFrame     int
	InstallOutput    string
	Action           models.ActionType
}

func BuildInstallationProgressMessage(params ProgressMessageParams) string {
	mb := NewMessageBuilder()

	// Spinner animation for in-progress installation
	spinnerColor := ANSIBlue
	spinnerFrames := []string{"⠋", "⠙", "⠹", "⠸", "⠼", "⠴", "⠦", "⠧", "⠇", "⠏"}
	spinnerFrame := spinnerFrames[params.SpinnerFrame%len(spinnerFrames)]
	actionVerb := getActionVerb(params.Action)

	if params.InstallationDone {
		mb.AddLine(fmt.Sprintf("%sInstallation completed!%s", colors.ANSIGreen, colors.ANSIReset))
	} else {
		progressLine := fmt.Sprintf(
			"%s%s %s %d/%d %s%s",
			spinnerColor,
			spinnerFrame,
			actionVerb,
			params.InstallingIndex+1,
			params.TotalTools,
			params.CurrentTool,
			colors.ANSIReset,
		)
		mb.AddLine(progressLine)

		// Show install output if available
		if params.InstallOutput != "" {
			outputLines := strings.Split(strings.TrimSpace(params.InstallOutput), "\n")
			for _, line := range outputLines {
				if strings.TrimSpace(line) != "" {
					mb.AddLine(fmt.Sprintf("  %s", strings.TrimSpace(line)))
				}
			}
		}
	}

	return mb.Build()
}

func BuildNewResultsMessage(results []models.InstallResult, action models.ActionType) string {
	mb := NewMessageBuilder()

	for i := len(results) - 1; i >= 0; i-- {
		result := results[i]
		verb := getActionVerb(action)

		if result.Success {
			successLine := fmt.Sprintf("%s✓ %s%s",
				colors.ANSIGreen, result.Tool, colors.ANSIReset)
			mb.AddLine(successLine)

			// For check action, show version output on success (white/default color)
			if action == models.ActionCheck && result.Error != "" {
				versionLines := strings.Split(result.Error, "\n")
				for _, vLine := range versionLines {
					if strings.TrimSpace(vLine) != "" {
						displayLine := fmt.Sprintf("  %s", strings.TrimSpace(vLine))
						mb.AddLine(displayLine)
					}
				}
			}
		} else {
			failedLine := fmt.Sprintf("%s✗ %s - %s failed (%ds)%s",
				colors.ANSIRed, result.Tool, verb, result.Duration, colors.ANSIReset)
			mb.AddLine(failedLine)

			if result.Error != "" {
				errorLines := strings.Split(result.Error, "\n")
				for _, errLine := range errorLines {
					if strings.TrimSpace(errLine) != "" {
						displayLine := fmt.Sprintf("%s  %s%s",
							colors.ANSIRed, strings.TrimSpace(errLine), colors.ANSIReset)
						mb.AddLine(displayLine)
						break
					}
				}
			}
		}
		mb.AddBlankLine()
	}

	return mb.Build()
}

func getActionVerb(action models.ActionType) string {
	switch action {
	case models.ActionInstall:
		return constants.ToolActionInstall
	case models.ActionUpdate:
		return constants.ToolActionUpdate
	case models.ActionUninstall:
		return constants.ToolActionUninstall
	case models.ActionCheck:
		return constants.ToolActionCheck
	default:
		return constants.ToolActionCheck
	}
}

package messages

import (
	"fmt"
	"strings"

	"github.com/youpele52/lazysetup/pkg/colors"
	"github.com/youpele52/lazysetup/pkg/models"
)

func BuildInstallationResultsMessage(results []models.InstallResult, action models.ActionType) string {
	mb := NewMessageBuilder()
	successCount := 0
	failureCount := 0
	isCheckAction := action == models.ActionCheck

	for i := len(results) - 1; i >= 0; i-- {
		result := results[i]
		if result.Success {
			if isCheckAction && result.Error != "" {
				mb.AddLine(fmt.Sprintf("%s✓ %s%s", colors.ANSIGreen, result.Tool, colors.ANSIReset))
				versionLines := strings.Split(strings.TrimSpace(result.Error), "\n")
				for _, versionLine := range versionLines {
					if strings.TrimSpace(versionLine) != "" {
						mb.AddLine(fmt.Sprintf("  %s", strings.TrimSpace(versionLine)))
					}
				}
			} else {
				verb := getActionVerb(action)
				successLine := fmt.Sprintf("%s✓ %s - %s successful (%ds)%s", colors.ANSIGreen, result.Tool, verb, result.Duration, colors.ANSIReset)
				mb.AddLine(successLine)
			}
			successCount++
		} else {
			verb := getActionVerb(action)
			failedLine := fmt.Sprintf("%s✗ %s - %s failed (%ds)%s", colors.ANSIRed, result.Tool, verb, result.Duration, colors.ANSIReset)
			mb.AddLine(failedLine)
			if result.Error != "" {
				errorLines := strings.Split(result.Error, "\n")
				for _, errLine := range errorLines {
					if strings.TrimSpace(errLine) != "" {
						displayLine := fmt.Sprintf("%s  %s%s", colors.ANSIRed, strings.TrimSpace(errLine), colors.ANSIReset)
						mb.AddLine(displayLine)
						break
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

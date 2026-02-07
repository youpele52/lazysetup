package handlers

import (
	"strings"

	"github.com/youpele52/lazysetup/pkg/constants"
)

// FilterTools returns tools matching the search query (case-insensitive substring match)
// Searches both tool internal name and display name
func FilterTools(tools []string, query string) []string {
	if query == "" {
		return tools
	}

	query = strings.ToLower(query)
	var filtered []string

	for _, tool := range tools {
		// Check internal name (e.g., "kubectl")
		if strings.Contains(strings.ToLower(tool), query) {
			filtered = append(filtered, tool)
			continue
		}

		// Check display name (e.g., "claude code" for "claude-code")
		displayName := constants.GetToolDisplayName(tool)
		if strings.Contains(strings.ToLower(displayName), query) {
			filtered = append(filtered, tool)
		}
	}

	return filtered
}

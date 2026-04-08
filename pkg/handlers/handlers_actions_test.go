package handlers

import (
	"sort"
	"testing"

	"github.com/youpele52/lazysetup/pkg/commands"
	"github.com/youpele52/lazysetup/pkg/models"
)

// TestMultiPanelSelectMethod_FiltersToolsByMethod tests method-specific tool availability.
// Priority: P1 - The visible tool list must match supported command pairs.
func TestMultiPanelSelectMethod_FiltersToolsByMethod(t *testing.T) {
	containsTool := func(tools []string, target string) bool {
		for _, tool := range tools {
			if tool == target {
				return true
			}
		}
		return false
	}

	t.Run("supported tool list can include codex", func(t *testing.T) {
		tools := commands.GetSupportedToolsForMethod("Homebrew")
		sort.Strings(tools)
		if !containsTool(tools, "codex") {
			t.Fatalf("expected Homebrew support list to include codex, got %v", tools)
		}
	})

	t.Run("unsupported tool list excludes codex", func(t *testing.T) {
		tools := commands.GetSupportedToolsForMethod("APT")
		sort.Strings(tools)
		if containsTool(tools, "codex") {
			t.Fatalf("did not expect APT support list to include codex, got %v", tools)
		}
	})

	t.Run("new state defaults to first method supported tools", func(t *testing.T) {
		state := models.NewState()
		if !containsTool(state.Tools, "codex") {
			t.Fatalf("expected default tool list to be populated from first method, got %v", state.Tools)
		}
	})
}

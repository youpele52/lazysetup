package tools

import (
	"testing"
)

// TestTools_NotEmpty tests that Tools slice is not empty.
// Priority: P3 - Configuration validation.
// Tests that Tools slice is not empty.
func TestTools_NotEmpty(t *testing.T) {
	t.Run("tools slice is not empty", func(t *testing.T) {
		if len(Tools) == 0 {
			t.Error("Tools slice should not be empty")
		}
	})

	t.Run("contains expected tools", func(t *testing.T) {
		expectedTools := map[string]bool{
			"git":        true,
			"docker":     true,
			"lazygit":    true,
			"lazydocker": true,
			"htop":       true,
		}
		for _, tool := range Tools {
			if !expectedTools[tool] {
				t.Errorf("Unexpected tool '%s' in Tools slice", tool)
			}
		}
		if len(Tools) != len(expectedTools) {
			t.Errorf("Expected %d tools, got %d", len(expectedTools), len(Tools))
		}
	})
}

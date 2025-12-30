package tools

import (
	"testing"
)

// TestTools_ContainsExpectedTools tests that Tools slice contains expected tools.
// Priority: P3 - Configuration validation.
// Tests that the Tools slice is not empty and contains known tools.
func TestTools_ContainsExpectedTools(t *testing.T) {
	t.Run("tools slice is not empty", func(t *testing.T) {
		if len(Tools) == 0 {
			t.Error("Tools slice should not be empty")
		}
	})

	t.Run("contains expected tools", func(t *testing.T) {
		expected := []string{"git", "docker", "lazygit", "lazydocker", "htop"}
		for _, tool := range expected {
			found := false
			for _, t2 := range Tools {
				if t2 == tool {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("Expected tool '%s' not found in Tools slice", tool)
			}
		}
	})
}

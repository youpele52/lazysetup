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
			// Version Control & Build
			"git":        true,
			"docker":     true,
			"lazygit":    true,
			"lazydocker": true,
			"gh":         true,
			"make":       true,
			"just":       true,
			// Shell & Terminal
			"zsh":  true,
			"tmux": true,
			// Development Environment
			"nvim":    true,
			"node":    true,
			"python3": true,
			"bun":     true,
			"pnpm":    true,
			"uv":      true,
			// AI Assistants
			"claude-code": true,
			"opencode":    true,
			// System Monitoring
			"btop": true,
			// Core Utilities
			"fzf":     true,
			"ripgrep": true,
			"fd":      true,
			"bat":     true,
			"eza":     true,
			"zoxide":  true,
			"tree":    true,
			"rsync":   true,
			// Shell Enhancement
			"starship": true,
			"delta":    true,
			// Data Processing
			"jq": true,
			// Cloud-Native & DevOps
			"kubectl":   true,
			"k9s":       true,
			"terraform": true,
			"helm":      true,
			// Network & Web
			"httpie": true,
			"wget":   true,
			// Documentation
			"tldr":    true,
			"lazysql": true,
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

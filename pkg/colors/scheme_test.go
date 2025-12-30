package colors

import (
	"strings"
	"testing"
)

// TestANSIColorCodes_ValidEscapeSequences tests that ANSI color codes are valid.
// Priority: P3 - Terminal display correctness.
// Tests that ANSI codes start with escape sequence and are non-empty.
func TestANSIColorCodes_ValidEscapeSequences(t *testing.T) {
	t.Run("ANSI codes start with escape sequence", func(t *testing.T) {
		codes := []string{ANSIGreen, ANSIRed, ANSIYellow, ANSIMagenta, ANSIReset}
		for _, code := range codes {
			if !strings.HasPrefix(code, "\033[") {
				t.Errorf("ANSI code '%s' should start with escape sequence", code)
			}
		}
	})

	t.Run("ANSI codes are non-empty", func(t *testing.T) {
		if ANSIGreen == "" || ANSIRed == "" || ANSIYellow == "" || ANSIMagenta == "" || ANSIReset == "" {
			t.Error("ANSI color codes should not be empty")
		}
	})

	t.Run("reset code ends with 0m", func(t *testing.T) {
		if !strings.HasSuffix(ANSIReset, "0m") {
			t.Errorf("ANSIReset should end with '0m', got '%s'", ANSIReset)
		}
	})
}

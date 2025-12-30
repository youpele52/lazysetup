package constants

import (
	"strings"
	"testing"
)

// TestViewConstants_AreNonEmpty tests that view constants are defined and non-empty.
// Priority: P3 - Configuration validation.
// Tests that key UI view constants have valid values.
func TestViewConstants_AreNonEmpty(t *testing.T) {
	tests := []struct {
		name  string
		value string
	}{
		{"ViewMenu", ViewMenu},
		{"ViewResult", ViewResult},
		{"ViewResults", ViewResults},
		{"ViewTools", ViewTools},
		{"ViewInstalling", ViewInstalling},
	}
	for _, tt := range tests {
		t.Run(tt.name+" is non-empty", func(t *testing.T) {
			if tt.value == "" {
				t.Errorf("Constant %s should not be empty", tt.name)
			}
		})
	}
}

// TestPanelConstants_AreNonEmpty tests that panel constants are defined and non-empty.
// Priority: P3 - Configuration validation.
// Tests that UI panel constants have valid values.
func TestPanelConstants_AreNonEmpty(t *testing.T) {
	tests := []struct {
		name  string
		value string
	}{
		{"PanelTools", PanelTools},
		{"PanelPackageManager", PanelPackageManager},
		{"PanelProgress", PanelProgress},
		{"PanelAction", PanelAction},
		{"PanelStatusView", PanelStatusView},
		{"PopupConfirm", PopupConfirm},
	}
	for _, tt := range tests {
		t.Run(tt.name+" is non-empty", func(t *testing.T) {
			if tt.value == "" {
				t.Errorf("Constant %s should not be empty", tt.name)
			}
		})
	}
}

// TestTitleConstants_AreNonEmpty tests that title constants are defined and non-empty.
// Priority: P3 - Configuration validation.
// Tests that UI title constants have valid values.
func TestTitleConstants_AreNonEmpty(t *testing.T) {
	tests := []struct {
		name  string
		value string
	}{
		{"TitlePackageManager", TitlePackageManager},
		{"TitleInstalling", TitleInstalling},
		{"TitleTools", TitleTools},
		{"TitleAction", TitleAction},
		{"TitleStatus", TitleStatus},
		{"TitleSelection", TitleSelection},
	}
	for _, tt := range tests {
		t.Run(tt.name+" is non-empty", func(t *testing.T) {
			if tt.value == "" {
				t.Errorf("Constant %s should not be empty", tt.name)
			}
		})
	}
}

// TestUpdateMessageConstants_AreNonEmpty tests that update message constants are defined and non-empty.
// Priority: P3 - Configuration validation.
// Tests that update notification constants have valid values.
func TestUpdateMessageConstants_AreNonEmpty(t *testing.T) {
	tests := []struct {
		name  string
		value string
	}{
		{"UpdateAvailable", UpdateAvailable},
		{"UpdateDownloading", UpdateDownloading},
		{"UpdateInstalling", UpdateInstalling},
		{"UpdateSuccess", UpdateSuccess},
		{"UpdateFailed", UpdateFailed},
		{"UpdateCheckFailed", UpdateCheckFailed},
		{"UpdateNotAvailable", UpdateNotAvailable},
	}
	for _, tt := range tests {
		t.Run(tt.name+" is non-empty", func(t *testing.T) {
			if tt.value == "" {
				t.Errorf("Constant %s should not be empty", tt.name)
			}
		})
	}
}

// TestLogo_IsNonEmpty tests that logo is defined and non-empty.
// Priority: P3 - Configuration validation.
// Tests that the ASCII art logo is present.
func TestLogo_IsNonEmpty(t *testing.T) {
	t.Run("logo is non-empty", func(t *testing.T) {
		if Logo == "" {
			t.Error("Logo constant should not be empty")
		}
	})

	t.Run("logo contains expected content", func(t *testing.T) {
		if !strings.Contains(Logo, "lazysetup") && !strings.Contains(Logo, "P.E.L.E.") {
			t.Error("Logo should contain 'lazysetup' or 'P.E.L.E.'")
		}
	})
}

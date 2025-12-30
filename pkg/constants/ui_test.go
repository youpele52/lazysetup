package constants

import (
	"testing"
)

// TestConstants_AllDefinedAndNonEmpty tests that UI constants are defined and non-empty.
// Priority: P3 - Configuration validation.
// Tests that key UI constants have valid values.
func TestConstants_AllDefinedAndNonEmpty(t *testing.T) {
	t.Run("view constants are non-empty", func(t *testing.T) {
		views := []string{ViewMenu, ViewResult, ViewResults, ViewTools, ViewInstalling}
		for _, v := range views {
			if v == "" {
				t.Error("View constant should not be empty")
			}
		}
	})

	t.Run("panel constants are non-empty", func(t *testing.T) {
		panels := []string{PanelTools, PanelPackageManager, PanelProgress, PanelAction, PanelStatusView}
		for _, p := range panels {
			if p == "" {
				t.Error("Panel constant should not be empty")
			}
		}
	})

	t.Run("title constants are non-empty", func(t *testing.T) {
		titles := []string{TitlePackageManager, TitleInstalling, TitleTools, TitleAction, TitleStatus}
		for _, title := range titles {
			if title == "" {
				t.Error("Title constant should not be empty")
			}
		}
	})

	t.Run("update message constants are non-empty", func(t *testing.T) {
		messages := []string{UpdateAvailable, UpdateDownloading, UpdateInstalling, UpdateSuccess}
		for _, msg := range messages {
			if msg == "" {
				t.Error("Update message constant should not be empty")
			}
		}
	})

	t.Run("logo is non-empty", func(t *testing.T) {
		if Logo == "" {
			t.Error("Logo should not be empty")
		}
	})
}

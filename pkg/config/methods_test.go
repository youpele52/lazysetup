package config

import (
	"testing"
)

// TestInstallMethods_Contains6Managers tests that InstallMethods contains 6 package managers.
// Priority: P3 - Configuration validation.
// Tests that InstallMethods slice has exactly 6 package managers.
func TestInstallMethods_Contains6Managers(t *testing.T) {
	t.Run("contains exactly 6 package managers", func(t *testing.T) {
		if len(InstallMethods) != 6 {
			t.Errorf("Expected 6 install methods, got %d", len(InstallMethods))
		}
	})

	t.Run("contains expected managers", func(t *testing.T) {
		expected := []string{"Homebrew", "APT", "Curl", "YUM", "Scoop", "Chocolatey"}
		for _, method := range expected {
			found := false
			for _, m := range InstallMethods {
				if m == method {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("Expected method '%s' not found", method)
			}
		}
	})
}

// TestActions_Contains4Actions tests that Actions contains 4 actions.
// Priority: P3 - Configuration validation.
// Tests that Actions slice has exactly 4 actions.
func TestActions_Contains4Actions(t *testing.T) {
	t.Run("contains exactly 4 actions", func(t *testing.T) {
		if len(Actions) != 4 {
			t.Errorf("Expected 4 actions, got %d", len(Actions))
		}
	})

	t.Run("contains expected actions", func(t *testing.T) {
		expected := []string{"Check", "Install", "Update", "Uninstall"}
		for _, action := range expected {
			found := false
			for _, a := range Actions {
				if a == action {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("Expected action '%s' not found", action)
			}
		}
	})
}

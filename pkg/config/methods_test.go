package config

import (
	"testing"
)

// TestInstallMethods_Contains9Managers tests that InstallMethods contains 9 package managers.
// Priority: P3 - Configuration validation.
// Tests that InstallMethods slice has exactly 9 package managers.
func TestInstallMethods_Contains9Managers(t *testing.T) {
	t.Run("contains exactly 9 package managers", func(t *testing.T) {
		if len(InstallMethods) != 9 {
			t.Errorf("Expected 9 install methods, got %d", len(InstallMethods))
		}
	})

	t.Run("contains expected managers", func(t *testing.T) {
		expectedManagers := map[string]bool{
			"Homebrew":   true,
			"APT":        true,
			"Curl":       true,
			"YUM":        true,
			"Scoop":      true,
			"Chocolatey": true,
			"Pacman":     true,
			"DNF":        true,
			"Nix":        true,
		}
		for _, method := range InstallMethods {
			if !expectedManagers[method] {
				t.Errorf("Unexpected method '%s' in InstallMethods slice", method)
			}
		}
		if len(InstallMethods) != len(expectedManagers) {
			t.Errorf("Expected %d methods, got %d", len(expectedManagers), len(InstallMethods))
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
		expectedActions := map[string]bool{
			"Check":     true,
			"Install":   true,
			"Update":    true,
			"Uninstall": true,
		}
		for _, action := range Actions {
			if !expectedActions[action] {
				t.Errorf("Unexpected action '%s' in Actions slice", action)
			}
		}
		if len(Actions) != len(expectedActions) {
			t.Errorf("Expected %d actions, got %d", len(expectedActions), len(Actions))
		}
	})
}

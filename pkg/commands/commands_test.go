package commands

import (
	"testing"
)

// TestGetInstallCommand_ValidReturnsCommand tests that GetInstallCommand returns valid commands.
// Priority: P1 - Incorrect commands cause all installations to fail.
// Tests all 6 package managers (Homebrew, APT, YUM, Curl, Scoop, Chocolatey) with all supported tools.
func TestGetInstallCommand_ValidReturnsCommand(t *testing.T) {
	t.Run("Homebrew returns valid commands for all tools", func(t *testing.T) {
		tools := []string{"git", "docker", "lazygit", "lazydocker", "htop"}
		for _, tool := range tools {
			cmd := GetInstallCommand("Homebrew", tool)
			if cmd == "" {
				t.Errorf("Expected command for Homebrew/%s, got empty string", tool)
			}
		}
	})

	t.Run("APT returns valid commands for all tools", func(t *testing.T) {
		tools := []string{"git", "docker", "lazygit", "lazydocker", "htop"}
		for _, tool := range tools {
			cmd := GetInstallCommand("APT", tool)
			if cmd == "" {
				t.Errorf("Expected command for APT/%s, got empty string", tool)
			}
		}
	})

	t.Run("YUM returns valid commands for all tools", func(t *testing.T) {
		tools := []string{"git", "docker", "lazygit", "lazydocker", "htop"}
		for _, tool := range tools {
			cmd := GetInstallCommand("YUM", tool)
			if cmd == "" {
				t.Errorf("Expected command for YUM/%s, got empty string", tool)
			}
		}
	})

	t.Run("Curl returns valid commands for all tools", func(t *testing.T) {
		tools := []string{"git", "docker", "lazygit", "lazydocker", "htop"}
		for _, tool := range tools {
			cmd := GetInstallCommand("Curl", tool)
			if cmd == "" {
				t.Errorf("Expected command for Curl/%s, got empty string", tool)
			}
		}
	})

	t.Run("Scoop returns valid commands for all tools", func(t *testing.T) {
		tools := []string{"git", "docker", "lazygit", "lazydocker", "htop"}
		for _, tool := range tools {
			cmd := GetInstallCommand("Scoop", tool)
			if cmd == "" {
				t.Errorf("Expected command for Scoop/%s, got empty string", tool)
			}
		}
	})

	t.Run("Chocolatey returns valid commands for all tools", func(t *testing.T) {
		tools := []string{"git", "docker", "lazygit", "lazydocker", "htop"}
		for _, tool := range tools {
			cmd := GetInstallCommand("Chocolatey", tool)
			if cmd == "" {
				t.Errorf("Expected command for Chocolatey/%s, got empty string", tool)
			}
		}
	})
}

// TestGetInstallCommand_InvalidReturnsEmpty tests that GetInstallCommand returns empty for invalid inputs.
// Priority: P1 - Empty commands cause confusing errors and must be handled gracefully.
// Tests invalid method, invalid tool, empty method, and empty tool scenarios.
func TestGetInstallCommand_InvalidReturnsEmpty(t *testing.T) {
	t.Run("invalid method returns empty", func(t *testing.T) {
		cmd := GetInstallCommand("InvalidMethod", "git")
		if cmd != "" {
			t.Errorf("Expected empty string for invalid method, got '%s'", cmd)
		}
	})

	t.Run("invalid tool returns empty", func(t *testing.T) {
		cmd := GetInstallCommand("Homebrew", "nonexistent-tool")
		if cmd != "" {
			t.Errorf("Expected empty string for invalid tool, got '%s'", cmd)
		}
	})

	t.Run("empty method returns empty", func(t *testing.T) {
		cmd := GetInstallCommand("", "git")
		if cmd != "" {
			t.Errorf("Expected empty string for empty method, got '%s'", cmd)
		}
	})

	t.Run("empty tool returns empty", func(t *testing.T) {
		cmd := GetInstallCommand("Homebrew", "")
		if cmd != "" {
			t.Errorf("Expected empty string for empty tool, got '%s'", cmd)
		}
	})
}

// TestGetUpdateCommand_ValidReturnsCommand tests that GetUpdateCommand returns valid update commands.
// Priority: P1 - Update functionality is critical for security patches.
// Tests Homebrew, APT, and YUM package managers with all supported tools.
func TestGetUpdateCommand_ValidReturnsCommand(t *testing.T) {
	t.Run("Homebrew returns valid update commands", func(t *testing.T) {
		tools := []string{"git", "docker", "lazygit", "lazydocker", "htop"}
		for _, tool := range tools {
			cmd := GetUpdateCommand("Homebrew", tool)
			if cmd == "" {
				t.Errorf("Expected update command for Homebrew/%s, got empty string", tool)
			}
		}
	})

	t.Run("APT returns valid update commands", func(t *testing.T) {
		tools := []string{"git", "docker", "lazygit", "lazydocker", "htop"}
		for _, tool := range tools {
			cmd := GetUpdateCommand("APT", tool)
			if cmd == "" {
				t.Errorf("Expected update command for APT/%s, got empty string", tool)
			}
		}
	})

	t.Run("YUM returns valid update commands", func(t *testing.T) {
		tools := []string{"git", "docker", "lazygit", "lazydocker", "htop"}
		for _, tool := range tools {
			cmd := GetUpdateCommand("YUM", tool)
			if cmd == "" {
				t.Errorf("Expected update command for YUM/%s, got empty string", tool)
			}
		}
	})
}

// TestGetUpdateCommand_InvalidReturnsEmpty tests that GetUpdateCommand returns empty for invalid inputs.
// Priority: P1 - Prevents silent update failures.
// Tests invalid method and invalid tool scenarios.
func TestGetUpdateCommand_InvalidReturnsEmpty(t *testing.T) {
	t.Run("invalid method returns empty", func(t *testing.T) {
		cmd := GetUpdateCommand("InvalidMethod", "git")
		if cmd != "" {
			t.Errorf("Expected empty string for invalid method, got '%s'", cmd)
		}
	})

	t.Run("invalid tool returns empty", func(t *testing.T) {
		cmd := GetUpdateCommand("Homebrew", "nonexistent-tool")
		if cmd != "" {
			t.Errorf("Expected empty string for invalid tool, got '%s'", cmd)
		}
	})
}

// TestGetUninstallCommand_ValidReturnsCommand tests that GetUninstallCommand returns valid uninstall commands.
// Priority: P1 - Users need to cleanly remove tools.
// Tests Homebrew, APT, and YUM package managers with all supported tools.
func TestGetUninstallCommand_ValidReturnsCommand(t *testing.T) {
	t.Run("Homebrew returns valid uninstall commands", func(t *testing.T) {
		tools := []string{"git", "docker", "lazygit", "lazydocker", "htop"}
		for _, tool := range tools {
			cmd := GetUninstallCommand("Homebrew", tool)
			if cmd == "" {
				t.Errorf("Expected uninstall command for Homebrew/%s, got empty string", tool)
			}
		}
	})

	t.Run("APT returns valid uninstall commands", func(t *testing.T) {
		tools := []string{"git", "docker", "lazygit", "lazydocker", "htop"}
		for _, tool := range tools {
			cmd := GetUninstallCommand("APT", tool)
			if cmd == "" {
				t.Errorf("Expected uninstall command for APT/%s, got empty string", tool)
			}
		}
	})

	t.Run("YUM returns valid uninstall commands", func(t *testing.T) {
		tools := []string{"git", "docker", "lazygit", "lazydocker", "htop"}
		for _, tool := range tools {
			cmd := GetUninstallCommand("YUM", tool)
			if cmd == "" {
				t.Errorf("Expected uninstall command for YUM/%s, got empty string", tool)
			}
		}
	})
}

// TestGetUninstallCommand_InvalidReturnsEmpty tests that GetUninstallCommand returns empty for invalid inputs.
// Priority: P1 - Prevents partial uninstalls that leave system dirty.
// Tests invalid method and invalid tool scenarios.
func TestGetUninstallCommand_InvalidReturnsEmpty(t *testing.T) {
	t.Run("invalid method returns empty", func(t *testing.T) {
		cmd := GetUninstallCommand("InvalidMethod", "git")
		if cmd != "" {
			t.Errorf("Expected empty string for invalid method, got '%s'", cmd)
		}
	})

	t.Run("invalid tool returns empty", func(t *testing.T) {
		cmd := GetUninstallCommand("Homebrew", "nonexistent-tool")
		if cmd != "" {
			t.Errorf("Expected empty string for invalid tool, got '%s'", cmd)
		}
	})
}

// TestGetCheckCommand_ReturnsCorrectCommand tests that GetCheckCommand returns correct package manager check commands.
// Priority: P1 - Package manager detection is essential for the application to work.
// Tests all package managers and verifies invalid managers return empty.
func TestGetCheckCommand_ReturnsCorrectCommand(t *testing.T) {
	t.Run("returns check commands for all package managers", func(t *testing.T) {
		managers := []string{"Homebrew", "APT", "YUM", "Scoop", "Chocolatey"}
		for _, manager := range managers {
			cmd := GetCheckCommand(manager)
			if cmd == "" {
				t.Errorf("Expected check command for %s, got empty string", manager)
			}
		}
	})

	t.Run("invalid manager returns empty", func(t *testing.T) {
		cmd := GetCheckCommand("InvalidManager")
		if cmd != "" {
			t.Errorf("Expected empty string for invalid manager, got '%s'", cmd)
		}
	})
}

// TestGetToolCheckCommand_ReturnsCorrectCommand tests that GetToolCheckCommand returns correct tool version commands.
// Priority: P1 - Tool version checking is a core feature for status display.
// Tests all supported tools and verifies invalid tools return empty.
func TestGetToolCheckCommand_ReturnsCorrectCommand(t *testing.T) {
	t.Run("returns check commands for all tools", func(t *testing.T) {
		tools := []string{"git", "docker", "lazygit", "lazydocker", "htop"}
		for _, tool := range tools {
			cmd := GetToolCheckCommand(tool)
			if cmd == "" {
				t.Errorf("Expected check command for %s, got empty string", tool)
			}
		}
	})

	t.Run("invalid tool returns empty", func(t *testing.T) {
		cmd := GetToolCheckCommand("nonexistent-tool")
		if cmd != "" {
			t.Errorf("Expected empty string for invalid tool, got '%s'", cmd)
		}
	})
}

// TestGetLifecycleCommand_ValidReturnsCommand tests the core GetLifecycleCommand function.
// Priority: P1 - Used by all install/update/uninstall operations.
// Tests that valid method and tool combinations return non-empty commands.
func TestGetLifecycleCommand_ValidReturnsCommand(t *testing.T) {
	t.Run("returns command for valid method and tool", func(t *testing.T) {
		params := GetLifecycleCommandType{
			method:                "Homebrew",
			tool:                  "git",
			LifecycleCommandsType: PackageManagerInstallCommands,
		}
		cmd := GetLifecycleCommand(params)
		if cmd == "" {
			t.Error("Expected non-empty command")
		}
	})
}

// TestGetLifecycleCommand_InvalidReturnsEmpty tests GetLifecycleCommand error handling.
// Priority: P1 - Error handling depends on empty string returns for invalid inputs.
// Tests invalid method and invalid tool scenarios.
func TestGetLifecycleCommand_InvalidReturnsEmpty(t *testing.T) {
	t.Run("returns empty for invalid method", func(t *testing.T) {
		params := GetLifecycleCommandType{
			method:                "InvalidMethod",
			tool:                  "git",
			LifecycleCommandsType: PackageManagerInstallCommands,
		}
		cmd := GetLifecycleCommand(params)
		if cmd != "" {
			t.Errorf("Expected empty string, got '%s'", cmd)
		}
	})

	t.Run("returns empty for invalid tool", func(t *testing.T) {
		params := GetLifecycleCommandType{
			method:                "Homebrew",
			tool:                  "invalid-tool",
			LifecycleCommandsType: PackageManagerInstallCommands,
		}
		cmd := GetLifecycleCommand(params)
		if cmd != "" {
			t.Errorf("Expected empty string, got '%s'", cmd)
		}
	})
}

// TestMergeMaps_MergesCorrectly tests the MergeMaps utility function.
// Priority: P2 - Data structure validation for command map merging.
// Tests merging two maps, override behavior, and empty map handling.
func TestMergeMaps_MergesCorrectly(t *testing.T) {
	t.Run("merges two maps", func(t *testing.T) {
		map1 := map[string]string{"a": "1", "b": "2"}
		map2 := map[string]string{"c": "3", "d": "4"}
		result := MergeMaps(map1, map2)

		if len(result) != 4 {
			t.Errorf("Expected 4 items, got %d", len(result))
		}
		if result["a"] != "1" || result["c"] != "3" {
			t.Error("Merged map has incorrect values")
		}
	})

	t.Run("later maps override earlier", func(t *testing.T) {
		map1 := map[string]string{"a": "1"}
		map2 := map[string]string{"a": "2"}
		result := MergeMaps(map1, map2)

		if result["a"] != "2" {
			t.Errorf("Expected 'a' to be '2', got '%s'", result["a"])
		}
	})

	t.Run("handles empty maps", func(t *testing.T) {
		map1 := map[string]string{}
		map2 := map[string]string{"a": "1"}
		result := MergeMaps(map1, map2)

		if len(result) != 1 {
			t.Errorf("Expected 1 item, got %d", len(result))
		}
	})
}

// TestCommandCorrectness_Homebrew tests that Homebrew commands have correct syntax.
// Priority: P1 - Every Homebrew command combination must work for users.
// Tests that install, update, and uninstall commands start with 'brew'.
func TestCommandCorrectness_Homebrew(t *testing.T) {
	t.Run("Homebrew install commands use brew install", func(t *testing.T) {
		tools := []string{"git", "docker", "lazygit", "lazydocker", "htop"}
		for _, tool := range tools {
			cmd := GetInstallCommand("Homebrew", tool)
			if cmd == "" {
				t.Errorf("No command for %s", tool)
				continue
			}
			if len(cmd) < 5 || cmd[:4] != "brew" {
				t.Errorf("Homebrew command for %s should start with 'brew', got '%s'", tool, cmd)
			}
		}
	})

	t.Run("Homebrew update commands use brew upgrade", func(t *testing.T) {
		tools := []string{"git", "docker", "lazygit", "lazydocker", "htop"}
		for _, tool := range tools {
			cmd := GetUpdateCommand("Homebrew", tool)
			if cmd == "" {
				t.Errorf("No update command for %s", tool)
				continue
			}
			if len(cmd) < 5 || cmd[:4] != "brew" {
				t.Errorf("Homebrew update command for %s should start with 'brew', got '%s'", tool, cmd)
			}
		}
	})

	t.Run("Homebrew uninstall commands use brew uninstall", func(t *testing.T) {
		tools := []string{"git", "docker", "lazygit", "lazydocker", "htop"}
		for _, tool := range tools {
			cmd := GetUninstallCommand("Homebrew", tool)
			if cmd == "" {
				t.Errorf("No uninstall command for %s", tool)
				continue
			}
			if len(cmd) < 5 || cmd[:4] != "brew" {
				t.Errorf("Homebrew uninstall command for %s should start with 'brew', got '%s'", tool, cmd)
			}
		}
	})
}

// TestCommandCorrectness_APT tests that APT commands have correct syntax.
// Priority: P1 - Every APT command combination must work for users.
// Tests that install commands start with 'apt-get'.
func TestCommandCorrectness_APT(t *testing.T) {
	t.Run("APT install commands use apt-get install", func(t *testing.T) {
		tools := []string{"git", "docker", "lazygit", "lazydocker", "htop"}
		for _, tool := range tools {
			cmd := GetInstallCommand("APT", tool)
			if cmd == "" {
				t.Errorf("No command for %s", tool)
				continue
			}
			if len(cmd) < 7 || cmd[:7] != "apt-get" {
				t.Errorf("APT command for %s should start with 'apt-get', got '%s'", tool, cmd)
			}
		}
	})
}

// TestCommandCorrectness_YUM tests that YUM commands have correct syntax.
// Priority: P1 - Every YUM command combination must work for users.
// Tests that install commands start with 'yum'.
func TestCommandCorrectness_YUM(t *testing.T) {
	t.Run("YUM install commands use yum install", func(t *testing.T) {
		tools := []string{"git", "docker", "lazygit", "lazydocker", "htop"}
		for _, tool := range tools {
			cmd := GetInstallCommand("YUM", tool)
			if cmd == "" {
				t.Errorf("No command for %s", tool)
				continue
			}
			if len(cmd) < 3 || cmd[:3] != "yum" {
				t.Errorf("YUM command for %s should start with 'yum', got '%s'", tool, cmd)
			}
		}
	})
}

package commands

import "testing"

var testedManagers = []string{"Homebrew", "APT", "YUM", "DNF", "Nix", "Curl", "Scoop", "Chocolatey", "Pacman"}

func supportedTools(method string) []string {
	return GetSupportedToolsForMethod(method)
}

// TestGetInstallCommand_ValidReturnsCommand tests that install commands exist for supported pairs.
// Priority: P1 - Incorrect commands cause installations to fail.
func TestGetInstallCommand_ValidReturnsCommand(t *testing.T) {
	for _, method := range testedManagers {
		method := method
		t.Run(method+" returns valid install commands", func(t *testing.T) {
			for _, tool := range supportedTools(method) {
				cmd := GetInstallCommand(method, tool)
				if cmd == "" {
					t.Errorf("Expected install command for %s/%s, got empty string", method, tool)
				}
			}
		})
	}
}

// TestGetInstallCommand_InvalidReturnsEmpty tests invalid or unsupported install lookups.
// Priority: P1 - Unsupported combinations must fail cleanly.
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

	t.Run("unsupported tool-method pair returns empty", func(t *testing.T) {
		cmd := GetInstallCommand("APT", "codex")
		if cmd != "" {
			t.Errorf("Expected empty string for unsupported pair, got '%s'", cmd)
		}
	})

	t.Run("openclaw uses curl fallback on apt", func(t *testing.T) {
		cmd := GetInstallCommand("APT", "openclaw")
		if cmd != openclawUnixInstallScript {
			t.Errorf("Expected openclaw apt install to use curl fallback, got '%s'", cmd)
		}
	})
}

// TestGetUpdateCommand_ValidReturnsCommand tests that update commands exist for supported pairs.
// Priority: P1 - Update functionality is critical for security patches.
func TestGetUpdateCommand_ValidReturnsCommand(t *testing.T) {
	for _, method := range testedManagers {
		method := method
		t.Run(method+" returns valid update commands", func(t *testing.T) {
			for _, tool := range supportedTools(method) {
				cmd := GetUpdateCommand(method, tool)
				if cmd == "" {
					t.Errorf("Expected update command for %s/%s, got empty string", method, tool)
				}
			}
		})
	}
}

// TestGetUpdateCommand_InvalidReturnsEmpty tests invalid or unsupported update lookups.
// Priority: P1 - Prevents silent update failures.
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

// TestGetUninstallCommand_ValidReturnsCommand tests that uninstall commands exist for supported pairs.
// Priority: P1 - Users need to cleanly remove tools.
func TestGetUninstallCommand_ValidReturnsCommand(t *testing.T) {
	for _, method := range testedManagers {
		method := method
		t.Run(method+" returns valid uninstall commands", func(t *testing.T) {
			for _, tool := range supportedTools(method) {
				cmd := GetUninstallCommand(method, tool)
				if cmd == "" {
					t.Errorf("Expected uninstall command for %s/%s, got empty string", method, tool)
				}
			}
		})
	}
}

// TestGetUninstallCommand_InvalidReturnsEmpty tests invalid or unsupported uninstall lookups.
// Priority: P1 - Prevents partial uninstalls.
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

// TestGetCheckCommand_ReturnsCorrectCommand tests package manager check commands.
// Priority: P1 - Package manager detection is essential.
func TestGetCheckCommand_ReturnsCorrectCommand(t *testing.T) {
	for _, manager := range testedManagers {
		cmd := GetCheckCommand(manager)
		if cmd == "" {
			t.Errorf("Expected check command for %s, got empty string", manager)
		}
	}

	cmd := GetCheckCommand("InvalidManager")
	if cmd != "" {
		t.Errorf("Expected empty string for invalid manager, got '%s'", cmd)
	}
}

// TestGetToolCheckCommand_ReturnsCorrectCommand tests tool version commands.
// Priority: P1 - Tool version checking is a core feature.
func TestGetToolCheckCommand_ReturnsCorrectCommand(t *testing.T) {
	for tool := range supportedMethodsByTool {
		cmd := GetToolCheckCommand(tool)
		if cmd == "" {
			t.Errorf("Expected check command for %s, got empty string", tool)
		}
	}

	cmd := GetToolCheckCommand("nonexistent-tool")
	if cmd != "" {
		t.Errorf("Expected empty string for invalid tool, got '%s'", cmd)
	}
}

// TestGetSupportedToolsForMethod_ReturnsOnlySupportedTools validates the support matrix.
// Priority: P1 - The UI and commands both depend on correct method-specific availability.
func TestGetSupportedToolsForMethod_ReturnsOnlySupportedTools(t *testing.T) {
	homebrewTools := supportedTools("Homebrew")
	if len(homebrewTools) == 0 {
		t.Fatal("Expected Homebrew supported tools")
	}

	if !IsToolSupportedByMethod("codex", "Homebrew") {
		t.Error("Expected codex to be supported on Homebrew")
	}

	if IsToolSupportedByMethod("codex", "APT") {
		t.Error("Did not expect codex to be supported on APT")
	}

	if !IsToolSupportedByMethod("openclaw", "Curl") {
		t.Error("Expected openclaw to be supported on Curl")
	}

	if !IsToolSupportedByMethod("openclaw", "Homebrew") {
		t.Error("Expected openclaw to be supported via Unix package-manager fallback")
	}
}

// TestGetLifecycleCommand_ValidReturnsCommand tests the core GetLifecycleCommand function.
// Priority: P1 - Used by all install/update/uninstall operations.
func TestGetLifecycleCommand_ValidReturnsCommand(t *testing.T) {
	params := GetLifecycleCommandType{
		method:                "Homebrew",
		tool:                  "git",
		LifecycleCommandsType: PackageManagerInstallCommands,
	}
	cmd := GetLifecycleCommand(params)
	if cmd == "" {
		t.Error("Expected non-empty command")
	}
}

// TestGetLifecycleCommand_InvalidReturnsEmpty tests GetLifecycleCommand error handling.
// Priority: P1 - Error handling depends on empty string returns for invalid inputs.
func TestGetLifecycleCommand_InvalidReturnsEmpty(t *testing.T) {
	params := GetLifecycleCommandType{
		method:                "InvalidMethod",
		tool:                  "git",
		LifecycleCommandsType: PackageManagerInstallCommands,
	}
	cmd := GetLifecycleCommand(params)
	if cmd != "" {
		t.Errorf("Expected empty string, got '%s'", cmd)
	}

	params = GetLifecycleCommandType{
		method:                "Homebrew",
		tool:                  "invalid-tool",
		LifecycleCommandsType: PackageManagerInstallCommands,
	}
	cmd = GetLifecycleCommand(params)
	if cmd != "" {
		t.Errorf("Expected empty string, got '%s'", cmd)
	}
}

// TestMergeMaps_MergesCorrectly tests the MergeMaps utility function.
// Priority: P2 - Data structure validation for command map merging.
func TestMergeMaps_MergesCorrectly(t *testing.T) {
	map1 := map[string]string{"a": "1", "b": "2"}
	map2 := map[string]string{"c": "3", "d": "4"}
	result := MergeMaps(map1, map2)

	if len(result) != 4 {
		t.Errorf("Expected 4 items, got %d", len(result))
	}
	if result["a"] != "1" || result["c"] != "3" {
		t.Error("Merged map has incorrect values")
	}

	result = MergeMaps(map[string]string{"a": "1"}, map[string]string{"a": "2"})
	if result["a"] != "2" {
		t.Errorf("Expected 'a' to be '2', got '%s'", result["a"])
	}
}

// TestCommandCorrectness_Homebrew tests Homebrew command syntax for supported tools.
// Priority: P1 - Every supported Homebrew combination must work.
func TestCommandCorrectness_Homebrew(t *testing.T) {
	for _, tool := range supportedTools("Homebrew") {
		cmd := GetInstallCommand("Homebrew", tool)
		if tool == "openclaw" {
			if cmd != openclawUnixInstallScript {
				t.Errorf("Homebrew install command for openclaw should use curl fallback, got '%s'", cmd)
			}
		} else if len(cmd) < 5 || cmd[:4] != "brew" {
			t.Errorf("Homebrew install command for %s should start with 'brew', got '%s'", tool, cmd)
		}

		cmd = GetUpdateCommand("Homebrew", tool)
		if tool == "openclaw" {
			if cmd != openclawUnixInstallScript {
				t.Errorf("Homebrew update command for openclaw should use curl fallback, got '%s'", cmd)
			}
		} else if len(cmd) < 5 || cmd[:4] != "brew" {
			t.Errorf("Homebrew update command for %s should start with 'brew', got '%s'", tool, cmd)
		}

		cmd = GetUninstallCommand("Homebrew", tool)
		if tool == "openclaw" {
			if cmd == "" {
				t.Error("Homebrew uninstall command for openclaw should not be empty")
			}
		} else if len(cmd) < 5 || cmd[:4] != "brew" {
			t.Errorf("Homebrew uninstall command for %s should start with 'brew', got '%s'", tool, cmd)
		}
	}
}

// TestCommandCorrectness_APT tests APT command syntax for supported tools.
// Priority: P1 - Every supported APT combination must work.
func TestCommandCorrectness_APT(t *testing.T) {
	for _, tool := range supportedTools("APT") {
		cmd := GetInstallCommand("APT", tool)
		if tool == "openclaw" {
			if cmd != openclawUnixInstallScript {
				t.Errorf("APT command for openclaw should use curl fallback, got '%s'", cmd)
			}
		} else if len(cmd) < 7 || cmd[:7] != "apt-get" {
			t.Errorf("APT command for %s should start with 'apt-get', got '%s'", tool, cmd)
		}
	}
}

// TestCommandCorrectness_YUM tests YUM command syntax for supported tools.
// Priority: P1 - Every supported YUM combination must work.
func TestCommandCorrectness_YUM(t *testing.T) {
	for _, tool := range supportedTools("YUM") {
		cmd := GetInstallCommand("YUM", tool)
		if tool == "openclaw" {
			if cmd != openclawUnixInstallScript {
				t.Errorf("YUM command for openclaw should use curl fallback, got '%s'", cmd)
			}
		} else if len(cmd) < 3 || cmd[:3] != "yum" {
			t.Errorf("YUM command for %s should start with 'yum', got '%s'", tool, cmd)
		}
	}
}

// TestCurlCommands_UseLatestVersions tests Curl commands that should track latest releases.
// Priority: P1 - Hardcoded versions become outdated.
func TestCurlCommands_UseLatestVersions(t *testing.T) {
	githubTools := []struct {
		tool          string
		shouldContain string
	}{
		{"nvim", "/releases/latest/"},
		{"ripgrep", "/releases/latest/"},
		{"fd", "/releases/latest/"},
		{"bat", "/releases/latest/"},
		{"gh", "/releases/latest/"},
		{"eza", "/releases/latest/"},
		{"delta", "/releases/latest/"},
		{"btop", "/releases/latest/"},
		{"lazysql", "/releases/latest/"},
		{"k9s", "/releases/latest/"},
		{"codex", "/releases/latest/"},
	}

	for _, test := range githubTools {
		cmd := GetInstallCommand("Curl", test.tool)
		if cmd == "" {
			t.Errorf("No Curl command for %s", test.tool)
			continue
		}
		if !contains(cmd, test.shouldContain) {
			t.Errorf("Tool %s should use '%s' pattern for auto-updates, got: %s", test.tool, test.shouldContain, cmd)
		}
	}

	cmd := GetInstallCommand("Curl", "node")
	if !contains(cmd, "latest-v20.x") {
		t.Errorf("Node should use latest-v20.x pattern for auto-updates, got: %s", cmd)
	}

	cmd = GetInstallCommand("Curl", "wget")
	if !contains(cmd, "wget-latest") {
		t.Errorf("Wget should use wget-latest pattern for auto-updates, got: %s", cmd)
	}
}

// TestCurlCommands_UseArchitectureDetection tests dynamic architecture detection for Curl commands.
// Priority: P1 - Hardcoded architectures break on ARM systems.
func TestCurlCommands_UseArchitectureDetection(t *testing.T) {
	archTools := []string{"ripgrep", "fd", "bat", "gh", "eza", "delta", "btop", "httpie", "lazysql", "tldr", "k9s", "kubectl", "codex"}

	for _, tool := range archTools {
		cmd := GetInstallCommand("Curl", tool)
		if cmd == "" {
			t.Errorf("No Curl command for %s", tool)
			continue
		}
		if !contains(cmd, "$(uname") {
			t.Errorf("Tool %s should use $(uname -m) or $(uname -s) for architecture detection, got: %s", tool, cmd)
		}
	}
}

// Helper function to check if a string contains a substring.
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(s) > len(substr) && containsAt(s, substr))
}

func containsAt(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}

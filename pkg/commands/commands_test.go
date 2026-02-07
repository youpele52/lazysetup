package commands

import (
	"testing"
)

// TestGetInstallCommand_ValidReturnsCommand tests that GetInstallCommand returns valid commands.
// Priority: P1 - Incorrect commands cause all installations to fail.
// Tests all 7 package managers (Homebrew, APT, YUM, Curl, Scoop, Chocolatey, Pacman) with all 29 supported tools.
func TestGetInstallCommand_ValidReturnsCommand(t *testing.T) {
	tools := []string{"git", "docker", "lazygit", "lazydocker", "nvim", "zsh", "tmux", "fzf", "ripgrep", "fd", "bat", "jq", "node", "gh", "eza", "zoxide", "starship", "python3", "bun", "pnpm", "uv", "delta", "tree", "make", "just", "btop", "wget", "httpie", "tldr", "lazysql", "claude-code", "opencode", "rsync", "kubectl", "k9s", "terraform", "helm"}

	t.Run("Homebrew returns valid commands for all tools", func(t *testing.T) {
		for _, tool := range tools {
			cmd := GetInstallCommand("Homebrew", tool)
			if cmd == "" {
				t.Errorf("Expected command for Homebrew/%s, got empty string", tool)
			}
		}
	})

	t.Run("APT returns valid commands for all tools", func(t *testing.T) {
		for _, tool := range tools {
			cmd := GetInstallCommand("APT", tool)
			if cmd == "" {
				t.Errorf("Expected command for APT/%s, got empty string", tool)
			}
		}
	})

	t.Run("YUM returns valid commands for all tools", func(t *testing.T) {
		for _, tool := range tools {
			cmd := GetInstallCommand("YUM", tool)
			if cmd == "" {
				t.Errorf("Expected command for YUM/%s, got empty string", tool)
			}
		}
	})

	t.Run("Curl returns valid commands for all tools", func(t *testing.T) {
		for _, tool := range tools {
			cmd := GetInstallCommand("Curl", tool)
			if cmd == "" {
				t.Errorf("Expected command for Curl/%s, got empty string", tool)
			}
		}
	})

	t.Run("Scoop returns valid commands for all tools", func(t *testing.T) {
		for _, tool := range tools {
			cmd := GetInstallCommand("Scoop", tool)
			if cmd == "" {
				t.Errorf("Expected command for Scoop/%s, got empty string", tool)
			}
		}
	})

	t.Run("Chocolatey returns valid commands for all tools", func(t *testing.T) {
		for _, tool := range tools {
			cmd := GetInstallCommand("Chocolatey", tool)
			if cmd == "" {
				t.Errorf("Expected command for Chocolatey/%s, got empty string", tool)
			}
		}
	})

	t.Run("Pacman returns valid commands for all tools", func(t *testing.T) {
		for _, tool := range tools {
			cmd := GetInstallCommand("Pacman", tool)
			if cmd == "" {
				t.Errorf("Expected command for Pacman/%s, got empty string", tool)
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
// Tests all 7 package managers (Homebrew, APT, YUM, Curl, Scoop, Chocolatey, Pacman) with all 29 supported tools.
func TestGetUpdateCommand_ValidReturnsCommand(t *testing.T) {
	tools := []string{"git", "docker", "lazygit", "lazydocker", "nvim", "zsh", "tmux", "fzf", "ripgrep", "fd", "bat", "jq", "node", "gh", "eza", "zoxide", "starship", "python3", "bun", "pnpm", "uv", "delta", "tree", "make", "just", "btop", "wget", "httpie", "tldr", "lazysql", "claude-code", "opencode", "rsync", "kubectl", "k9s", "terraform", "helm"}

	t.Run("Homebrew returns valid update commands", func(t *testing.T) {
		for _, tool := range tools {
			cmd := GetUpdateCommand("Homebrew", tool)
			if cmd == "" {
				t.Errorf("Expected update command for Homebrew/%s, got empty string", tool)
			}
		}
	})

	t.Run("APT returns valid update commands", func(t *testing.T) {
		for _, tool := range tools {
			cmd := GetUpdateCommand("APT", tool)
			if cmd == "" {
				t.Errorf("Expected update command for APT/%s, got empty string", tool)
			}
		}
	})

	t.Run("YUM returns valid update commands", func(t *testing.T) {
		for _, tool := range tools {
			cmd := GetUpdateCommand("YUM", tool)
			if cmd == "" {
				t.Errorf("Expected update command for YUM/%s, got empty string", tool)
			}
		}
	})

	t.Run("Curl returns valid update commands", func(t *testing.T) {
		for _, tool := range tools {
			cmd := GetUpdateCommand("Curl", tool)
			if cmd == "" {
				t.Errorf("Expected update command for Curl/%s, got empty string", tool)
			}
		}
	})

	t.Run("Scoop returns valid update commands", func(t *testing.T) {
		for _, tool := range tools {
			cmd := GetUpdateCommand("Scoop", tool)
			if cmd == "" {
				t.Errorf("Expected update command for Scoop/%s, got empty string", tool)
			}
		}
	})

	t.Run("Chocolatey returns valid update commands", func(t *testing.T) {
		for _, tool := range tools {
			cmd := GetUpdateCommand("Chocolatey", tool)
			if cmd == "" {
				t.Errorf("Expected update command for Chocolatey/%s, got empty string", tool)
			}
		}
	})

	t.Run("Pacman returns valid update commands", func(t *testing.T) {
		for _, tool := range tools {
			cmd := GetUpdateCommand("Pacman", tool)
			if cmd == "" {
				t.Errorf("Expected update command for Pacman/%s, got empty string", tool)
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
// Tests all 7 package managers (Homebrew, APT, YUM, Curl, Scoop, Chocolatey, Pacman) with all 29 supported tools.
func TestGetUninstallCommand_ValidReturnsCommand(t *testing.T) {
	tools := []string{"git", "docker", "lazygit", "lazydocker", "nvim", "zsh", "tmux", "fzf", "ripgrep", "fd", "bat", "jq", "node", "gh", "eza", "zoxide", "starship", "python3", "bun", "pnpm", "uv", "delta", "tree", "make", "just", "btop", "wget", "httpie", "tldr", "lazysql", "claude-code", "opencode", "rsync", "kubectl", "k9s", "terraform", "helm"}

	t.Run("Homebrew returns valid uninstall commands", func(t *testing.T) {
		for _, tool := range tools {
			cmd := GetUninstallCommand("Homebrew", tool)
			if cmd == "" {
				t.Errorf("Expected uninstall command for Homebrew/%s, got empty string", tool)
			}
		}
	})

	t.Run("APT returns valid uninstall commands", func(t *testing.T) {
		for _, tool := range tools {
			cmd := GetUninstallCommand("APT", tool)
			if cmd == "" {
				t.Errorf("Expected uninstall command for APT/%s, got empty string", tool)
			}
		}
	})

	t.Run("YUM returns valid uninstall commands", func(t *testing.T) {
		for _, tool := range tools {
			cmd := GetUninstallCommand("YUM", tool)
			if cmd == "" {
				t.Errorf("Expected uninstall command for YUM/%s, got empty string", tool)
			}
		}
	})

	t.Run("Curl returns valid uninstall commands", func(t *testing.T) {
		for _, tool := range tools {
			cmd := GetUninstallCommand("Curl", tool)
			if cmd == "" {
				t.Errorf("Expected uninstall command for Curl/%s, got empty string", tool)
			}
		}
	})

	t.Run("Scoop returns valid uninstall commands", func(t *testing.T) {
		for _, tool := range tools {
			cmd := GetUninstallCommand("Scoop", tool)
			if cmd == "" {
				t.Errorf("Expected uninstall command for Scoop/%s, got empty string", tool)
			}
		}
	})

	t.Run("Chocolatey returns valid uninstall commands", func(t *testing.T) {
		for _, tool := range tools {
			cmd := GetUninstallCommand("Chocolatey", tool)
			if cmd == "" {
				t.Errorf("Expected uninstall command for Chocolatey/%s, got empty string", tool)
			}
		}
	})

	t.Run("Pacman returns valid uninstall commands", func(t *testing.T) {
		for _, tool := range tools {
			cmd := GetUninstallCommand("Pacman", tool)
			if cmd == "" {
				t.Errorf("Expected uninstall command for Pacman/%s, got empty string", tool)
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
// Tests all 9 package managers (Homebrew, APT, YUM, Scoop, Chocolatey, Pacman, DNF, Nix, Curl) and verifies invalid managers return empty.
func TestGetCheckCommand_ReturnsCorrectCommand(t *testing.T) {
	t.Run("returns check commands for all package managers", func(t *testing.T) {
		managers := []string{"Homebrew", "APT", "YUM", "Scoop", "Chocolatey", "Pacman", "DNF", "Nix", "Curl"}
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
// Tests all 29 supported tools and verifies invalid tools return empty.
func TestGetToolCheckCommand_ReturnsCorrectCommand(t *testing.T) {
	t.Run("returns check commands for all tools", func(t *testing.T) {
		tools := []string{"git", "docker", "lazygit", "lazydocker", "nvim", "zsh", "tmux", "fzf", "ripgrep", "fd", "bat", "jq", "node", "gh", "eza", "zoxide", "starship", "python3", "bun", "pnpm", "uv", "delta", "tree", "make", "just", "btop", "wget", "httpie", "tldr", "lazysql", "claude-code", "opencode", "rsync", "kubectl", "k9s", "terraform", "helm"}
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
	tools := []string{"git", "docker", "lazygit", "lazydocker", "nvim", "zsh", "tmux", "fzf", "ripgrep", "fd", "bat", "jq", "node", "gh", "eza", "zoxide", "starship", "python3", "bun", "pnpm", "uv", "delta", "tree", "make", "just", "btop", "wget", "httpie", "tldr", "lazysql", "claude-code", "opencode", "rsync", "kubectl", "k9s", "terraform", "helm"}

	t.Run("Homebrew install commands use brew install", func(t *testing.T) {
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
	tools := []string{"git", "docker", "lazygit", "lazydocker", "nvim", "zsh", "tmux", "fzf", "ripgrep", "fd", "bat", "jq", "node", "gh", "eza", "zoxide", "starship", "python3", "bun", "pnpm", "uv", "delta", "tree", "make", "just", "btop", "wget", "httpie", "tldr", "lazysql", "claude-code", "opencode", "rsync", "kubectl", "k9s", "terraform", "helm"}

	t.Run("APT install commands use apt-get install", func(t *testing.T) {
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
	tools := []string{"git", "docker", "lazygit", "lazydocker", "nvim", "zsh", "tmux", "fzf", "ripgrep", "fd", "bat", "jq", "node", "gh", "eza", "zoxide", "starship", "python3", "bun", "pnpm", "uv", "delta", "tree", "make", "just", "btop", "wget", "httpie", "tldr", "lazysql", "claude-code", "opencode", "rsync", "kubectl", "k9s", "terraform", "helm"}

	t.Run("YUM install commands use yum install", func(t *testing.T) {
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

// TestCurlCommands_UseLatestVersions tests that Curl commands use /latest/ or latest-stable URLs where possible.
// Priority: P1 - Hardcoded versions become outdated and cause security vulnerabilities.
// GitHub-hosted tools should use /releases/latest/ pattern. Official sites should use latest or stable URLs.
func TestCurlCommands_UseLatestVersions(t *testing.T) {
	t.Run("GitHub tools use latest releases", func(t *testing.T) {
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
	})

	t.Run("Node uses latest-v20.x pattern", func(t *testing.T) {
		cmd := GetInstallCommand("Curl", "node")
		if cmd == "" {
			t.Error("No Curl command for node")
			return
		}
		if !contains(cmd, "latest-v20.x") {
			t.Errorf("Node should use latest-v20.x pattern for auto-updates, got: %s", cmd)
		}
	})

	t.Run("Wget uses latest pattern", func(t *testing.T) {
		cmd := GetInstallCommand("Curl", "wget")
		if cmd == "" {
			t.Error("No Curl command for wget")
			return
		}
		if !contains(cmd, "wget-latest") {
			t.Errorf("Wget should use wget-latest pattern for auto-updates, got: %s", cmd)
		}
	})

	t.Run("Commands use wildcard patterns for version-agnostic extraction", func(t *testing.T) {
		wildcardTools := []struct {
			tool          string
			shouldContain string
		}{
			{"tmux", "tmux-*"},
			{"node", "node-*"},
			{"python3", "Python-*"},
			{"tree", "tree-*"},
			{"make", "make-*"},
			{"wget", "wget-*"},
			{"rsync", "rsync-*"},
		}

		for _, test := range wildcardTools {
			cmd := GetInstallCommand("Curl", test.tool)
			if cmd == "" {
				t.Errorf("No Curl command for %s", test.tool)
				continue
			}
			if !contains(cmd, test.shouldContain) {
				t.Errorf("Tool %s should use '%s' wildcard pattern for version-agnostic extraction, got: %s", test.tool, test.shouldContain, cmd)
			}
		}
	})
}

// TestCurlCommands_UseArchitectureDetection tests that Curl commands use dynamic architecture detection.
// Priority: P1 - Hardcoded architectures break on ARM servers.
// Commands should use $(uname -m) and $(uname -s) for cross-platform compatibility.
func TestCurlCommands_UseArchitectureDetection(t *testing.T) {
	t.Run("Architecture-specific tools use uname detection", func(t *testing.T) {
		archTools := []string{"ripgrep", "fd", "bat", "gh", "eza", "delta", "btop", "httpie", "lazysql", "tldr", "k9s", "kubectl"}

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
	})
}

// Helper function to check if a string contains a substring
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

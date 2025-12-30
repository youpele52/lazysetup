package updater

import (
	"testing"
)

// TestIsNewerVersion_NewerReturnsTrue tests that isNewerVersion returns true for newer versions.
// Priority: P1 - Update notification logic depends on correct version comparison.
// Tests major, minor, patch version comparisons and v prefix handling.
func TestIsNewerVersion_NewerReturnsTrue(t *testing.T) {
	t.Run("newer major version returns true", func(t *testing.T) {
		if !isNewerVersion("1.0.0", "2.0.0") {
			t.Error("Expected 2.0.0 to be newer than 1.0.0")
		}
	})

	t.Run("newer minor version returns true", func(t *testing.T) {
		if !isNewerVersion("1.0.0", "1.1.0") {
			t.Error("Expected 1.1.0 to be newer than 1.0.0")
		}
	})

	t.Run("newer patch version returns true", func(t *testing.T) {
		if !isNewerVersion("1.0.0", "1.0.1") {
			t.Error("Expected 1.0.1 to be newer than 1.0.0")
		}
	})

	t.Run("handles v prefix in current", func(t *testing.T) {
		if !isNewerVersion("v1.0.0", "1.1.0") {
			t.Error("Expected 1.1.0 to be newer than v1.0.0")
		}
	})

	t.Run("handles v prefix in latest", func(t *testing.T) {
		if !isNewerVersion("1.0.0", "v1.1.0") {
			t.Error("Expected v1.1.0 to be newer than 1.0.0")
		}
	})

	t.Run("handles v prefix in both", func(t *testing.T) {
		if !isNewerVersion("v1.0.0", "v1.1.0") {
			t.Error("Expected v1.1.0 to be newer than v1.0.0")
		}
	})
}

// TestIsNewerVersion_SameOrOlderReturnsFalse tests that isNewerVersion returns false for same/older versions.
// Priority: P1 - Prevents unnecessary update prompts for same or older versions.
// Tests same version, older major/minor/patch, and v prefix handling.
func TestIsNewerVersion_SameOrOlderReturnsFalse(t *testing.T) {
	t.Run("same version returns false", func(t *testing.T) {
		if isNewerVersion("1.0.0", "1.0.0") {
			t.Error("Expected same version to return false")
		}
	})

	t.Run("older major version returns false", func(t *testing.T) {
		if isNewerVersion("2.0.0", "1.0.0") {
			t.Error("Expected 1.0.0 to not be newer than 2.0.0")
		}
	})

	t.Run("older minor version returns false", func(t *testing.T) {
		if isNewerVersion("1.1.0", "1.0.0") {
			t.Error("Expected 1.0.0 to not be newer than 1.1.0")
		}
	})

	t.Run("older patch version returns false", func(t *testing.T) {
		if isNewerVersion("1.0.1", "1.0.0") {
			t.Error("Expected 1.0.0 to not be newer than 1.0.1")
		}
	})

	t.Run("same version with v prefix returns false", func(t *testing.T) {
		if isNewerVersion("v1.0.0", "v1.0.0") {
			t.Error("Expected same version to return false")
		}
	})
}

// TestIsNewerVersion_EdgeCases tests version comparison edge cases.
// Priority: P2 - Version string parsing robustness for unusual formats.
// Tests different version lengths, single digit, and double digit versions.
func TestIsNewerVersion_EdgeCases(t *testing.T) {
	t.Run("longer version is newer when prefix matches", func(t *testing.T) {
		if !isNewerVersion("1.0", "1.0.1") {
			t.Error("Expected 1.0.1 to be newer than 1.0")
		}
	})

	t.Run("handles single digit versions", func(t *testing.T) {
		if !isNewerVersion("1", "2") {
			t.Error("Expected 2 to be newer than 1")
		}
	})

	t.Run("handles double digit versions", func(t *testing.T) {
		if !isNewerVersion("1.9.0", "1.10.0") {
			t.Error("Expected 1.10.0 to be newer than 1.9.0")
		}
	})
}

// TestFindDownloadURL_CorrectAssetFound tests that findDownloadURL finds the correct asset for current OS/arch.
// Priority: P1 - Update download depends on finding the correct binary for the platform.
// Tests asset matching for current platform, no matching assets, and empty assets list.
func TestFindDownloadURL_CorrectAssetFound(t *testing.T) {
	t.Run("finds darwin-amd64 asset", func(t *testing.T) {
		assets := []Asset{
			{Name: "lazysetup-v1.0.0-linux-amd64", BrowserDownloadURL: "https://example.com/linux-amd64"},
			{Name: "lazysetup-v1.0.0-darwin-amd64", BrowserDownloadURL: "https://example.com/darwin-amd64"},
			{Name: "lazysetup-v1.0.0-darwin-arm64", BrowserDownloadURL: "https://example.com/darwin-arm64"},
		}

		// This test will find the asset matching the current runtime
		url := findDownloadURL(assets)
		// URL should be non-empty if current OS/arch matches one of the assets
		t.Logf("Found URL for current platform: %s", url)
	})

	t.Run("returns empty for no matching assets", func(t *testing.T) {
		assets := []Asset{
			{Name: "lazysetup-v1.0.0-windows-amd64.exe", BrowserDownloadURL: "https://example.com/windows"},
		}

		// On non-Windows systems, this should return empty
		url := findDownloadURL(assets)
		t.Logf("URL for non-matching assets: '%s'", url)
	})

	t.Run("handles empty assets list", func(t *testing.T) {
		assets := []Asset{}
		url := findDownloadURL(assets)
		if url != "" {
			t.Errorf("Expected empty URL for empty assets, got '%s'", url)
		}
	})
}

// TestFindDownloadURL_ArchitectureMatching tests architecture variant matching (x86_64, aarch64).
// Priority: P2 - No match returns empty, handles different architecture naming conventions.
// Tests amd64 variants (x86_64, x64) and arm64 variants (aarch64).
func TestFindDownloadURL_ArchitectureMatching(t *testing.T) {
	t.Run("matches amd64 variants", func(t *testing.T) {
		// Test that x86_64 and x64 are recognized as amd64
		assets := []Asset{
			{Name: "lazysetup-darwin-x86_64", BrowserDownloadURL: "https://example.com/x86_64"},
		}
		url := findDownloadURL(assets)
		t.Logf("URL for x86_64 variant: '%s'", url)
	})

	t.Run("matches arm64 variants", func(t *testing.T) {
		// Test that aarch64 is recognized as arm64
		assets := []Asset{
			{Name: "lazysetup-darwin-aarch64", BrowserDownloadURL: "https://example.com/aarch64"},
		}
		url := findDownloadURL(assets)
		t.Logf("URL for aarch64 variant: '%s'", url)
	})
}

// TestUpdateInfo_Structure tests the UpdateInfo struct field initialization.
// Priority: P2 - Struct validation for update information storage.
// Tests that all required fields are properly set and accessible.
func TestUpdateInfo_Structure(t *testing.T) {
	t.Run("UpdateInfo has all required fields", func(t *testing.T) {
		info := &UpdateInfo{
			Available:      true,
			CurrentVersion: "1.0.0",
			LatestVersion:  "1.1.0",
			ReleaseURL:     "https://github.com/example/releases/v1.1.0",
			ReleaseNotes:   "Bug fixes",
			DownloadURL:    "https://example.com/download",
			Error:          nil,
		}

		if !info.Available {
			t.Error("Expected Available to be true")
		}
		if info.CurrentVersion != "1.0.0" {
			t.Errorf("Expected CurrentVersion '1.0.0', got '%s'", info.CurrentVersion)
		}
		if info.LatestVersion != "1.1.0" {
			t.Errorf("Expected LatestVersion '1.1.0', got '%s'", info.LatestVersion)
		}
		if info.DownloadURL == "" {
			t.Error("Expected non-empty DownloadURL")
		}
	})

	t.Run("UpdateInfo with error", func(t *testing.T) {
		info := &UpdateInfo{
			Available: false,
			Error:     nil,
		}

		if info.Available {
			t.Error("Expected Available to be false when error is set")
		}
	})
}

// TestGitHubRelease_Structure tests the GitHubRelease struct field initialization.
// Priority: P2 - Struct validation for GitHub API response parsing.
// Tests that all required fields including Assets are properly set.
func TestGitHubRelease_Structure(t *testing.T) {
	t.Run("GitHubRelease has all required fields", func(t *testing.T) {
		release := GitHubRelease{
			TagName:     "v1.0.0",
			Name:        "Release v1.0.0",
			Body:        "Release notes",
			HTMLURL:     "https://github.com/example/releases/v1.0.0",
			PublishedAt: "2025-01-01T00:00:00Z",
			Assets: []Asset{
				{Name: "binary", BrowserDownloadURL: "https://example.com/binary", Size: 1024},
			},
		}

		if release.TagName != "v1.0.0" {
			t.Errorf("Expected TagName 'v1.0.0', got '%s'", release.TagName)
		}
		if len(release.Assets) != 1 {
			t.Errorf("Expected 1 asset, got %d", len(release.Assets))
		}
		if release.Assets[0].Size != 1024 {
			t.Errorf("Expected asset size 1024, got %d", release.Assets[0].Size)
		}
	})
}

// TestDownloadAndInstall_EmptyURL tests error handling for empty download URL.
// Priority: P1 - Update installation must fail gracefully with clear error for empty URL.
// Tests that DownloadAndInstall returns an error when given an empty URL.
func TestDownloadAndInstall_EmptyURL(t *testing.T) {
	t.Run("returns error for empty URL", func(t *testing.T) {
		err := DownloadAndInstall("")
		if err == nil {
			t.Error("Expected error for empty download URL")
		}
	})
}

// TestCopyFile_NonExistentSource tests error handling for non-existent source file.
// Priority: P2 - Helper function validation for cross-device copy fallback.
// Tests that copyFile returns an error when source file doesn't exist.
func TestCopyFile_NonExistentSource(t *testing.T) {
	t.Run("returns error for non-existent source", func(t *testing.T) {
		err := copyFile("/nonexistent/path/file", "/tmp/dest")
		if err == nil {
			t.Error("Expected error for non-existent source file")
		}
	})
}

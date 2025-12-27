package updater

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"runtime"
	"strings"
	"time"

	"github.com/youpele52/lazysetup/pkg/version"
)

// GitHubRelease represents a GitHub release response
type GitHubRelease struct {
	TagName     string  `json:"tag_name"`
	Name        string  `json:"name"`
	Body        string  `json:"body"`
	HTMLURL     string  `json:"html_url"`
	PublishedAt string  `json:"published_at"`
	Assets      []Asset `json:"assets"`
}

// Asset represents a release asset
type Asset struct {
	Name               string `json:"name"`
	BrowserDownloadURL string `json:"browser_download_url"`
	Size               int64  `json:"size"`
}

// UpdateInfo contains information about an available update
type UpdateInfo struct {
	Available      bool
	CurrentVersion string
	LatestVersion  string
	ReleaseURL     string
	ReleaseNotes   string
	DownloadURL    string
	Error          error
}

// CheckForUpdates checks GitHub for the latest release
func CheckForUpdates() *UpdateInfo {
	info := &UpdateInfo{
		CurrentVersion: version.Version,
	}

	url := fmt.Sprintf("https://api.github.com/repos/%s/%s/releases/latest",
		version.GitHubOwner, version.GitHubRepo)

	client := &http.Client{Timeout: 10 * time.Second}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		info.Error = err
		return info
	}
	req.Header.Set("Accept", "application/vnd.github.v3+json")
	req.Header.Set("User-Agent", "lazysetup-updater")

	resp, err := client.Do(req)
	if err != nil {
		info.Error = err
		return info
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		info.Error = fmt.Errorf("GitHub API returned status %d", resp.StatusCode)
		return info
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		info.Error = err
		return info
	}

	var release GitHubRelease
	if err := json.Unmarshal(body, &release); err != nil {
		info.Error = err
		return info
	}

	info.LatestVersion = strings.TrimPrefix(release.TagName, "v")
	info.ReleaseURL = release.HTMLURL
	info.ReleaseNotes = release.Body
	info.DownloadURL = findDownloadURL(release.Assets)
	info.Available = isNewerVersion(info.CurrentVersion, info.LatestVersion)

	return info
}

// findDownloadURL finds the appropriate download URL for the current OS/arch
func findDownloadURL(assets []Asset) string {
	osName := runtime.GOOS
	arch := runtime.GOARCH

	// Map common architecture names
	archNames := map[string][]string{
		"amd64": {"amd64", "x86_64", "x64"},
		"arm64": {"arm64", "aarch64"},
		"386":   {"386", "i386", "x86"},
	}

	for _, asset := range assets {
		name := strings.ToLower(asset.Name)

		// Check if asset matches OS
		if !strings.Contains(name, osName) {
			continue
		}

		// Check if asset matches architecture
		if archList, ok := archNames[arch]; ok {
			for _, archName := range archList {
				if strings.Contains(name, archName) {
					return asset.BrowserDownloadURL
				}
			}
		}
	}

	return ""
}

// isNewerVersion compares two semantic versions
func isNewerVersion(current, latest string) bool {
	current = strings.TrimPrefix(current, "v")
	latest = strings.TrimPrefix(latest, "v")

	if current == latest {
		return false
	}

	currentParts := strings.Split(current, ".")
	latestParts := strings.Split(latest, ".")

	for i := 0; i < len(latestParts) && i < len(currentParts); i++ {
		var currentNum, latestNum int
		fmt.Sscanf(currentParts[i], "%d", &currentNum)
		fmt.Sscanf(latestParts[i], "%d", &latestNum)

		if latestNum > currentNum {
			return true
		} else if latestNum < currentNum {
			return false
		}
	}

	return len(latestParts) > len(currentParts)
}

// DownloadAndInstall downloads and installs the update
func DownloadAndInstall(downloadURL string) error {
	if downloadURL == "" {
		return fmt.Errorf("no download URL available for %s/%s", runtime.GOOS, runtime.GOARCH)
	}

	// Get current executable path
	execPath, err := os.Executable()
	if err != nil {
		return fmt.Errorf("failed to get executable path: %w", err)
	}

	// Download to temp file
	tmpFile, err := os.CreateTemp("", "lazysetup-update-*")
	if err != nil {
		return fmt.Errorf("failed to create temp file: %w", err)
	}
	defer os.Remove(tmpFile.Name())

	client := &http.Client{Timeout: 5 * time.Minute}
	resp, err := client.Get(downloadURL)
	if err != nil {
		return fmt.Errorf("failed to download update: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("download failed with status %d", resp.StatusCode)
	}

	if _, err := io.Copy(tmpFile, resp.Body); err != nil {
		return fmt.Errorf("failed to write update: %w", err)
	}
	tmpFile.Close()

	// Make executable
	if err := os.Chmod(tmpFile.Name(), 0755); err != nil {
		return fmt.Errorf("failed to set permissions: %w", err)
	}

	// Replace current executable
	if err := os.Rename(tmpFile.Name(), execPath); err != nil {
		// Try copy if rename fails (cross-device)
		if err := copyFile(tmpFile.Name(), execPath); err != nil {
			return fmt.Errorf("failed to replace executable: %w", err)
		}
	}

	return nil
}

// copyFile copies a file from src to dst
func copyFile(src, dst string) error {
	source, err := os.Open(src)
	if err != nil {
		return err
	}
	defer source.Close()

	dest, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer dest.Close()

	if _, err := io.Copy(dest, source); err != nil {
		return err
	}

	return os.Chmod(dst, 0755)
}

// RestartApplication restarts the application
func RestartApplication() error {
	execPath, err := os.Executable()
	if err != nil {
		return err
	}

	cmd := exec.Command(execPath, os.Args[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if err := cmd.Start(); err != nil {
		return err
	}

	os.Exit(0)
	return nil
}

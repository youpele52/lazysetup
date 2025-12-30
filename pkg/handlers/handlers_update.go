package handlers

import (
	"fmt"
	"time"

	"github.com/youpele52/lazysetup/pkg/models"
	"github.com/youpele52/lazysetup/pkg/updater"
	"github.com/youpele52/lazysetup/pkg/version"
)

// PerformUpdate downloads and installs the available update
func PerformUpdate(state *models.State) {
	if !state.UpdateAvailable || state.UpdateDownloadURL == "" {
		state.UpdateMessage = "No update available"
		return
	}

	// Show update prompt dialog
	state.UpdateMessage = fmt.Sprintf("Press 'u' to update from v%s to %s", version.Version, state.UpdateVersion)
	state.UpdateMessageTime = time.Now().Unix()
}

// ExecuteUpdate performs the actual download and installation
func ExecuteUpdate(state *models.State) {
	if !state.UpdateAvailable || state.UpdateDownloadURL == "" {
		state.UpdateMessage = "No update available"
		return
	}

	// Show downloading status
	state.UpdateMessage = fmt.Sprintf("%s Downloading update to %s...", "⠋", state.UpdateVersion)
	state.UpdateMessageTime = time.Now().Unix()

	err := updater.DownloadAndInstall(state.UpdateDownloadURL)
	if err != nil {
		state.UpdateMessage = fmt.Sprintf("✗ Update failed: %s", err.Error())
		state.UpdateMessageTime = time.Now().Unix()
		return
	}

	// Show success message
	state.UpdateMessage = fmt.Sprintf("✓ Update successful! Restarting with %s...", state.UpdateVersion)
	state.UpdateMessageTime = time.Now().Unix()
	state.UpdateAvailable = false
}

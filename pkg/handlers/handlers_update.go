package handlers

import (
	"fmt"

	"github.com/youpele52/lazysetup/pkg/constants"
	"github.com/youpele52/lazysetup/pkg/models"
	"github.com/youpele52/lazysetup/pkg/updater"
)

// PerformUpdate downloads and installs the available update
func PerformUpdate(state *models.State) {
	if !state.UpdateAvailable || state.UpdateDownloadURL == "" {
		state.UpdateMessage = "No update available"
		return
	}

	state.UpdateMessage = constants.UpdateDownloading

	err := updater.DownloadAndInstall(state.UpdateDownloadURL)
	if err != nil {
		state.UpdateMessage = fmt.Sprintf(constants.UpdateFailed, err.Error())
		return
	}

	state.UpdateMessage = constants.UpdateSuccess
	state.UpdateAvailable = false
}

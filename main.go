package main

import (
	"fmt"
	"log"
	"time"

	"github.com/jesseduffield/gocui"
	"github.com/youpele52/lazysetup/pkg/constants"
	"github.com/youpele52/lazysetup/pkg/models"
	"github.com/youpele52/lazysetup/pkg/ui"
	"github.com/youpele52/lazysetup/pkg/updater"
)

func main() {
	state := models.NewState()

	g := gocui.NewGui()
	if err := g.Init(); err != nil {
		log.Panicln(err)
	}
	defer g.Close()

	g.SetLayout(ui.Layout(state))
	ui.SetupKeybindings(g, state)

	// Start UI refresh goroutine for animations and status updates
	go refreshUI(g, state)

	// Check for updates on startup (in background)
	go checkForUpdates(state)

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

// checkForUpdates checks for available updates on startup and automatically installs them
func checkForUpdates(state *models.State) {
	info := updater.CheckForUpdates()
	if info.Error != nil {
		return
	}
	if info.Available {
		// Automatically download and install the update
		state.UpdateMessage = fmt.Sprintf(constants.UpdateDownloading)
		state.UpdateMessageTime = time.Now().Unix()

		err := updater.DownloadAndInstall(info.DownloadURL)
		if err != nil {
			// If update fails, show error message
			state.UpdateMessage = fmt.Sprintf(constants.UpdateFailed, err.Error())
			state.UpdateMessageTime = time.Now().Unix()
			return
		}

		// Update successful, restart the application
		state.UpdateMessage = constants.UpdateSuccess
		state.UpdateMessageTime = time.Now().Unix()
		time.Sleep(1 * time.Second) // Brief delay to show success message
		updater.RestartApplication()
	}
}

// refreshUI periodically updates the UI to show spinner animations and status changes
// This is necessary because goroutines updating state need to trigger UI redraws
func refreshUI(g *gocui.Gui, state *models.State) {
	ticker := time.NewTicker(100 * time.Millisecond)
	defer ticker.Stop()

	for range ticker.C {
		g.Execute(func(g *gocui.Gui) error {
			// Force layout refresh to pick up state changes
			return nil
		})
	}
}

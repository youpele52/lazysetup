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

// checkForUpdates checks for available updates on startup
func checkForUpdates(state *models.State) {
	info := updater.CheckForUpdates()
	if info.Error != nil {
		return
	}
	if info.Available {
		state.UpdateAvailable = true
		state.UpdateVersion = info.LatestVersion
		state.UpdateDownloadURL = info.DownloadURL
		state.UpdateMessage = fmt.Sprintf(constants.UpdateAvailable, info.CurrentVersion, info.LatestVersion)
		state.UpdateMessageTime = time.Now().Unix()
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

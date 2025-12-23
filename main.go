package main

import (
	"log"
	"time"

	"github.com/jesseduffield/gocui"
	"github.com/youpele52/lazysetup/pkg/models"
	"github.com/youpele52/lazysetup/pkg/ui"
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

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
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

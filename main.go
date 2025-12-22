package main

import (
	"log"

	"github.com/jesseduffield/gocui"
	"github.com/youpele52/lazysetup/pkg/handlers"
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

	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, handlers.Quit); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("menu", gocui.KeyArrowUp, gocui.ModNone, handlers.CursorUp(state)); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("menu", gocui.KeyArrowDown, gocui.ModNone, handlers.CursorDown(state)); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("menu", gocui.KeyEnter, gocui.ModNone, handlers.SelectMethod(state)); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("menu", gocui.KeyEsc, gocui.ModNone, handlers.GoBack(state)); err != nil {
		log.Panicln(err)
	}

	if err := g.SetKeybinding("result", gocui.KeyEsc, gocui.ModNone, handlers.GoBack(state)); err != nil {
		log.Panicln(err)
	}

	if err := g.MainLoop(); err != nil && err != gocui.ErrQuit {
		log.Panicln(err)
	}
}

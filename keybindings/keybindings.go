package keybindings

import (
	"log"

	"github.com/jroimartin/gocui"
)

var gui *gocui.Gui

func Set(g *gocui.Gui) {
	gui = g
	addBinding(gocui.KeyCtrlC, quit)
	addBinding(gocui.KeyCtrlD, quit)
	addBinding('q', quit)
	addBinding('h', help)
	addBinding('j', cursorDown)
	addBinding('k', cursorUp)
	addBinding('a', upvote)
	addBinding('z', downvote)
	addBinding(gocui.KeyEnter, enter)
	addBinding('l', link)
	addBinding('f', front)
	addBinding('c', comments)
	addBinding('r', refresh)
	addBinding('i', info)
	addBinding('y', yank)
}

func addBinding(c interface{}, f gocui.KeybindingHandler) {
	if err := gui.SetKeybinding("", c, gocui.ModNone, f); err != nil {
		log.Panicln(err)
	}
}

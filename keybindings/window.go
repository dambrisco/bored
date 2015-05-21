package keybindings

import (
	"github.com/dambrisco/bored/layout"
	"github.com/dambrisco/bored/reddit"
	"github.com/jroimartin/gocui"
)

func help(g *gocui.Gui, v *gocui.View) error {
	layout.SetPage(layout.Help)
	layout.Clear(g, v)
	return nil
}

func front(g *gocui.Gui, v *gocui.View) error {
	layout.SetPage(layout.List)
	layout.Clear(g, v)
	return nil
}

func refresh(g *gocui.Gui, v *gocui.View) error {
	if layout.GetPage() == layout.List {
		layout.SetPage(layout.Empty)
		reddit.SetCurrentIndex(0)
		layout.Clear(g, v)
		reddit.Load()
		return front(g, v)
	} else if layout.GetPage() == layout.Comments || layout.GetPage() == layout.Help {
		layout.Clear(g, v)
	}
	return nil
}

func quit(g *gocui.Gui, _ *gocui.View) error {
	return gocui.Quit
}

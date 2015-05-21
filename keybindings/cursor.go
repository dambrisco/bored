package keybindings

import (
	"github.com/dambrisco/bored/layout"
	"github.com/dambrisco/bored/reddit"
	"github.com/jroimartin/gocui"
)

func cursorDown(g *gocui.Gui, v *gocui.View) error {
	if layout.GetPage() == layout.List {
		current := v.Name()
		next := 0
		views := layout.GetViews()
		for i, v := range views {
			if v.Name() == current {
				next = i + 1
			}
		}
		if next >= len(views) {
			next = len(views) - 1
		}
		reddit.SetCurrentIndex(next)
		return g.SetCurrentView(views[next].Name())
	} else if layout.GetPage() == layout.Comments || layout.GetPage() == layout.Help {
		x, y := v.Origin()
		v.SetOrigin(x, y+1)
	}
	return nil
}

func cursorUp(g *gocui.Gui, v *gocui.View) error {
	if layout.GetPage() == layout.List {
		current := v.Name()
		prev := 0
		views := layout.GetViews()
		for i, v := range views {
			if v.Name() == current {
				prev = i - 1
			}
		}
		if prev < 0 {
			prev = 0
		}
		reddit.SetCurrentIndex(prev)
		return g.SetCurrentView(views[prev].Name())
	} else if layout.GetPage() == layout.Comments || layout.GetPage() == layout.Help {
		x, y := v.Origin()
		v.SetOrigin(x, y-1)
	}
	return nil
}

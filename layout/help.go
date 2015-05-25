package layout

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/jroimartin/gocui"
)

func HelpLayout(g *gocui.Gui, allViews []*gocui.View) []*gocui.View {
	maxX, maxY := g.Size()
	if help, err := g.SetView("help", -1, 0, maxX, maxY); err != nil && help != nil {
		allViews = append(allViews, help)
		help.Frame = false
		help.Wrap = true
		h, err := ioutil.ReadFile("help.txt")
		l, err := ioutil.ReadFile("LICENSE")
		if err != nil {
			log.Panicln(err)
		}
		fmt.Fprintln(help, string(h))
		fmt.Fprint(help, string(l))
		g.SetCurrentView("help")
	}

	return allViews
}

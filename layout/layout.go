package layout

import (
	"fmt"

	"github.com/dambrisco/bored/reddit"
	"github.com/jroimartin/gocui"
)

type pagetype int

const (
	List pagetype = iota
	Comments
	Empty
	Help
)

var allViews []*gocui.View
var currentPageType pagetype

func Layout(g *gocui.Gui) error {
	maxX, count := g.Size()
	count = count - 1
	if title, err := g.SetView("title", -1, -1, maxX, 1); err != nil {
		allViews = append(allViews, title)
		title.Frame = false
		title.BgColor = gocui.ColorBlue
		title.FgColor = gocui.ColorWhite
		fmt.Fprintln(title, "Bored v0.0.1")
	}

	submissions := reddit.GetSubmissions()
	current := reddit.GetCurrentIndex()

	if currentPageType == List {
		allViews = PostsLayout(g, allViews, submissions, current)
	} else if currentPageType == Comments {
		allViews = CommentsLayout(g, allViews, submissions[current])
	} else if currentPageType == Help {
		allViews = HelpLayout(g, allViews)
	}

	return nil
}

func GetPage() pagetype {
	return currentPageType
}

func SetPage(page pagetype) {
	currentPageType = page
}

func Clear(g *gocui.Gui, v *gocui.View) {
	for _, w := range allViews {
		g.DeleteView(w.Name())
	}
	SetViews(nil)
	votes = nil
	allViews = nil
	g.Flush()
}

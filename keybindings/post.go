package keybindings

import (
	"fmt"
	"strings"

	"github.com/dambrisco/bored/layout"
	"github.com/dambrisco/bored/reddit"
	"github.com/jroimartin/gocui"
	"github.com/toqueteos/webbrowser"
)

func enter(g *gocui.Gui, v *gocui.View) error {
	submission := reddit.GetCurrentSubmission()
	webbrowser.Open("https://www.reddit.com/" + submission.Permalink)
	return nil
}

func link(g *gocui.Gui, v *gocui.View) error {
	submission := reddit.GetCurrentSubmission()
	webbrowser.Open(submission.URL)
	return nil
}

func comments(g *gocui.Gui, v *gocui.View) error {
	layout.SetPage(layout.Comments)
	layout.Clear(g, v)
	return nil
}

func info(g *gocui.Gui, v *gocui.View) error {
	if layout.GetPage() == layout.List {
		s := reddit.GetCurrentSubmission()
		b := v.Buffer()
		v.Clear()
		if strings.HasPrefix(b, "S") {
			fmt.Fprint(v, layout.BuildTitleTag(s))
		} else {
			fmt.Fprintf(v, "Score: %d | Subreddit: %s", s.Score, s.Subreddit)
		}
	}
	return nil
}

package layout

import (
	"fmt"

	"github.com/dambrisco/bored/reddit"
	"github.com/dambrisco/geddit"
	"github.com/jroimartin/gocui"
)

var views []*gocui.View
var votes []*gocui.View

func PostsLayout(g *gocui.Gui, allViews []*gocui.View, submissions []*geddit.Submission, currentIndex int) []*gocui.View {
	maxX, maxY := g.Size()
	i := 0
	for i < maxY-1 && i < len(submissions) {
		name := fmt.Sprintf("submission-%d", i)
		if v, err := g.SetView(fmt.Sprintf("vote-%d", i), -1, i, 1, i+2); err != nil && v != nil {
			allViews = append(allViews, v)
			votes = append(votes, v)
			v.Frame = false
			fmt.Fprint(v, "•")
		}
		if s, err := g.SetView(name, 1, i, maxX, i+2); err != nil && s != nil {
			if i == currentIndex {
				g.SetCurrentView(name)
			}
			views = append(views, s)
			s.Frame = false
			subm := submissions[i]
			reddit.SetAfter(subm.Title)
			fmt.Fprint(s, BuildTitleTag(subm))
			allViews = append(allViews, s)
		}
		i += 1
	}

	setColor(g.CurrentView())

	return allViews
}

func setColor(v *gocui.View) {
	for _, w := range views {
		w.FgColor = gocui.ColorDefault
	}
	v.FgColor = gocui.ColorBlue
}

func GetViews() []*gocui.View {
	return views
}

func SetViews(v []*gocui.View) {
	views = v
}

func GetVotes() []*gocui.View {
	return votes
}

func SetVotes(v []*gocui.View) {
	votes = v
}

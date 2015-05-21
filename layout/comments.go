package layout

import (
	"fmt"
	"html"
	"strings"

	"github.com/dambrisco/geddit"
	"github.com/dambrisco/prose"
	"github.com/jroimartin/gocui"
)

func CommentsLayout(g *gocui.Gui, allViews []*gocui.View, submission *geddit.Submission) []*gocui.View {
	maxX, maxY := g.Size()
	if vote, err := g.SetView("post-vote", -1, 0, 1, 2); err != nil && vote != nil {
		allViews = append(allViews, vote)
		vote.Frame = false
		fmt.Fprint(vote, "•")
	}
	if title, err := g.SetView("post-title", 1, 0, maxX, 2); err != nil && title != nil {
		allViews = append(allViews, title)
		title.Frame = false
		fmt.Fprint(title, BuildTitleTag(submission))
	}
	if meta, err := g.SetView("post-meta", -1, 2, maxX, 4); err != nil && meta != nil {
		allViews = append(allViews, meta)
		meta.Frame = false
		fmt.Fprintf(meta, "Score: %d | Subreddit: %s", submission.Score, submission.Subreddit)
	}
	if text, err := g.SetView("post-text", -1, 4, maxX, maxY); err != nil && text != nil {
		allViews = append(allViews, text)
		text.Frame = false
		text.Wrap = true
		fmt.Fprintf(text, "%s", prose.Wrap(strings.Replace(html.UnescapeString(submission.Selftext), "\n\n", "\n", -1), maxX))
		g.SetCurrentView("post-text")
	}

	return allViews
}

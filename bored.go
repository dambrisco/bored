package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/dambrisco/geddit"
	"github.com/jroimartin/gocui"
	"github.com/toqueteos/webbrowser"

	"github.com/dambrisco/bored/auth"
	"github.com/dambrisco/bored/flag-parser"
	"github.com/dambrisco/bored/layout"
	"github.com/dambrisco/bored/reddit"
)

var submissions []*geddit.Submission

var session *geddit.LoginSession
var opts geddit.ListingOptions
var count int
var subreddit string

func quit(g *gocui.Gui, _ *gocui.View) error {
	return gocui.Quit
}

func help(g *gocui.Gui, v *gocui.View) error {
	layout.SetPage(layout.Help)
	layout.Clear(g, v)
	return nil
}

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

func upvote(g *gocui.Gui, v *gocui.View) error {
	s := reddit.GetCurrentSubmission()
	currentIndex := reddit.GetCurrentIndex()
	if layout.GetPage() == layout.List {
		votes := layout.GetVotes()
		vote := votes[currentIndex]
		if vote.FgColor == gocui.ColorGreen {
			vote.FgColor = gocui.ColorDefault
			session.Vote(s, geddit.RemoveVote)
		} else {
			vote.FgColor = gocui.ColorGreen
			session.Vote(s, geddit.UpVote)
		}
	} else if layout.GetPage() == layout.Comments {
		vote, _ := g.View("post-vote")
		if vote.FgColor == gocui.ColorGreen {
			vote.FgColor = gocui.ColorDefault
			session.Vote(s, geddit.RemoveVote)
		} else {
			vote.FgColor = gocui.ColorGreen
			session.Vote(s, geddit.UpVote)
		}
	}
	return nil
}

func downvote(g *gocui.Gui, v *gocui.View) error {
	s := reddit.GetCurrentSubmission()
	currentIndex := reddit.GetCurrentIndex()
	if layout.GetPage() == layout.List {
		votes := layout.GetVotes()
		vote := votes[currentIndex]
		if vote.FgColor == gocui.ColorRed {
			vote.FgColor = gocui.ColorDefault
			session.Vote(s, geddit.RemoveVote)
		} else {
			vote.FgColor = gocui.ColorRed
			session.Vote(s, geddit.DownVote)
		}
	} else if layout.GetPage() == layout.Comments {
		vote, _ := g.View("post-vote")
		if vote.FgColor == gocui.ColorRed {
			vote.FgColor = gocui.ColorDefault
			session.Vote(s, geddit.RemoveVote)
		} else {
			vote.FgColor = gocui.ColorRed
			session.Vote(s, geddit.DownVote)
		}
	}
	return nil
}

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

func front(g *gocui.Gui, v *gocui.View) error {
	layout.SetPage(layout.List)
	layout.Clear(g, v)
	return nil
}

func comments(g *gocui.Gui, v *gocui.View) error {
	layout.SetPage(layout.Comments)
	layout.Clear(g, v)
	return nil
}

func refresh(g *gocui.Gui, v *gocui.View) error {
	if layout.GetPage() == layout.List {
		layout.SetPage(layout.Empty)
		reddit.SetCurrentIndex(0)
		layout.Clear(g, v)
		submissions = reddit.Load(count, subreddit, session)
		return front(g, v)
	} else if layout.GetPage() == layout.Comments || layout.GetPage() == layout.Help {
		layout.Clear(g, v)
	}
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

func setKeybinds(g *gocui.Gui) {
	if err := g.SetKeybinding("", gocui.KeyCtrlC, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("", gocui.KeyCtrlD, gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("", 'q', gocui.ModNone, quit); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("", 'h', gocui.ModNone, help); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("", 'j', gocui.ModNone, cursorDown); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("", 'k', gocui.ModNone, cursorUp); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("", 'a', gocui.ModNone, upvote); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("", 'z', gocui.ModNone, downvote); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("", gocui.KeyEnter, gocui.ModNone, enter); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("", 'l', gocui.ModNone, link); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("", 'f', gocui.ModNone, front); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("", 'c', gocui.ModNone, comments); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("", 'r', gocui.ModNone, refresh); err != nil {
		log.Panicln(err)
	}
	if err := g.SetKeybinding("", 'i', gocui.ModNone, info); err != nil {
		log.Panicln(err)
	}
}

func main() {
	var err error
	username, password, s := flagParser.Parse()
	subreddit = s
	session = auth.Login(username, password)
	reddit.SetAfter("")

	g := gocui.NewGui()
	if err := g.Init(); err != nil {
		log.Panicln(err)
	}
	_, count = g.Size()
	count = count - 1
	submissions = reddit.Load(count, subreddit, session)
	layout.SetPage(layout.List)
	g.BgColor = gocui.ColorDefault
	g.FgColor = gocui.ColorDefault
	defer g.Close()
	g.SetLayout(layout.Layout)
	setKeybinds(g)

	err = g.MainLoop()
	if err != nil && err != gocui.Quit {
		log.Panicln(err)
	}
}

package main

import (
	"log"

	"github.com/dambrisco/geddit"
	"github.com/jroimartin/gocui"

	"github.com/dambrisco/bored/auth"
	"github.com/dambrisco/bored/flag-parser"
	"github.com/dambrisco/bored/keybindings"
	"github.com/dambrisco/bored/layout"
	"github.com/dambrisco/bored/reddit"
)

var submissions []*geddit.Submission

var session *geddit.LoginSession
var opts geddit.ListingOptions

func main() {
	var err error
	username, password, subreddit := flagParser.Parse()
	session = auth.Login(username, password)
	reddit.SetAfter("")

	g := gocui.NewGui()
	if err := g.Init(); err != nil {
		log.Panicln(err)
	}

	_, count := g.Size()

	reddit.SetOptions(count-1, subreddit, session)

	layout.SetPage(layout.List)
	g.BgColor = gocui.ColorDefault
	g.FgColor = gocui.ColorDefault
	defer g.Close()
	g.SetLayout(layout.Layout)
	keybindings.Set(g)

	err = g.MainLoop()
	if err != nil && err != gocui.Quit {
		log.Panicln(err)
	}
}

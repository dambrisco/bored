package main

import (
	"fmt"
	"log"

	"github.com/dambrisco/geddit"
	"github.com/jroimartin/gocui"

	"github.com/dambrisco/bored/auth"
	"github.com/dambrisco/bored/crypto"
	"github.com/dambrisco/bored/flag-parser"
	"github.com/dambrisco/bored/keybindings"
	"github.com/dambrisco/bored/layout"
	"github.com/dambrisco/bored/reddit"
)

var submissions []*geddit.Submission

var session *geddit.LoginSession
var opts geddit.ListingOptions

func main() {
	username, password, subreddit, encrypted, onlyEncrypt := flagParser.Parse()
	switch {
	case onlyEncrypt:
		encrypt(username, password)
	default:
		bored(username, password, subreddit, encrypted)
	}
}

func bored(username, password, subreddit string, encrypted bool) {
	session = auth.Login(username, password, encrypted)
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

	if err := g.MainLoop(); err != nil && err != gocui.Quit {
		log.Panicln(err)
	}
}

func encrypt(username, password string) {
	if username != "" {
		username = crypto.Encrypt(username)
		fmt.Println(username)
	}
	if password != "" {
		password = crypto.Encrypt(password)
		fmt.Println(password)
	}
}

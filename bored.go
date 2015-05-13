package main

import (
	"flag"
	"fmt"
	"github.com/dambrisco/geddit"
	"github.com/dambrisco/prose"
	"github.com/jroimartin/gocui"
	"github.com/toqueteos/webbrowser"
	"html"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type pagetype int

const (
	List pagetype = iota
	Comments
	Empty
	Help
)

var submissions []*geddit.Submission
var views []*gocui.View
var votes []*gocui.View
var allViews []*gocui.View
var session *geddit.LoginSession
var opts geddit.ListingOptions
var after string
var count int
var subreddit string
var currentIndex int
var currentPageType pagetype

func getCredentials() (string, string) {
	username := flag.String("u", "", "Username")
	password := flag.String("p", "", "Password")
	flag.StringVar(&subreddit, "s", "", "Subreddit")
	flag.Parse()
	if *username == "" {
		*username = os.Getenv("BORED_USERNAME")
	}
	if *password == "" {
		*password = os.Getenv("BORED_PASSWORD")
	}
	if subreddit == "" {
		subreddit = os.Getenv("BORED_SUBREDDIT")
	}
	if *username == "" || *password == "" {
		log.Panicln("bored requires a username and password")
		os.Exit(2)
	}

	return *username, *password
}

func login() *geddit.LoginSession {
	username, password := getCredentials()
	session, _ := geddit.NewLoginSession(username, password, "geddit")
	return session
}

func load(limit int) []*geddit.Submission {
	opts = geddit.ListingOptions{
		Limit: limit,
		After: after,
	}
	if subreddit == "" {
		submissions, _ = session.Frontpage(geddit.DefaultPopularity, opts)
	} else {
		submissions, _ = session.SubredditSubmissions(subreddit, geddit.DefaultPopularity, opts)
	}
	return submissions
}

func layoutList(g *gocui.Gui) {
	maxX, maxY := g.Size()
	i := 0
	y := 0
	for y < maxY-1 && i < len(submissions) {
		name := fmt.Sprintf("submission-%d", i)
		if v, err := g.SetView(fmt.Sprintf("vote-%d", i), -1, y, 1, y+2); err != nil && v != nil {
			allViews = append(allViews, v)
			votes = append(votes, v)
			v.Frame = false
			fmt.Fprint(v, "•")
		}
		if s, err := g.SetView(name, 1, y, maxX, y+2); err != nil && s != nil {
			if i == currentIndex {
				g.SetCurrentView(name)
			}
			views = append(views, s)
			s.Frame = false
			subm := submissions[i]
			after = subm.Title
			fmt.Fprint(s, buildTitleTag(subm))
			allViews = append(allViews, s)
		}
		y += 1
		i += 1
	}
	setColor(g.CurrentView())
}

func layoutComments(g *gocui.Gui) {
	s := submissions[currentIndex]
	maxX, maxY := g.Size()
	if vote, err := g.SetView("post-vote", -1, 0, 1, 2); err != nil && vote != nil {
		allViews = append(allViews, vote)
		vote.Frame = false
		fmt.Fprint(vote, "•")
	}
	if title, err := g.SetView("post-title", 1, 0, maxX, 2); err != nil && title != nil {
		allViews = append(allViews, title)
		title.Frame = false
		fmt.Fprint(title, buildTitleTag(s))
	}
	if meta, err := g.SetView("post-meta", -1, 2, maxX, 4); err != nil && meta != nil {
		allViews = append(allViews, meta)
		meta.Frame = false
		fmt.Fprintf(meta, "Score: %d | Subreddit: %s", s.Score, s.Subreddit)
	}
	if text, err := g.SetView("post-text", -1, 4, maxX, maxY); err != nil && text != nil {
		allViews = append(allViews, text)
		text.Frame = false
		text.Wrap = true
		fmt.Fprintf(text, "%s", prose.Wrap(strings.Replace(html.UnescapeString(s.Selftext), "\n\n", "\n", -1), maxX))
		g.SetCurrentView("post-text")
	}
}

func layoutHelp(g *gocui.Gui) {
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
}

func layout(g *gocui.Gui) error {
	maxX, count := g.Size()
	count = count - 1
	if title, err := g.SetView("title", -1, -1, maxX, 1); err != nil {
		allViews = append(allViews, title)
		title.Frame = false
		title.BgColor = gocui.ColorBlue
		title.FgColor = gocui.ColorWhite
		fmt.Fprintln(title, "Bored v0.0.1")
	}
	if currentPageType == List {
		layoutList(g)
	} else if currentPageType == Comments {
		layoutComments(g)
	} else if currentPageType == Help {
		layoutHelp(g)
	}
	return nil
}

func quit(g *gocui.Gui, _ *gocui.View) error {
	return gocui.Quit
}

func help(g *gocui.Gui, v *gocui.View) error {
	currentPageType = Help
	clearViews(g, v)
	return nil
}

func cursorDown(g *gocui.Gui, v *gocui.View) error {
	if currentPageType == List {
		current := v.Name()
		next := 0
		for i, v := range views {
			if v.Name() == current {
				next = i + 1
			}
		}
		if next >= len(views) {
			next = len(views) - 1
		}
		currentIndex = next
		return g.SetCurrentView(views[next].Name())
	} else if currentPageType == Comments || currentPageType == Help {
		x, y := v.Origin()
		v.SetOrigin(x, y+1)
	}
	return nil
}

func cursorUp(g *gocui.Gui, v *gocui.View) error {
	if currentPageType == List {
		current := v.Name()
		prev := 0
		for i, v := range views {
			if v.Name() == current {
				prev = i - 1
			}
		}
		if prev < 0 {
			prev = 0
		}
		currentIndex = prev
		return g.SetCurrentView(views[prev].Name())
	} else if currentPageType == Comments || currentPageType == Help {
		x, y := v.Origin()
		v.SetOrigin(x, y-1)
	}
	return nil
}

func upvote(g *gocui.Gui, v *gocui.View) error {
	s := submissions[currentIndex]
	if currentPageType == List {
		vote := votes[currentIndex]
		if vote.FgColor == gocui.ColorGreen {
			vote.FgColor = gocui.ColorDefault
			session.Vote(s, geddit.RemoveVote)
		} else {
			vote.FgColor = gocui.ColorGreen
			session.Vote(s, geddit.UpVote)
		}
	} else if currentPageType == Comments {
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
	s := submissions[currentIndex]
	if currentPageType == List {
		vote := votes[currentIndex]
		if vote.FgColor == gocui.ColorRed {
			vote.FgColor = gocui.ColorDefault
			session.Vote(s, geddit.RemoveVote)
		} else {
			vote.FgColor = gocui.ColorRed
			session.Vote(s, geddit.DownVote)
		}
	} else if currentPageType == Comments {
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
	webbrowser.Open("https://www.reddit.com/" + submissions[currentIndex].Permalink)
	return nil
}

func link(g *gocui.Gui, v *gocui.View) error {
	webbrowser.Open(submissions[currentIndex].URL)
	return nil
}

func front(g *gocui.Gui, v *gocui.View) error {
	currentPageType = List
	clearViews(g, v)
	return nil
}

func comments(g *gocui.Gui, v *gocui.View) error {
	currentPageType = Comments
	clearViews(g, v)
	return nil
}

func refresh(g *gocui.Gui, v *gocui.View) error {
	if currentPageType == List {
		currentPageType = Empty
		currentIndex = 0
		clearViews(g, v)
		load(count)
		return front(g, v)
	} else if currentPageType == Comments || currentPageType == Help {
		clearViews(g, v)
	}
	return nil
}

func info(g *gocui.Gui, v *gocui.View) error {
	if currentPageType == List {
		s := submissions[currentIndex]
		b := v.Buffer()
		v.Clear()
		if strings.HasPrefix(b, "S") {
			fmt.Fprint(v, buildTitleTag(s))
		} else {
			fmt.Fprintf(v, "Score: %d | Subreddit: %s", s.Score, s.Subreddit)
		}
	}
	return nil
}

func clearViews(g *gocui.Gui, v *gocui.View) {
	for _, w := range allViews {
		g.DeleteView(w.Name())
	}
	views = nil
	votes = nil
	allViews = nil
	g.Flush()
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

func setColor(v *gocui.View) {
	for _, w := range views {
		w.FgColor = gocui.ColorDefault
	}
	v.FgColor = gocui.ColorBlue
}

func buildTitleTag(s *geddit.Submission) string {
	title := html.UnescapeString(s.Title)
	tag := "LINK"
	if s.IsSelf {
		tag = "SELF"
	}
	if s.IsNSFW {
		tag = "NSFW+" + tag
	}
	if s.IsSticky {
		tag = "STICKY+" + tag
	}
	return fmt.Sprintf("[%s] %s", tag, title)
}

func main() {
	var err error
	session = login()
	after = ""

	g := gocui.NewGui()
	if err := g.Init(); err != nil {
		log.Panicln(err)
	}
	_, count = g.Size()
	count = count - 1
	submissions = load(count)
	currentPageType = List
	g.BgColor = gocui.ColorDefault
	g.FgColor = gocui.ColorDefault
	defer g.Close()
	g.SetLayout(layout)
	setKeybinds(g)

	err = g.MainLoop()
	if err != nil && err != gocui.Quit {
		log.Panicln(err)
	}
}

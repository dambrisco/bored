package keybindings

import (
	"github.com/dambrisco/bored/auth"
	"github.com/dambrisco/bored/layout"
	"github.com/dambrisco/bored/reddit"
	"github.com/dambrisco/geddit"
	"github.com/jroimartin/gocui"
)

func upvote(g *gocui.Gui, v *gocui.View) error {
	session := auth.GetSession()

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
	session := auth.GetSession()

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

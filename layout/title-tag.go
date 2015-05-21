package layout

import (
	"fmt"
	"html"

	"github.com/dambrisco/geddit"
)

func BuildTitleTag(s *geddit.Submission) string {
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

package reddit

import "github.com/dambrisco/geddit"

var currentIndex int
var submissions []*geddit.Submission
var opts geddit.ListingOptions

var limit int
var after string
var session *geddit.LoginSession
var subreddit string

func Load() {
	opts = geddit.ListingOptions{
		Limit: limit,
		After: after,
	}

	if subreddit == "" {
		submissions, _ = session.Frontpage(geddit.DefaultPopularity, opts)
	} else {
		submissions, _ = session.SubredditSubmissions(subreddit, geddit.DefaultPopularity, opts)
	}
}

func SetOptions(count int, sub string, sess *geddit.LoginSession) {
	limit = count
	subreddit = sub
	session = sess
	Load()
}

func SetAfter(a string) {
	after = a
}

func GetSubmissions() []*geddit.Submission {
	return submissions
}

func GetCurrentSubmission() *geddit.Submission {
	return submissions[currentIndex]
}

func GetCurrentIndex() int {
	return currentIndex
}

func SetCurrentIndex(idx int) {
	currentIndex = idx
}

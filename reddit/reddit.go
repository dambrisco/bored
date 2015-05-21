package reddit

import "github.com/dambrisco/geddit"

var currentIndex int
var loadedSubmissions []*geddit.Submission
var opts geddit.ListingOptions
var after string

func Load(limit int, subreddit string, session *geddit.LoginSession) []*geddit.Submission {
	var submissions []*geddit.Submission

	opts = geddit.ListingOptions{
		Limit: limit,
		After: after,
	}

	if subreddit == "" {
		submissions, _ = session.Frontpage(geddit.DefaultPopularity, opts)
	} else {
		submissions, _ = session.SubredditSubmissions(subreddit, geddit.DefaultPopularity, opts)
	}

	loadedSubmissions = submissions
	return submissions
}

func SetAfter(a string) {
	after = a
}

func GetSubmissions() []*geddit.Submission {
	return loadedSubmissions
}

func GetCurrentSubmission() *geddit.Submission {
	return loadedSubmissions[currentIndex]
}

func GetCurrentIndex() int {
	return currentIndex
}

func SetCurrentIndex(idx int) {
	currentIndex = idx
}

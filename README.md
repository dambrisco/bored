#Bored ![Travis CI status](https://travis-ci.org/dambrisco/bored.svg?branch=master "Oh god why aren't there any tests")
Welcome to Bored, the world's premier termbox-based Reddit CLI. It has *numerous* features including keybinds, opening links in your browser, and even the ability to upvote and downvote links!

Speaking of keybinds:

    h - help screen
    q - quit
    j - navigate/scroll down
    k - navigate/scroll up
    i - show submission info (only on list view)
    r - refresh
    f - front page
    c - comments (currently just a self-text view)
    l - open link url
    y - copy link url
    enter - open reddit permalink

Params (these override the environment variables below):

    -u username (MUST NOT be encrypted)
    -p password (MUST NOT be encrypted)
    -s subreddit (using this loads submissions from ONLY the given subreddit)
    -e encrypt-only
        This param will encrypt the given username and/or password and exit,
        allowing you to set up your environment variables with the printed
        values

Environment variables to make your life easier:

    BORED_USERNAME - same as specifying -u (MUST be encrypted)
    BORED_PASSWORD - same as specifying -p (MUST be encrypted)
    BORED_SUBREDDIT - same as specifying -s

To install, clone the repo and run `go install .`

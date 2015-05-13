#Bored
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
    enter - open reddit permalink

Params:

    -u username
    -p password
    -s subreddit (using this loads submissions from ONLY the given subreddit)

Environment variables to make your life easier:

    BORED_USERNAME - same as specifying -u
    BORED_PASSWORD - same as specifying -p
    BORED_SUBREDDIT - same as specifying -s

To install, clone the repo and run `go install .`

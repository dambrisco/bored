package flagParser

import (
	"flag"
	"log"
	"os"
)

func Parse() (string, string, string) {
	username := flag.String("u", "", "Username")
	password := flag.String("p", "", "Password")
	subreddit := flag.String("s", "", "Subreddit")

	flag.Parse()

	if *username == "" {
		*username = os.Getenv("BORED_USERNAME")
	}
	if *password == "" {
		*password = os.Getenv("BORED_PASSWORD")
	}
	if *subreddit == "" {
		*subreddit = os.Getenv("BORED_SUBREDDIT")
	}
	if *username == "" || *password == "" {
		log.Panicln("bored requires a username and password")
		os.Exit(2)
	}

	return *username, *password, *subreddit
}

package flagParser

import (
	"flag"
	"log"
	"os"
)

func Parse() (string, string, string, bool, bool) {
	username := flag.String("u", "", "Username")
	password := flag.String("p", "", "Password")
	subreddit := flag.String("s", "", "Subreddit")
	onlyEncrypt := flag.Bool("e", false, "Only perform encryption - will exit after printing the encrypted username and/or password provided")
	encrypted := false

	flag.Parse()

	if *username == "" {
		*username = os.Getenv("BORED_USERNAME")
		encrypted = true
	}
	if *password == "" {
		*password = os.Getenv("BORED_PASSWORD")
		encrypted = true
	}
	if *subreddit == "" {
		*subreddit = os.Getenv("BORED_SUBREDDIT")
	}
	switch {
	case *onlyEncrypt:
	default:
		if *username == "" || *password == "" {
			log.Panicln("bored requires a username and password")
			os.Exit(2)
		}
	}

	return *username, *password, *subreddit, encrypted, *onlyEncrypt
}

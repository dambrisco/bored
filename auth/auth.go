package auth

import "github.com/dambrisco/geddit"

func Login(username string, password string) *geddit.LoginSession {
	session, _ := geddit.NewLoginSession(username, password, "geddit")
	return session
}

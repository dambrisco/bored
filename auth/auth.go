package auth

import "github.com/dambrisco/geddit"

var session *geddit.LoginSession

func Login(username string, password string) *geddit.LoginSession {
	currentSession, _ := geddit.NewLoginSession(username, password, "geddit")
	return currentSession
}

func GetSession() *geddit.LoginSession {
	return session
}

package auth

import (
	"github.com/dambrisco/bored/crypto"
	"github.com/dambrisco/geddit"
)

var session *geddit.LoginSession

func Login(username, password string, encrypted bool) *geddit.LoginSession {
	if encrypted {
		username = crypto.Decrypt(username)
		password = crypto.Decrypt(password)
	}
	currentSession, _ := geddit.NewLoginSession(username, password, "geddit")
	return currentSession
}

func GetSession() *geddit.LoginSession {
	return session
}

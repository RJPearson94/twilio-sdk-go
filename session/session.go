package session

import (
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

type Session struct {
	*credentials.Credentials
}

func New(creds *credentials.Credentials) *Session {
	return &Session{
		Credentials: creds,
	}
}

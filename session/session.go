package session

import (
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

// Session represents a session object that can be used to make requests against the Twilio APIs
type Session struct {
	*credentials.Credentials
}

// NewWithCredentials creates a new session instance using the credentials supplied
func New(creds *credentials.Credentials) *Session {
	return &Session{
		Credentials: creds,
	}
}

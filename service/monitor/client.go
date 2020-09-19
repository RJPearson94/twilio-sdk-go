package monitor

import (
	v1 "github.com/RJPearson94/twilio-sdk-go/service/monitor/v1"
	"github.com/RJPearson94/twilio-sdk-go/session"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

// Monitor client is used to manage versioned resources for Twilio Monitor
type Monitor struct {
	V1 *v1.Monitor
}

// New creates a new instance of the client using session data
func New(sess *session.Session) *Monitor {
	return &Monitor{
		V1: v1.New(sess),
	}
}

// NewWithCredentials creates a new instance of the client with credentials
func NewWithCredentials(creds *credentials.Credentials) *Monitor {
	return New(session.New(creds))
}

package sync

import (
	v1 "github.com/RJPearson94/twilio-sdk-go/service/sync/v1"
	"github.com/RJPearson94/twilio-sdk-go/session"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

// Sync client is used to manage versioned resources for Twilio Sync
// See https://www.twilio.com/docs/sync for more details on the API
// See https://www.twilio.com/sync for more details on the product
type Sync struct {
	V1 *v1.Sync
}

// New creates a new instance of the client using session data
func New(sess *session.Session) *Sync {
	return &Sync{
		V1: v1.New(sess),
	}
}

// NewWithCredentials creates a new instance of the client with credentials
func NewWithCredentials(creds *credentials.Credentials) *Sync {
	return New(session.New(creds))
}

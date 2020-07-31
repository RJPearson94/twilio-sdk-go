package messaging

import (
	v1 "github.com/RJPearson94/twilio-sdk-go/service/messaging/v1"
	"github.com/RJPearson94/twilio-sdk-go/session"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

// Messaging client is used to manage versioned resources for Twilio Messaging
// See https://www.twilio.com/docs/messaging for more details on the API
// See https://www.twilio.com/messaging for more details on the product
type Messaging struct {
	V1 *v1.Messaging
}

// New creates a new instance of the client using session data
func New(sess *session.Session) *Messaging {
	return &Messaging{
		V1: v1.New(sess),
	}
}

// NewWithCredentials creates a new instance of the client with credentials
func NewWithCredentials(creds *credentials.Credentials) *Messaging {
	return New(session.New(creds))
}

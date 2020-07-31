package conversations

import (
	v1 "github.com/RJPearson94/twilio-sdk-go/service/conversations/v1"
	"github.com/RJPearson94/twilio-sdk-go/session"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

// Conversations client is used to manage versioned resources for Twilio Conversations
// See https://www.twilio.com/docs/conversations for more details on the API
// See https://www.twilio.com/conversations for more details on the product
type Conversations struct {
	V1 *v1.Conversations
}

// New creates a new instance of the client using session data
func New(sess *session.Session) *Conversations {
	return &Conversations{
		V1: v1.New(sess),
	}
}

// NewWithCredentials creates a new instance of the client with credentials
func NewWithCredentials(creds *credentials.Credentials) *Conversations {
	return New(session.New(creds))
}

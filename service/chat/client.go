package chat

import (
	v2 "github.com/RJPearson94/twilio-sdk-go/service/chat/v2"
	"github.com/RJPearson94/twilio-sdk-go/session"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

// Chat client is used to manage versioned resources for Programmable Chat
// See https://www.twilio.com/docs/chat for more details on the API
// See https://www.twilio.com/chat for more details on the product
type Chat struct {
	V2 *v2.Chat
}

// New creates a new instance of the client using session data
func New(sess *session.Session) *Chat {
	return &Chat{
		V2: v2.New(sess),
	}
}

// NewWithCredentials creates a new instance of the client with credentials
func NewWithCredentials(creds *credentials.Credentials) *Chat {
	return New(session.New(creds))
}

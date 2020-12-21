package conversations

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	v1 "github.com/RJPearson94/twilio-sdk-go/service/conversations/v1"
	"github.com/RJPearson94/twilio-sdk-go/session"
)

// Conversations client is used to manage versioned resources for Twilio Conversations
// See https://www.twilio.com/docs/conversations for more details on the API
// See https://www.twilio.com/conversations for more details on the product
type Conversations struct {
	V1 *v1.Conversations
}

// New creates a new instance of the client using session data and config
func New(sess *session.Session, config *client.Config) *Conversations {
	return &Conversations{
		V1: v1.New(sess, config),
	}
}

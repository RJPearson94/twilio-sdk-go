package chat

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	v2 "github.com/RJPearson94/twilio-sdk-go/service/chat/v2"
	"github.com/RJPearson94/twilio-sdk-go/session"
)

// Chat client is used to manage versioned resources for Programmable Chat
// See https://www.twilio.com/docs/chat for more details on the API
// See https://www.twilio.com/chat for more details on the product
type Chat struct {
	V2 *v2.Chat
}

// New creates a new instance of the client using session data and config
func New(sess *session.Session, config *client.Config) *Chat {
	return &Chat{
		V2: v2.New(sess, config),
	}
}

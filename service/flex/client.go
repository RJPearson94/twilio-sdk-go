package flex

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	v1 "github.com/RJPearson94/twilio-sdk-go/service/flex/v1"
	"github.com/RJPearson94/twilio-sdk-go/session"
)

// Flex client is used to manage versioned resources for Twilio Flex
// See https://www.twilio.com/docs/flex for more details on the API
// See https://www.twilio.com/flex for more details on the product
type Flex struct {
	V1 *v1.Flex
}

// New creates a new instance of the client using session data and config
func New(sess *session.Session, config *client.Config) *Flex {
	return &Flex{
		V1: v1.New(sess, config),
	}
}

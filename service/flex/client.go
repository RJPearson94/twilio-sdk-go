package flex

import (
	v1 "github.com/RJPearson94/twilio-sdk-go/service/flex/v1"
	"github.com/RJPearson94/twilio-sdk-go/session"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

// Flex client is used to manage versioned resources for Twilio Flex
// See https://www.twilio.com/docs/flex for more details on the API
// See https://www.twilio.com/flex for more details on the product
type Flex struct {
	V1 *v1.Flex
}

// New creates a new instance of the client using session data
func New(sess *session.Session) *Flex {
	return &Flex{
		V1: v1.New(sess),
	}
}

// NewWithCredentials creates a new instance of the client with credentials
func NewWithCredentials(creds *credentials.Credentials) *Flex {
	return New(session.New(creds))
}

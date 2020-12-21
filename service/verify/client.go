package verify

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	v2 "github.com/RJPearson94/twilio-sdk-go/service/verify/v2"
	"github.com/RJPearson94/twilio-sdk-go/session"
)

// Verify client is used to manage versioned resources for Twilio Verify
// See https://www.twilio.com/docs/verify for more details on the API
// See https://www.twilio.com/verify for more details on the product
type Verify struct {
	V2 *v2.Verify
}

// New creates a new instance of the client using session data and config
func New(sess *session.Session, config *client.Config) *Verify {
	return &Verify{
		V2: v2.New(sess, config),
	}
}

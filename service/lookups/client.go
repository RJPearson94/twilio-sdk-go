package lookups

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	v1 "github.com/RJPearson94/twilio-sdk-go/service/lookups/v1"
	"github.com/RJPearson94/twilio-sdk-go/session"
)

// Lookups client is used to manage versioned resources for Programmable Lookups
// See https://www.twilio.com/lookup for more details on the API
// See https://www.twilio.com/docs/lookup for more details on the product
type Lookups struct {
	V1 *v1.Lookups
}

// New creates a new instance of the client using session data and config
func New(sess *session.Session, config *client.Config) *Lookups {
	return &Lookups{
		V1: v1.New(sess, config),
	}
}

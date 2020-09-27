package lookups

import (
	v1 "github.com/RJPearson94/twilio-sdk-go/service/lookups/v1"
	"github.com/RJPearson94/twilio-sdk-go/session"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

// Lookups client is used to manage versioned resources for Programmable Lookups
// See https://www.twilio.com/lookup for more details on the API
// See https://www.twilio.com/docs/lookup for more details on the product
type Lookups struct {
	V1 *v1.Lookups
}

// New creates a new instance of the client using session data
func New(sess *session.Session) *Lookups {
	return &Lookups{
		V1: v1.New(sess),
	}
}

// NewWithCredentials creates a new instance of the client with credentials
func NewWithCredentials(creds *credentials.Credentials) *Lookups {
	return New(session.New(creds))
}

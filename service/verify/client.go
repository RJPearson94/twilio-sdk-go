package verify

import (
	v2 "github.com/RJPearson94/twilio-sdk-go/service/verify/v2"
	"github.com/RJPearson94/twilio-sdk-go/session"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

// Verify client is used to manage versioned resources for Twilio Verify
// See https://www.twilio.com/docs/verify for more details on the API
// See https://www.twilio.com/verify for more details on the product
type Verify struct {
	V2 *v2.Verify
}

// New creates a new instance of the client using session data
func New(sess *session.Session) *Verify {
	return &Verify{
		V2: v2.New(sess),
	}
}

// NewWithCredentials creates a new instance of the client with credentials
func NewWithCredentials(creds *credentials.Credentials) *Verify {
	return New(session.New(creds))
}

package studio

import (
	v2 "github.com/RJPearson94/twilio-sdk-go/service/studio/v2"
	"github.com/RJPearson94/twilio-sdk-go/session"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

// Studio client is used to manage versioned resources for Twilio Studio
// See https://www.twilio.com/docs/studio for more details on the API
// See https://www.twilio.com/studio for more details on the product
type Studio struct {
	V2 *v2.Studio
}

// New creates a new instance of the client using session data
func New(sess *session.Session) *Studio {
	return &Studio{
		V2: v2.New(sess),
	}
}

// NewWithCredentials creates a new instance of the client with credentials
func NewWithCredentials(creds *credentials.Credentials) *Studio {
	return New(session.New(creds))
}

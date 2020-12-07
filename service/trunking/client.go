package trunking

import (
	v1 "github.com/RJPearson94/twilio-sdk-go/service/trunking/v1"
	"github.com/RJPearson94/twilio-sdk-go/session"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

// Trunking client is used to manage versioned resources for Twilio Trunking
// See https://www.twilio.com/docs/sip-trunking for more details on the API
// See https://www.twilio.com/sip-trunking for more details on the product
type Trunking struct {
	V1 *v1.Trunking
}

// New creates a new instance of the client using session data
func New(sess *session.Session) *Trunking {
	return &Trunking{
		V1: v1.New(sess),
	}
}

// NewWithCredentials creates a new instance of the client with credentials
func NewWithCredentials(creds *credentials.Credentials) *Trunking {
	return New(session.New(creds))
}

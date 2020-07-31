package fax

import (
	v1 "github.com/RJPearson94/twilio-sdk-go/service/fax/v1"
	"github.com/RJPearson94/twilio-sdk-go/session"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

// Fax client is used to manage versioned resources for Programmable Fax
// See https://www.twilio.com/fax for more details on the API
// See https://www.twilio.com/docs/fax for more details on the product
type Fax struct {
	V1 *v1.Fax
}

// New creates a new instance of the client using session data
func New(sess *session.Session) *Fax {
	return &Fax{
		V1: v1.New(sess),
	}
}

// NewWithCredentials creates a new instance of the client with credentials
func NewWithCredentials(creds *credentials.Credentials) *Fax {
	return New(session.New(creds))
}

package proxy

import (
	v1 "github.com/RJPearson94/twilio-sdk-go/service/proxy/v1"
	"github.com/RJPearson94/twilio-sdk-go/session"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

// Proxy client is used to manage versioned resources for Twilio Proxy
// See https://www.twilio.com/docs/proxy for more details on the API
// See https://www.twilio.com/proxy for more details on the product
type Proxy struct {
	V1 *v1.Proxy
}

// New creates a new instance of the client using session data
func New(sess *session.Session) *Proxy {
	return &Proxy{
		V1: v1.New(sess),
	}
}

// NewWithCredentials creates a new instance of the client with credentials
func NewWithCredentials(creds *credentials.Credentials) *Proxy {
	return New(session.New(creds))
}

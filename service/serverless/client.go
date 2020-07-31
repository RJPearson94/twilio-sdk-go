package serverless

import (
	v1 "github.com/RJPearson94/twilio-sdk-go/service/serverless/v1"
	"github.com/RJPearson94/twilio-sdk-go/session"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

// Serverless client is used to manage versioned resources for Twilio Serverless
// See https://www.twilio.com/docs/runtime for more details on the API
// See https://www.twilio.com/runtime for more details on the product
type Serverless struct {
	V1 *v1.Serverless
}

// New creates a new instance of the client using session data
func New(sess *session.Session) *Serverless {
	return &Serverless{
		V1: v1.New(sess),
	}
}

// NewWithCredentials creates a new instance of the client with credentials
func NewWithCredentials(creds *credentials.Credentials) *Serverless {
	return New(session.New(creds))
}

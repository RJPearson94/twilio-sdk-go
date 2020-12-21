package fax

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	v1 "github.com/RJPearson94/twilio-sdk-go/service/fax/v1"
	"github.com/RJPearson94/twilio-sdk-go/session"
)

// Fax client is used to manage versioned resources for Programmable Fax
// See https://www.twilio.com/fax for more details on the API
// See https://www.twilio.com/docs/fax for more details on the product
type Fax struct {
	V1 *v1.Fax
}

// New creates a new instance of the client using session data and config
func New(sess *session.Session, config *client.Config) *Fax {
	return &Fax{
		V1: v1.New(sess, config),
	}
}

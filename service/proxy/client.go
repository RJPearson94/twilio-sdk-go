package proxy

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	v1 "github.com/RJPearson94/twilio-sdk-go/service/proxy/v1"
	"github.com/RJPearson94/twilio-sdk-go/session"
)

// Proxy client is used to manage versioned resources for Twilio Proxy
// See https://www.twilio.com/docs/proxy for more details on the API
// See https://www.twilio.com/proxy for more details on the product
type Proxy struct {
	V1 *v1.Proxy
}

// New creates a new instance of the client using session data and config
func New(sess *session.Session, config *client.Config) *Proxy {
	return &Proxy{
		V1: v1.New(sess, config),
	}
}

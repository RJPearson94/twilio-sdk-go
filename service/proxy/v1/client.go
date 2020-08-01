package v1

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/proxy/v1/service"
	"github.com/RJPearson94/twilio-sdk-go/service/proxy/v1/services"
	"github.com/RJPearson94/twilio-sdk-go/session"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

// Proxy client is used to manage resources for Twilio Proxy
// See https://www.twilio.com/docs/proxy for more details
type Proxy struct {
	client   *client.Client
	Service  func(string) *service.Client
	Services *services.Client
}

// Used for testing purposes only
func (s Proxy) GetClient() *client.Client {
	return s.client
}

// New creates a new instance of the client using session data
func New(sess *session.Session) *Proxy {
	config := client.GetDefaultConfig()
	config.Beta = true
	config.SubDomain = "proxy"
	config.APIVersion = "v1"

	return NewWithClient(client.New(sess, config))
}

// NewWithClient creates a new instance of the client with a HTTP client
func NewWithClient(client *client.Client) *Proxy {
	return &Proxy{
		client:   client,
		Services: services.New(client),
		Service: func(sid string) *service.Client {
			return service.New(client, service.ClientProperties{
				Sid: sid,
			})
		},
	}
}

// NewWithCredentials creates a new instance of the client with credentials
func NewWithCredentials(creds *credentials.Credentials) *Proxy {
	return New(session.New(creds))
}

package v1

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/messaging/v1/service"
	"github.com/RJPearson94/twilio-sdk-go/service/messaging/v1/services"
	"github.com/RJPearson94/twilio-sdk-go/session"
	"github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

// Messaging client is used to manage resources for Twilio Messaging
// See https://www.twilio.com/docs/messaging for more details
type Messaging struct {
	client *client.Client

	Service  func(string) *service.Client
	Services *services.Client
}

// Used for testing purposes only
func (s Messaging) GetClient() *client.Client {
	return s.client
}

// New creates a new instance of the client using session data
func New(sess *session.Session) *Messaging {
	config := client.GetDefaultConfig()
	config.Beta = true
	config.SubDomain = "messaging"
	config.APIVersion = "v1"

	return NewWithClient(client.New(sess, config))
}

// NewWithClient creates a new instance of the client with a HTTP client
func NewWithClient(client *client.Client) *Messaging {
	return &Messaging{
		client: client,
		Service: func(sid string) *service.Client {
			return service.New(client, service.ClientProperties{
				Sid: sid,
			})
		},
		Services: services.New(client),
	}
}

// NewWithCredentials creates a new instance of the client with credentials
func NewWithCredentials(creds *credentials.Credentials) *Messaging {
	return New(session.New(creds))
}

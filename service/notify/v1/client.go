// Package v1 contains auto-generated files. DO NOT MODIFY
package v1

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/notify/v1/credential"
	"github.com/RJPearson94/twilio-sdk-go/service/notify/v1/credentials"
	"github.com/RJPearson94/twilio-sdk-go/service/notify/v1/service"
	"github.com/RJPearson94/twilio-sdk-go/service/notify/v1/services"
	"github.com/RJPearson94/twilio-sdk-go/session"
)

// Notify client is used to manage resources for Twilio Notify
// See https://www.twilio.com/docs/notify for more details
type Notify struct {
	client *client.Client

	Credential  func(string) *credential.Client
	Credentials *credentials.Client
	Service     func(string) *service.Client
	Services    *services.Client
}

// NewWithClient creates a new instance of the client with a HTTP client
func NewWithClient(client *client.Client) *Notify {
	return &Notify{
		client: client,

		Credential: func(credentialSid string) *credential.Client {
			return credential.New(client, credential.ClientProperties{
				Sid: credentialSid,
			})
		},
		Credentials: credentials.New(client),
		Service: func(serviceSid string) *service.Client {
			return service.New(client, service.ClientProperties{
				Sid: serviceSid,
			})
		},
		Services: services.New(client),
	}
}

// GetClient is used for testing purposes only
func (s Notify) GetClient() *client.Client {
	return s.client
}

// New creates a new instance of the client using session data and config
func New(sess *session.Session, clientConfig *client.Config) *Notify {
	config := client.NewAPIClientConfig(clientConfig)
	config.Beta = false
	config.SubDomain = "notify"
	config.APIVersion = "v1"

	return NewWithClient(client.New(sess, config))
}

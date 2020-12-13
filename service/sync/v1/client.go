// Package v1 contains auto-generated files. DO NOT MODIFY
package v1

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/sync/v1/service"
	"github.com/RJPearson94/twilio-sdk-go/service/sync/v1/services"
	"github.com/RJPearson94/twilio-sdk-go/session"
	sessionCredentials "github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

// Sync client is used to manage resources for Twilio Sync
// See https://www.twilio.com/docs/sync for more details
// This client is currently in beta and subject to change. Please use with caution
type Sync struct {
	client *client.Client

	Service  func(string) *service.Client
	Services *services.Client
}

// NewWithClient creates a new instance of the client with a HTTP client
func NewWithClient(client *client.Client) *Sync {
	return &Sync{
		client: client,

		Service: func(serviceSid string) *service.Client {
			return service.New(client, service.ClientProperties{
				Sid: serviceSid,
			})
		},
		Services: services.New(client),
	}
}

// GetClient is used for testing purposes only
func (s Sync) GetClient() *client.Client {
	return s.client
}

// NewWithCredentials creates a new instance of the client with credentials
func NewWithCredentials(creds *sessionCredentials.Credentials) *Sync {
	return New(session.New(creds))
}

// New creates a new instance of the client using session data
func New(sess *session.Session) *Sync {
	config := client.GetDefaultConfig()
	config.Beta = true
	config.SubDomain = "sync"
	config.APIVersion = "v1"

	return NewWithClient(client.New(sess, config))
}

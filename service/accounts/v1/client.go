// Package v1 contains auto-generated files. DO NOT MODIFY
package v1

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/accounts/v1/credentials"
	"github.com/RJPearson94/twilio-sdk-go/session"
)

// Client for managing account resources
type Accounts struct {
	client *client.Client

	Credentials *credentials.Client
}

// NewWithClient creates a new instance of the client with a HTTP client
func NewWithClient(client *client.Client) *Accounts {
	return &Accounts{
		client: client,

		Credentials: credentials.New(client),
	}
}

// GetClient is used for testing purposes only
func (s Accounts) GetClient() *client.Client {
	return s.client
}

// New creates a new instance of the client using session data and config
func New(sess *session.Session, clientConfig *client.Config) *Accounts {
	config := client.NewAPIClientConfig(clientConfig)
	config.Beta = false
	config.SubDomain = "accounts"
	config.APIVersion = "v1"

	return NewWithClient(client.New(sess, config))
}

// Package v1 contains auto-generated files. DO NOT MODIFY
package v1

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/lookups/v1/phone_number"
	"github.com/RJPearson94/twilio-sdk-go/session"
	sessionCredentials "github.com/RJPearson94/twilio-sdk-go/session/credentials"
)

// Lookups client is used to manage resources for Lookups
// See https://www.twilio.com/docs/lookup for more details
type Lookups struct {
	client *client.Client

	PhoneNumber func(string) *phone_number.Client
}

// NewWithClient creates a new instance of the client with a HTTP client
func NewWithClient(client *client.Client) *Lookups {
	return &Lookups{
		client: client,

		PhoneNumber: func(phoneNumber string) *phone_number.Client {
			return phone_number.New(client, phone_number.ClientProperties{
				PhoneNumber: phoneNumber,
			})
		},
	}
}

// GetClient is used for testing purposes only
func (s Lookups) GetClient() *client.Client {
	return s.client
}

// NewWithCredentials creates a new instance of the client with credentials
func NewWithCredentials(creds *sessionCredentials.Credentials) *Lookups {
	return New(session.New(creds))
}

// New creates a new instance of the client using session data
func New(sess *session.Session) *Lookups {
	config := client.GetDefaultConfig()
	config.Beta = false
	config.SubDomain = "lookups"
	config.APIVersion = "v1"

	return NewWithClient(client.New(sess, config))
}

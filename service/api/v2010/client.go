// Package v2010 contains auto-generated files. DO NOT MODIFY
package v2010

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/accounts"
	"github.com/RJPearson94/twilio-sdk-go/session"
)

// V2010 client to manage resources that are part of the Twilio API
type V2010 struct {
	client *client.Client

	Account  func(string) *account.Client
	Accounts *accounts.Client
}

// NewWithClient creates a new instance of the client with a HTTP client
func NewWithClient(client *client.Client) *V2010 {
	return &V2010{
		client: client,

		Account: func(accountSid string) *account.Client {
			return account.New(client, account.ClientProperties{
				Sid: accountSid,
			})
		},
		Accounts: accounts.New(client),
	}
}

// GetClient is used for testing purposes only
func (s V2010) GetClient() *client.Client {
	return s.client
}

// New creates a new instance of the client using session data and config
func New(sess *session.Session, clientConfig *client.Config) *V2010 {
	config := client.NewAPIClientConfig(clientConfig)
	config.Beta = false
	config.SubDomain = "api"
	config.APIVersion = "2010-04-01"

	return NewWithClient(client.New(sess, config))
}

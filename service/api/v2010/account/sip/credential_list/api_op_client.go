// Package credential_list contains auto-generated files. DO NOT MODIFY
package credential_list

import (
	"github.com/RJPearson94/twilio-sdk-go/client"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/sip/credential_list/credential"
	"github.com/RJPearson94/twilio-sdk-go/service/api/v2010/account/sip/credential_list/credentials"
)

// Client for managing a specific Credential List resource
type Client struct {
	client *client.Client

	accountSid string
	sid        string

	Credential  func(string) *credential.Client
	Credentials *credentials.Client
}

// ClientProperties are the properties required to manage the credential list resources
type ClientProperties struct {
	AccountSid string
	Sid        string
}

// New creates a new instance of the credential list client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		accountSid: properties.AccountSid,
		sid:        properties.Sid,

		Credential: func(credentialSid string) *credential.Client {
			return credential.New(client, credential.ClientProperties{
				AccountSid:        properties.AccountSid,
				CredentialListSid: properties.Sid,
				Sid:               credentialSid,
			})
		},
		Credentials: credentials.New(client, credentials.ClientProperties{
			AccountSid:        properties.AccountSid,
			CredentialListSid: properties.Sid,
		}),
	}
}

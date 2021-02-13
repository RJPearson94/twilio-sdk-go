// Package credentials contains auto-generated files. DO NOT MODIFY
package credentials

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing credential resources
type Client struct {
	client *client.Client

	accountSid        string
	credentialListSid string
}

// ClientProperties are the properties required to manage the credentials resources
type ClientProperties struct {
	AccountSid        string
	CredentialListSid string
}

// New creates a new instance of the credentials client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		accountSid:        properties.AccountSid,
		credentialListSid: properties.CredentialListSid,
	}
}

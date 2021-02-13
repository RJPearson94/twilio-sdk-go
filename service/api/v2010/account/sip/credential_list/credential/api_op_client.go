// Package credential contains auto-generated files. DO NOT MODIFY
package credential

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing a specific credential resource
type Client struct {
	client *client.Client

	accountSid        string
	credentialListSid string
	sid               string
}

// ClientProperties are the properties required to manage the credential resources
type ClientProperties struct {
	AccountSid        string
	CredentialListSid string
	Sid               string
}

// New creates a new instance of the credential client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		accountSid:        properties.AccountSid,
		credentialListSid: properties.CredentialListSid,
		sid:               properties.Sid,
	}
}

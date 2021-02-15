// Package credential_list_mappings contains auto-generated files. DO NOT MODIFY
package credential_list_mappings

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing credential list mapping resources
type Client struct {
	client *client.Client

	accountSid string
	domainSid  string
}

// ClientProperties are the properties required to manage the credential list mappings resources
type ClientProperties struct {
	AccountSid string
	DomainSid  string
}

// New creates a new instance of the credential list mappings client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		accountSid: properties.AccountSid,
		domainSid:  properties.DomainSid,
	}
}

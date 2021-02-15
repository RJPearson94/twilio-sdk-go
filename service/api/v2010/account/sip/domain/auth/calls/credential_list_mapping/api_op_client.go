// Package credential_list_mapping contains auto-generated files. DO NOT MODIFY
package credential_list_mapping

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing a specific Credential List Mapping resource
type Client struct {
	client *client.Client

	accountSid string
	domainSid  string
	sid        string
}

// ClientProperties are the properties required to manage the credential list mapping resources
type ClientProperties struct {
	AccountSid string
	DomainSid  string
	Sid        string
}

// New creates a new instance of the credential list mapping client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		accountSid: properties.AccountSid,
		domainSid:  properties.DomainSid,
		sid:        properties.Sid,
	}
}

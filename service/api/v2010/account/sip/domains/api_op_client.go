// Package domains contains auto-generated files. DO NOT MODIFY
package domains

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing SIP domain resources
type Client struct {
	client *client.Client

	accountSid string
}

// ClientProperties are the properties required to manage the domains resources
type ClientProperties struct {
	AccountSid string
}

// New creates a new instance of the domains client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		accountSid: properties.AccountSid,
	}
}

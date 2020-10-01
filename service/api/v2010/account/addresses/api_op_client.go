// Package addresses contains auto-generated files. DO NOT MODIFY
package addresses

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing address resources
// See https://www.twilio.com/docs/usage/api/address for more details
type Client struct {
	client *client.Client

	accountSid string
}

// ClientProperties are the properties required to manage the addresses resources
type ClientProperties struct {
	AccountSid string
}

// New creates a new instance of the addresses client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		accountSid: properties.AccountSid,
	}
}

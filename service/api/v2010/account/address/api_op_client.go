// Package address contains auto-generated files. DO NOT MODIFY
package address

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing a specific address resource
// See https://www.twilio.com/docs/usage/api/address for more details
type Client struct {
	client *client.Client

	accountSid string
	sid        string
}

// ClientProperties are the properties required to manage the address resources
type ClientProperties struct {
	AccountSid string
	Sid        string
}

// New creates a new instance of the address client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		accountSid: properties.AccountSid,
		sid:        properties.Sid,
	}
}

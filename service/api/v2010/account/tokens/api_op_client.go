// Package tokens contains auto-generated files. DO NOT MODIFY
package tokens

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing token resources
type Client struct {
	client *client.Client

	accountSid string
}

// ClientProperties are the properties required to manage the tokens resources
type ClientProperties struct {
	AccountSid string
}

// New creates a new instance of the tokens client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		accountSid: properties.AccountSid,
	}
}

// Package applications contains auto-generated files. DO NOT MODIFY
package applications

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing application resources
// See https://www.twilio.com/docs/usage/api/applications for more details
type Client struct {
	client *client.Client

	accountSid string
}

// ClientProperties are the properties required to manage the applications resources
type ClientProperties struct {
	AccountSid string
}

// New creates a new instance of the applications client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		accountSid: properties.AccountSid,
	}
}

// Package conferences contains auto-generated files. DO NOT MODIFY
package conferences

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing conference resources
// See https://www.twilio.com/docs/voice/api/conference-resource for more details
type Client struct {
	client *client.Client

	accountSid string
}

// ClientProperties are the properties required to manage the conferences resources
type ClientProperties struct {
	AccountSid string
}

// New creates a new instance of the conferences client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		accountSid: properties.AccountSid,
	}
}

// Package call contains auto-generated files. DO NOT MODIFY
package call

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing a specific call resource
// See https://www.twilio.com/docs/voice/api/call-resource for more details
type Client struct {
	client *client.Client

	accountSid string
	sid        string
}

// ClientProperties are the properties required to manage the call resources
type ClientProperties struct {
	AccountSid string
	Sid        string
}

// New creates a new instance of the call client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		accountSid: properties.AccountSid,
		sid:        properties.Sid,
	}
}

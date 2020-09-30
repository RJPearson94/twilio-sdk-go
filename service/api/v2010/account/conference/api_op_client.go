// Package conference contains auto-generated files. DO NOT MODIFY
package conference

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing a specific conference resource
// See https://www.twilio.com/docs/voice/api/conference-resource for more details
type Client struct {
	client *client.Client

	accountSid string
	sid        string
}

// ClientProperties are the properties required to manage the conference resources
type ClientProperties struct {
	AccountSid string
	Sid        string
}

// New creates a new instance of the conference client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		accountSid: properties.AccountSid,
		sid:        properties.Sid,
	}
}

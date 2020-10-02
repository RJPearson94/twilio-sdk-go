// Package recording contains auto-generated files. DO NOT MODIFY
package recording

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing a specific recording resource
// See https://www.twilio.com/docs/voice/api/recording for more details
type Client struct {
	client *client.Client

	accountSid    string
	conferenceSid string
	sid           string
}

// ClientProperties are the properties required to manage the recording resources
type ClientProperties struct {
	AccountSid    string
	ConferenceSid string
	Sid           string
}

// New creates a new instance of the recording client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		accountSid:    properties.AccountSid,
		conferenceSid: properties.ConferenceSid,
		sid:           properties.Sid,
	}
}

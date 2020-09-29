// Package media_attachment contains auto-generated files. DO NOT MODIFY
package media_attachment

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing a specific message media attachment resource
// See https://www.twilio.com/docs/sms/api/media-resource for more details
type Client struct {
	client *client.Client

	accountSid string
	messageSid string
	sid        string
}

// ClientProperties are the properties required to manage the media attachment resources
type ClientProperties struct {
	AccountSid string
	MessageSid string
	Sid        string
}

// New creates a new instance of the media attachment client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		accountSid: properties.AccountSid,
		messageSid: properties.MessageSid,
		sid:        properties.Sid,
	}
}

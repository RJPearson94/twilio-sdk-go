// Package member contains auto-generated files. DO NOT MODIFY
package member

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing a specific member resource
// See https://www.twilio.com/docs/voice/api/member-resource for more details
type Client struct {
	client *client.Client

	accountSid string
	queueSid   string
	sid        string
}

// ClientProperties are the properties required to manage the member resources
type ClientProperties struct {
	AccountSid string
	QueueSid   string
	Sid        string
}

// New creates a new instance of the member client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		accountSid: properties.AccountSid,
		queueSid:   properties.QueueSid,
		sid:        properties.Sid,
	}
}

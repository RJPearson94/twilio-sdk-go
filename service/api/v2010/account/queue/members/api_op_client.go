// Package members contains auto-generated files. DO NOT MODIFY
package members

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing member resources
// See https://www.twilio.com/docs/voice/api/member-resource for more details
type Client struct {
	client *client.Client

	accountSid string
	queueSid   string
}

// ClientProperties are the properties required to manage the members resources
type ClientProperties struct {
	AccountSid string
	QueueSid   string
}

// New creates a new instance of the members client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		accountSid: properties.AccountSid,
		queueSid:   properties.QueueSid,
	}
}

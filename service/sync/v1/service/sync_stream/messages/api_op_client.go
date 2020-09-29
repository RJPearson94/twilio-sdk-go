// Package messages contains auto-generated files. DO NOT MODIFY
package messages

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing stream message resources
// See https://www.twilio.com/docs/sync/api/stream-message-resource for more details
type Client struct {
	client *client.Client

	serviceSid    string
	syncStreamSid string
}

// ClientProperties are the properties required to manage the messages resources
type ClientProperties struct {
	ServiceSid    string
	SyncStreamSid string
}

// New creates a new instance of the messages client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		serviceSid:    properties.ServiceSid,
		syncStreamSid: properties.SyncStreamSid,
	}
}

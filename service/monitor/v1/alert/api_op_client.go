// Package alert contains auto-generated files. DO NOT MODIFY
package alert

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing a specific alert resource
// See https://www.twilio.com/docs/usage/monitor-alert for more details
type Client struct {
	client *client.Client

	sid string
}

// ClientProperties are the properties required to manage the alert resources
type ClientProperties struct {
	Sid string
}

// New creates a new instance of the alert client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		sid: properties.Sid,
	}
}

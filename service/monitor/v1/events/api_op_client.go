// Package events contains auto-generated files. DO NOT MODIFY
package events

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing event resources
// See https://www.twilio.com/docs/usage/monitor-events for more details
type Client struct {
	client *client.Client
}

// New creates a new instance of the events client
func New(client *client.Client) *Client {
	return &Client{
		client: client,
	}
}

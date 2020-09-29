// Package assistants contains auto-generated files. DO NOT MODIFY
package assistants

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing assistant resources
// See https://www.twilio.com/docs/autopilot/api/assistant for more details
type Client struct {
	client *client.Client
}

// New creates a new instance of the assistants client
func New(client *client.Client) *Client {
	return &Client{
		client: client,
	}
}

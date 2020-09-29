// Package configuration contains auto-generated files. DO NOT MODIFY
package configuration

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing a specific configuration resource
// See https://www.twilio.com/docs/flex/ui/configuration for more details
type Client struct {
	client *client.Client
}

// New creates a new instance of the configuration client
func New(client *client.Client) *Client {
	return &Client{
		client: client,
	}
}

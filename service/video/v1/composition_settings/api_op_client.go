// Package composition_settings contains auto-generated files. DO NOT MODIFY
package composition_settings

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing the default composition settings
// See https://www.twilio.com/docs/video/api/encrypted-compositions for more details
type Client struct {
	client *client.Client
}

// New creates a new instance of the composition settings client
func New(client *client.Client) *Client {
	return &Client{
		client: client,
	}
}

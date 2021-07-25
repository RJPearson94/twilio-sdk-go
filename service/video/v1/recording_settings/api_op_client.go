// Package recording_settings contains auto-generated files. DO NOT MODIFY
package recording_settings

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing the default recording settings
// See https://www.twilio.com/docs/video/api/encrypted-recordings for more details
type Client struct {
	client *client.Client
}

// New creates a new instance of the recording settings client
func New(client *client.Client) *Client {
	return &Client{
		client: client,
	}
}

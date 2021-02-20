// Package rooms contains auto-generated files. DO NOT MODIFY
package rooms

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing room resources
// See https://www.twilio.com/docs/video/api/rooms-resource for more details
type Client struct {
	client *client.Client
}

// New creates a new instance of the rooms client
func New(client *client.Client) *Client {
	return &Client{
		client: client,
	}
}

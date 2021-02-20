// Package compositions contains auto-generated files. DO NOT MODIFY
package compositions

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing composition resources
// See https://www.twilio.com/docs/video/api/compositions-resource for more details
type Client struct {
	client *client.Client
}

// New creates a new instance of the compositions client
func New(client *client.Client) *Client {
	return &Client{
		client: client,
	}
}

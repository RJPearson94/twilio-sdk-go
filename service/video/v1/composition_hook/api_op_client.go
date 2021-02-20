// Package composition_hook contains auto-generated files. DO NOT MODIFY
package composition_hook

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing a specific composition hook resource
// See https://www.twilio.com/docs/video/api/composition-hooks for more details
type Client struct {
	client *client.Client

	sid string
}

// ClientProperties are the properties required to manage the composition hook resources
type ClientProperties struct {
	Sid string
}

// New creates a new instance of the composition hook client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		sid: properties.Sid,
	}
}

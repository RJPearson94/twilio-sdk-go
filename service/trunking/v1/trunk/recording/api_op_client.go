// Package recording contains auto-generated files. DO NOT MODIFY
package recording

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing a specific recording resource
type Client struct {
	client *client.Client

	trunkSid string
}

// ClientProperties are the properties required to manage the recording resources
type ClientProperties struct {
	TrunkSid string
}

// New creates a new instance of the recording client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		trunkSid: properties.TrunkSid,
	}
}

// Package defaults contains auto-generated files. DO NOT MODIFY
package defaults

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing default resources
// See https://www.twilio.com/docs/autopilot/api/assistant/defaults for more details
type Client struct {
	client *client.Client

	assistantSid string
}

// ClientProperties are the properties required to manage the defaults resources
type ClientProperties struct {
	AssistantSid string
}

// New creates a new instance of the defaults client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		assistantSid: properties.AssistantSid,
	}
}

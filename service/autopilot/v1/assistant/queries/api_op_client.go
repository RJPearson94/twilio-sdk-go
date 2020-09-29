// Package queries contains auto-generated files. DO NOT MODIFY
package queries

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing query resources
// See https://www.twilio.com/docs/autopilot/api/query for more details
type Client struct {
	client *client.Client

	assistantSid string
}

// ClientProperties are the properties required to manage the queries resources
type ClientProperties struct {
	AssistantSid string
}

// New creates a new instance of the queries client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		assistantSid: properties.AssistantSid,
	}
}

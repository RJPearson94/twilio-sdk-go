// Package dialogue contains auto-generated files. DO NOT MODIFY
package dialogue

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing a specific dialogue resource
// See https://www.twilio.com/docs/autopilot/api/dialogue for more details
type Client struct {
	client *client.Client

	assistantSid string
	sid          string
}

// ClientProperties are the properties required to manage the dialogue resources
type ClientProperties struct {
	AssistantSid string
	Sid          string
}

// New creates a new instance of the dialogue client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		assistantSid: properties.AssistantSid,
		sid:          properties.Sid,
	}
}

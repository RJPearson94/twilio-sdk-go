// Package participant contains auto-generated files. DO NOT MODIFY
package participant

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing a specific participant resource
// See https://www.twilio.com/docs/video/api/participants for more details
type Client struct {
	client *client.Client

	roomSid string
	sid     string
}

// ClientProperties are the properties required to manage the participant resources
type ClientProperties struct {
	RoomSid string
	Sid     string
}

// New creates a new instance of the participant client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		roomSid: properties.RoomSid,
		sid:     properties.Sid,
	}
}

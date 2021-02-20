// Package recording contains auto-generated files. DO NOT MODIFY
package recording

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing a specific recording resource
// See https://www.twilio.com/docs/video/api/rooms-resource for more details
type Client struct {
	client *client.Client

	roomSid string
	sid     string
}

// ClientProperties are the properties required to manage the recording resources
type ClientProperties struct {
	RoomSid string
	Sid     string
}

// New creates a new instance of the recording client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		roomSid: properties.RoomSid,
		sid:     properties.Sid,
	}
}

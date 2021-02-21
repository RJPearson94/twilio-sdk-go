// Package published_track contains auto-generated files. DO NOT MODIFY
package published_track

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing a specific published track resource
// See https://www.twilio.com/docs/video/api/participants for more details
type Client struct {
	client *client.Client

	participantSid string
	roomSid        string
	sid            string
}

// ClientProperties are the properties required to manage the published track resources
type ClientProperties struct {
	ParticipantSid string
	RoomSid        string
	Sid            string
}

// New creates a new instance of the published track client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		participantSid: properties.ParticipantSid,
		roomSid:        properties.RoomSid,
		sid:            properties.Sid,
	}
}

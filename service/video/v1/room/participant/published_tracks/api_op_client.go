// Package published_tracks contains auto-generated files. DO NOT MODIFY
package published_tracks

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing published track resources
// See https://www.twilio.com/docs/video/api/track-subscriptions for more details
type Client struct {
	client *client.Client

	participantSid string
	roomSid        string
}

// ClientProperties are the properties required to manage the published tracks resources
type ClientProperties struct {
	ParticipantSid string
	RoomSid        string
}

// New creates a new instance of the published tracks client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		participantSid: properties.ParticipantSid,
		roomSid:        properties.RoomSid,
	}
}

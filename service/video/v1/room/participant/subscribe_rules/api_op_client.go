// Package subscribe_rules contains auto-generated files. DO NOT MODIFY
package subscribe_rules

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing a specific subscribe rules resource
// See https://www.twilio.com/docs/video/api/track-subscriptions#sl-resource for more details
type Client struct {
	client *client.Client

	participantSid string
	roomSid        string
}

// ClientProperties are the properties required to manage the subscribe rules resources
type ClientProperties struct {
	ParticipantSid string
	RoomSid        string
}

// New creates a new instance of the subscribe rules client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		participantSid: properties.ParticipantSid,
		roomSid:        properties.RoomSid,
	}
}

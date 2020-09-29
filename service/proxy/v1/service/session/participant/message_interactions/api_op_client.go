// Package message_interactions contains auto-generated files. DO NOT MODIFY
package message_interactions

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing message interaction resources
// See https://www.twilio.com/docs/proxy/api/sending-messages for more details
type Client struct {
	client *client.Client

	participantSid string
	serviceSid     string
	sessionSid     string
}

// ClientProperties are the properties required to manage the message interactions resources
type ClientProperties struct {
	ParticipantSid string
	ServiceSid     string
	SessionSid     string
}

// New creates a new instance of the message interactions client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		participantSid: properties.ParticipantSid,
		serviceSid:     properties.ServiceSid,
		sessionSid:     properties.SessionSid,
	}
}

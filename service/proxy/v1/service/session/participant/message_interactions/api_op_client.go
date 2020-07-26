// This is an autogenerated file. DO NOT MODIFY
package message_interactions

import "github.com/RJPearson94/twilio-sdk-go/client"

type Client struct {
	client *client.Client

	participantSid string
	serviceSid     string
	sessionSid     string
}

// The properties required to manage the message interactions resources
type ClientProperties struct {
	ParticipantSid string
	ServiceSid     string
	SessionSid     string
}

// Create a new instance of the client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		participantSid: properties.ParticipantSid,
		serviceSid:     properties.ServiceSid,
		sessionSid:     properties.SessionSid,
	}
}

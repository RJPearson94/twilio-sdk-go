// Package participant contains auto-generated files. DO NOT MODIFY
package participant

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing a specific participant resource
// See https://www.twilio.com/docs/conversations/api/conversation-participant-resource for more details
type Client struct {
	client *client.Client

	conversationSid string
	serviceSid      string
	sid             string
}

// ClientProperties are the properties required to manage the participant resources
type ClientProperties struct {
	ConversationSid string
	ServiceSid      string
	Sid             string
}

// New creates a new instance of the participant client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		conversationSid: properties.ConversationSid,
		serviceSid:      properties.ServiceSid,
		sid:             properties.Sid,
	}
}

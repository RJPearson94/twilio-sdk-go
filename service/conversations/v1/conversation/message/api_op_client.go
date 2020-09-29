// Package message contains auto-generated files. DO NOT MODIFY
package message

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing a specific message resource
// See https://www.twilio.com/docs/conversations/api/conversation-message-resource for more details
type Client struct {
	client *client.Client

	conversationSid string
	sid             string
}

// ClientProperties are the properties required to manage the message resources
type ClientProperties struct {
	ConversationSid string
	Sid             string
}

// New creates a new instance of the message client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		conversationSid: properties.ConversationSid,
		sid:             properties.Sid,
	}
}

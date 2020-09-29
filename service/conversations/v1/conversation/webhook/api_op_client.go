// Package webhook contains auto-generated files. DO NOT MODIFY
package webhook

import "github.com/RJPearson94/twilio-sdk-go/client"

// Client for managing a specific webhook resource
// See https://www.twilio.com/docs/conversations/api/conversation-scoped-webhook-resource for more details
type Client struct {
	client *client.Client

	conversationSid string
	sid             string
}

// ClientProperties are the properties required to manage the webhook resources
type ClientProperties struct {
	ConversationSid string
	Sid             string
}

// New creates a new instance of the webhook client
func New(client *client.Client, properties ClientProperties) *Client {
	return &Client{
		client: client,

		conversationSid: properties.ConversationSid,
		sid:             properties.Sid,
	}
}
